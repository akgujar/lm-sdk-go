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

// DeleteAdminByIDReader is a Reader for the DeleteAdminByID structure.
type DeleteAdminByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAdminByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAdminByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewDeleteAdminByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAdminByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAdminByIDOK creates a DeleteAdminByIDOK with default headers values
func NewDeleteAdminByIDOK() *DeleteAdminByIDOK {
	return &DeleteAdminByIDOK{}
}

/* DeleteAdminByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteAdminByIDOK struct {
	Payload interface{}
}

func (o *DeleteAdminByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /setting/admins/{id}][%d] deleteAdminByIdOK  %+v", 200, o.Payload)
}
func (o *DeleteAdminByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAdminByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAdminByIDTooManyRequests creates a DeleteAdminByIDTooManyRequests with default headers values
func NewDeleteAdminByIDTooManyRequests() *DeleteAdminByIDTooManyRequests {
	return &DeleteAdminByIDTooManyRequests{}
}

/* DeleteAdminByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteAdminByIDTooManyRequests struct {

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

func (o *DeleteAdminByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /setting/admins/{id}][%d] deleteAdminByIdTooManyRequests ", 429)
}

func (o *DeleteAdminByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAdminByIDDefault creates a DeleteAdminByIDDefault with default headers values
func NewDeleteAdminByIDDefault(code int) *DeleteAdminByIDDefault {
	return &DeleteAdminByIDDefault{
		_statusCode: code,
	}
}

/* DeleteAdminByIDDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteAdminByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete admin by Id default response
func (o *DeleteAdminByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAdminByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /setting/admins/{id}][%d] deleteAdminById default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteAdminByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteAdminByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
