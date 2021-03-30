// Code generated by go-swagger; DO NOT EDIT.

package file_payments

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

// GetFilePaymentConsentsConsentIDFileReader is a Reader for the GetFilePaymentConsentsConsentIDFile structure.
type GetFilePaymentConsentsConsentIDFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilePaymentConsentsConsentIDFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFilePaymentConsentsConsentIDFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFilePaymentConsentsConsentIDFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFilePaymentConsentsConsentIDFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFilePaymentConsentsConsentIDFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFilePaymentConsentsConsentIDFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetFilePaymentConsentsConsentIDFileMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetFilePaymentConsentsConsentIDFileNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFilePaymentConsentsConsentIDFileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFilePaymentConsentsConsentIDFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFilePaymentConsentsConsentIDFileOK creates a GetFilePaymentConsentsConsentIDFileOK with default headers values
func NewGetFilePaymentConsentsConsentIDFileOK() *GetFilePaymentConsentsConsentIDFileOK {
	return &GetFilePaymentConsentsConsentIDFileOK{}
}

/* GetFilePaymentConsentsConsentIDFileOK describes a response with status code 200, with default header values.

File Payment Consents Read
*/
type GetFilePaymentConsentsConsentIDFileOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload models.File
}

func (o *GetFilePaymentConsentsConsentIDFileOK) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}/file][%d] getFilePaymentConsentsConsentIdFileOK  %+v", 200, o.Payload)
}
func (o *GetFilePaymentConsentsConsentIDFileOK) GetPayload() models.File {
	return o.Payload
}

func (o *GetFilePaymentConsentsConsentIDFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDFileBadRequest creates a GetFilePaymentConsentsConsentIDFileBadRequest with default headers values
func NewGetFilePaymentConsentsConsentIDFileBadRequest() *GetFilePaymentConsentsConsentIDFileBadRequest {
	return &GetFilePaymentConsentsConsentIDFileBadRequest{}
}

/* GetFilePaymentConsentsConsentIDFileBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetFilePaymentConsentsConsentIDFileBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetFilePaymentConsentsConsentIDFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}/file][%d] getFilePaymentConsentsConsentIdFileBadRequest  %+v", 400, o.Payload)
}
func (o *GetFilePaymentConsentsConsentIDFileBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetFilePaymentConsentsConsentIDFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFilePaymentConsentsConsentIDFileUnauthorized creates a GetFilePaymentConsentsConsentIDFileUnauthorized with default headers values
func NewGetFilePaymentConsentsConsentIDFileUnauthorized() *GetFilePaymentConsentsConsentIDFileUnauthorized {
	return &GetFilePaymentConsentsConsentIDFileUnauthorized{}
}

/* GetFilePaymentConsentsConsentIDFileUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFilePaymentConsentsConsentIDFileUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentConsentsConsentIDFileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}/file][%d] getFilePaymentConsentsConsentIdFileUnauthorized ", 401)
}

func (o *GetFilePaymentConsentsConsentIDFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDFileForbidden creates a GetFilePaymentConsentsConsentIDFileForbidden with default headers values
func NewGetFilePaymentConsentsConsentIDFileForbidden() *GetFilePaymentConsentsConsentIDFileForbidden {
	return &GetFilePaymentConsentsConsentIDFileForbidden{}
}

/* GetFilePaymentConsentsConsentIDFileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFilePaymentConsentsConsentIDFileForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetFilePaymentConsentsConsentIDFileForbidden) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}/file][%d] getFilePaymentConsentsConsentIdFileForbidden  %+v", 403, o.Payload)
}
func (o *GetFilePaymentConsentsConsentIDFileForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetFilePaymentConsentsConsentIDFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFilePaymentConsentsConsentIDFileNotFound creates a GetFilePaymentConsentsConsentIDFileNotFound with default headers values
func NewGetFilePaymentConsentsConsentIDFileNotFound() *GetFilePaymentConsentsConsentIDFileNotFound {
	return &GetFilePaymentConsentsConsentIDFileNotFound{}
}

/* GetFilePaymentConsentsConsentIDFileNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetFilePaymentConsentsConsentIDFileNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentConsentsConsentIDFileNotFound) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}/file][%d] getFilePaymentConsentsConsentIdFileNotFound ", 404)
}

func (o *GetFilePaymentConsentsConsentIDFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDFileMethodNotAllowed creates a GetFilePaymentConsentsConsentIDFileMethodNotAllowed with default headers values
func NewGetFilePaymentConsentsConsentIDFileMethodNotAllowed() *GetFilePaymentConsentsConsentIDFileMethodNotAllowed {
	return &GetFilePaymentConsentsConsentIDFileMethodNotAllowed{}
}

/* GetFilePaymentConsentsConsentIDFileMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetFilePaymentConsentsConsentIDFileMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentConsentsConsentIDFileMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}/file][%d] getFilePaymentConsentsConsentIdFileMethodNotAllowed ", 405)
}

func (o *GetFilePaymentConsentsConsentIDFileMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDFileNotAcceptable creates a GetFilePaymentConsentsConsentIDFileNotAcceptable with default headers values
func NewGetFilePaymentConsentsConsentIDFileNotAcceptable() *GetFilePaymentConsentsConsentIDFileNotAcceptable {
	return &GetFilePaymentConsentsConsentIDFileNotAcceptable{}
}

/* GetFilePaymentConsentsConsentIDFileNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetFilePaymentConsentsConsentIDFileNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentConsentsConsentIDFileNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}/file][%d] getFilePaymentConsentsConsentIdFileNotAcceptable ", 406)
}

func (o *GetFilePaymentConsentsConsentIDFileNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDFileTooManyRequests creates a GetFilePaymentConsentsConsentIDFileTooManyRequests with default headers values
func NewGetFilePaymentConsentsConsentIDFileTooManyRequests() *GetFilePaymentConsentsConsentIDFileTooManyRequests {
	return &GetFilePaymentConsentsConsentIDFileTooManyRequests{}
}

/* GetFilePaymentConsentsConsentIDFileTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetFilePaymentConsentsConsentIDFileTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentConsentsConsentIDFileTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}/file][%d] getFilePaymentConsentsConsentIdFileTooManyRequests ", 429)
}

func (o *GetFilePaymentConsentsConsentIDFileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFilePaymentConsentsConsentIDFileInternalServerError creates a GetFilePaymentConsentsConsentIDFileInternalServerError with default headers values
func NewGetFilePaymentConsentsConsentIDFileInternalServerError() *GetFilePaymentConsentsConsentIDFileInternalServerError {
	return &GetFilePaymentConsentsConsentIDFileInternalServerError{}
}

/* GetFilePaymentConsentsConsentIDFileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetFilePaymentConsentsConsentIDFileInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetFilePaymentConsentsConsentIDFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}/file][%d] getFilePaymentConsentsConsentIdFileInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFilePaymentConsentsConsentIDFileInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetFilePaymentConsentsConsentIDFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
