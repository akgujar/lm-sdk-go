// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// DeleteDashboardByIDReader is a Reader for the DeleteDashboardByID structure.
type DeleteDashboardByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDashboardByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDashboardByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewDeleteDashboardByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteDashboardByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDashboardByIDOK creates a DeleteDashboardByIDOK with default headers values
func NewDeleteDashboardByIDOK() *DeleteDashboardByIDOK {
	return &DeleteDashboardByIDOK{}
}

/*
DeleteDashboardByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteDashboardByIDOK struct {
	Payload interface{}
}

// IsSuccess returns true when this delete dashboard by Id o k response has a 2xx status code
func (o *DeleteDashboardByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete dashboard by Id o k response has a 3xx status code
func (o *DeleteDashboardByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard by Id o k response has a 4xx status code
func (o *DeleteDashboardByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete dashboard by Id o k response has a 5xx status code
func (o *DeleteDashboardByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete dashboard by Id o k response a status code equal to that given
func (o *DeleteDashboardByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete dashboard by Id o k response
func (o *DeleteDashboardByIDOK) Code() int {
	return 200
}

func (o *DeleteDashboardByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /dashboard/dashboards/{id}][%d] deleteDashboardByIdOK %s", 200, payload)
}

func (o *DeleteDashboardByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /dashboard/dashboards/{id}][%d] deleteDashboardByIdOK %s", 200, payload)
}

func (o *DeleteDashboardByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteDashboardByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardByIDTooManyRequests creates a DeleteDashboardByIDTooManyRequests with default headers values
func NewDeleteDashboardByIDTooManyRequests() *DeleteDashboardByIDTooManyRequests {
	return &DeleteDashboardByIDTooManyRequests{}
}

/*
DeleteDashboardByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteDashboardByIDTooManyRequests struct {

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

// IsSuccess returns true when this delete dashboard by Id too many requests response has a 2xx status code
func (o *DeleteDashboardByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete dashboard by Id too many requests response has a 3xx status code
func (o *DeleteDashboardByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard by Id too many requests response has a 4xx status code
func (o *DeleteDashboardByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete dashboard by Id too many requests response has a 5xx status code
func (o *DeleteDashboardByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete dashboard by Id too many requests response a status code equal to that given
func (o *DeleteDashboardByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete dashboard by Id too many requests response
func (o *DeleteDashboardByIDTooManyRequests) Code() int {
	return 429
}

func (o *DeleteDashboardByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /dashboard/dashboards/{id}][%d] deleteDashboardByIdTooManyRequests", 429)
}

func (o *DeleteDashboardByIDTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /dashboard/dashboards/{id}][%d] deleteDashboardByIdTooManyRequests", 429)
}

func (o *DeleteDashboardByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteDashboardByIDDefault creates a DeleteDashboardByIDDefault with default headers values
func NewDeleteDashboardByIDDefault(code int) *DeleteDashboardByIDDefault {
	return &DeleteDashboardByIDDefault{
		_statusCode: code,
	}
}

/*
DeleteDashboardByIDDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteDashboardByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete dashboard by Id default response has a 2xx status code
func (o *DeleteDashboardByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete dashboard by Id default response has a 3xx status code
func (o *DeleteDashboardByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete dashboard by Id default response has a 4xx status code
func (o *DeleteDashboardByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete dashboard by Id default response has a 5xx status code
func (o *DeleteDashboardByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete dashboard by Id default response a status code equal to that given
func (o *DeleteDashboardByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete dashboard by Id default response
func (o *DeleteDashboardByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDashboardByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /dashboard/dashboards/{id}][%d] deleteDashboardById default %s", o._statusCode, payload)
}

func (o *DeleteDashboardByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /dashboard/dashboards/{id}][%d] deleteDashboardById default %s", o._statusCode, payload)
}

func (o *DeleteDashboardByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteDashboardByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
