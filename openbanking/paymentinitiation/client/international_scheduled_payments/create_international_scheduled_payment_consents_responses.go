// Code generated by go-swagger; DO NOT EDIT.

package international_scheduled_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/openbanking-sample-apps/openbanking/paymentinitiation/models"
)

// CreateInternationalScheduledPaymentConsentsReader is a Reader for the CreateInternationalScheduledPaymentConsents structure.
type CreateInternationalScheduledPaymentConsentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInternationalScheduledPaymentConsentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateInternationalScheduledPaymentConsentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInternationalScheduledPaymentConsentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateInternationalScheduledPaymentConsentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateInternationalScheduledPaymentConsentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateInternationalScheduledPaymentConsentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateInternationalScheduledPaymentConsentsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateInternationalScheduledPaymentConsentsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateInternationalScheduledPaymentConsentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateInternationalScheduledPaymentConsentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateInternationalScheduledPaymentConsentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateInternationalScheduledPaymentConsentsCreated creates a CreateInternationalScheduledPaymentConsentsCreated with default headers values
func NewCreateInternationalScheduledPaymentConsentsCreated() *CreateInternationalScheduledPaymentConsentsCreated {
	return &CreateInternationalScheduledPaymentConsentsCreated{}
}

/* CreateInternationalScheduledPaymentConsentsCreated describes a response with status code 201, with default header values.

International Scheduled Payment Consents Created
*/
type CreateInternationalScheduledPaymentConsentsCreated struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteInternationalScheduledConsentResponse6
}

func (o *CreateInternationalScheduledPaymentConsentsCreated) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payment-consents][%d] createInternationalScheduledPaymentConsentsCreated  %+v", 201, o.Payload)
}
func (o *CreateInternationalScheduledPaymentConsentsCreated) GetPayload() *models.OBWriteInternationalScheduledConsentResponse6 {
	return o.Payload
}

func (o *CreateInternationalScheduledPaymentConsentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBWriteInternationalScheduledConsentResponse6)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalScheduledPaymentConsentsBadRequest creates a CreateInternationalScheduledPaymentConsentsBadRequest with default headers values
func NewCreateInternationalScheduledPaymentConsentsBadRequest() *CreateInternationalScheduledPaymentConsentsBadRequest {
	return &CreateInternationalScheduledPaymentConsentsBadRequest{}
}

/* CreateInternationalScheduledPaymentConsentsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateInternationalScheduledPaymentConsentsBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateInternationalScheduledPaymentConsentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payment-consents][%d] createInternationalScheduledPaymentConsentsBadRequest  %+v", 400, o.Payload)
}
func (o *CreateInternationalScheduledPaymentConsentsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateInternationalScheduledPaymentConsentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalScheduledPaymentConsentsUnauthorized creates a CreateInternationalScheduledPaymentConsentsUnauthorized with default headers values
func NewCreateInternationalScheduledPaymentConsentsUnauthorized() *CreateInternationalScheduledPaymentConsentsUnauthorized {
	return &CreateInternationalScheduledPaymentConsentsUnauthorized{}
}

/* CreateInternationalScheduledPaymentConsentsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateInternationalScheduledPaymentConsentsUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalScheduledPaymentConsentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payment-consents][%d] createInternationalScheduledPaymentConsentsUnauthorized ", 401)
}

func (o *CreateInternationalScheduledPaymentConsentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalScheduledPaymentConsentsForbidden creates a CreateInternationalScheduledPaymentConsentsForbidden with default headers values
func NewCreateInternationalScheduledPaymentConsentsForbidden() *CreateInternationalScheduledPaymentConsentsForbidden {
	return &CreateInternationalScheduledPaymentConsentsForbidden{}
}

/* CreateInternationalScheduledPaymentConsentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateInternationalScheduledPaymentConsentsForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateInternationalScheduledPaymentConsentsForbidden) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payment-consents][%d] createInternationalScheduledPaymentConsentsForbidden  %+v", 403, o.Payload)
}
func (o *CreateInternationalScheduledPaymentConsentsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateInternationalScheduledPaymentConsentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalScheduledPaymentConsentsNotFound creates a CreateInternationalScheduledPaymentConsentsNotFound with default headers values
func NewCreateInternationalScheduledPaymentConsentsNotFound() *CreateInternationalScheduledPaymentConsentsNotFound {
	return &CreateInternationalScheduledPaymentConsentsNotFound{}
}

/* CreateInternationalScheduledPaymentConsentsNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateInternationalScheduledPaymentConsentsNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalScheduledPaymentConsentsNotFound) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payment-consents][%d] createInternationalScheduledPaymentConsentsNotFound ", 404)
}

func (o *CreateInternationalScheduledPaymentConsentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalScheduledPaymentConsentsMethodNotAllowed creates a CreateInternationalScheduledPaymentConsentsMethodNotAllowed with default headers values
func NewCreateInternationalScheduledPaymentConsentsMethodNotAllowed() *CreateInternationalScheduledPaymentConsentsMethodNotAllowed {
	return &CreateInternationalScheduledPaymentConsentsMethodNotAllowed{}
}

/* CreateInternationalScheduledPaymentConsentsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type CreateInternationalScheduledPaymentConsentsMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalScheduledPaymentConsentsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payment-consents][%d] createInternationalScheduledPaymentConsentsMethodNotAllowed ", 405)
}

func (o *CreateInternationalScheduledPaymentConsentsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalScheduledPaymentConsentsNotAcceptable creates a CreateInternationalScheduledPaymentConsentsNotAcceptable with default headers values
func NewCreateInternationalScheduledPaymentConsentsNotAcceptable() *CreateInternationalScheduledPaymentConsentsNotAcceptable {
	return &CreateInternationalScheduledPaymentConsentsNotAcceptable{}
}

/* CreateInternationalScheduledPaymentConsentsNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type CreateInternationalScheduledPaymentConsentsNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalScheduledPaymentConsentsNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payment-consents][%d] createInternationalScheduledPaymentConsentsNotAcceptable ", 406)
}

func (o *CreateInternationalScheduledPaymentConsentsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalScheduledPaymentConsentsUnsupportedMediaType creates a CreateInternationalScheduledPaymentConsentsUnsupportedMediaType with default headers values
func NewCreateInternationalScheduledPaymentConsentsUnsupportedMediaType() *CreateInternationalScheduledPaymentConsentsUnsupportedMediaType {
	return &CreateInternationalScheduledPaymentConsentsUnsupportedMediaType{}
}

/* CreateInternationalScheduledPaymentConsentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type
*/
type CreateInternationalScheduledPaymentConsentsUnsupportedMediaType struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalScheduledPaymentConsentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payment-consents][%d] createInternationalScheduledPaymentConsentsUnsupportedMediaType ", 415)
}

func (o *CreateInternationalScheduledPaymentConsentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalScheduledPaymentConsentsTooManyRequests creates a CreateInternationalScheduledPaymentConsentsTooManyRequests with default headers values
func NewCreateInternationalScheduledPaymentConsentsTooManyRequests() *CreateInternationalScheduledPaymentConsentsTooManyRequests {
	return &CreateInternationalScheduledPaymentConsentsTooManyRequests{}
}

/* CreateInternationalScheduledPaymentConsentsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateInternationalScheduledPaymentConsentsTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalScheduledPaymentConsentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payment-consents][%d] createInternationalScheduledPaymentConsentsTooManyRequests ", 429)
}

func (o *CreateInternationalScheduledPaymentConsentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateInternationalScheduledPaymentConsentsInternalServerError creates a CreateInternationalScheduledPaymentConsentsInternalServerError with default headers values
func NewCreateInternationalScheduledPaymentConsentsInternalServerError() *CreateInternationalScheduledPaymentConsentsInternalServerError {
	return &CreateInternationalScheduledPaymentConsentsInternalServerError{}
}

/* CreateInternationalScheduledPaymentConsentsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateInternationalScheduledPaymentConsentsInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateInternationalScheduledPaymentConsentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payment-consents][%d] createInternationalScheduledPaymentConsentsInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateInternationalScheduledPaymentConsentsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateInternationalScheduledPaymentConsentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
