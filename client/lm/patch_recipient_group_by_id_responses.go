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

// PatchRecipientGroupByIDReader is a Reader for the PatchRecipientGroupByID structure.
type PatchRecipientGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRecipientGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRecipientGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPatchRecipientGroupByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchRecipientGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchRecipientGroupByIDOK creates a PatchRecipientGroupByIDOK with default headers values
func NewPatchRecipientGroupByIDOK() *PatchRecipientGroupByIDOK {
	return &PatchRecipientGroupByIDOK{}
}

/* PatchRecipientGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchRecipientGroupByIDOK struct {
	Payload *models.RecipientGroup
}

func (o *PatchRecipientGroupByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /setting/recipientgroups/{id}][%d] patchRecipientGroupByIdOK  %+v", 200, o.Payload)
}
func (o *PatchRecipientGroupByIDOK) GetPayload() *models.RecipientGroup {
	return o.Payload
}

func (o *PatchRecipientGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecipientGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecipientGroupByIDTooManyRequests creates a PatchRecipientGroupByIDTooManyRequests with default headers values
func NewPatchRecipientGroupByIDTooManyRequests() *PatchRecipientGroupByIDTooManyRequests {
	return &PatchRecipientGroupByIDTooManyRequests{}
}

/* PatchRecipientGroupByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PatchRecipientGroupByIDTooManyRequests struct {

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

func (o *PatchRecipientGroupByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /setting/recipientgroups/{id}][%d] patchRecipientGroupByIdTooManyRequests ", 429)
}

func (o *PatchRecipientGroupByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchRecipientGroupByIDDefault creates a PatchRecipientGroupByIDDefault with default headers values
func NewPatchRecipientGroupByIDDefault(code int) *PatchRecipientGroupByIDDefault {
	return &PatchRecipientGroupByIDDefault{
		_statusCode: code,
	}
}

/* PatchRecipientGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchRecipientGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch recipient group by Id default response
func (o *PatchRecipientGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchRecipientGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /setting/recipientgroups/{id}][%d] patchRecipientGroupById default  %+v", o._statusCode, o.Payload)
}
func (o *PatchRecipientGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchRecipientGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
