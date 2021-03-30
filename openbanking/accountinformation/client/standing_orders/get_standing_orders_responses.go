// Code generated by go-swagger; DO NOT EDIT.

package standing_orders

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

// GetStandingOrdersReader is a Reader for the GetStandingOrders structure.
type GetStandingOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStandingOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStandingOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStandingOrdersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetStandingOrdersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetStandingOrdersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStandingOrdersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetStandingOrdersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetStandingOrdersNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetStandingOrdersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStandingOrdersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStandingOrdersOK creates a GetStandingOrdersOK with default headers values
func NewGetStandingOrdersOK() *GetStandingOrdersOK {
	return &GetStandingOrdersOK{}
}

/* GetStandingOrdersOK describes a response with status code 200, with default header values.

Standing Orders Read
*/
type GetStandingOrdersOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadStandingOrder6
}

func (o *GetStandingOrdersOK) Error() string {
	return fmt.Sprintf("[GET /standing-orders][%d] getStandingOrdersOK  %+v", 200, o.Payload)
}
func (o *GetStandingOrdersOK) GetPayload() *models.OBReadStandingOrder6 {
	return o.Payload
}

func (o *GetStandingOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBReadStandingOrder6)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStandingOrdersBadRequest creates a GetStandingOrdersBadRequest with default headers values
func NewGetStandingOrdersBadRequest() *GetStandingOrdersBadRequest {
	return &GetStandingOrdersBadRequest{}
}

/* GetStandingOrdersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetStandingOrdersBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetStandingOrdersBadRequest) Error() string {
	return fmt.Sprintf("[GET /standing-orders][%d] getStandingOrdersBadRequest  %+v", 400, o.Payload)
}
func (o *GetStandingOrdersBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetStandingOrdersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetStandingOrdersUnauthorized creates a GetStandingOrdersUnauthorized with default headers values
func NewGetStandingOrdersUnauthorized() *GetStandingOrdersUnauthorized {
	return &GetStandingOrdersUnauthorized{}
}

/* GetStandingOrdersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetStandingOrdersUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetStandingOrdersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /standing-orders][%d] getStandingOrdersUnauthorized ", 401)
}

func (o *GetStandingOrdersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetStandingOrdersForbidden creates a GetStandingOrdersForbidden with default headers values
func NewGetStandingOrdersForbidden() *GetStandingOrdersForbidden {
	return &GetStandingOrdersForbidden{}
}

/* GetStandingOrdersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetStandingOrdersForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetStandingOrdersForbidden) Error() string {
	return fmt.Sprintf("[GET /standing-orders][%d] getStandingOrdersForbidden  %+v", 403, o.Payload)
}
func (o *GetStandingOrdersForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetStandingOrdersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetStandingOrdersNotFound creates a GetStandingOrdersNotFound with default headers values
func NewGetStandingOrdersNotFound() *GetStandingOrdersNotFound {
	return &GetStandingOrdersNotFound{}
}

/* GetStandingOrdersNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetStandingOrdersNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetStandingOrdersNotFound) Error() string {
	return fmt.Sprintf("[GET /standing-orders][%d] getStandingOrdersNotFound ", 404)
}

func (o *GetStandingOrdersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetStandingOrdersMethodNotAllowed creates a GetStandingOrdersMethodNotAllowed with default headers values
func NewGetStandingOrdersMethodNotAllowed() *GetStandingOrdersMethodNotAllowed {
	return &GetStandingOrdersMethodNotAllowed{}
}

/* GetStandingOrdersMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetStandingOrdersMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetStandingOrdersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /standing-orders][%d] getStandingOrdersMethodNotAllowed ", 405)
}

func (o *GetStandingOrdersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetStandingOrdersNotAcceptable creates a GetStandingOrdersNotAcceptable with default headers values
func NewGetStandingOrdersNotAcceptable() *GetStandingOrdersNotAcceptable {
	return &GetStandingOrdersNotAcceptable{}
}

/* GetStandingOrdersNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetStandingOrdersNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetStandingOrdersNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /standing-orders][%d] getStandingOrdersNotAcceptable ", 406)
}

func (o *GetStandingOrdersNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetStandingOrdersTooManyRequests creates a GetStandingOrdersTooManyRequests with default headers values
func NewGetStandingOrdersTooManyRequests() *GetStandingOrdersTooManyRequests {
	return &GetStandingOrdersTooManyRequests{}
}

/* GetStandingOrdersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetStandingOrdersTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetStandingOrdersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /standing-orders][%d] getStandingOrdersTooManyRequests ", 429)
}

func (o *GetStandingOrdersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetStandingOrdersInternalServerError creates a GetStandingOrdersInternalServerError with default headers values
func NewGetStandingOrdersInternalServerError() *GetStandingOrdersInternalServerError {
	return &GetStandingOrdersInternalServerError{}
}

/* GetStandingOrdersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetStandingOrdersInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetStandingOrdersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /standing-orders][%d] getStandingOrdersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetStandingOrdersInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetStandingOrdersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
