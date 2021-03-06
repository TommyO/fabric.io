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

// NewTimeInAppParams creates a new TimeInAppParams object
// with the default values initialized.
func NewTimeInAppParams() *TimeInAppParams {
	var (
		buildDefault = string("all")
	)
	return &TimeInAppParams{
		Build: &buildDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTimeInAppParamsWithTimeout creates a new TimeInAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimeInAppParamsWithTimeout(timeout time.Duration) *TimeInAppParams {
	var (
		buildDefault = string("all")
	)
	return &TimeInAppParams{
		Build: &buildDefault,

		timeout: timeout,
	}
}

// NewTimeInAppParamsWithContext creates a new TimeInAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimeInAppParamsWithContext(ctx context.Context) *TimeInAppParams {
	var (
		buildDefault = string("all")
	)
	return &TimeInAppParams{
		Build: &buildDefault,

		Context: ctx,
	}
}

// NewTimeInAppParamsWithHTTPClient creates a new TimeInAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimeInAppParamsWithHTTPClient(client *http.Client) *TimeInAppParams {
	var (
		buildDefault = string("all")
	)
	return &TimeInAppParams{
		Build:      &buildDefault,
		HTTPClient: client,
	}
}

/*TimeInAppParams contains all the parameters to send to the API endpoint
for the time in app operation typically these are written to a http.Request
*/
type TimeInAppParams struct {

	/*Authorization*/
	Authorization string
	/*AppID*/
	AppID string
	/*Build*/
	Build *string
	/*End*/
	End *int64
	/*OrganizationID*/
	OrganizationID string
	/*Start*/
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the time in app params
func (o *TimeInAppParams) WithTimeout(timeout time.Duration) *TimeInAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the time in app params
func (o *TimeInAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the time in app params
func (o *TimeInAppParams) WithContext(ctx context.Context) *TimeInAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the time in app params
func (o *TimeInAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the time in app params
func (o *TimeInAppParams) WithHTTPClient(client *http.Client) *TimeInAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the time in app params
func (o *TimeInAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the time in app params
func (o *TimeInAppParams) WithAuthorization(authorization string) *TimeInAppParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the time in app params
func (o *TimeInAppParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAppID adds the appID to the time in app params
func (o *TimeInAppParams) WithAppID(appID string) *TimeInAppParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the time in app params
func (o *TimeInAppParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithBuild adds the build to the time in app params
func (o *TimeInAppParams) WithBuild(build *string) *TimeInAppParams {
	o.SetBuild(build)
	return o
}

// SetBuild adds the build to the time in app params
func (o *TimeInAppParams) SetBuild(build *string) {
	o.Build = build
}

// WithEnd adds the end to the time in app params
func (o *TimeInAppParams) WithEnd(end *int64) *TimeInAppParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the time in app params
func (o *TimeInAppParams) SetEnd(end *int64) {
	o.End = end
}

// WithOrganizationID adds the organizationID to the time in app params
func (o *TimeInAppParams) WithOrganizationID(organizationID string) *TimeInAppParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the time in app params
func (o *TimeInAppParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithStart adds the start to the time in app params
func (o *TimeInAppParams) WithStart(start *int64) *TimeInAppParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the time in app params
func (o *TimeInAppParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *TimeInAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
		var qrEnd int64
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
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
		var qrStart int64
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
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
