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

// GetSiteMonitorCheckPointListJSONReader is a Reader for the GetSiteMonitorCheckPointListJSON structure.
type GetSiteMonitorCheckPointListJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSiteMonitorCheckPointListJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSiteMonitorCheckPointListJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetSiteMonitorCheckPointListJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSiteMonitorCheckPointListJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSiteMonitorCheckPointListJSONOK creates a GetSiteMonitorCheckPointListJSONOK with default headers values
func NewGetSiteMonitorCheckPointListJSONOK() *GetSiteMonitorCheckPointListJSONOK {
	return &GetSiteMonitorCheckPointListJSONOK{}
}

/*
GetSiteMonitorCheckPointListJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSiteMonitorCheckPointListJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get site monitor check point list Json o k response has a 2xx status code
func (o *GetSiteMonitorCheckPointListJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get site monitor check point list Json o k response has a 3xx status code
func (o *GetSiteMonitorCheckPointListJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get site monitor check point list Json o k response has a 4xx status code
func (o *GetSiteMonitorCheckPointListJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get site monitor check point list Json o k response has a 5xx status code
func (o *GetSiteMonitorCheckPointListJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get site monitor check point list Json o k response a status code equal to that given
func (o *GetSiteMonitorCheckPointListJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get site monitor check point list Json o k response
func (o *GetSiteMonitorCheckPointListJSONOK) Code() int {
	return 200
}

func (o *GetSiteMonitorCheckPointListJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/smcheckpoints?__json=][%d] getSiteMonitorCheckPointListJsonOK %s", 200, payload)
}

func (o *GetSiteMonitorCheckPointListJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/smcheckpoints?__json=][%d] getSiteMonitorCheckPointListJsonOK %s", 200, payload)
}

func (o *GetSiteMonitorCheckPointListJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetSiteMonitorCheckPointListJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteMonitorCheckPointListJSONTooManyRequests creates a GetSiteMonitorCheckPointListJSONTooManyRequests with default headers values
func NewGetSiteMonitorCheckPointListJSONTooManyRequests() *GetSiteMonitorCheckPointListJSONTooManyRequests {
	return &GetSiteMonitorCheckPointListJSONTooManyRequests{}
}

/*
GetSiteMonitorCheckPointListJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetSiteMonitorCheckPointListJSONTooManyRequests struct {

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

// IsSuccess returns true when this get site monitor check point list Json too many requests response has a 2xx status code
func (o *GetSiteMonitorCheckPointListJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get site monitor check point list Json too many requests response has a 3xx status code
func (o *GetSiteMonitorCheckPointListJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get site monitor check point list Json too many requests response has a 4xx status code
func (o *GetSiteMonitorCheckPointListJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get site monitor check point list Json too many requests response has a 5xx status code
func (o *GetSiteMonitorCheckPointListJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get site monitor check point list Json too many requests response a status code equal to that given
func (o *GetSiteMonitorCheckPointListJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get site monitor check point list Json too many requests response
func (o *GetSiteMonitorCheckPointListJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetSiteMonitorCheckPointListJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /website/smcheckpoints?__json=][%d] getSiteMonitorCheckPointListJsonTooManyRequests", 429)
}

func (o *GetSiteMonitorCheckPointListJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /website/smcheckpoints?__json=][%d] getSiteMonitorCheckPointListJsonTooManyRequests", 429)
}

func (o *GetSiteMonitorCheckPointListJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSiteMonitorCheckPointListJSONDefault creates a GetSiteMonitorCheckPointListJSONDefault with default headers values
func NewGetSiteMonitorCheckPointListJSONDefault(code int) *GetSiteMonitorCheckPointListJSONDefault {
	return &GetSiteMonitorCheckPointListJSONDefault{
		_statusCode: code,
	}
}

/*
GetSiteMonitorCheckPointListJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetSiteMonitorCheckPointListJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get site monitor check point list Json default response has a 2xx status code
func (o *GetSiteMonitorCheckPointListJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get site monitor check point list Json default response has a 3xx status code
func (o *GetSiteMonitorCheckPointListJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get site monitor check point list Json default response has a 4xx status code
func (o *GetSiteMonitorCheckPointListJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get site monitor check point list Json default response has a 5xx status code
func (o *GetSiteMonitorCheckPointListJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get site monitor check point list Json default response a status code equal to that given
func (o *GetSiteMonitorCheckPointListJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get site monitor check point list Json default response
func (o *GetSiteMonitorCheckPointListJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetSiteMonitorCheckPointListJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/smcheckpoints?__json=][%d] getSiteMonitorCheckPointListJson default %s", o._statusCode, payload)
}

func (o *GetSiteMonitorCheckPointListJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/smcheckpoints?__json=][%d] getSiteMonitorCheckPointListJson default %s", o._statusCode, payload)
}

func (o *GetSiteMonitorCheckPointListJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSiteMonitorCheckPointListJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}