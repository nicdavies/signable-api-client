package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nicdavies/signable-api-client/types"
	"net/http"
)

func (c *Client) GetContacts(ctx context.Context, options *types.ListContactsOptions) (*types.ListContactsResponse, error) {
	offset := 20
	limit := 20

	if options != nil {
		offset = options.Offset
		limit = options.Limit
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/contacts?offset=%d&limit=%d", c.BaseURL, offset, limit), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := types.ListContactsResponse{}
	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetContact(ctx context.Context, options *types.ReadContactOptions) (*types.ReadContactResponse, error) {
	if options == nil {
		return nil, errors.New("ID is required")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/contacts/%d", c.BaseURL, options.Id), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := types.ReadContactResponse{}
	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateContact(ctx context.Context, options *types.CreateContactsOptions) (*types.CreateContactsResponse, error) {
	b, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/contacts", c.BaseURL), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := types.CreateContactsResponse{}
	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateContact(ctx context.Context, id int, options *types.UpdateContactOptions) (*types.UpdateContactResponse, error) {
	b, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/contacts/%d", c.BaseURL, id), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := types.UpdateContactResponse{}
	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteContact(ctx context.Context, options *types.DeleteContactOptions) (*types.DeleteContactResponse, error) {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/contacts/%d", c.BaseURL, options.Id), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := types.DeleteContactResponse{}
	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return &res, nil
}
