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

// DeleteReportGroupByIDReader is a Reader for the DeleteReportGroupByID structure.
type DeleteReportGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReportGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteReportGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewDeleteReportGroupByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteReportGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteReportGroupByIDOK creates a DeleteReportGroupByIDOK with default headers values
func NewDeleteReportGroupByIDOK() *DeleteReportGroupByIDOK {
	return &DeleteReportGroupByIDOK{}
}

/*
DeleteReportGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteReportGroupByIDOK struct {
	Payload interface{}
}

// IsSuccess returns true when this delete report group by Id o k response has a 2xx status code
func (o *DeleteReportGroupByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete report group by Id o k response has a 3xx status code
func (o *DeleteReportGroupByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete report group by Id o k response has a 4xx status code
func (o *DeleteReportGroupByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete report group by Id o k response has a 5xx status code
func (o *DeleteReportGroupByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete report group by Id o k response a status code equal to that given
func (o *DeleteReportGroupByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete report group by Id o k response
func (o *DeleteReportGroupByIDOK) Code() int {
	return 200
}

func (o *DeleteReportGroupByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /report/groups/{id}][%d] deleteReportGroupByIdOK %s", 200, payload)
}

func (o *DeleteReportGroupByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /report/groups/{id}][%d] deleteReportGroupByIdOK %s", 200, payload)
}

func (o *DeleteReportGroupByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteReportGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReportGroupByIDTooManyRequests creates a DeleteReportGroupByIDTooManyRequests with default headers values
func NewDeleteReportGroupByIDTooManyRequests() *DeleteReportGroupByIDTooManyRequests {
	return &DeleteReportGroupByIDTooManyRequests{}
}

/*
DeleteReportGroupByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteReportGroupByIDTooManyRequests struct {

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

// IsSuccess returns true when this delete report group by Id too many requests response has a 2xx status code
func (o *DeleteReportGroupByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete report group by Id too many requests response has a 3xx status code
func (o *DeleteReportGroupByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete report group by Id too many requests response has a 4xx status code
func (o *DeleteReportGroupByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete report group by Id too many requests response has a 5xx status code
func (o *DeleteReportGroupByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete report group by Id too many requests response a status code equal to that given
func (o *DeleteReportGroupByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete report group by Id too many requests response
func (o *DeleteReportGroupByIDTooManyRequests) Code() int {
	return 429
}

func (o *DeleteReportGroupByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /report/groups/{id}][%d] deleteReportGroupByIdTooManyRequests", 429)
}

func (o *DeleteReportGroupByIDTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /report/groups/{id}][%d] deleteReportGroupByIdTooManyRequests", 429)
}

func (o *DeleteReportGroupByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteReportGroupByIDDefault creates a DeleteReportGroupByIDDefault with default headers values
func NewDeleteReportGroupByIDDefault(code int) *DeleteReportGroupByIDDefault {
	return &DeleteReportGroupByIDDefault{
		_statusCode: code,
	}
}

/*
DeleteReportGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteReportGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete report group by Id default response has a 2xx status code
func (o *DeleteReportGroupByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete report group by Id default response has a 3xx status code
func (o *DeleteReportGroupByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete report group by Id default response has a 4xx status code
func (o *DeleteReportGroupByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete report group by Id default response has a 5xx status code
func (o *DeleteReportGroupByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete report group by Id default response a status code equal to that given
func (o *DeleteReportGroupByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete report group by Id default response
func (o *DeleteReportGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteReportGroupByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /report/groups/{id}][%d] deleteReportGroupById default %s", o._statusCode, payload)
}

func (o *DeleteReportGroupByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /report/groups/{id}][%d] deleteReportGroupById default %s", o._statusCode, payload)
}

func (o *DeleteReportGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteReportGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
