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

// PostCertificatesScrollReader is a Reader for the PostCertificatesScroll structure.
type PostCertificatesScrollReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCertificatesScrollReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCertificatesScrollOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400,402,401:
		result := NewPostCertificatesScrollBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostCertificatesScrollRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCertificatesScrollInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /certificates/scroll] PostCertificatesScroll", response, response.Code())
	}
}

// NewPostCertificatesScrollOK creates a PostCertificatesScrollOK with default headers values
func NewPostCertificatesScrollOK() *PostCertificatesScrollOK {
	return &PostCertificatesScrollOK{}
}

/*
PostCertificatesScrollOK describes a response with status code 200, with default header values.

OK
*/
type PostCertificatesScrollOK struct {
	Payload *PostCertificatesScrollOKBody
}

// IsSuccess returns true when this post certificates scroll o k response has a 2xx status code
func (o *PostCertificatesScrollOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post certificates scroll o k response has a 3xx status code
func (o *PostCertificatesScrollOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post certificates scroll o k response has a 4xx status code
func (o *PostCertificatesScrollOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post certificates scroll o k response has a 5xx status code
func (o *PostCertificatesScrollOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post certificates scroll o k response a status code equal to that given
func (o *PostCertificatesScrollOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post certificates scroll o k response
func (o *PostCertificatesScrollOK) Code() int {
	return 200
}

func (o *PostCertificatesScrollOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/scroll][%d] postCertificatesScrollOK %s", 200, payload)
}

func (o *PostCertificatesScrollOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/scroll][%d] postCertificatesScrollOK %s", 200, payload)
}

func (o *PostCertificatesScrollOK) GetPayload() *PostCertificatesScrollOKBody {
	return o.Payload
}

func (o *PostCertificatesScrollOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCertificatesScrollOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificatesScrollBadRequest creates a PostCertificatesScrollBadRequest with default headers values
func NewPostCertificatesScrollBadRequest() *PostCertificatesScrollBadRequest {
	return &PostCertificatesScrollBadRequest{}
}

/*
PostCertificatesScrollBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostCertificatesScrollBadRequest struct {
	Payload *models.CertificateErrorResponse
}

// IsSuccess returns true when this post certificates scroll bad request response has a 2xx status code
func (o *PostCertificatesScrollBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post certificates scroll bad request response has a 3xx status code
func (o *PostCertificatesScrollBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post certificates scroll bad request response has a 4xx status code
func (o *PostCertificatesScrollBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post certificates scroll bad request response has a 5xx status code
func (o *PostCertificatesScrollBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post certificates scroll bad request response a status code equal to that given
func (o *PostCertificatesScrollBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post certificates scroll bad request response
func (o *PostCertificatesScrollBadRequest) Code() int {
	return 400
}

func (o *PostCertificatesScrollBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/scroll][%d] postCertificatesScrollBadRequest %s", 400, payload)
}

func (o *PostCertificatesScrollBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/scroll][%d] postCertificatesScrollBadRequest %s", 400, payload)
}

func (o *PostCertificatesScrollBadRequest) GetPayload() *models.CertificateErrorResponse {
	return o.Payload
}

func (o *PostCertificatesScrollBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificatesScrollRequestTimeout creates a PostCertificatesScrollRequestTimeout with default headers values
func NewPostCertificatesScrollRequestTimeout() *PostCertificatesScrollRequestTimeout {
	return &PostCertificatesScrollRequestTimeout{}
}

/*
PostCertificatesScrollRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PostCertificatesScrollRequestTimeout struct {
	Payload *models.CertificateErrorResponse
}

// IsSuccess returns true when this post certificates scroll request timeout response has a 2xx status code
func (o *PostCertificatesScrollRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post certificates scroll request timeout response has a 3xx status code
func (o *PostCertificatesScrollRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post certificates scroll request timeout response has a 4xx status code
func (o *PostCertificatesScrollRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post certificates scroll request timeout response has a 5xx status code
func (o *PostCertificatesScrollRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post certificates scroll request timeout response a status code equal to that given
func (o *PostCertificatesScrollRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the post certificates scroll request timeout response
func (o *PostCertificatesScrollRequestTimeout) Code() int {
	return 408
}

func (o *PostCertificatesScrollRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/scroll][%d] postCertificatesScrollRequestTimeout %s", 408, payload)
}

func (o *PostCertificatesScrollRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/scroll][%d] postCertificatesScrollRequestTimeout %s", 408, payload)
}

func (o *PostCertificatesScrollRequestTimeout) GetPayload() *models.CertificateErrorResponse {
	return o.Payload
}

func (o *PostCertificatesScrollRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificatesScrollInternalServerError creates a PostCertificatesScrollInternalServerError with default headers values
func NewPostCertificatesScrollInternalServerError() *PostCertificatesScrollInternalServerError {
	return &PostCertificatesScrollInternalServerError{}
}

/*
PostCertificatesScrollInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostCertificatesScrollInternalServerError struct {
	Payload *models.CertificateErrorResponse
}

// IsSuccess returns true when this post certificates scroll internal server error response has a 2xx status code
func (o *PostCertificatesScrollInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post certificates scroll internal server error response has a 3xx status code
func (o *PostCertificatesScrollInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post certificates scroll internal server error response has a 4xx status code
func (o *PostCertificatesScrollInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post certificates scroll internal server error response has a 5xx status code
func (o *PostCertificatesScrollInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post certificates scroll internal server error response a status code equal to that given
func (o *PostCertificatesScrollInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post certificates scroll internal server error response
func (o *PostCertificatesScrollInternalServerError) Code() int {
	return 500
}

func (o *PostCertificatesScrollInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/scroll][%d] postCertificatesScrollInternalServerError %s", 500, payload)
}

func (o *PostCertificatesScrollInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /certificates/scroll][%d] postCertificatesScrollInternalServerError %s", 500, payload)
}

func (o *PostCertificatesScrollInternalServerError) GetPayload() *models.CertificateErrorResponse {
	return o.Payload
}

func (o *PostCertificatesScrollInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostCertificatesScrollOKBody post certificates scroll o k body
swagger:model PostCertificatesScrollOKBody
*/
type PostCertificatesScrollOKBody struct {
	models.CertificateAPIResponse

	// data
	Data *models.CertificateCertScroll `json:"data,omitempty"`

	// pagination
	Pagination *models.CertificateSearchPagination `json:"pagination,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostCertificatesScrollOKBody) UnmarshalJSON(raw []byte) error {
	// PostCertificatesScrollOKBodyAO0
	var postCertificatesScrollOKBodyAO0 models.CertificateAPIResponse
	if err := swag.ReadJSON(raw, &postCertificatesScrollOKBodyAO0); err != nil {
		return err
	}
	o.CertificateAPIResponse = postCertificatesScrollOKBodyAO0

	// PostCertificatesScrollOKBodyAO1
	var dataPostCertificatesScrollOKBodyAO1 struct {
		Data *models.CertificateCertScroll `json:"data,omitempty"`

		Pagination *models.CertificateSearchPagination `json:"pagination,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostCertificatesScrollOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostCertificatesScrollOKBodyAO1.Data

	o.Pagination = dataPostCertificatesScrollOKBodyAO1.Pagination

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostCertificatesScrollOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postCertificatesScrollOKBodyAO0, err := swag.WriteJSON(o.CertificateAPIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postCertificatesScrollOKBodyAO0)
	var dataPostCertificatesScrollOKBodyAO1 struct {
		Data *models.CertificateCertScroll `json:"data,omitempty"`

		Pagination *models.CertificateSearchPagination `json:"pagination,omitempty"`
	}

	dataPostCertificatesScrollOKBodyAO1.Data = o.Data

	dataPostCertificatesScrollOKBodyAO1.Pagination = o.Pagination

	jsonDataPostCertificatesScrollOKBodyAO1, errPostCertificatesScrollOKBodyAO1 := swag.WriteJSON(dataPostCertificatesScrollOKBodyAO1)
	if errPostCertificatesScrollOKBodyAO1 != nil {
		return nil, errPostCertificatesScrollOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostCertificatesScrollOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post certificates scroll o k body
func (o *PostCertificatesScrollOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.CertificateAPIResponse
	if err := o.CertificateAPIResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCertificatesScrollOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCertificatesScrollOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCertificatesScrollOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

func (o *PostCertificatesScrollOKBody) validatePagination(formats strfmt.Registry) error {

	if swag.IsZero(o.Pagination) { // not required
		return nil
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCertificatesScrollOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCertificatesScrollOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post certificates scroll o k body based on the context it is used
func (o *PostCertificatesScrollOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.CertificateAPIResponse
	if err := o.CertificateAPIResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCertificatesScrollOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCertificatesScrollOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCertificatesScrollOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

func (o *PostCertificatesScrollOKBody) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if o.Pagination != nil {

		if swag.IsZero(o.Pagination) { // not required
			return nil
		}

		if err := o.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCertificatesScrollOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCertificatesScrollOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCertificatesScrollOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCertificatesScrollOKBody) UnmarshalBinary(b []byte) error {
	var res PostCertificatesScrollOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
