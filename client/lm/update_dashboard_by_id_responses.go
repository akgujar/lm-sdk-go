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

// UpdateDashboardByIDReader is a Reader for the UpdateDashboardByID structure.
type UpdateDashboardByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDashboardByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDashboardByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewUpdateDashboardByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateDashboardByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDashboardByIDOK creates a UpdateDashboardByIDOK with default headers values
func NewUpdateDashboardByIDOK() *UpdateDashboardByIDOK {
	return &UpdateDashboardByIDOK{}
}

/* UpdateDashboardByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateDashboardByIDOK struct {
	Payload *models.Dashboard
}

func (o *UpdateDashboardByIDOK) Error() string {
	return fmt.Sprintf("[PUT /dashboard/dashboards/{id}][%d] updateDashboardByIdOK  %+v", 200, o.Payload)
}
func (o *UpdateDashboardByIDOK) GetPayload() *models.Dashboard {
	return o.Payload
}

func (o *UpdateDashboardByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardByIDTooManyRequests creates a UpdateDashboardByIDTooManyRequests with default headers values
func NewUpdateDashboardByIDTooManyRequests() *UpdateDashboardByIDTooManyRequests {
	return &UpdateDashboardByIDTooManyRequests{}
}

/* UpdateDashboardByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateDashboardByIDTooManyRequests struct {

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

func (o *UpdateDashboardByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /dashboard/dashboards/{id}][%d] updateDashboardByIdTooManyRequests ", 429)
}

func (o *UpdateDashboardByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDashboardByIDDefault creates a UpdateDashboardByIDDefault with default headers values
func NewUpdateDashboardByIDDefault(code int) *UpdateDashboardByIDDefault {
	return &UpdateDashboardByIDDefault{
		_statusCode: code,
	}
}

/* UpdateDashboardByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateDashboardByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update dashboard by Id default response
func (o *UpdateDashboardByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDashboardByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /dashboard/dashboards/{id}][%d] updateDashboardById default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateDashboardByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDashboardByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
