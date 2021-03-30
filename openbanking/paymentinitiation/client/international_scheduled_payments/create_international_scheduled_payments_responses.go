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

// CreateInternationalScheduledPaymentsReader is a Reader for the CreateInternationalScheduledPayments structure.
type CreateInternationalScheduledPaymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInternationalScheduledPaymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateInternationalScheduledPaymentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInternationalScheduledPaymentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateInternationalScheduledPaymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateInternationalScheduledPaymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateInternationalScheduledPaymentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateInternationalScheduledPaymentsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateInternationalScheduledPaymentsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateInternationalScheduledPaymentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateInternationalScheduledPaymentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateInternationalScheduledPaymentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateInternationalScheduledPaymentsCreated creates a CreateInternationalScheduledPaymentsCreated with default headers values
func NewCreateInternationalScheduledPaymentsCreated() *CreateInternationalScheduledPaymentsCreated {
	return &CreateInternationalScheduledPaymentsCreated{}
}

/* CreateInternationalScheduledPaymentsCreated describes a response with status code 201, with default header values.

International Scheduled Payments Created
*/
type CreateInternationalScheduledPaymentsCreated struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteInternationalScheduledResponse6
}

func (o *CreateInternationalScheduledPaymentsCreated) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payments][%d] createInternationalScheduledPaymentsCreated  %+v", 201, o.Payload)
}
func (o *CreateInternationalScheduledPaymentsCreated) GetPayload() *models.OBWriteInternationalScheduledResponse6 {
	return o.Payload
}

func (o *CreateInternationalScheduledPaymentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.OBWriteInternationalScheduledResponse6)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalScheduledPaymentsBadRequest creates a CreateInternationalScheduledPaymentsBadRequest with default headers values
func NewCreateInternationalScheduledPaymentsBadRequest() *CreateInternationalScheduledPaymentsBadRequest {
	return &CreateInternationalScheduledPaymentsBadRequest{}
}

/* CreateInternationalScheduledPaymentsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateInternationalScheduledPaymentsBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateInternationalScheduledPaymentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payments][%d] createInternationalScheduledPaymentsBadRequest  %+v", 400, o.Payload)
}
func (o *CreateInternationalScheduledPaymentsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateInternationalScheduledPaymentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateInternationalScheduledPaymentsUnauthorized creates a CreateInternationalScheduledPaymentsUnauthorized with default headers values
func NewCreateInternationalScheduledPaymentsUnauthorized() *CreateInternationalScheduledPaymentsUnauthorized {
	return &CreateInternationalScheduledPaymentsUnauthorized{}
}

/* CreateInternationalScheduledPaymentsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateInternationalScheduledPaymentsUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalScheduledPaymentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payments][%d] createInternationalScheduledPaymentsUnauthorized ", 401)
}

func (o *CreateInternationalScheduledPaymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalScheduledPaymentsForbidden creates a CreateInternationalScheduledPaymentsForbidden with default headers values
func NewCreateInternationalScheduledPaymentsForbidden() *CreateInternationalScheduledPaymentsForbidden {
	return &CreateInternationalScheduledPaymentsForbidden{}
}

/* CreateInternationalScheduledPaymentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateInternationalScheduledPaymentsForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateInternationalScheduledPaymentsForbidden) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payments][%d] createInternationalScheduledPaymentsForbidden  %+v", 403, o.Payload)
}
func (o *CreateInternationalScheduledPaymentsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateInternationalScheduledPaymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateInternationalScheduledPaymentsNotFound creates a CreateInternationalScheduledPaymentsNotFound with default headers values
func NewCreateInternationalScheduledPaymentsNotFound() *CreateInternationalScheduledPaymentsNotFound {
	return &CreateInternationalScheduledPaymentsNotFound{}
}

/* CreateInternationalScheduledPaymentsNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateInternationalScheduledPaymentsNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalScheduledPaymentsNotFound) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payments][%d] createInternationalScheduledPaymentsNotFound ", 404)
}

func (o *CreateInternationalScheduledPaymentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalScheduledPaymentsMethodNotAllowed creates a CreateInternationalScheduledPaymentsMethodNotAllowed with default headers values
func NewCreateInternationalScheduledPaymentsMethodNotAllowed() *CreateInternationalScheduledPaymentsMethodNotAllowed {
	return &CreateInternationalScheduledPaymentsMethodNotAllowed{}
}

/* CreateInternationalScheduledPaymentsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type CreateInternationalScheduledPaymentsMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalScheduledPaymentsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payments][%d] createInternationalScheduledPaymentsMethodNotAllowed ", 405)
}

func (o *CreateInternationalScheduledPaymentsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalScheduledPaymentsNotAcceptable creates a CreateInternationalScheduledPaymentsNotAcceptable with default headers values
func NewCreateInternationalScheduledPaymentsNotAcceptable() *CreateInternationalScheduledPaymentsNotAcceptable {
	return &CreateInternationalScheduledPaymentsNotAcceptable{}
}

/* CreateInternationalScheduledPaymentsNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type CreateInternationalScheduledPaymentsNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalScheduledPaymentsNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payments][%d] createInternationalScheduledPaymentsNotAcceptable ", 406)
}

func (o *CreateInternationalScheduledPaymentsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalScheduledPaymentsUnsupportedMediaType creates a CreateInternationalScheduledPaymentsUnsupportedMediaType with default headers values
func NewCreateInternationalScheduledPaymentsUnsupportedMediaType() *CreateInternationalScheduledPaymentsUnsupportedMediaType {
	return &CreateInternationalScheduledPaymentsUnsupportedMediaType{}
}

/* CreateInternationalScheduledPaymentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type
*/
type CreateInternationalScheduledPaymentsUnsupportedMediaType struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalScheduledPaymentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payments][%d] createInternationalScheduledPaymentsUnsupportedMediaType ", 415)
}

func (o *CreateInternationalScheduledPaymentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalScheduledPaymentsTooManyRequests creates a CreateInternationalScheduledPaymentsTooManyRequests with default headers values
func NewCreateInternationalScheduledPaymentsTooManyRequests() *CreateInternationalScheduledPaymentsTooManyRequests {
	return &CreateInternationalScheduledPaymentsTooManyRequests{}
}

/* CreateInternationalScheduledPaymentsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateInternationalScheduledPaymentsTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalScheduledPaymentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payments][%d] createInternationalScheduledPaymentsTooManyRequests ", 429)
}

func (o *CreateInternationalScheduledPaymentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateInternationalScheduledPaymentsInternalServerError creates a CreateInternationalScheduledPaymentsInternalServerError with default headers values
func NewCreateInternationalScheduledPaymentsInternalServerError() *CreateInternationalScheduledPaymentsInternalServerError {
	return &CreateInternationalScheduledPaymentsInternalServerError{}
}

/* CreateInternationalScheduledPaymentsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateInternationalScheduledPaymentsInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateInternationalScheduledPaymentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /international-scheduled-payments][%d] createInternationalScheduledPaymentsInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateInternationalScheduledPaymentsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateInternationalScheduledPaymentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
