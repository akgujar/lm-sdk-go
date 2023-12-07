// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/logicmonitor/lm-sdk-go/models"
)

// DeleteEscalationChainByIDReader is a Reader for the DeleteEscalationChainByID structure.
type DeleteEscalationChainByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEscalationChainByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEscalationChainByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteEscalationChainByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteEscalationChainByIDOK creates a DeleteEscalationChainByIDOK with default headers values
func NewDeleteEscalationChainByIDOK() *DeleteEscalationChainByIDOK {
	return &DeleteEscalationChainByIDOK{}
}

/* DeleteEscalationChainByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteEscalationChainByIDOK struct {
	Payload interface{}
}

func (o *DeleteEscalationChainByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /setting/alert/chains/{id}][%d] deleteEscalationChainByIdOK  %+v", 200, o.Payload)
}
func (o *DeleteEscalationChainByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteEscalationChainByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEscalationChainByIDDefault creates a DeleteEscalationChainByIDDefault with default headers values
func NewDeleteEscalationChainByIDDefault(code int) *DeleteEscalationChainByIDDefault {
	return &DeleteEscalationChainByIDDefault{
		_statusCode: code,
	}
}

/* DeleteEscalationChainByIDDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteEscalationChainByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete escalation chain by Id default response
func (o *DeleteEscalationChainByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEscalationChainByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /setting/alert/chains/{id}][%d] deleteEscalationChainById default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteEscalationChainByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteEscalationChainByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
