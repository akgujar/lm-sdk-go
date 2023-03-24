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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// GetMetricsUsageReader is a Reader for the GetMetricsUsage structure.
type GetMetricsUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricsUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetricsUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetMetricsUsageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMetricsUsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMetricsUsageOK creates a GetMetricsUsageOK with default headers values
func NewGetMetricsUsageOK() *GetMetricsUsageOK {
	return &GetMetricsUsageOK{}
}

/* GetMetricsUsageOK describes a response with status code 200, with default header values.

successful operation
*/
type GetMetricsUsageOK struct {
	Payload *models.Usage
}

func (o *GetMetricsUsageOK) Error() string {
	return fmt.Sprintf("[GET /metrics/usage][%d] getMetricsUsageOK  %+v", 200, o.Payload)
}
func (o *GetMetricsUsageOK) GetPayload() *models.Usage {
	return o.Payload
}

func (o *GetMetricsUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Usage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsUsageTooManyRequests creates a GetMetricsUsageTooManyRequests with default headers values
func NewGetMetricsUsageTooManyRequests() *GetMetricsUsageTooManyRequests {
	return &GetMetricsUsageTooManyRequests{}
}

/* GetMetricsUsageTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetMetricsUsageTooManyRequests struct {

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

func (o *GetMetricsUsageTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /metrics/usage][%d] getMetricsUsageTooManyRequests ", 429)
}

func (o *GetMetricsUsageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMetricsUsageDefault creates a GetMetricsUsageDefault with default headers values
func NewGetMetricsUsageDefault(code int) *GetMetricsUsageDefault {
	return &GetMetricsUsageDefault{
		_statusCode: code,
	}
}

/* GetMetricsUsageDefault describes a response with status code -1, with default header values.

Error
*/
type GetMetricsUsageDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get metrics usage default response
func (o *GetMetricsUsageDefault) Code() int {
	return o._statusCode
}

func (o *GetMetricsUsageDefault) Error() string {
	return fmt.Sprintf("[GET /metrics/usage][%d] getMetricsUsage default  %+v", o._statusCode, o.Payload)
}
func (o *GetMetricsUsageDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetMetricsUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}