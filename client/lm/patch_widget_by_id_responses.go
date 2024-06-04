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

// PatchWidgetByIDReader is a Reader for the PatchWidgetByID structure.
type PatchWidgetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWidgetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchWidgetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPatchWidgetByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchWidgetByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchWidgetByIDOK creates a PatchWidgetByIDOK with default headers values
func NewPatchWidgetByIDOK() *PatchWidgetByIDOK {
	return &PatchWidgetByIDOK{}
}

/*
PatchWidgetByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchWidgetByIDOK struct {
	Payload models.Widget
}

// IsSuccess returns true when this patch widget by Id o k response has a 2xx status code
func (o *PatchWidgetByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch widget by Id o k response has a 3xx status code
func (o *PatchWidgetByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch widget by Id o k response has a 4xx status code
func (o *PatchWidgetByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch widget by Id o k response has a 5xx status code
func (o *PatchWidgetByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch widget by Id o k response a status code equal to that given
func (o *PatchWidgetByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch widget by Id o k response
func (o *PatchWidgetByIDOK) Code() int {
	return 200
}

func (o *PatchWidgetByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboard/widgets/{id}][%d] patchWidgetByIdOK %s", 200, payload)
}

func (o *PatchWidgetByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboard/widgets/{id}][%d] patchWidgetByIdOK %s", 200, payload)
}

func (o *PatchWidgetByIDOK) GetPayload() models.Widget {
	return o.Payload
}

func (o *PatchWidgetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalWidget(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchWidgetByIDTooManyRequests creates a PatchWidgetByIDTooManyRequests with default headers values
func NewPatchWidgetByIDTooManyRequests() *PatchWidgetByIDTooManyRequests {
	return &PatchWidgetByIDTooManyRequests{}
}

/*
PatchWidgetByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PatchWidgetByIDTooManyRequests struct {

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

// IsSuccess returns true when this patch widget by Id too many requests response has a 2xx status code
func (o *PatchWidgetByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch widget by Id too many requests response has a 3xx status code
func (o *PatchWidgetByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch widget by Id too many requests response has a 4xx status code
func (o *PatchWidgetByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch widget by Id too many requests response has a 5xx status code
func (o *PatchWidgetByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch widget by Id too many requests response a status code equal to that given
func (o *PatchWidgetByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the patch widget by Id too many requests response
func (o *PatchWidgetByIDTooManyRequests) Code() int {
	return 429
}

func (o *PatchWidgetByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /dashboard/widgets/{id}][%d] patchWidgetByIdTooManyRequests", 429)
}

func (o *PatchWidgetByIDTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /dashboard/widgets/{id}][%d] patchWidgetByIdTooManyRequests", 429)
}

func (o *PatchWidgetByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchWidgetByIDDefault creates a PatchWidgetByIDDefault with default headers values
func NewPatchWidgetByIDDefault(code int) *PatchWidgetByIDDefault {
	return &PatchWidgetByIDDefault{
		_statusCode: code,
	}
}

/*
PatchWidgetByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchWidgetByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this patch widget by Id default response has a 2xx status code
func (o *PatchWidgetByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch widget by Id default response has a 3xx status code
func (o *PatchWidgetByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch widget by Id default response has a 4xx status code
func (o *PatchWidgetByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch widget by Id default response has a 5xx status code
func (o *PatchWidgetByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch widget by Id default response a status code equal to that given
func (o *PatchWidgetByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch widget by Id default response
func (o *PatchWidgetByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchWidgetByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboard/widgets/{id}][%d] patchWidgetById default %s", o._statusCode, payload)
}

func (o *PatchWidgetByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboard/widgets/{id}][%d] patchWidgetById default %s", o._statusCode, payload)
}

func (o *PatchWidgetByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchWidgetByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
