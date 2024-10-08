// Code generated by go-swagger; DO NOT EDIT.

package hosts

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

// PostHostsCountReader is a Reader for the PostHostsCount structure.
type PostHostsCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHostsCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostHostsCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400,402,401:
		result := NewPostHostsCountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostHostsCountRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostHostsCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /hosts/count] PostHostsCount", response, response.Code())
	}
}

// NewPostHostsCountOK creates a PostHostsCountOK with default headers values
func NewPostHostsCountOK() *PostHostsCountOK {
	return &PostHostsCountOK{}
}

/*
PostHostsCountOK describes a response with status code 200, with default header values.

OK
*/
type PostHostsCountOK struct {
	Payload *PostHostsCountOKBody
}

// IsSuccess returns true when this post hosts count o k response has a 2xx status code
func (o *PostHostsCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post hosts count o k response has a 3xx status code
func (o *PostHostsCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post hosts count o k response has a 4xx status code
func (o *PostHostsCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post hosts count o k response has a 5xx status code
func (o *PostHostsCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post hosts count o k response a status code equal to that given
func (o *PostHostsCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post hosts count o k response
func (o *PostHostsCountOK) Code() int {
	return 200
}

func (o *PostHostsCountOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /hosts/count][%d] postHostsCountOK %s", 200, payload)
}

func (o *PostHostsCountOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /hosts/count][%d] postHostsCountOK %s", 200, payload)
}

func (o *PostHostsCountOK) GetPayload() *PostHostsCountOKBody {
	return o.Payload
}

func (o *PostHostsCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostHostsCountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostHostsCountBadRequest creates a PostHostsCountBadRequest with default headers values
func NewPostHostsCountBadRequest() *PostHostsCountBadRequest {
	return &PostHostsCountBadRequest{}
}

/*
PostHostsCountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostHostsCountBadRequest struct {
	Payload *models.IpservicesErrorResponse
}

// IsSuccess returns true when this post hosts count bad request response has a 2xx status code
func (o *PostHostsCountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post hosts count bad request response has a 3xx status code
func (o *PostHostsCountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post hosts count bad request response has a 4xx status code
func (o *PostHostsCountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post hosts count bad request response has a 5xx status code
func (o *PostHostsCountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post hosts count bad request response a status code equal to that given
func (o *PostHostsCountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post hosts count bad request response
func (o *PostHostsCountBadRequest) Code() int {
	return 400
}

func (o *PostHostsCountBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /hosts/count][%d] postHostsCountBadRequest %s", 400, payload)
}

func (o *PostHostsCountBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /hosts/count][%d] postHostsCountBadRequest %s", 400, payload)
}

func (o *PostHostsCountBadRequest) GetPayload() *models.IpservicesErrorResponse {
	return o.Payload
}

func (o *PostHostsCountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpservicesErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostHostsCountRequestTimeout creates a PostHostsCountRequestTimeout with default headers values
func NewPostHostsCountRequestTimeout() *PostHostsCountRequestTimeout {
	return &PostHostsCountRequestTimeout{}
}

/*
PostHostsCountRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PostHostsCountRequestTimeout struct {
	Payload *models.IpservicesErrorResponse
}

// IsSuccess returns true when this post hosts count request timeout response has a 2xx status code
func (o *PostHostsCountRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post hosts count request timeout response has a 3xx status code
func (o *PostHostsCountRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post hosts count request timeout response has a 4xx status code
func (o *PostHostsCountRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post hosts count request timeout response has a 5xx status code
func (o *PostHostsCountRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post hosts count request timeout response a status code equal to that given
func (o *PostHostsCountRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the post hosts count request timeout response
func (o *PostHostsCountRequestTimeout) Code() int {
	return 408
}

func (o *PostHostsCountRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /hosts/count][%d] postHostsCountRequestTimeout %s", 408, payload)
}

func (o *PostHostsCountRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /hosts/count][%d] postHostsCountRequestTimeout %s", 408, payload)
}

func (o *PostHostsCountRequestTimeout) GetPayload() *models.IpservicesErrorResponse {
	return o.Payload
}

func (o *PostHostsCountRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpservicesErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostHostsCountInternalServerError creates a PostHostsCountInternalServerError with default headers values
func NewPostHostsCountInternalServerError() *PostHostsCountInternalServerError {
	return &PostHostsCountInternalServerError{}
}

/*
PostHostsCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostHostsCountInternalServerError struct {
	Payload *models.IpservicesErrorResponse
}

// IsSuccess returns true when this post hosts count internal server error response has a 2xx status code
func (o *PostHostsCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post hosts count internal server error response has a 3xx status code
func (o *PostHostsCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post hosts count internal server error response has a 4xx status code
func (o *PostHostsCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post hosts count internal server error response has a 5xx status code
func (o *PostHostsCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post hosts count internal server error response a status code equal to that given
func (o *PostHostsCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post hosts count internal server error response
func (o *PostHostsCountInternalServerError) Code() int {
	return 500
}

func (o *PostHostsCountInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /hosts/count][%d] postHostsCountInternalServerError %s", 500, payload)
}

func (o *PostHostsCountInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /hosts/count][%d] postHostsCountInternalServerError %s", 500, payload)
}

func (o *PostHostsCountInternalServerError) GetPayload() *models.IpservicesErrorResponse {
	return o.Payload
}

func (o *PostHostsCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpservicesErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostHostsCountOKBody post hosts count o k body
swagger:model PostHostsCountOKBody
*/
type PostHostsCountOKBody struct {
	models.IpservicesAPIResponse

	// data
	Data *models.IpservicesCertCount `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostHostsCountOKBody) UnmarshalJSON(raw []byte) error {
	// PostHostsCountOKBodyAO0
	var postHostsCountOKBodyAO0 models.IpservicesAPIResponse
	if err := swag.ReadJSON(raw, &postHostsCountOKBodyAO0); err != nil {
		return err
	}
	o.IpservicesAPIResponse = postHostsCountOKBodyAO0

	// PostHostsCountOKBodyAO1
	var dataPostHostsCountOKBodyAO1 struct {
		Data *models.IpservicesCertCount `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostHostsCountOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostHostsCountOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostHostsCountOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postHostsCountOKBodyAO0, err := swag.WriteJSON(o.IpservicesAPIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postHostsCountOKBodyAO0)
	var dataPostHostsCountOKBodyAO1 struct {
		Data *models.IpservicesCertCount `json:"data,omitempty"`
	}

	dataPostHostsCountOKBodyAO1.Data = o.Data

	jsonDataPostHostsCountOKBodyAO1, errPostHostsCountOKBodyAO1 := swag.WriteJSON(dataPostHostsCountOKBodyAO1)
	if errPostHostsCountOKBodyAO1 != nil {
		return nil, errPostHostsCountOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostHostsCountOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post hosts count o k body
func (o *PostHostsCountOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.IpservicesAPIResponse
	if err := o.IpservicesAPIResponse.Validate(formats); err != nil {
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

func (o *PostHostsCountOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postHostsCountOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postHostsCountOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post hosts count o k body based on the context it is used
func (o *PostHostsCountOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.IpservicesAPIResponse
	if err := o.IpservicesAPIResponse.ContextValidate(ctx, formats); err != nil {
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

func (o *PostHostsCountOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postHostsCountOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postHostsCountOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostHostsCountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostHostsCountOKBody) UnmarshalBinary(b []byte) error {
	var res PostHostsCountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
