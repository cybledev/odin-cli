// Code generated by go-swagger; DO NOT EDIT.

package exposed_files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cybledev/odin-cli/models"
)

// PostExposedFilesSummaryReader is a Reader for the PostExposedFilesSummary structure.
type PostExposedFilesSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExposedFilesSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExposedFilesSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400,402,401:
		result := NewPostExposedFilesSummaryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExposedFilesSummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /exposed/files/summary] PostExposedFilesSummary", response, response.Code())
	}
}

// NewPostExposedFilesSummaryOK creates a PostExposedFilesSummaryOK with default headers values
func NewPostExposedFilesSummaryOK() *PostExposedFilesSummaryOK {
	return &PostExposedFilesSummaryOK{}
}

/*
PostExposedFilesSummaryOK describes a response with status code 200, with default header values.

OK
*/
type PostExposedFilesSummaryOK struct {
	Payload *PostExposedFilesSummaryOKBody
}

// IsSuccess returns true when this post exposed files summary o k response has a 2xx status code
func (o *PostExposedFilesSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post exposed files summary o k response has a 3xx status code
func (o *PostExposedFilesSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post exposed files summary o k response has a 4xx status code
func (o *PostExposedFilesSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post exposed files summary o k response has a 5xx status code
func (o *PostExposedFilesSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post exposed files summary o k response a status code equal to that given
func (o *PostExposedFilesSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post exposed files summary o k response
func (o *PostExposedFilesSummaryOK) Code() int {
	return 200
}

func (o *PostExposedFilesSummaryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /exposed/files/summary][%d] postExposedFilesSummaryOK %s", 200, payload)
}

func (o *PostExposedFilesSummaryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /exposed/files/summary][%d] postExposedFilesSummaryOK %s", 200, payload)
}

func (o *PostExposedFilesSummaryOK) GetPayload() *PostExposedFilesSummaryOKBody {
	return o.Payload
}

func (o *PostExposedFilesSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostExposedFilesSummaryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExposedFilesSummaryBadRequest creates a PostExposedFilesSummaryBadRequest with default headers values
func NewPostExposedFilesSummaryBadRequest() *PostExposedFilesSummaryBadRequest {
	return &PostExposedFilesSummaryBadRequest{}
}

/*
PostExposedFilesSummaryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostExposedFilesSummaryBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this post exposed files summary bad request response has a 2xx status code
func (o *PostExposedFilesSummaryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post exposed files summary bad request response has a 3xx status code
func (o *PostExposedFilesSummaryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post exposed files summary bad request response has a 4xx status code
func (o *PostExposedFilesSummaryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post exposed files summary bad request response has a 5xx status code
func (o *PostExposedFilesSummaryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post exposed files summary bad request response a status code equal to that given
func (o *PostExposedFilesSummaryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post exposed files summary bad request response
func (o *PostExposedFilesSummaryBadRequest) Code() int {
	return 400
}

func (o *PostExposedFilesSummaryBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /exposed/files/summary][%d] postExposedFilesSummaryBadRequest %s", 400, payload)
}

func (o *PostExposedFilesSummaryBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /exposed/files/summary][%d] postExposedFilesSummaryBadRequest %s", 400, payload)
}

func (o *PostExposedFilesSummaryBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PostExposedFilesSummaryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExposedFilesSummaryInternalServerError creates a PostExposedFilesSummaryInternalServerError with default headers values
func NewPostExposedFilesSummaryInternalServerError() *PostExposedFilesSummaryInternalServerError {
	return &PostExposedFilesSummaryInternalServerError{}
}

/*
PostExposedFilesSummaryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostExposedFilesSummaryInternalServerError struct {
	Payload interface{}
}

// IsSuccess returns true when this post exposed files summary internal server error response has a 2xx status code
func (o *PostExposedFilesSummaryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post exposed files summary internal server error response has a 3xx status code
func (o *PostExposedFilesSummaryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post exposed files summary internal server error response has a 4xx status code
func (o *PostExposedFilesSummaryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post exposed files summary internal server error response has a 5xx status code
func (o *PostExposedFilesSummaryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post exposed files summary internal server error response a status code equal to that given
func (o *PostExposedFilesSummaryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post exposed files summary internal server error response
func (o *PostExposedFilesSummaryInternalServerError) Code() int {
	return 500
}

func (o *PostExposedFilesSummaryInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /exposed/files/summary][%d] postExposedFilesSummaryInternalServerError %s", 500, payload)
}

func (o *PostExposedFilesSummaryInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /exposed/files/summary][%d] postExposedFilesSummaryInternalServerError %s", 500, payload)
}

func (o *PostExposedFilesSummaryInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *PostExposedFilesSummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostExposedFilesSummaryOKBody post exposed files summary o k body
swagger:model PostExposedFilesSummaryOKBody
*/
type PostExposedFilesSummaryOKBody struct {
	models.ExposedAPIResponse

	// data
	Data *models.ExposedAggregate `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostExposedFilesSummaryOKBody) UnmarshalJSON(raw []byte) error {
	// PostExposedFilesSummaryOKBodyAO0
	var postExposedFilesSummaryOKBodyAO0 models.ExposedAPIResponse
	if err := swag.ReadJSON(raw, &postExposedFilesSummaryOKBodyAO0); err != nil {
		return err
	}
	o.ExposedAPIResponse = postExposedFilesSummaryOKBodyAO0

	// PostExposedFilesSummaryOKBodyAO1
	var dataPostExposedFilesSummaryOKBodyAO1 struct {
		Data *models.ExposedAggregate `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostExposedFilesSummaryOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostExposedFilesSummaryOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostExposedFilesSummaryOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postExposedFilesSummaryOKBodyAO0, err := swag.WriteJSON(o.ExposedAPIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postExposedFilesSummaryOKBodyAO0)
	var dataPostExposedFilesSummaryOKBodyAO1 struct {
		Data *models.ExposedAggregate `json:"data,omitempty"`
	}

	dataPostExposedFilesSummaryOKBodyAO1.Data = o.Data

	jsonDataPostExposedFilesSummaryOKBodyAO1, errPostExposedFilesSummaryOKBodyAO1 := swag.WriteJSON(dataPostExposedFilesSummaryOKBodyAO1)
	if errPostExposedFilesSummaryOKBodyAO1 != nil {
		return nil, errPostExposedFilesSummaryOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostExposedFilesSummaryOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post exposed files summary o k body
func (o *PostExposedFilesSummaryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ExposedAPIResponse
	if err := o.ExposedAPIResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostExposedFilesSummaryOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postExposedFilesSummaryOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postExposedFilesSummaryOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post exposed files summary o k body based on the context it is used
func (o *PostExposedFilesSummaryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ExposedAPIResponse
	if err := o.ExposedAPIResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostExposedFilesSummaryOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postExposedFilesSummaryOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postExposedFilesSummaryOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostExposedFilesSummaryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostExposedFilesSummaryOKBody) UnmarshalBinary(b []byte) error {
	var res PostExposedFilesSummaryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
