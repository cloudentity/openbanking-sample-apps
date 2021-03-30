// Code generated by go-swagger; DO NOT EDIT.

package international_payments

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

// GetInternationalPaymentsInternationalPaymentIDReader is a Reader for the GetInternationalPaymentsInternationalPaymentID structure.
type GetInternationalPaymentsInternationalPaymentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInternationalPaymentsInternationalPaymentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInternationalPaymentsInternationalPaymentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInternationalPaymentsInternationalPaymentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInternationalPaymentsInternationalPaymentIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInternationalPaymentsInternationalPaymentIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInternationalPaymentsInternationalPaymentIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetInternationalPaymentsInternationalPaymentIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetInternationalPaymentsInternationalPaymentIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInternationalPaymentsInternationalPaymentIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInternationalPaymentsInternationalPaymentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInternationalPaymentsInternationalPaymentIDOK creates a GetInternationalPaymentsInternationalPaymentIDOK with default headers values
func NewGetInternationalPaymentsInternationalPaymentIDOK() *GetInternationalPaymentsInternationalPaymentIDOK {
	return &GetInternationalPaymentsInternationalPaymentIDOK{}
}

/* GetInternationalPaymentsInternationalPaymentIDOK describes a response with status code 200, with default header values.

International Payments Read
*/
type GetInternationalPaymentsInternationalPaymentIDOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteInternationalResponse5
}

func (o *GetInternationalPaymentsInternationalPaymentIDOK) Error() string {
	return fmt.Sprintf("[GET /international-payments/{InternationalPaymentId}][%d] getInternationalPaymentsInternationalPaymentIdOK  %+v", 200, o.Payload)
}
func (o *GetInternationalPaymentsInternationalPaymentIDOK) GetPayload() *models.OBWriteInternationalResponse5 {
	return o.Payload
}

func (o *GetInternationalPaymentsInternationalPaymentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.OBWriteInternationalResponse5)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternationalPaymentsInternationalPaymentIDBadRequest creates a GetInternationalPaymentsInternationalPaymentIDBadRequest with default headers values
func NewGetInternationalPaymentsInternationalPaymentIDBadRequest() *GetInternationalPaymentsInternationalPaymentIDBadRequest {
	return &GetInternationalPaymentsInternationalPaymentIDBadRequest{}
}

/* GetInternationalPaymentsInternationalPaymentIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetInternationalPaymentsInternationalPaymentIDBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetInternationalPaymentsInternationalPaymentIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /international-payments/{InternationalPaymentId}][%d] getInternationalPaymentsInternationalPaymentIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetInternationalPaymentsInternationalPaymentIDBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetInternationalPaymentsInternationalPaymentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInternationalPaymentsInternationalPaymentIDUnauthorized creates a GetInternationalPaymentsInternationalPaymentIDUnauthorized with default headers values
func NewGetInternationalPaymentsInternationalPaymentIDUnauthorized() *GetInternationalPaymentsInternationalPaymentIDUnauthorized {
	return &GetInternationalPaymentsInternationalPaymentIDUnauthorized{}
}

/* GetInternationalPaymentsInternationalPaymentIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInternationalPaymentsInternationalPaymentIDUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalPaymentsInternationalPaymentIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /international-payments/{InternationalPaymentId}][%d] getInternationalPaymentsInternationalPaymentIdUnauthorized ", 401)
}

func (o *GetInternationalPaymentsInternationalPaymentIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetInternationalPaymentsInternationalPaymentIDForbidden creates a GetInternationalPaymentsInternationalPaymentIDForbidden with default headers values
func NewGetInternationalPaymentsInternationalPaymentIDForbidden() *GetInternationalPaymentsInternationalPaymentIDForbidden {
	return &GetInternationalPaymentsInternationalPaymentIDForbidden{}
}

/* GetInternationalPaymentsInternationalPaymentIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInternationalPaymentsInternationalPaymentIDForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetInternationalPaymentsInternationalPaymentIDForbidden) Error() string {
	return fmt.Sprintf("[GET /international-payments/{InternationalPaymentId}][%d] getInternationalPaymentsInternationalPaymentIdForbidden  %+v", 403, o.Payload)
}
func (o *GetInternationalPaymentsInternationalPaymentIDForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetInternationalPaymentsInternationalPaymentIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInternationalPaymentsInternationalPaymentIDNotFound creates a GetInternationalPaymentsInternationalPaymentIDNotFound with default headers values
func NewGetInternationalPaymentsInternationalPaymentIDNotFound() *GetInternationalPaymentsInternationalPaymentIDNotFound {
	return &GetInternationalPaymentsInternationalPaymentIDNotFound{}
}

/* GetInternationalPaymentsInternationalPaymentIDNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetInternationalPaymentsInternationalPaymentIDNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalPaymentsInternationalPaymentIDNotFound) Error() string {
	return fmt.Sprintf("[GET /international-payments/{InternationalPaymentId}][%d] getInternationalPaymentsInternationalPaymentIdNotFound ", 404)
}

func (o *GetInternationalPaymentsInternationalPaymentIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetInternationalPaymentsInternationalPaymentIDMethodNotAllowed creates a GetInternationalPaymentsInternationalPaymentIDMethodNotAllowed with default headers values
func NewGetInternationalPaymentsInternationalPaymentIDMethodNotAllowed() *GetInternationalPaymentsInternationalPaymentIDMethodNotAllowed {
	return &GetInternationalPaymentsInternationalPaymentIDMethodNotAllowed{}
}

/* GetInternationalPaymentsInternationalPaymentIDMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetInternationalPaymentsInternationalPaymentIDMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalPaymentsInternationalPaymentIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /international-payments/{InternationalPaymentId}][%d] getInternationalPaymentsInternationalPaymentIdMethodNotAllowed ", 405)
}

func (o *GetInternationalPaymentsInternationalPaymentIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetInternationalPaymentsInternationalPaymentIDNotAcceptable creates a GetInternationalPaymentsInternationalPaymentIDNotAcceptable with default headers values
func NewGetInternationalPaymentsInternationalPaymentIDNotAcceptable() *GetInternationalPaymentsInternationalPaymentIDNotAcceptable {
	return &GetInternationalPaymentsInternationalPaymentIDNotAcceptable{}
}

/* GetInternationalPaymentsInternationalPaymentIDNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetInternationalPaymentsInternationalPaymentIDNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalPaymentsInternationalPaymentIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /international-payments/{InternationalPaymentId}][%d] getInternationalPaymentsInternationalPaymentIdNotAcceptable ", 406)
}

func (o *GetInternationalPaymentsInternationalPaymentIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetInternationalPaymentsInternationalPaymentIDTooManyRequests creates a GetInternationalPaymentsInternationalPaymentIDTooManyRequests with default headers values
func NewGetInternationalPaymentsInternationalPaymentIDTooManyRequests() *GetInternationalPaymentsInternationalPaymentIDTooManyRequests {
	return &GetInternationalPaymentsInternationalPaymentIDTooManyRequests{}
}

/* GetInternationalPaymentsInternationalPaymentIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetInternationalPaymentsInternationalPaymentIDTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalPaymentsInternationalPaymentIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /international-payments/{InternationalPaymentId}][%d] getInternationalPaymentsInternationalPaymentIdTooManyRequests ", 429)
}

func (o *GetInternationalPaymentsInternationalPaymentIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInternationalPaymentsInternationalPaymentIDInternalServerError creates a GetInternationalPaymentsInternationalPaymentIDInternalServerError with default headers values
func NewGetInternationalPaymentsInternationalPaymentIDInternalServerError() *GetInternationalPaymentsInternationalPaymentIDInternalServerError {
	return &GetInternationalPaymentsInternationalPaymentIDInternalServerError{}
}

/* GetInternationalPaymentsInternationalPaymentIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInternationalPaymentsInternationalPaymentIDInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetInternationalPaymentsInternationalPaymentIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /international-payments/{InternationalPaymentId}][%d] getInternationalPaymentsInternationalPaymentIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetInternationalPaymentsInternationalPaymentIDInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetInternationalPaymentsInternationalPaymentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
