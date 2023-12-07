// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// GetWebsiteListReader is a Reader for the GetWebsiteList structure.
type GetWebsiteListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebsiteListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetWebsiteListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWebsiteListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteListOK creates a GetWebsiteListOK with default headers values
func NewGetWebsiteListOK() *GetWebsiteListOK {
	return &GetWebsiteListOK{}
}

/* GetWebsiteListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebsiteListOK struct {
	Payload *models.WebsitePaginationResponse
}

func (o *GetWebsiteListOK) Error() string {
	return fmt.Sprintf("[GET /website/websites][%d] getWebsiteListOK  %+v", 200, o.Payload)
}
func (o *GetWebsiteListOK) GetPayload() *models.WebsitePaginationResponse {
	return o.Payload
}

func (o *GetWebsiteListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebsitePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsiteListTooManyRequests creates a GetWebsiteListTooManyRequests with default headers values
func NewGetWebsiteListTooManyRequests() *GetWebsiteListTooManyRequests {
	return &GetWebsiteListTooManyRequests{}
}

/* GetWebsiteListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetWebsiteListTooManyRequests struct {

	/* Request limit per X-Rate-Limit-Window
	 */
	XRateLimitLimit int64

	/* The number of requests left for the time window
	 */
	XRateLimitRemaining int64

	/* The rolling time window length with the unit of second
	 */
	XRateLimitWindow int64
}

func (o *GetWebsiteListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /website/websites][%d] getWebsiteListTooManyRequests ", 429)
}

func (o *GetWebsiteListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-rate-limit-limit
	hdrXRateLimitLimit := response.GetHeader("x-rate-limit-limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("x-rate-limit-limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header x-rate-limit-remaining
	hdrXRateLimitRemaining := response.GetHeader("x-rate-limit-remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("x-rate-limit-remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header x-rate-limit-window
	hdrXRateLimitWindow := response.GetHeader("x-rate-limit-window")

	if hdrXRateLimitWindow != "" {
		valxRateLimitWindow, err := swag.ConvertInt64(hdrXRateLimitWindow)
		if err != nil {
			return errors.InvalidType("x-rate-limit-window", "header", "int64", hdrXRateLimitWindow)
		}
		o.XRateLimitWindow = valxRateLimitWindow
	}

	return nil
}

// NewGetWebsiteListDefault creates a GetWebsiteListDefault with default headers values
func NewGetWebsiteListDefault(code int) *GetWebsiteListDefault {
	return &GetWebsiteListDefault{
		_statusCode: code,
	}
}

/* GetWebsiteListDefault describes a response with status code -1, with default header values.

Error
*/
type GetWebsiteListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get website list default response
func (o *GetWebsiteListDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsiteListDefault) Error() string {
	return fmt.Sprintf("[GET /website/websites][%d] getWebsiteList default  %+v", o._statusCode, o.Payload)
}
func (o *GetWebsiteListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWebsiteListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
