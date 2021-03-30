// Code generated by go-swagger; DO NOT EDIT.

package domestic_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetDomesticPaymentsDomesticPaymentIDParams creates a new GetDomesticPaymentsDomesticPaymentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomesticPaymentsDomesticPaymentIDParams() *GetDomesticPaymentsDomesticPaymentIDParams {
	return &GetDomesticPaymentsDomesticPaymentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomesticPaymentsDomesticPaymentIDParamsWithTimeout creates a new GetDomesticPaymentsDomesticPaymentIDParams object
// with the ability to set a timeout on a request.
func NewGetDomesticPaymentsDomesticPaymentIDParamsWithTimeout(timeout time.Duration) *GetDomesticPaymentsDomesticPaymentIDParams {
	return &GetDomesticPaymentsDomesticPaymentIDParams{
		timeout: timeout,
	}
}

// NewGetDomesticPaymentsDomesticPaymentIDParamsWithContext creates a new GetDomesticPaymentsDomesticPaymentIDParams object
// with the ability to set a context for a request.
func NewGetDomesticPaymentsDomesticPaymentIDParamsWithContext(ctx context.Context) *GetDomesticPaymentsDomesticPaymentIDParams {
	return &GetDomesticPaymentsDomesticPaymentIDParams{
		Context: ctx,
	}
}

// NewGetDomesticPaymentsDomesticPaymentIDParamsWithHTTPClient creates a new GetDomesticPaymentsDomesticPaymentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomesticPaymentsDomesticPaymentIDParamsWithHTTPClient(client *http.Client) *GetDomesticPaymentsDomesticPaymentIDParams {
	return &GetDomesticPaymentsDomesticPaymentIDParams{
		HTTPClient: client,
	}
}

/* GetDomesticPaymentsDomesticPaymentIDParams contains all the parameters to send to the API endpoint
   for the get domestic payments domestic payment Id operation.

   Typically these are written to a http.Request.
*/
type GetDomesticPaymentsDomesticPaymentIDParams struct {

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

	/* DomesticPaymentID.

	   DomesticPaymentId
	*/
	DomesticPaymentID string

	/* XCustomerUserAgent.

	   Indicates the user-agent that the PSU is using.
	*/
	XCustomerUserAgent *string

	/* XFapiAuthDate.

	     The time when the PSU last logged in with the TPP.
	All dates in the HTTP headers are represented as RFC 7231 Full Dates. An example is below:
	Sun, 10 Sep 2017 19:43:31 UTC
	*/
	XFapiAuthDate *string

	/* XFapiCustomerIPAddress.

	   The PSU's IP address if the PSU is currently logged in with the TPP.
	*/
	XFapiCustomerIPAddress *string

	/* XFapiInteractionID.

	   An RFC4122 UID used as a correlation id.
	*/
	XFapiInteractionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get domestic payments domestic payment Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomesticPaymentsDomesticPaymentIDParams) WithDefaults() *GetDomesticPaymentsDomesticPaymentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domestic payments domestic payment Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomesticPaymentsDomesticPaymentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) WithTimeout(timeout time.Duration) *GetDomesticPaymentsDomesticPaymentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) WithContext(ctx context.Context) *GetDomesticPaymentsDomesticPaymentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) WithHTTPClient(client *http.Client) *GetDomesticPaymentsDomesticPaymentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) WithAuthorization(authorization string) *GetDomesticPaymentsDomesticPaymentIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDomesticPaymentID adds the domesticPaymentID to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) WithDomesticPaymentID(domesticPaymentID string) *GetDomesticPaymentsDomesticPaymentIDParams {
	o.SetDomesticPaymentID(domesticPaymentID)
	return o
}

// SetDomesticPaymentID adds the domesticPaymentId to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) SetDomesticPaymentID(domesticPaymentID string) {
	o.DomesticPaymentID = domesticPaymentID
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetDomesticPaymentsDomesticPaymentIDParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetDomesticPaymentsDomesticPaymentIDParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetDomesticPaymentsDomesticPaymentIDParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetDomesticPaymentsDomesticPaymentIDParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get domestic payments domestic payment Id params
func (o *GetDomesticPaymentsDomesticPaymentIDParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomesticPaymentsDomesticPaymentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param DomesticPaymentId
	if err := r.SetPathParam("DomesticPaymentId", o.DomesticPaymentID); err != nil {
		return err
	}

	if o.XCustomerUserAgent != nil {

		// header param x-customer-user-agent
		if err := r.SetHeaderParam("x-customer-user-agent", *o.XCustomerUserAgent); err != nil {
			return err
		}
	}

	if o.XFapiAuthDate != nil {

		// header param x-fapi-auth-date
		if err := r.SetHeaderParam("x-fapi-auth-date", *o.XFapiAuthDate); err != nil {
			return err
		}
	}

	if o.XFapiCustomerIPAddress != nil {

		// header param x-fapi-customer-ip-address
		if err := r.SetHeaderParam("x-fapi-customer-ip-address", *o.XFapiCustomerIPAddress); err != nil {
			return err
		}
	}

	if o.XFapiInteractionID != nil {

		// header param x-fapi-interaction-id
		if err := r.SetHeaderParam("x-fapi-interaction-id", *o.XFapiInteractionID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
