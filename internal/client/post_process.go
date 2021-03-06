package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type PostProcess struct {
	Id            string                 `json:"id,omitempty"`
	Name          string                 `json:"name"`
	Service       string                 `json:"service"`
	GroupId       int                    `json:"group_id,omitempty"`
	SharedByGroup *bool                  `json:"shared_by_group,omitempty"`
	Parameters    map[string]interface{} `json:"parameters"`
}

type PostProcessGetResponse struct {
	Data json.RawMessage `json:"data"`
}

type PostProcessPostResponse struct {
	Data json.RawMessage `json:"data"`
}

type PostProcessPatchResponse struct {
	Data json.RawMessage `json:"data"`
}

type PostProcessListResponse struct {
	Data  []json.RawMessage `json:"data"`
	Links struct {
		Self  string `json:"self"`
		First string `json:"first"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Total int `json:"total"`
	} `json:"meta"`
}

type RawPostProcessData struct {
	Id         string                `json:"id"`
	Type       string                `json:"type"`
	Attributes PostProcessAttributes `json:"attributes"`
}

type PostProcessAttributes struct {
	Name          string                 `json:"name"`
	Service       string                 `json:"service"`
	GroupId       int                    `json:"group_id"`
	SharedByGroup *bool                  `json:"shared_by_group"`
	Parameters    map[string]interface{} `json:"parameters"`
}

func (c *Client) GetPostProcess(postProcessId string) (*PostProcess, *http.Response, error) {
	requestUrl := fmt.Sprintf("post_processes/%s", postProcessId)
	getResponse := new(PostProcessGetResponse)
	resp, err := c.requestWithRetry("GET", requestUrl, nil, getResponse, defaultRetryCount)
	if err != nil {
		return nil, resp, err
	}

	postProcess := new(PostProcess)
	json.Unmarshal(getResponse.Data, &postProcess)

	return postProcess, resp, nil
}

func (c *Client) GetPostProcesses() (*[]PostProcess, *http.Response, error) {
	postProcesses := []PostProcess{}
	requestUrl := "post_processes"

	for len(requestUrl) > 0 {
		rel, err := url.Parse(requestUrl)
		if err != nil {
			return nil, nil, err
		}

		q := rel.Query()
		q.Set("page[size]", "100")
		rel.RawQuery = q.Encode()

		listResponse := new(PostProcessListResponse)
		resp, err := c.requestWithRetry("GET", rel.String(), nil, listResponse, defaultRetryCount)
		if err != nil {
			return nil, resp, err
		}

		for _, r := range listResponse.Data {
			postProcess := new(PostProcess)
			if err := json.Unmarshal(r, &postProcess); err != nil {
				return nil, nil, errors.New("unmarshal failed")
			}
			postProcesses = append(postProcesses, *postProcess)
		}

		requestUrl = listResponse.Links.Next
	}

	return &postProcesses, nil, nil
}

func (c *Client) CreatePostProcess(postProcess *PostProcess) (*PostProcess, *http.Response, error) {
	requestUrl := "post_processes"
	postResponse := new(PostProcessPostResponse)
	resp, err := c.requestWithRetry("POST", requestUrl, postProcess, postResponse, defaultRetryCount)
	if err != nil {
		return nil, resp, err
	}

	p := new(PostProcess)
	if err := json.Unmarshal(postResponse.Data, &p); err != nil {
		return nil, nil, errors.New("unmarshal failed")
	}

	return p, resp, nil
}

func (c *Client) UpdatePostProcess(postProcess *PostProcess) (*PostProcess, *http.Response, error) {
	requestUrl := fmt.Sprintf("post_processes/%s", postProcess.Id)
	patchResponse := new(PostProcessPatchResponse)
	resp, err := c.requestWithRetry("PATCH", requestUrl, postProcess, patchResponse, defaultRetryCount)
	if err != nil {
		return nil, resp, err
	}

	p := new(PostProcess)
	if err := json.Unmarshal(patchResponse.Data, &p); err != nil {
		return nil, resp, errors.New("unmarshal failed")
	}

	return p, resp, nil
}

func (c *Client) DeletePostProcess(postProcessId string) (*http.Response, error) {
	requestUrl := fmt.Sprintf("post_processes/%s", postProcessId)
	resp, err := c.requestWithRetry("DELETE", requestUrl, nil, nil, defaultRetryCount)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func readParameters(rawPostProcess *PostProcessAttributes) map[string]interface{} {
	parameters := rawPostProcess.Parameters

	switch rawPostProcess.Service {
	case "email":
		parameters["email_recipient"] = parameters["recipients"].([]interface{})[0]
		delete(parameters, "recipients")
	case "sqs":
		parameters["sqs_aws_account_id"] = parameters["aws_account_id"]
		parameters["sqs_queue"] = parameters["queue"]
		parameters["sqs_region"] = parameters["region"]

		delete(parameters, "aws_account_id")
		delete(parameters, "queue")
		delete(parameters, "region")
	case "slack":
		parameters["slack_channel_name"] = parameters["channel_name"]
		parameters["slack_language"] = parameters["language"]
		parameters["slack_time_zone"] = parameters["time_zone"]

		delete(parameters, "channel_name")
		delete(parameters, "channel_id")
		delete(parameters, "language")
		delete(parameters, "time_zone")
	case "webhook":
		parameters["webhook_authorization_header"] = parameters["authorization_header"]
		parameters["webhook_url"] = parameters["url"]

		delete(parameters, "authorization_header")
		delete(parameters, "url")
	}

	return parameters
}

func (p *PostProcess) UnmarshalJSON(data []byte) error {
	rp := RawPostProcessData{}
	if err := json.Unmarshal(data, &rp); err != nil {
		return errors.New("unmarshal failed")
	}

	p.Id = rp.Id
	p.Name = rp.Attributes.Name
	p.Service = rp.Attributes.Service
	p.GroupId = rp.Attributes.GroupId
	p.SharedByGroup = rp.Attributes.SharedByGroup
	p.Parameters = readParameters(&rp.Attributes)

	return nil
}
