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

// PatchAppliesToFunctionReader is a Reader for the PatchAppliesToFunction structure.
type PatchAppliesToFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAppliesToFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAppliesToFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPatchAppliesToFunctionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchAppliesToFunctionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchAppliesToFunctionOK creates a PatchAppliesToFunctionOK with default headers values
func NewPatchAppliesToFunctionOK() *PatchAppliesToFunctionOK {
	return &PatchAppliesToFunctionOK{}
}

/* PatchAppliesToFunctionOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchAppliesToFunctionOK struct {
	Payload *models.AppliesToFunction
}

func (o *PatchAppliesToFunctionOK) Error() string {
	return fmt.Sprintf("[PATCH /setting/functions/{id}][%d] patchAppliesToFunctionOK  %+v", 200, o.Payload)
}
func (o *PatchAppliesToFunctionOK) GetPayload() *models.AppliesToFunction {
	return o.Payload
}

func (o *PatchAppliesToFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppliesToFunction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAppliesToFunctionTooManyRequests creates a PatchAppliesToFunctionTooManyRequests with default headers values
func NewPatchAppliesToFunctionTooManyRequests() *PatchAppliesToFunctionTooManyRequests {
	return &PatchAppliesToFunctionTooManyRequests{}
}

/* PatchAppliesToFunctionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PatchAppliesToFunctionTooManyRequests struct {

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

func (o *PatchAppliesToFunctionTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /setting/functions/{id}][%d] patchAppliesToFunctionTooManyRequests ", 429)
}

func (o *PatchAppliesToFunctionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchAppliesToFunctionDefault creates a PatchAppliesToFunctionDefault with default headers values
func NewPatchAppliesToFunctionDefault(code int) *PatchAppliesToFunctionDefault {
	return &PatchAppliesToFunctionDefault{
		_statusCode: code,
	}
}

/* PatchAppliesToFunctionDefault describes a response with status code -1, with default header values.

Error
*/
type PatchAppliesToFunctionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch applies to function default response
func (o *PatchAppliesToFunctionDefault) Code() int {
	return o._statusCode
}

func (o *PatchAppliesToFunctionDefault) Error() string {
	return fmt.Sprintf("[PATCH /setting/functions/{id}][%d] patchAppliesToFunction default  %+v", o._statusCode, o.Payload)
}
func (o *PatchAppliesToFunctionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchAppliesToFunctionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
