// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSingleAppParams creates a new GetSingleAppParams object
// with the default values initialized.
func NewGetSingleAppParams() *GetSingleAppParams {
	var (
		authorizationDefault = string("Bearer {access_token}")
	)
	return &GetSingleAppParams{
		Authorization: &authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSingleAppParamsWithTimeout creates a new GetSingleAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSingleAppParamsWithTimeout(timeout time.Duration) *GetSingleAppParams {
	var (
		authorizationDefault = string("Bearer {access_token}")
	)
	return &GetSingleAppParams{
		Authorization: &authorizationDefault,

		timeout: timeout,
	}
}

// NewGetSingleAppParamsWithContext creates a new GetSingleAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSingleAppParamsWithContext(ctx context.Context) *GetSingleAppParams {
	var (
		authorizationDefault = string("Bearer {access_token}")
	)
	return &GetSingleAppParams{
		Authorization: &authorizationDefault,

		Context: ctx,
	}
}

// NewGetSingleAppParamsWithHTTPClient creates a new GetSingleAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSingleAppParamsWithHTTPClient(client *http.Client) *GetSingleAppParams {
	var (
		authorizationDefault = string("Bearer {access_token}")
	)
	return &GetSingleAppParams{
		Authorization: &authorizationDefault,
		HTTPClient:    client,
	}
}

/*GetSingleAppParams contains all the parameters to send to the API endpoint
for the get single app operation typically these are written to a http.Request
*/
type GetSingleAppParams struct {

	/*Authorization*/
	Authorization *string
	/*AppID*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get single app params
func (o *GetSingleAppParams) WithTimeout(timeout time.Duration) *GetSingleAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get single app params
func (o *GetSingleAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get single app params
func (o *GetSingleAppParams) WithContext(ctx context.Context) *GetSingleAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get single app params
func (o *GetSingleAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get single app params
func (o *GetSingleAppParams) WithHTTPClient(client *http.Client) *GetSingleAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get single app params
func (o *GetSingleAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get single app params
func (o *GetSingleAppParams) WithAuthorization(authorization *string) *GetSingleAppParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get single app params
func (o *GetSingleAppParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithAppID adds the appID to the get single app params
func (o *GetSingleAppParams) WithAppID(appID string) *GetSingleAppParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get single app params
func (o *GetSingleAppParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSingleAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}

	}

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}