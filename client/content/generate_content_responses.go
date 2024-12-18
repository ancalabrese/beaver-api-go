// Code generated by go-swagger; DO NOT EDIT.

package content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GenerateContentReader is a Reader for the GenerateContent structure.
type GenerateContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGenerateContentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGenerateContentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGenerateContentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /content/generate] generateContent", response, response.Code())
	}
}

// NewGenerateContentOK creates a GenerateContentOK with default headers values
func NewGenerateContentOK() *GenerateContentOK {
	return &GenerateContentOK{}
}

/*
GenerateContentOK describes a response with status code 200, with default header values.

OK
*/
type GenerateContentOK struct {
	Payload string
}

// IsSuccess returns true when this generate content o k response has a 2xx status code
func (o *GenerateContentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate content o k response has a 3xx status code
func (o *GenerateContentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate content o k response has a 4xx status code
func (o *GenerateContentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate content o k response has a 5xx status code
func (o *GenerateContentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate content o k response a status code equal to that given
func (o *GenerateContentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the generate content o k response
func (o *GenerateContentOK) Code() int {
	return 200
}

func (o *GenerateContentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /content/generate][%d] generateContentOK %s", 200, payload)
}

func (o *GenerateContentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /content/generate][%d] generateContentOK %s", 200, payload)
}

func (o *GenerateContentOK) GetPayload() string {
	return o.Payload
}

func (o *GenerateContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateContentBadRequest creates a GenerateContentBadRequest with default headers values
func NewGenerateContentBadRequest() *GenerateContentBadRequest {
	return &GenerateContentBadRequest{}
}

/*
GenerateContentBadRequest describes a response with status code 400, with default header values.

Malformed request
*/
type GenerateContentBadRequest struct {
}

// IsSuccess returns true when this generate content bad request response has a 2xx status code
func (o *GenerateContentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate content bad request response has a 3xx status code
func (o *GenerateContentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate content bad request response has a 4xx status code
func (o *GenerateContentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate content bad request response has a 5xx status code
func (o *GenerateContentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this generate content bad request response a status code equal to that given
func (o *GenerateContentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the generate content bad request response
func (o *GenerateContentBadRequest) Code() int {
	return 400
}

func (o *GenerateContentBadRequest) Error() string {
	return fmt.Sprintf("[POST /content/generate][%d] generateContentBadRequest", 400)
}

func (o *GenerateContentBadRequest) String() string {
	return fmt.Sprintf("[POST /content/generate][%d] generateContentBadRequest", 400)
}

func (o *GenerateContentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGenerateContentUnauthorized creates a GenerateContentUnauthorized with default headers values
func NewGenerateContentUnauthorized() *GenerateContentUnauthorized {
	return &GenerateContentUnauthorized{}
}

/*
GenerateContentUnauthorized describes a response with status code 401, with default header values.

API key is missing or invalid
*/
type GenerateContentUnauthorized struct {
	WWWAuthenticate string
}

// IsSuccess returns true when this generate content unauthorized response has a 2xx status code
func (o *GenerateContentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate content unauthorized response has a 3xx status code
func (o *GenerateContentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate content unauthorized response has a 4xx status code
func (o *GenerateContentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate content unauthorized response has a 5xx status code
func (o *GenerateContentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this generate content unauthorized response a status code equal to that given
func (o *GenerateContentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the generate content unauthorized response
func (o *GenerateContentUnauthorized) Code() int {
	return 401
}

func (o *GenerateContentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /content/generate][%d] generateContentUnauthorized", 401)
}

func (o *GenerateContentUnauthorized) String() string {
	return fmt.Sprintf("[POST /content/generate][%d] generateContentUnauthorized", 401)
}

func (o *GenerateContentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header WWW_Authenticate
	hdrWWWAuthenticate := response.GetHeader("WWW_Authenticate")

	if hdrWWWAuthenticate != "" {
		o.WWWAuthenticate = hdrWWWAuthenticate
	}

	return nil
}

// NewGenerateContentInternalServerError creates a GenerateContentInternalServerError with default headers values
func NewGenerateContentInternalServerError() *GenerateContentInternalServerError {
	return &GenerateContentInternalServerError{}
}

/*
GenerateContentInternalServerError describes a response with status code 500, with default header values.

Something didin't work on our side. Not your fault!
*/
type GenerateContentInternalServerError struct {
	WWWAuthenticate string
}

// IsSuccess returns true when this generate content internal server error response has a 2xx status code
func (o *GenerateContentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate content internal server error response has a 3xx status code
func (o *GenerateContentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate content internal server error response has a 4xx status code
func (o *GenerateContentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate content internal server error response has a 5xx status code
func (o *GenerateContentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this generate content internal server error response a status code equal to that given
func (o *GenerateContentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the generate content internal server error response
func (o *GenerateContentInternalServerError) Code() int {
	return 500
}

func (o *GenerateContentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /content/generate][%d] generateContentInternalServerError", 500)
}

func (o *GenerateContentInternalServerError) String() string {
	return fmt.Sprintf("[POST /content/generate][%d] generateContentInternalServerError", 500)
}

func (o *GenerateContentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header WWW_Authenticate
	hdrWWWAuthenticate := response.GetHeader("WWW_Authenticate")

	if hdrWWWAuthenticate != "" {
		o.WWWAuthenticate = hdrWWWAuthenticate
	}

	return nil
}
