package powerdns

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	ser "github.com/reconquest/ser-go"
)

const (
	apiDefaultVersion string = "v1"

	apiResourceServers string = "servers"
	apiResourceZones   string = "zones"
)

type PowerDNSClient struct {
	apiKey     string
	dnsDSN     string
	apiVer     string
	httpClient *http.Client
}

func NewPowerDNSClient(
	dnsDSN string,
	apiKey string,
	apiVer string,
) *PowerDNSClient {

	if apiVer == "" {
		apiVer = apiDefaultVersion
	}

	return &PowerDNSClient{
		apiKey, dnsDSN, apiVer, &http.Client{},
	}
}

func (client *PowerDNSClient) GetServers() ([]*Server, error) {
	request, err := client.makeRequest(
		http.MethodGet,
		client.makeRequestURL(apiResourceServers),
		nil,
	)
	if err != nil {
		return nil, err
	}

	response, err := client.executeRequest(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	servers := []*Server{}
	err = client.checkAndDecodeResponse(
		response,
		&servers,
	)

	return servers, err
}

func (client *PowerDNSClient) GetServer(
	name string,
) (*Server, error) {
	request, err := client.makeRequest(
		http.MethodGet,
		client.makeRequestURL(apiResourceServers, name),
		nil,
	)
	if err != nil {
		return nil, err
	}

	response, err := client.executeRequest(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	server := &Server{}
	err = client.checkAndDecodeResponse(
		response,
		server,
	)

	return server, err
}

func (client *PowerDNSClient) GetZones(
	server string,
) ([]*BasicZoneInfo, error) {
	request, err := client.makeRequest(
		http.MethodGet,
		client.makeRequestURL(
			apiResourceServers,
			server,
			apiResourceZones,
		),
		nil,
	)
	if err != nil {
		return nil, err
	}

	response, err := client.executeRequest(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	zones := []*BasicZoneInfo{}

	err = client.checkAndDecodeResponse(
		response,
		&zones,
	)

	return zones, err
}

func (client *PowerDNSClient) GetZone(
	server string,
	zone string,
) (*Zone, error) {
	request, err := client.makeRequest(
		http.MethodGet,
		client.makeRequestURL(
			apiResourceServers,
			server,
			apiResourceZones,
			zone,
		),
		nil,
	)
	if err != nil {
		return nil, err
	}

	response, err := client.executeRequest(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	dnsZone := &Zone{}

	err = client.checkAndDecodeResponse(
		response,
		dnsZone,
	)

	return dnsZone, err
}

func (client *PowerDNSClient) UpdateZone(
	server string,
	zone string,
	rrSetsPayload []*RRSet,
) error {
	var payload []byte
	payloadBuffer := bytes.NewBuffer(payload)

	err := json.NewEncoder(payloadBuffer).Encode(
		map[string][]*RRSet{
			"rrsets": rrSetsPayload,
		},
	)
	if err != nil {
		return err
	}

	request, err := client.makeRequest(
		http.MethodPatch,
		client.makeRequestURL(
			apiResourceServers,
			server,
			apiResourceZones,
			zone,
		),
		payloadBuffer,
	)
	if err != nil {
		return err
	}

	response, err := client.executeRequest(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return client.checkAndDecodeResponse(
		response,
		nil,
	)

}

func (client *PowerDNSClient) makeRequestURL(
	resources ...string,
) string {

	resultResource := ""
	for _, resource := range resources {
		resultResource += fmt.Sprintf("/%s", resource)
	}

	return fmt.Sprintf(
		"http://%s/api/%s%s",
		client.dnsDSN,
		client.apiVer,
		resultResource,
	)
}

func (client *PowerDNSClient) makeRequest(
	method string,
	url string,
	payload io.Reader,
) (*http.Request, error) {
	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, ser.Errorf(
			err,
			"can`t create HTTP request",
		)
	}

	request.Header.Add("X-API-Key", client.apiKey)

	return request, nil
}

func (client *PowerDNSClient) executeRequest(
	request *http.Request,
) (*http.Response, error) {
	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, ser.Errorf(
			err,
			"can`t execute HTTP request %v, reason: %s",
			request,
			err.Error(),
		)
	}

	return response, err
}

func (client *PowerDNSClient) checkAndDecodeResponse(
	response *http.Response,
	successAnswer interface{},
) error {
	errorAnswer := &PowerDNSAPIError{}

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated:
		err := json.NewDecoder(response.Body).Decode(successAnswer)
		if err != nil {
			return ser.Errorf(
				err,
				"can`t decode success answer from response body, reason: %s",
				err.Error(),
			)
		}

	case http.StatusNoContent:
		return nil

	case http.StatusInternalServerError, http.StatusBadRequest,
		http.StatusUnprocessableEntity:
		err := json.NewDecoder(response.Body).Decode(errorAnswer)
		if err != nil {
			return ser.Errorf(
				err,
				"can`t decode error answer from response body, reason: %s",
				err.Error(),
			)
		}

		return errors.New(errorAnswer.Error)

	case http.StatusNotFound:
		return fmt.Errorf(
			"requested URL %s was not found",
			response.Request.URL.String(),
		)

	default:
		return fmt.Errorf(
			"unsupported HTTP code returned: %d",
			response.StatusCode,
		)
	}

	return nil
}
