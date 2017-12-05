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

// NewOAuthParams creates a new OAuthParams object
// with the default values initialized.
func NewOAuthParams() *OAuthParams {
	var ()
	return &OAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOAuthParamsWithTimeout creates a new OAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOAuthParamsWithTimeout(timeout time.Duration) *OAuthParams {
	var ()
	return &OAuthParams{

		timeout: timeout,
	}
}

// NewOAuthParamsWithContext creates a new OAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewOAuthParamsWithContext(ctx context.Context) *OAuthParams {
	var ()
	return &OAuthParams{

		Context: ctx,
	}
}

// NewOAuthParamsWithHTTPClient creates a new OAuthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOAuthParamsWithHTTPClient(client *http.Client) *OAuthParams {
	var ()
	return &OAuthParams{
		HTTPClient: client,
	}
}

/*OAuthParams contains all the parameters to send to the API endpoint
for the o auth operation typically these are written to a http.Request
*/
type OAuthParams struct {

	/*Body*/
	Body *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the o auth params
func (o *OAuthParams) WithTimeout(timeout time.Duration) *OAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the o auth params
func (o *OAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the o auth params
func (o *OAuthParams) WithContext(ctx context.Context) *OAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the o auth params
func (o *OAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the o auth params
func (o *OAuthParams) WithHTTPClient(client *http.Client) *OAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the o auth params
func (o *OAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the o auth params
func (o *OAuthParams) WithBody(body *string) *OAuthParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the o auth params
func (o *OAuthParams) SetBody(body *string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *OAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
