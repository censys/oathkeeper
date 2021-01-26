// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ory/oathkeeper/internal/httpclient/models"
)

// GetWellKnownJSONWebKeysReader is a Reader for the GetWellKnownJSONWebKeys structure.
type GetWellKnownJSONWebKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWellKnownJSONWebKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWellKnownJSONWebKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetWellKnownJSONWebKeysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWellKnownJSONWebKeysOK creates a GetWellKnownJSONWebKeysOK with default headers values
func NewGetWellKnownJSONWebKeysOK() *GetWellKnownJSONWebKeysOK {
	return &GetWellKnownJSONWebKeysOK{}
}

/* GetWellKnownJSONWebKeysOK describes a response with status code 200, with default header values.

jsonWebKeySet
*/
type GetWellKnownJSONWebKeysOK struct {
	Payload *models.JSONWebKeySet
}

func (o *GetWellKnownJSONWebKeysOK) Error() string {
	return fmt.Sprintf("[GET /.well-known/jwks.json][%d] getWellKnownJsonWebKeysOK  %+v", 200, o.Payload)
}
func (o *GetWellKnownJSONWebKeysOK) GetPayload() *models.JSONWebKeySet {
	return o.Payload
}

func (o *GetWellKnownJSONWebKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONWebKeySet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWellKnownJSONWebKeysInternalServerError creates a GetWellKnownJSONWebKeysInternalServerError with default headers values
func NewGetWellKnownJSONWebKeysInternalServerError() *GetWellKnownJSONWebKeysInternalServerError {
	return &GetWellKnownJSONWebKeysInternalServerError{}
}

/* GetWellKnownJSONWebKeysInternalServerError describes a response with status code 500, with default header values.

The standard error format
*/
type GetWellKnownJSONWebKeysInternalServerError struct {
	Payload *GetWellKnownJSONWebKeysInternalServerErrorBody
}

func (o *GetWellKnownJSONWebKeysInternalServerError) Error() string {
	return fmt.Sprintf("[GET /.well-known/jwks.json][%d] getWellKnownJsonWebKeysInternalServerError  %+v", 500, o.Payload)
}
func (o *GetWellKnownJSONWebKeysInternalServerError) GetPayload() *GetWellKnownJSONWebKeysInternalServerErrorBody {
	return o.Payload
}

func (o *GetWellKnownJSONWebKeysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWellKnownJSONWebKeysInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWellKnownJSONWebKeysInternalServerErrorBody get well known JSON web keys internal server error body
swagger:model GetWellKnownJSONWebKeysInternalServerErrorBody
*/
type GetWellKnownJSONWebKeysInternalServerErrorBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this get well known JSON web keys internal server error body
func (o *GetWellKnownJSONWebKeysInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get well known JSON web keys internal server error body based on context it is used
func (o *GetWellKnownJSONWebKeysInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWellKnownJSONWebKeysInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWellKnownJSONWebKeysInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetWellKnownJSONWebKeysInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
