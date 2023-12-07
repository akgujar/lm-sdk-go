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

// PatchReportByIDReader is a Reader for the PatchReportByID structure.
type PatchReportByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchReportByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchReportByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPatchReportByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchReportByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchReportByIDOK creates a PatchReportByIDOK with default headers values
func NewPatchReportByIDOK() *PatchReportByIDOK {
	return &PatchReportByIDOK{}
}

/* PatchReportByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchReportByIDOK struct {
	Payload models.ReportBase
}

func (o *PatchReportByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /report/reports/{id}][%d] patchReportByIdOK  %+v", 200, o.Payload)
}
func (o *PatchReportByIDOK) GetPayload() models.ReportBase {
	return o.Payload
}

func (o *PatchReportByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalReportBase(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchReportByIDTooManyRequests creates a PatchReportByIDTooManyRequests with default headers values
func NewPatchReportByIDTooManyRequests() *PatchReportByIDTooManyRequests {
	return &PatchReportByIDTooManyRequests{}
}

/* PatchReportByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PatchReportByIDTooManyRequests struct {

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

func (o *PatchReportByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /report/reports/{id}][%d] patchReportByIdTooManyRequests ", 429)
}

func (o *PatchReportByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchReportByIDDefault creates a PatchReportByIDDefault with default headers values
func NewPatchReportByIDDefault(code int) *PatchReportByIDDefault {
	return &PatchReportByIDDefault{
		_statusCode: code,
	}
}

/* PatchReportByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchReportByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch report by Id default response
func (o *PatchReportByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchReportByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /report/reports/{id}][%d] patchReportById default  %+v", o._statusCode, o.Payload)
}
func (o *PatchReportByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchReportByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
