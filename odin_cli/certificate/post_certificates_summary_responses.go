// Code generated by go-swagger; DO NOT EDIT.

package certificate

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

// PostCertificatesSummaryReader is a Reader for the PostCertificatesSummary structure.
type PostCertificatesSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCertificatesSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCertificatesSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400,402,401:
		result := NewPostCertificatesSummaryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostCertificatesSummaryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCertificatesSummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /certificates/summary] PostCertificatesSummary", response, response.Code())
	}
}

// NewPostCertificatesSummaryOK creates a PostCertificatesSummaryOK with default headers values
func NewPostCertificatesSummaryOK() *PostCertificatesSummaryOK {
	return &PostCertificatesSummaryOK{}
}

/*
PostCertificatesSummaryOK describes a response with status code 200, with default header values.

OK
*/
type PostCertificatesSummaryOK struct {
	Payload *PostCertificatesSummaryOKBody
}

// IsSuccess returns true when this post certificates summary o k response has a 2xx status code
func (o *PostCertificatesSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post certificates summary o k response has a 3xx status code
func (o *PostCertificatesSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post certificates summary o k response has a 4xx status code
func (o *PostCertificatesSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post certificates summary o k response has a 5xx status code
func (o *PostCertificatesSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post certificates summary o k response a status code equal to that given
func (o *PostCertificatesSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post certificates summary o k response
func (o *PostCertificatesSummaryOK) Code() int {
	return 200
}

func (o *PostCertificatesSummaryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/summary][%d] postCertificatesSummaryOK %s", 200, payload)
}

func (o *PostCertificatesSummaryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/summary][%d] postCertificatesSummaryOK %s", 200, payload)
}

func (o *PostCertificatesSummaryOK) GetPayload() *PostCertificatesSummaryOKBody {
	return o.Payload
}

func (o *PostCertificatesSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCertificatesSummaryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificatesSummaryBadRequest creates a PostCertificatesSummaryBadRequest with default headers values
func NewPostCertificatesSummaryBadRequest() *PostCertificatesSummaryBadRequest {
	return &PostCertificatesSummaryBadRequest{}
}

/*
PostCertificatesSummaryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostCertificatesSummaryBadRequest struct {
	Payload *models.CertificateErrorResponse
}

// IsSuccess returns true when this post certificates summary bad request response has a 2xx status code
func (o *PostCertificatesSummaryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post certificates summary bad request response has a 3xx status code
func (o *PostCertificatesSummaryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post certificates summary bad request response has a 4xx status code
func (o *PostCertificatesSummaryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post certificates summary bad request response has a 5xx status code
func (o *PostCertificatesSummaryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post certificates summary bad request response a status code equal to that given
func (o *PostCertificatesSummaryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post certificates summary bad request response
func (o *PostCertificatesSummaryBadRequest) Code() int {
	return 400
}

func (o *PostCertificatesSummaryBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/summary][%d] postCertificatesSummaryBadRequest %s", 400, payload)
}

func (o *PostCertificatesSummaryBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/summary][%d] postCertificatesSummaryBadRequest %s", 400, payload)
}

func (o *PostCertificatesSummaryBadRequest) GetPayload() *models.CertificateErrorResponse {
	return o.Payload
}

func (o *PostCertificatesSummaryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificatesSummaryRequestTimeout creates a PostCertificatesSummaryRequestTimeout with default headers values
func NewPostCertificatesSummaryRequestTimeout() *PostCertificatesSummaryRequestTimeout {
	return &PostCertificatesSummaryRequestTimeout{}
}

/*
PostCertificatesSummaryRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PostCertificatesSummaryRequestTimeout struct {
	Payload *models.CertificateErrorResponse
}

// IsSuccess returns true when this post certificates summary request timeout response has a 2xx status code
func (o *PostCertificatesSummaryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post certificates summary request timeout response has a 3xx status code
func (o *PostCertificatesSummaryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post certificates summary request timeout response has a 4xx status code
func (o *PostCertificatesSummaryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post certificates summary request timeout response has a 5xx status code
func (o *PostCertificatesSummaryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post certificates summary request timeout response a status code equal to that given
func (o *PostCertificatesSummaryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the post certificates summary request timeout response
func (o *PostCertificatesSummaryRequestTimeout) Code() int {
	return 408
}

func (o *PostCertificatesSummaryRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/summary][%d] postCertificatesSummaryRequestTimeout %s", 408, payload)
}

func (o *PostCertificatesSummaryRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/summary][%d] postCertificatesSummaryRequestTimeout %s", 408, payload)
}

func (o *PostCertificatesSummaryRequestTimeout) GetPayload() *models.CertificateErrorResponse {
	return o.Payload
}

func (o *PostCertificatesSummaryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificatesSummaryInternalServerError creates a PostCertificatesSummaryInternalServerError with default headers values
func NewPostCertificatesSummaryInternalServerError() *PostCertificatesSummaryInternalServerError {
	return &PostCertificatesSummaryInternalServerError{}
}

/*
PostCertificatesSummaryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostCertificatesSummaryInternalServerError struct {
	Payload *models.CertificateErrorResponse
}

// IsSuccess returns true when this post certificates summary internal server error response has a 2xx status code
func (o *PostCertificatesSummaryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post certificates summary internal server error response has a 3xx status code
func (o *PostCertificatesSummaryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post certificates summary internal server error response has a 4xx status code
func (o *PostCertificatesSummaryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post certificates summary internal server error response has a 5xx status code
func (o *PostCertificatesSummaryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post certificates summary internal server error response a status code equal to that given
func (o *PostCertificatesSummaryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post certificates summary internal server error response
func (o *PostCertificatesSummaryInternalServerError) Code() int {
	return 500
}

func (o *PostCertificatesSummaryInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/summary][%d] postCertificatesSummaryInternalServerError %s", 500, payload)
}

func (o *PostCertificatesSummaryInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/summary][%d] postCertificatesSummaryInternalServerError %s", 500, payload)
}

func (o *PostCertificatesSummaryInternalServerError) GetPayload() *models.CertificateErrorResponse {
	return o.Payload
}

func (o *PostCertificatesSummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostCertificatesSummaryOKBody post certificates summary o k body
swagger:model PostCertificatesSummaryOKBody
*/
type PostCertificatesSummaryOKBody struct {
	models.CertificateAPIResponse

	// data
	Data *models.EshandlerAggregate `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostCertificatesSummaryOKBody) UnmarshalJSON(raw []byte) error {
	// PostCertificatesSummaryOKBodyAO0
	var postCertificatesSummaryOKBodyAO0 models.CertificateAPIResponse
	if err := swag.ReadJSON(raw, &postCertificatesSummaryOKBodyAO0); err != nil {
		return err
	}
	o.CertificateAPIResponse = postCertificatesSummaryOKBodyAO0

	// PostCertificatesSummaryOKBodyAO1
	var dataPostCertificatesSummaryOKBodyAO1 struct {
		Data *models.EshandlerAggregate `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostCertificatesSummaryOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostCertificatesSummaryOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostCertificatesSummaryOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postCertificatesSummaryOKBodyAO0, err := swag.WriteJSON(o.CertificateAPIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postCertificatesSummaryOKBodyAO0)
	var dataPostCertificatesSummaryOKBodyAO1 struct {
		Data *models.EshandlerAggregate `json:"data,omitempty"`
	}

	dataPostCertificatesSummaryOKBodyAO1.Data = o.Data

	jsonDataPostCertificatesSummaryOKBodyAO1, errPostCertificatesSummaryOKBodyAO1 := swag.WriteJSON(dataPostCertificatesSummaryOKBodyAO1)
	if errPostCertificatesSummaryOKBodyAO1 != nil {
		return nil, errPostCertificatesSummaryOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostCertificatesSummaryOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post certificates summary o k body
func (o *PostCertificatesSummaryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.CertificateAPIResponse
	if err := o.CertificateAPIResponse.Validate(formats); err != nil {
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

func (o *PostCertificatesSummaryOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCertificatesSummaryOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCertificatesSummaryOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post certificates summary o k body based on the context it is used
func (o *PostCertificatesSummaryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.CertificateAPIResponse
	if err := o.CertificateAPIResponse.ContextValidate(ctx, formats); err != nil {
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

func (o *PostCertificatesSummaryOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCertificatesSummaryOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCertificatesSummaryOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCertificatesSummaryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCertificatesSummaryOKBody) UnmarshalBinary(b []byte) error {
	var res PostCertificatesSummaryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
