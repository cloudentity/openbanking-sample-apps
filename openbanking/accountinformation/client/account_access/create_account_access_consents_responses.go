// Code generated by go-swagger; DO NOT EDIT.

package account_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/openbanking-sample-apps/openbanking/accountinformation/models"
)

// CreateAccountAccessConsentsReader is a Reader for the CreateAccountAccessConsents structure.
type CreateAccountAccessConsentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccountAccessConsentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAccountAccessConsentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAccountAccessConsentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAccountAccessConsentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAccountAccessConsentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateAccountAccessConsentsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateAccountAccessConsentsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateAccountAccessConsentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateAccountAccessConsentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAccountAccessConsentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAccountAccessConsentsCreated creates a CreateAccountAccessConsentsCreated with default headers values
func NewCreateAccountAccessConsentsCreated() *CreateAccountAccessConsentsCreated {
	return &CreateAccountAccessConsentsCreated{}
}

/* CreateAccountAccessConsentsCreated describes a response with status code 201, with default header values.

Account Access Consents Created
*/
type CreateAccountAccessConsentsCreated struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadConsentResponse1
}

func (o *CreateAccountAccessConsentsCreated) Error() string {
	return fmt.Sprintf("[POST /account-access-consents][%d] createAccountAccessConsentsCreated  %+v", 201, o.Payload)
}
func (o *CreateAccountAccessConsentsCreated) GetPayload() *models.OBReadConsentResponse1 {
	return o.Payload
}

func (o *CreateAccountAccessConsentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBReadConsentResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAccessConsentsBadRequest creates a CreateAccountAccessConsentsBadRequest with default headers values
func NewCreateAccountAccessConsentsBadRequest() *CreateAccountAccessConsentsBadRequest {
	return &CreateAccountAccessConsentsBadRequest{}
}

/* CreateAccountAccessConsentsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateAccountAccessConsentsBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *CreateAccountAccessConsentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /account-access-consents][%d] createAccountAccessConsentsBadRequest  %+v", 400, o.Payload)
}
func (o *CreateAccountAccessConsentsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateAccountAccessConsentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAccessConsentsUnauthorized creates a CreateAccountAccessConsentsUnauthorized with default headers values
func NewCreateAccountAccessConsentsUnauthorized() *CreateAccountAccessConsentsUnauthorized {
	return &CreateAccountAccessConsentsUnauthorized{}
}

/* CreateAccountAccessConsentsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateAccountAccessConsentsUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateAccountAccessConsentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /account-access-consents][%d] createAccountAccessConsentsUnauthorized ", 401)
}

func (o *CreateAccountAccessConsentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateAccountAccessConsentsForbidden creates a CreateAccountAccessConsentsForbidden with default headers values
func NewCreateAccountAccessConsentsForbidden() *CreateAccountAccessConsentsForbidden {
	return &CreateAccountAccessConsentsForbidden{}
}

/* CreateAccountAccessConsentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateAccountAccessConsentsForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *CreateAccountAccessConsentsForbidden) Error() string {
	return fmt.Sprintf("[POST /account-access-consents][%d] createAccountAccessConsentsForbidden  %+v", 403, o.Payload)
}
func (o *CreateAccountAccessConsentsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateAccountAccessConsentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAccessConsentsMethodNotAllowed creates a CreateAccountAccessConsentsMethodNotAllowed with default headers values
func NewCreateAccountAccessConsentsMethodNotAllowed() *CreateAccountAccessConsentsMethodNotAllowed {
	return &CreateAccountAccessConsentsMethodNotAllowed{}
}

/* CreateAccountAccessConsentsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type CreateAccountAccessConsentsMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateAccountAccessConsentsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /account-access-consents][%d] createAccountAccessConsentsMethodNotAllowed ", 405)
}

func (o *CreateAccountAccessConsentsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateAccountAccessConsentsNotAcceptable creates a CreateAccountAccessConsentsNotAcceptable with default headers values
func NewCreateAccountAccessConsentsNotAcceptable() *CreateAccountAccessConsentsNotAcceptable {
	return &CreateAccountAccessConsentsNotAcceptable{}
}

/* CreateAccountAccessConsentsNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type CreateAccountAccessConsentsNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateAccountAccessConsentsNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /account-access-consents][%d] createAccountAccessConsentsNotAcceptable ", 406)
}

func (o *CreateAccountAccessConsentsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateAccountAccessConsentsUnsupportedMediaType creates a CreateAccountAccessConsentsUnsupportedMediaType with default headers values
func NewCreateAccountAccessConsentsUnsupportedMediaType() *CreateAccountAccessConsentsUnsupportedMediaType {
	return &CreateAccountAccessConsentsUnsupportedMediaType{}
}

/* CreateAccountAccessConsentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type
*/
type CreateAccountAccessConsentsUnsupportedMediaType struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateAccountAccessConsentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /account-access-consents][%d] createAccountAccessConsentsUnsupportedMediaType ", 415)
}

func (o *CreateAccountAccessConsentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateAccountAccessConsentsTooManyRequests creates a CreateAccountAccessConsentsTooManyRequests with default headers values
func NewCreateAccountAccessConsentsTooManyRequests() *CreateAccountAccessConsentsTooManyRequests {
	return &CreateAccountAccessConsentsTooManyRequests{}
}

/* CreateAccountAccessConsentsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateAccountAccessConsentsTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateAccountAccessConsentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /account-access-consents][%d] createAccountAccessConsentsTooManyRequests ", 429)
}

func (o *CreateAccountAccessConsentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Retry-After
	hdrRetryAfter := response.GetHeader("Retry-After")

	if hdrRetryAfter != "" {
		valretryAfter, err := swag.ConvertInt64(hdrRetryAfter)
		if err != nil {
			return errors.InvalidType("Retry-After", "header", "int64", hdrRetryAfter)
		}
		o.RetryAfter = valretryAfter
	}

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateAccountAccessConsentsInternalServerError creates a CreateAccountAccessConsentsInternalServerError with default headers values
func NewCreateAccountAccessConsentsInternalServerError() *CreateAccountAccessConsentsInternalServerError {
	return &CreateAccountAccessConsentsInternalServerError{}
}

/* CreateAccountAccessConsentsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateAccountAccessConsentsInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *CreateAccountAccessConsentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /account-access-consents][%d] createAccountAccessConsentsInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateAccountAccessConsentsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateAccountAccessConsentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
