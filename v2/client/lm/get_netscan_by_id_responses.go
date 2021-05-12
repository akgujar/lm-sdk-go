// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/v2/models"
)

// GetNetscanByIDReader is a Reader for the GetNetscanByID structure.
type GetNetscanByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetscanByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetscanByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNetscanByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetscanByIDOK creates a GetNetscanByIDOK with default headers values
func NewGetNetscanByIDOK() *GetNetscanByIDOK {
	return &GetNetscanByIDOK{}
}

/*GetNetscanByIDOK handles this case with default header values.

successful operation
*/
type GetNetscanByIDOK struct {
	Payload models.Netscan
}

func (o *GetNetscanByIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/netscans/{id}][%d] getNetscanByIdOK  %+v", 200, o.Payload)
}

func (o *GetNetscanByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalNetscan(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetNetscanByIDDefault creates a GetNetscanByIDDefault with default headers values
func NewGetNetscanByIDDefault(code int) *GetNetscanByIDDefault {
	return &GetNetscanByIDDefault{
		_statusCode: code,
	}
}

/*GetNetscanByIDDefault handles this case with default header values.

Error
*/
type GetNetscanByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get netscan by Id default response
func (o *GetNetscanByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetNetscanByIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/netscans/{id}][%d] getNetscanById default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetscanByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}