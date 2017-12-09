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

// ActiveNowReader is a Reader for the ActiveNow structure.
type ActiveNowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActiveNowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewActiveNowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActiveNowOK creates a ActiveNowOK with default headers values
func NewActiveNowOK() *ActiveNowOK {
	return &ActiveNowOK{}
}

/*ActiveNowOK handles this case with default header values.

Report
*/
type ActiveNowOK struct {
	Payload *models.ActiveNowOKBody
}

func (o *ActiveNowOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{organization_id}/apps/{app_id}/growth_analytics/active_now.json][%d] activeNowOK  %+v", 200, o.Payload)
}

func (o *ActiveNowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActiveNowOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
