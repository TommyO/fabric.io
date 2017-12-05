// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetTopBuildsReader is a Reader for the GetTopBuilds structure.
type GetTopBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTopBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetTopBuildsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetTopBuildsDefault creates a GetTopBuildsDefault with default headers values
func NewGetTopBuildsDefault(code int) *GetTopBuildsDefault {
	return &GetTopBuildsDefault{
		_statusCode: code,
	}
}

/*GetTopBuildsDefault handles this case with default header values.

stub description for swagger compliance
*/
type GetTopBuildsDefault struct {
	_statusCode int
}

// Code gets the status code for the get top builds default response
func (o *GetTopBuildsDefault) Code() int {
	return o._statusCode
}

func (o *GetTopBuildsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{ResponseBodyPath}/apps/{app_id}/growth_analytics/top_builds][%d] GetTopBuilds default ", o._statusCode)
}

func (o *GetTopBuildsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
