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

// GetCollectorInstallerJSONReader is a Reader for the GetCollectorInstallerJSON structure.
type GetCollectorInstallerJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCollectorInstallerJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCollectorInstallerJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetCollectorInstallerJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCollectorInstallerJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCollectorInstallerJSONOK creates a GetCollectorInstallerJSONOK with default headers values
func NewGetCollectorInstallerJSONOK() *GetCollectorInstallerJSONOK {
	return &GetCollectorInstallerJSONOK{}
}

/*
GetCollectorInstallerJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCollectorInstallerJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get collector installer Json o k response has a 2xx status code
func (o *GetCollectorInstallerJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get collector installer Json o k response has a 3xx status code
func (o *GetCollectorInstallerJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get collector installer Json o k response has a 4xx status code
func (o *GetCollectorInstallerJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get collector installer Json o k response has a 5xx status code
func (o *GetCollectorInstallerJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get collector installer Json o k response a status code equal to that given
func (o *GetCollectorInstallerJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get collector installer Json o k response
func (o *GetCollectorInstallerJSONOK) Code() int {
	return 200
}

func (o *GetCollectorInstallerJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/collector/collectors/{collectorId}/installers/{osAndArch}?__json=][%d] getCollectorInstallerJsonOK %s", 200, payload)
}

func (o *GetCollectorInstallerJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/collector/collectors/{collectorId}/installers/{osAndArch}?__json=][%d] getCollectorInstallerJsonOK %s", 200, payload)
}

func (o *GetCollectorInstallerJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetCollectorInstallerJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCollectorInstallerJSONTooManyRequests creates a GetCollectorInstallerJSONTooManyRequests with default headers values
func NewGetCollectorInstallerJSONTooManyRequests() *GetCollectorInstallerJSONTooManyRequests {
	return &GetCollectorInstallerJSONTooManyRequests{}
}

/*
GetCollectorInstallerJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCollectorInstallerJSONTooManyRequests struct {

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

// IsSuccess returns true when this get collector installer Json too many requests response has a 2xx status code
func (o *GetCollectorInstallerJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get collector installer Json too many requests response has a 3xx status code
func (o *GetCollectorInstallerJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get collector installer Json too many requests response has a 4xx status code
func (o *GetCollectorInstallerJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get collector installer Json too many requests response has a 5xx status code
func (o *GetCollectorInstallerJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get collector installer Json too many requests response a status code equal to that given
func (o *GetCollectorInstallerJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get collector installer Json too many requests response
func (o *GetCollectorInstallerJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetCollectorInstallerJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/collector/collectors/{collectorId}/installers/{osAndArch}?__json=][%d] getCollectorInstallerJsonTooManyRequests", 429)
}

func (o *GetCollectorInstallerJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /setting/collector/collectors/{collectorId}/installers/{osAndArch}?__json=][%d] getCollectorInstallerJsonTooManyRequests", 429)
}

func (o *GetCollectorInstallerJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCollectorInstallerJSONDefault creates a GetCollectorInstallerJSONDefault with default headers values
func NewGetCollectorInstallerJSONDefault(code int) *GetCollectorInstallerJSONDefault {
	return &GetCollectorInstallerJSONDefault{
		_statusCode: code,
	}
}

/*
GetCollectorInstallerJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetCollectorInstallerJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get collector installer Json default response has a 2xx status code
func (o *GetCollectorInstallerJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get collector installer Json default response has a 3xx status code
func (o *GetCollectorInstallerJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get collector installer Json default response has a 4xx status code
func (o *GetCollectorInstallerJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get collector installer Json default response has a 5xx status code
func (o *GetCollectorInstallerJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get collector installer Json default response a status code equal to that given
func (o *GetCollectorInstallerJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get collector installer Json default response
func (o *GetCollectorInstallerJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetCollectorInstallerJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/collector/collectors/{collectorId}/installers/{osAndArch}?__json=][%d] getCollectorInstallerJson default %s", o._statusCode, payload)
}

func (o *GetCollectorInstallerJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/collector/collectors/{collectorId}/installers/{osAndArch}?__json=][%d] getCollectorInstallerJson default %s", o._statusCode, payload)
}

func (o *GetCollectorInstallerJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCollectorInstallerJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
