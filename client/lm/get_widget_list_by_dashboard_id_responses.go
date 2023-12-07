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

// GetWidgetListByDashboardIDReader is a Reader for the GetWidgetListByDashboardID structure.
type GetWidgetListByDashboardIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWidgetListByDashboardIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWidgetListByDashboardIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetWidgetListByDashboardIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWidgetListByDashboardIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWidgetListByDashboardIDOK creates a GetWidgetListByDashboardIDOK with default headers values
func NewGetWidgetListByDashboardIDOK() *GetWidgetListByDashboardIDOK {
	return &GetWidgetListByDashboardIDOK{}
}

/* GetWidgetListByDashboardIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWidgetListByDashboardIDOK struct {
	Payload *models.WidgetPaginationResponse
}

func (o *GetWidgetListByDashboardIDOK) Error() string {
	return fmt.Sprintf("[GET /dashboard/dashboards/{id}/widgets][%d] getWidgetListByDashboardIdOK  %+v", 200, o.Payload)
}
func (o *GetWidgetListByDashboardIDOK) GetPayload() *models.WidgetPaginationResponse {
	return o.Payload
}

func (o *GetWidgetListByDashboardIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WidgetPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetListByDashboardIDTooManyRequests creates a GetWidgetListByDashboardIDTooManyRequests with default headers values
func NewGetWidgetListByDashboardIDTooManyRequests() *GetWidgetListByDashboardIDTooManyRequests {
	return &GetWidgetListByDashboardIDTooManyRequests{}
}

/* GetWidgetListByDashboardIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetWidgetListByDashboardIDTooManyRequests struct {

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

func (o *GetWidgetListByDashboardIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /dashboard/dashboards/{id}/widgets][%d] getWidgetListByDashboardIdTooManyRequests ", 429)
}

func (o *GetWidgetListByDashboardIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWidgetListByDashboardIDDefault creates a GetWidgetListByDashboardIDDefault with default headers values
func NewGetWidgetListByDashboardIDDefault(code int) *GetWidgetListByDashboardIDDefault {
	return &GetWidgetListByDashboardIDDefault{
		_statusCode: code,
	}
}

/* GetWidgetListByDashboardIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetWidgetListByDashboardIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get widget list by dashboard Id default response
func (o *GetWidgetListByDashboardIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWidgetListByDashboardIDDefault) Error() string {
	return fmt.Sprintf("[GET /dashboard/dashboards/{id}/widgets][%d] getWidgetListByDashboardId default  %+v", o._statusCode, o.Payload)
}
func (o *GetWidgetListByDashboardIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWidgetListByDashboardIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
