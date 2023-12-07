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

// GetReportByIDReader is a Reader for the GetReportByID structure.
type GetReportByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetReportByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetReportByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReportByIDOK creates a GetReportByIDOK with default headers values
func NewGetReportByIDOK() *GetReportByIDOK {
	return &GetReportByIDOK{}
}

/* GetReportByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetReportByIDOK struct {
	Payload models.ReportBase
}

func (o *GetReportByIDOK) Error() string {
	return fmt.Sprintf("[GET /report/reports/{id}][%d] getReportByIdOK  %+v", 200, o.Payload)
}
func (o *GetReportByIDOK) GetPayload() models.ReportBase {
	return o.Payload
}

func (o *GetReportByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalReportBase(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetReportByIDTooManyRequests creates a GetReportByIDTooManyRequests with default headers values
func NewGetReportByIDTooManyRequests() *GetReportByIDTooManyRequests {
	return &GetReportByIDTooManyRequests{}
}

/* GetReportByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetReportByIDTooManyRequests struct {

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

func (o *GetReportByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /report/reports/{id}][%d] getReportByIdTooManyRequests ", 429)
}

func (o *GetReportByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportByIDDefault creates a GetReportByIDDefault with default headers values
func NewGetReportByIDDefault(code int) *GetReportByIDDefault {
	return &GetReportByIDDefault{
		_statusCode: code,
	}
}

/* GetReportByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetReportByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get report by Id default response
func (o *GetReportByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetReportByIDDefault) Error() string {
	return fmt.Sprintf("[GET /report/reports/{id}][%d] getReportById default  %+v", o._statusCode, o.Payload)
}
func (o *GetReportByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetReportByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
