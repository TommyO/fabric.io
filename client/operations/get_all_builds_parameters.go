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

// NewGetAllBuildsParams creates a new GetAllBuildsParams object
// with the default values initialized.
func NewGetAllBuildsParams() *GetAllBuildsParams {
	var (
		authorizationDefault    = string("Bearer {access_token}")
		responseBodyPathDefault = string("{organization_id}")
	)
	return &GetAllBuildsParams{
		Authorization:    &authorizationDefault,
		ResponseBodyPath: responseBodyPathDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllBuildsParamsWithTimeout creates a new GetAllBuildsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllBuildsParamsWithTimeout(timeout time.Duration) *GetAllBuildsParams {
	var (
		authorizationDefault    = string("Bearer {access_token}")
		responseBodyPathDefault = string("{organization_id}")
	)
	return &GetAllBuildsParams{
		Authorization:    &authorizationDefault,
		ResponseBodyPath: responseBodyPathDefault,

		timeout: timeout,
	}
}

// NewGetAllBuildsParamsWithContext creates a new GetAllBuildsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllBuildsParamsWithContext(ctx context.Context) *GetAllBuildsParams {
	var (
		authorizationDefault    = string("Bearer {access_token}")
		responseBodyPathDefault = string("{organization_id}")
	)
	return &GetAllBuildsParams{
		Authorization:    &authorizationDefault,
		ResponseBodyPath: responseBodyPathDefault,

		Context: ctx,
	}
}

// NewGetAllBuildsParamsWithHTTPClient creates a new GetAllBuildsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllBuildsParamsWithHTTPClient(client *http.Client) *GetAllBuildsParams {
	var (
		authorizationDefault    = string("Bearer {access_token}")
		responseBodyPathDefault = string("{organization_id}")
	)
	return &GetAllBuildsParams{
		Authorization:    &authorizationDefault,
		ResponseBodyPath: responseBodyPathDefault,
		HTTPClient:       client,
	}
}

/*GetAllBuildsParams contains all the parameters to send to the API endpoint
for the get all builds operation typically these are written to a http.Request
*/
type GetAllBuildsParams struct {

	/*Authorization*/
	Authorization *string
	/*ResponseBodyPath*/
	ResponseBodyPath string
	/*AppID*/
	AppID string
	/*ReleaseID*/
	ReleaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all builds params
func (o *GetAllBuildsParams) WithTimeout(timeout time.Duration) *GetAllBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all builds params
func (o *GetAllBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all builds params
func (o *GetAllBuildsParams) WithContext(ctx context.Context) *GetAllBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all builds params
func (o *GetAllBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all builds params
func (o *GetAllBuildsParams) WithHTTPClient(client *http.Client) *GetAllBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all builds params
func (o *GetAllBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get all builds params
func (o *GetAllBuildsParams) WithAuthorization(authorization *string) *GetAllBuildsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get all builds params
func (o *GetAllBuildsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithResponseBodyPath adds the responseBodyPath to the get all builds params
func (o *GetAllBuildsParams) WithResponseBodyPath(responseBodyPath string) *GetAllBuildsParams {
	o.SetResponseBodyPath(responseBodyPath)
	return o
}

// SetResponseBodyPath adds the responseBodyPath to the get all builds params
func (o *GetAllBuildsParams) SetResponseBodyPath(responseBodyPath string) {
	o.ResponseBodyPath = responseBodyPath
}

// WithAppID adds the appID to the get all builds params
func (o *GetAllBuildsParams) WithAppID(appID string) *GetAllBuildsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get all builds params
func (o *GetAllBuildsParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithReleaseID adds the releaseID to the get all builds params
func (o *GetAllBuildsParams) WithReleaseID(releaseID string) *GetAllBuildsParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the get all builds params
func (o *GetAllBuildsParams) SetReleaseID(releaseID string) {
	o.ReleaseID = releaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param ResponseBodyPath
	if err := r.SetPathParam("ResponseBodyPath", o.ResponseBodyPath); err != nil {
		return err
	}

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
		return err
	}

	// path param release_id
	if err := r.SetPathParam("release_id", o.ReleaseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
