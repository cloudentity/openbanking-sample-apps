// Code generated by go-swagger; DO NOT EDIT.

package domestic_standing_orders

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

// CreateDomesticStandingOrdersReader is a Reader for the CreateDomesticStandingOrders structure.
type CreateDomesticStandingOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDomesticStandingOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDomesticStandingOrdersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDomesticStandingOrdersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateDomesticStandingOrdersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDomesticStandingOrdersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDomesticStandingOrdersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateDomesticStandingOrdersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateDomesticStandingOrdersNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateDomesticStandingOrdersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateDomesticStandingOrdersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDomesticStandingOrdersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDomesticStandingOrdersCreated creates a CreateDomesticStandingOrdersCreated with default headers values
func NewCreateDomesticStandingOrdersCreated() *CreateDomesticStandingOrdersCreated {
	return &CreateDomesticStandingOrdersCreated{}
}

/* CreateDomesticStandingOrdersCreated describes a response with status code 201, with default header values.

Domestic Standing Orders Created
*/
type CreateDomesticStandingOrdersCreated struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteDomesticStandingOrderResponse6
}

func (o *CreateDomesticStandingOrdersCreated) Error() string {
	return fmt.Sprintf("[POST /domestic-standing-orders][%d] createDomesticStandingOrdersCreated  %+v", 201, o.Payload)
}
func (o *CreateDomesticStandingOrdersCreated) GetPayload() *models.OBWriteDomesticStandingOrderResponse6 {
	return o.Payload
}

func (o *CreateDomesticStandingOrdersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.OBWriteDomesticStandingOrderResponse6)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomesticStandingOrdersBadRequest creates a CreateDomesticStandingOrdersBadRequest with default headers values
func NewCreateDomesticStandingOrdersBadRequest() *CreateDomesticStandingOrdersBadRequest {
	return &CreateDomesticStandingOrdersBadRequest{}
}

/* CreateDomesticStandingOrdersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateDomesticStandingOrdersBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateDomesticStandingOrdersBadRequest) Error() string {
	return fmt.Sprintf("[POST /domestic-standing-orders][%d] createDomesticStandingOrdersBadRequest  %+v", 400, o.Payload)
}
func (o *CreateDomesticStandingOrdersBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateDomesticStandingOrdersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDomesticStandingOrdersUnauthorized creates a CreateDomesticStandingOrdersUnauthorized with default headers values
func NewCreateDomesticStandingOrdersUnauthorized() *CreateDomesticStandingOrdersUnauthorized {
	return &CreateDomesticStandingOrdersUnauthorized{}
}

/* CreateDomesticStandingOrdersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateDomesticStandingOrdersUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticStandingOrdersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /domestic-standing-orders][%d] createDomesticStandingOrdersUnauthorized ", 401)
}

func (o *CreateDomesticStandingOrdersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticStandingOrdersForbidden creates a CreateDomesticStandingOrdersForbidden with default headers values
func NewCreateDomesticStandingOrdersForbidden() *CreateDomesticStandingOrdersForbidden {
	return &CreateDomesticStandingOrdersForbidden{}
}

/* CreateDomesticStandingOrdersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateDomesticStandingOrdersForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateDomesticStandingOrdersForbidden) Error() string {
	return fmt.Sprintf("[POST /domestic-standing-orders][%d] createDomesticStandingOrdersForbidden  %+v", 403, o.Payload)
}
func (o *CreateDomesticStandingOrdersForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateDomesticStandingOrdersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDomesticStandingOrdersNotFound creates a CreateDomesticStandingOrdersNotFound with default headers values
func NewCreateDomesticStandingOrdersNotFound() *CreateDomesticStandingOrdersNotFound {
	return &CreateDomesticStandingOrdersNotFound{}
}

/* CreateDomesticStandingOrdersNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateDomesticStandingOrdersNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticStandingOrdersNotFound) Error() string {
	return fmt.Sprintf("[POST /domestic-standing-orders][%d] createDomesticStandingOrdersNotFound ", 404)
}

func (o *CreateDomesticStandingOrdersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticStandingOrdersMethodNotAllowed creates a CreateDomesticStandingOrdersMethodNotAllowed with default headers values
func NewCreateDomesticStandingOrdersMethodNotAllowed() *CreateDomesticStandingOrdersMethodNotAllowed {
	return &CreateDomesticStandingOrdersMethodNotAllowed{}
}

/* CreateDomesticStandingOrdersMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type CreateDomesticStandingOrdersMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticStandingOrdersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /domestic-standing-orders][%d] createDomesticStandingOrdersMethodNotAllowed ", 405)
}

func (o *CreateDomesticStandingOrdersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticStandingOrdersNotAcceptable creates a CreateDomesticStandingOrdersNotAcceptable with default headers values
func NewCreateDomesticStandingOrdersNotAcceptable() *CreateDomesticStandingOrdersNotAcceptable {
	return &CreateDomesticStandingOrdersNotAcceptable{}
}

/* CreateDomesticStandingOrdersNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type CreateDomesticStandingOrdersNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticStandingOrdersNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /domestic-standing-orders][%d] createDomesticStandingOrdersNotAcceptable ", 406)
}

func (o *CreateDomesticStandingOrdersNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticStandingOrdersUnsupportedMediaType creates a CreateDomesticStandingOrdersUnsupportedMediaType with default headers values
func NewCreateDomesticStandingOrdersUnsupportedMediaType() *CreateDomesticStandingOrdersUnsupportedMediaType {
	return &CreateDomesticStandingOrdersUnsupportedMediaType{}
}

/* CreateDomesticStandingOrdersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type
*/
type CreateDomesticStandingOrdersUnsupportedMediaType struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticStandingOrdersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /domestic-standing-orders][%d] createDomesticStandingOrdersUnsupportedMediaType ", 415)
}

func (o *CreateDomesticStandingOrdersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticStandingOrdersTooManyRequests creates a CreateDomesticStandingOrdersTooManyRequests with default headers values
func NewCreateDomesticStandingOrdersTooManyRequests() *CreateDomesticStandingOrdersTooManyRequests {
	return &CreateDomesticStandingOrdersTooManyRequests{}
}

/* CreateDomesticStandingOrdersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateDomesticStandingOrdersTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticStandingOrdersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /domestic-standing-orders][%d] createDomesticStandingOrdersTooManyRequests ", 429)
}

func (o *CreateDomesticStandingOrdersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDomesticStandingOrdersInternalServerError creates a CreateDomesticStandingOrdersInternalServerError with default headers values
func NewCreateDomesticStandingOrdersInternalServerError() *CreateDomesticStandingOrdersInternalServerError {
	return &CreateDomesticStandingOrdersInternalServerError{}
}

/* CreateDomesticStandingOrdersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateDomesticStandingOrdersInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateDomesticStandingOrdersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /domestic-standing-orders][%d] createDomesticStandingOrdersInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateDomesticStandingOrdersInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateDomesticStandingOrdersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
