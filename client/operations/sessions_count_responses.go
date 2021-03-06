// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/TommyO/fabric.io/models"
)

// SessionsCountReader is a Reader for the SessionsCount structure.
type SessionsCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SessionsCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSessionsCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSessionsCountOK creates a SessionsCountOK with default headers values
func NewSessionsCountOK() *SessionsCountOK {
	return &SessionsCountOK{}
}

/*SessionsCountOK handles this case with default header values.

Report
*/
type SessionsCountOK struct {
	Payload *models.SessionsCountOKBody
}

func (o *SessionsCountOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{organization_id}/apps/{app_id}/growth_analytics/total_sessions_scalar.json][%d] sessionsCountOK  %+v", 200, o.Payload)
}

func (o *SessionsCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionsCountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
