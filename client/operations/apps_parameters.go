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

// NewAppsParams creates a new AppsParams object
// with the default values initialized.
func NewAppsParams() *AppsParams {
	var ()
	return &AppsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppsParamsWithTimeout creates a new AppsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppsParamsWithTimeout(timeout time.Duration) *AppsParams {
	var ()
	return &AppsParams{

		timeout: timeout,
	}
}

// NewAppsParamsWithContext creates a new AppsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppsParamsWithContext(ctx context.Context) *AppsParams {
	var ()
	return &AppsParams{

		Context: ctx,
	}
}

// NewAppsParamsWithHTTPClient creates a new AppsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppsParamsWithHTTPClient(client *http.Client) *AppsParams {
	var ()
	return &AppsParams{
		HTTPClient: client,
	}
}

/*AppsParams contains all the parameters to send to the API endpoint
for the apps operation typically these are written to a http.Request
*/
type AppsParams struct {

	/*Authorization*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the apps params
func (o *AppsParams) WithTimeout(timeout time.Duration) *AppsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the apps params
func (o *AppsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the apps params
func (o *AppsParams) WithContext(ctx context.Context) *AppsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the apps params
func (o *AppsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the apps params
func (o *AppsParams) WithHTTPClient(client *http.Client) *AppsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the apps params
func (o *AppsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the apps params
func (o *AppsParams) WithAuthorization(authorization string) *AppsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the apps params
func (o *AppsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *AppsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
