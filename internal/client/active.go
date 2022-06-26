package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Active struct {
	Id     string `json:"id,omitempty"`
	Active *bool  `json:"active,omitempty"`
}

type ActivePostResponse struct {
	Data json.RawMessage `json:"data"`
}

type ActiveDeleteResponse struct {
	Data json.RawMessage `json:"data"`
}

type RawActiveData struct {
	Id         string           `json:"id"`
	Type       string           `json:"type"`
	Attributes ActiveAttributes `json:"attributes"`
}

type ActiveAttributes struct {
	Active *bool `json:"active"`
}

func (c *Client) ActivateJob(jobId string) (*Active, *http.Response, error) {
	requestUrl := fmt.Sprintf("/jobs/%s/active", jobId)
	req, err := c.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	postResponse := new(ActivePostResponse)
	resp, err := c.Do(req, &postResponse)
	if err != nil {
		return nil, resp, err
	}

	a := new(Active)
	if err := json.Unmarshal(postResponse.Data, &a); err != nil {
		return nil, nil, errors.New("unmarshal failed")
	}

	return a, resp, nil
}

func (c *Client) DeactivateJob(jobId string) (*Active, *http.Response, error) {
	requestUrl := fmt.Sprintf("/jobs/%s/active", jobId)
	req, err := c.NewRequest("DELETE", requestUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	deleteResponse := new(ActiveDeleteResponse)
	resp, err := c.Do(req, &deleteResponse)
	if err != nil {
		return nil, resp, err
	}

	a := new(Active)
	if err := json.Unmarshal(deleteResponse.Data, &a); err != nil {
		return nil, nil, errors.New("unmarshal failed")
	}

	return a, resp, nil
}

func (a *Active) UnmarshalJSON(data []byte) error {
	ra := RawActiveData{}
	if err := json.Unmarshal(data, &ra); err != nil {
		return errors.New("unmarshal failed")
	}

	a.Id = ra.Id
	a.Active = ra.Attributes.Active

	return nil
}
