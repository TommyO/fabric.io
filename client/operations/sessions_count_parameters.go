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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSessionsCountParams creates a new SessionsCountParams object
// with the default values initialized.
func NewSessionsCountParams() *SessionsCountParams {
	var (
		buildDefault = string("all")
	)
	return &SessionsCountParams{
		Build: &buildDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSessionsCountParamsWithTimeout creates a new SessionsCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSessionsCountParamsWithTimeout(timeout time.Duration) *SessionsCountParams {
	var (
		buildDefault = string("all")
	)
	return &SessionsCountParams{
		Build: &buildDefault,

		timeout: timeout,
	}
}

// NewSessionsCountParamsWithContext creates a new SessionsCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewSessionsCountParamsWithContext(ctx context.Context) *SessionsCountParams {
	var (
		buildDefault = string("all")
	)
	return &SessionsCountParams{
		Build: &buildDefault,

		Context: ctx,
	}
}

// NewSessionsCountParamsWithHTTPClient creates a new SessionsCountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSessionsCountParamsWithHTTPClient(client *http.Client) *SessionsCountParams {
	var (
		buildDefault = string("all")
	)
	return &SessionsCountParams{
		Build:      &buildDefault,
		HTTPClient: client,
	}
}

/*SessionsCountParams contains all the parameters to send to the API endpoint
for the sessions count operation typically these are written to a http.Request
*/
type SessionsCountParams struct {

	/*Authorization*/
	Authorization string
	/*AppID*/
	AppID string
	/*Build*/
	Build *string
	/*End*/
	End *float64
	/*OrganizationID*/
	OrganizationID string
	/*Start*/
	Start *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sessions count params
func (o *SessionsCountParams) WithTimeout(timeout time.Duration) *SessionsCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sessions count params
func (o *SessionsCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sessions count params
func (o *SessionsCountParams) WithContext(ctx context.Context) *SessionsCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sessions count params
func (o *SessionsCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sessions count params
func (o *SessionsCountParams) WithHTTPClient(client *http.Client) *SessionsCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sessions count params
func (o *SessionsCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the sessions count params
func (o *SessionsCountParams) WithAuthorization(authorization string) *SessionsCountParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the sessions count params
func (o *SessionsCountParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAppID adds the appID to the sessions count params
func (o *SessionsCountParams) WithAppID(appID string) *SessionsCountParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the sessions count params
func (o *SessionsCountParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithBuild adds the build to the sessions count params
func (o *SessionsCountParams) WithBuild(build *string) *SessionsCountParams {
	o.SetBuild(build)
	return o
}

// SetBuild adds the build to the sessions count params
func (o *SessionsCountParams) SetBuild(build *string) {
	o.Build = build
}

// WithEnd adds the end to the sessions count params
func (o *SessionsCountParams) WithEnd(end *float64) *SessionsCountParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the sessions count params
func (o *SessionsCountParams) SetEnd(end *float64) {
	o.End = end
}

// WithOrganizationID adds the organizationID to the sessions count params
func (o *SessionsCountParams) WithOrganizationID(organizationID string) *SessionsCountParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the sessions count params
func (o *SessionsCountParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithStart adds the start to the sessions count params
func (o *SessionsCountParams) WithStart(start *float64) *SessionsCountParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the sessions count params
func (o *SessionsCountParams) SetStart(start *float64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *SessionsCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
		return err
	}

	if o.Build != nil {

		// query param build
		var qrBuild string
		if o.Build != nil {
			qrBuild = *o.Build
		}
		qBuild := qrBuild
		if qBuild != "" {
			if err := r.SetQueryParam("build", qBuild); err != nil {
				return err
			}
		}

	}

	if o.End != nil {

		// query param end
		var qrEnd float64
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatFloat64(qrEnd)
		if qEnd != "" {
			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}

	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart float64
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatFloat64(qrStart)
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}