// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostContentGenerateReader is a Reader for the PostContentGenerate structure.
type PostContentGenerateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostContentGenerateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostContentGenerateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostContentGenerateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostContentGenerateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostContentGenerateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /content/generate] PostContentGenerate", response, response.Code())
	}
}

// NewPostContentGenerateOK creates a PostContentGenerateOK with default headers values
func NewPostContentGenerateOK() *PostContentGenerateOK {
	return &PostContentGenerateOK{}
}

/*
PostContentGenerateOK describes a response with status code 200, with default header values.

OK
*/
type PostContentGenerateOK struct {
}

// IsSuccess returns true when this post content generate o k response has a 2xx status code
func (o *PostContentGenerateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post content generate o k response has a 3xx status code
func (o *PostContentGenerateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post content generate o k response has a 4xx status code
func (o *PostContentGenerateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post content generate o k response has a 5xx status code
func (o *PostContentGenerateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post content generate o k response a status code equal to that given
func (o *PostContentGenerateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post content generate o k response
func (o *PostContentGenerateOK) Code() int {
	return 200
}

func (o *PostContentGenerateOK) Error() string {
	return fmt.Sprintf("[POST /content/generate][%d] postContentGenerateOK", 200)
}

func (o *PostContentGenerateOK) String() string {
	return fmt.Sprintf("[POST /content/generate][%d] postContentGenerateOK", 200)
}

func (o *PostContentGenerateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostContentGenerateBadRequest creates a PostContentGenerateBadRequest with default headers values
func NewPostContentGenerateBadRequest() *PostContentGenerateBadRequest {
	return &PostContentGenerateBadRequest{}
}

/*
PostContentGenerateBadRequest describes a response with status code 400, with default header values.

Malformed request
*/
type PostContentGenerateBadRequest struct {
}

// IsSuccess returns true when this post content generate bad request response has a 2xx status code
func (o *PostContentGenerateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post content generate bad request response has a 3xx status code
func (o *PostContentGenerateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post content generate bad request response has a 4xx status code
func (o *PostContentGenerateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post content generate bad request response has a 5xx status code
func (o *PostContentGenerateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post content generate bad request response a status code equal to that given
func (o *PostContentGenerateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post content generate bad request response
func (o *PostContentGenerateBadRequest) Code() int {
	return 400
}

func (o *PostContentGenerateBadRequest) Error() string {
	return fmt.Sprintf("[POST /content/generate][%d] postContentGenerateBadRequest", 400)
}

func (o *PostContentGenerateBadRequest) String() string {
	return fmt.Sprintf("[POST /content/generate][%d] postContentGenerateBadRequest", 400)
}

func (o *PostContentGenerateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostContentGenerateUnauthorized creates a PostContentGenerateUnauthorized with default headers values
func NewPostContentGenerateUnauthorized() *PostContentGenerateUnauthorized {
	return &PostContentGenerateUnauthorized{}
}

/*
PostContentGenerateUnauthorized describes a response with status code 401, with default header values.

API key is missing or invalid
*/
type PostContentGenerateUnauthorized struct {
	WWWAuthenticate string
}

// IsSuccess returns true when this post content generate unauthorized response has a 2xx status code
func (o *PostContentGenerateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post content generate unauthorized response has a 3xx status code
func (o *PostContentGenerateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post content generate unauthorized response has a 4xx status code
func (o *PostContentGenerateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post content generate unauthorized response has a 5xx status code
func (o *PostContentGenerateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post content generate unauthorized response a status code equal to that given
func (o *PostContentGenerateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post content generate unauthorized response
func (o *PostContentGenerateUnauthorized) Code() int {
	return 401
}

func (o *PostContentGenerateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /content/generate][%d] postContentGenerateUnauthorized", 401)
}

func (o *PostContentGenerateUnauthorized) String() string {
	return fmt.Sprintf("[POST /content/generate][%d] postContentGenerateUnauthorized", 401)
}

func (o *PostContentGenerateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header WWW_Authenticate
	hdrWWWAuthenticate := response.GetHeader("WWW_Authenticate")

	if hdrWWWAuthenticate != "" {
		o.WWWAuthenticate = hdrWWWAuthenticate
	}

	return nil
}

// NewPostContentGenerateInternalServerError creates a PostContentGenerateInternalServerError with default headers values
func NewPostContentGenerateInternalServerError() *PostContentGenerateInternalServerError {
	return &PostContentGenerateInternalServerError{}
}

/*
PostContentGenerateInternalServerError describes a response with status code 500, with default header values.

Something didin't work on our side. Not your fault!
*/
type PostContentGenerateInternalServerError struct {
	WWWAuthenticate string
}

// IsSuccess returns true when this post content generate internal server error response has a 2xx status code
func (o *PostContentGenerateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post content generate internal server error response has a 3xx status code
func (o *PostContentGenerateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post content generate internal server error response has a 4xx status code
func (o *PostContentGenerateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post content generate internal server error response has a 5xx status code
func (o *PostContentGenerateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post content generate internal server error response a status code equal to that given
func (o *PostContentGenerateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post content generate internal server error response
func (o *PostContentGenerateInternalServerError) Code() int {
	return 500
}

func (o *PostContentGenerateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /content/generate][%d] postContentGenerateInternalServerError", 500)
}

func (o *PostContentGenerateInternalServerError) String() string {
	return fmt.Sprintf("[POST /content/generate][%d] postContentGenerateInternalServerError", 500)
}

func (o *PostContentGenerateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header WWW_Authenticate
	hdrWWWAuthenticate := response.GetHeader("WWW_Authenticate")

	if hdrWWWAuthenticate != "" {
		o.WWWAuthenticate = hdrWWWAuthenticate
	}

	return nil
}
