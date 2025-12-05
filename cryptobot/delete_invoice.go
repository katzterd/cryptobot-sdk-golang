package cryptobot

import (
	"fmt"
	"net/url"
)

type DeleteInvoiceRequest struct {
	// Invoice ID to be deleted.
	InvoiceId string `json:"invoice_id"`
}

type deleteInvoiceResponse struct {
	response
	Result bool `json:"result"`
}

// Use this method to delete invoices created by your app. Returns True on success.
func (c *Client) DeleteInvoice(deleteInvoiceRequest DeleteInvoiceRequest) (bool, error) {
	responseBodyReader, err := c.request("deleteInvoice", func(q url.Values) url.Values {
		q.Add("invoice_id", deleteInvoiceRequest.InvoiceId)
		return q
	})
	if err != nil {
		return false, err
	}
	defer responseBodyReader.Close()

	var response deleteInvoiceResponse
	if err := c.decodeResponse(responseBodyReader, &response); err != nil {
		return false, err
	}

	if response.Ok {
		return response.Result, nil
	} else {
		return false, fmt.Errorf("deleteInvoice request error: code - %v, name - %s", response.Error.Code, response.Error.Name)
	}
}
