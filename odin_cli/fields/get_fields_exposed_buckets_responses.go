// Code generated by go-swagger; DO NOT EDIT.

package fields

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cybledev/odin-cli/models"
)

// GetFieldsExposedBucketsReader is a Reader for the GetFieldsExposedBuckets structure.
type GetFieldsExposedBucketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFieldsExposedBucketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFieldsExposedBucketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400,401:
		result := NewGetFieldsExposedBucketsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewGetFieldsExposedBucketsPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFieldsExposedBucketsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFieldsExposedBucketsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /fields/exposed/buckets/] GetFieldsExposedBuckets", response, response.Code())
	}
}

// NewGetFieldsExposedBucketsOK creates a GetFieldsExposedBucketsOK with default headers values
func NewGetFieldsExposedBucketsOK() *GetFieldsExposedBucketsOK {
	return &GetFieldsExposedBucketsOK{}
}

/*
GetFieldsExposedBucketsOK describes a response with status code 200, with default header values.

OK
*/
type GetFieldsExposedBucketsOK struct {
	Payload *GetFieldsExposedBucketsOKBody
}

// IsSuccess returns true when this get fields exposed buckets o k response has a 2xx status code
func (o *GetFieldsExposedBucketsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get fields exposed buckets o k response has a 3xx status code
func (o *GetFieldsExposedBucketsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fields exposed buckets o k response has a 4xx status code
func (o *GetFieldsExposedBucketsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fields exposed buckets o k response has a 5xx status code
func (o *GetFieldsExposedBucketsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get fields exposed buckets o k response a status code equal to that given
func (o *GetFieldsExposedBucketsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get fields exposed buckets o k response
func (o *GetFieldsExposedBucketsOK) Code() int {
	return 200
}

func (o *GetFieldsExposedBucketsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/buckets/][%d] getFieldsExposedBucketsOK %s", 200, payload)
}

func (o *GetFieldsExposedBucketsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/buckets/][%d] getFieldsExposedBucketsOK %s", 200, payload)
}

func (o *GetFieldsExposedBucketsOK) GetPayload() *GetFieldsExposedBucketsOKBody {
	return o.Payload
}

func (o *GetFieldsExposedBucketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetFieldsExposedBucketsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFieldsExposedBucketsBadRequest creates a GetFieldsExposedBucketsBadRequest with default headers values
func NewGetFieldsExposedBucketsBadRequest() *GetFieldsExposedBucketsBadRequest {
	return &GetFieldsExposedBucketsBadRequest{}
}

/*
GetFieldsExposedBucketsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetFieldsExposedBucketsBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get fields exposed buckets bad request response has a 2xx status code
func (o *GetFieldsExposedBucketsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fields exposed buckets bad request response has a 3xx status code
func (o *GetFieldsExposedBucketsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fields exposed buckets bad request response has a 4xx status code
func (o *GetFieldsExposedBucketsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fields exposed buckets bad request response has a 5xx status code
func (o *GetFieldsExposedBucketsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get fields exposed buckets bad request response a status code equal to that given
func (o *GetFieldsExposedBucketsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get fields exposed buckets bad request response
func (o *GetFieldsExposedBucketsBadRequest) Code() int {
	return 400
}

func (o *GetFieldsExposedBucketsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/buckets/][%d] getFieldsExposedBucketsBadRequest %s", 400, payload)
}

func (o *GetFieldsExposedBucketsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/buckets/][%d] getFieldsExposedBucketsBadRequest %s", 400, payload)
}

func (o *GetFieldsExposedBucketsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFieldsExposedBucketsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFieldsExposedBucketsPaymentRequired creates a GetFieldsExposedBucketsPaymentRequired with default headers values
func NewGetFieldsExposedBucketsPaymentRequired() *GetFieldsExposedBucketsPaymentRequired {
	return &GetFieldsExposedBucketsPaymentRequired{}
}

/*
GetFieldsExposedBucketsPaymentRequired describes a response with status code 402, with default header values.

Payment Required
*/
type GetFieldsExposedBucketsPaymentRequired struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get fields exposed buckets payment required response has a 2xx status code
func (o *GetFieldsExposedBucketsPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fields exposed buckets payment required response has a 3xx status code
func (o *GetFieldsExposedBucketsPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fields exposed buckets payment required response has a 4xx status code
func (o *GetFieldsExposedBucketsPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fields exposed buckets payment required response has a 5xx status code
func (o *GetFieldsExposedBucketsPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this get fields exposed buckets payment required response a status code equal to that given
func (o *GetFieldsExposedBucketsPaymentRequired) IsCode(code int) bool {
	return code == 402
}

// Code gets the status code for the get fields exposed buckets payment required response
func (o *GetFieldsExposedBucketsPaymentRequired) Code() int {
	return 402
}

func (o *GetFieldsExposedBucketsPaymentRequired) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/buckets/][%d] getFieldsExposedBucketsPaymentRequired %s", 402, payload)
}

func (o *GetFieldsExposedBucketsPaymentRequired) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/buckets/][%d] getFieldsExposedBucketsPaymentRequired %s", 402, payload)
}

func (o *GetFieldsExposedBucketsPaymentRequired) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFieldsExposedBucketsPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFieldsExposedBucketsRequestTimeout creates a GetFieldsExposedBucketsRequestTimeout with default headers values
func NewGetFieldsExposedBucketsRequestTimeout() *GetFieldsExposedBucketsRequestTimeout {
	return &GetFieldsExposedBucketsRequestTimeout{}
}

/*
GetFieldsExposedBucketsRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetFieldsExposedBucketsRequestTimeout struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get fields exposed buckets request timeout response has a 2xx status code
func (o *GetFieldsExposedBucketsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fields exposed buckets request timeout response has a 3xx status code
func (o *GetFieldsExposedBucketsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fields exposed buckets request timeout response has a 4xx status code
func (o *GetFieldsExposedBucketsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fields exposed buckets request timeout response has a 5xx status code
func (o *GetFieldsExposedBucketsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get fields exposed buckets request timeout response a status code equal to that given
func (o *GetFieldsExposedBucketsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get fields exposed buckets request timeout response
func (o *GetFieldsExposedBucketsRequestTimeout) Code() int {
	return 408
}

func (o *GetFieldsExposedBucketsRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/buckets/][%d] getFieldsExposedBucketsRequestTimeout %s", 408, payload)
}

func (o *GetFieldsExposedBucketsRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/buckets/][%d] getFieldsExposedBucketsRequestTimeout %s", 408, payload)
}

func (o *GetFieldsExposedBucketsRequestTimeout) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFieldsExposedBucketsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFieldsExposedBucketsInternalServerError creates a GetFieldsExposedBucketsInternalServerError with default headers values
func NewGetFieldsExposedBucketsInternalServerError() *GetFieldsExposedBucketsInternalServerError {
	return &GetFieldsExposedBucketsInternalServerError{}
}

/*
GetFieldsExposedBucketsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetFieldsExposedBucketsInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get fields exposed buckets internal server error response has a 2xx status code
func (o *GetFieldsExposedBucketsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fields exposed buckets internal server error response has a 3xx status code
func (o *GetFieldsExposedBucketsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fields exposed buckets internal server error response has a 4xx status code
func (o *GetFieldsExposedBucketsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fields exposed buckets internal server error response has a 5xx status code
func (o *GetFieldsExposedBucketsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get fields exposed buckets internal server error response a status code equal to that given
func (o *GetFieldsExposedBucketsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get fields exposed buckets internal server error response
func (o *GetFieldsExposedBucketsInternalServerError) Code() int {
	return 500
}

func (o *GetFieldsExposedBucketsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/buckets/][%d] getFieldsExposedBucketsInternalServerError %s", 500, payload)
}

func (o *GetFieldsExposedBucketsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/buckets/][%d] getFieldsExposedBucketsInternalServerError %s", 500, payload)
}

func (o *GetFieldsExposedBucketsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFieldsExposedBucketsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetFieldsExposedBucketsOKBody get fields exposed buckets o k body
swagger:model GetFieldsExposedBucketsOKBody
*/
type GetFieldsExposedBucketsOKBody struct {
	models.APIResponse

	// data
	Data []*models.Field `json:"Data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetFieldsExposedBucketsOKBody) UnmarshalJSON(raw []byte) error {
	// GetFieldsExposedBucketsOKBodyAO0
	var getFieldsExposedBucketsOKBodyAO0 models.APIResponse
	if err := swag.ReadJSON(raw, &getFieldsExposedBucketsOKBodyAO0); err != nil {
		return err
	}
	o.APIResponse = getFieldsExposedBucketsOKBodyAO0

	// GetFieldsExposedBucketsOKBodyAO1
	var dataGetFieldsExposedBucketsOKBodyAO1 struct {
		Data []*models.Field `json:"Data"`
	}
	if err := swag.ReadJSON(raw, &dataGetFieldsExposedBucketsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetFieldsExposedBucketsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetFieldsExposedBucketsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getFieldsExposedBucketsOKBodyAO0, err := swag.WriteJSON(o.APIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getFieldsExposedBucketsOKBodyAO0)
	var dataGetFieldsExposedBucketsOKBodyAO1 struct {
		Data []*models.Field `json:"Data"`
	}

	dataGetFieldsExposedBucketsOKBodyAO1.Data = o.Data

	jsonDataGetFieldsExposedBucketsOKBodyAO1, errGetFieldsExposedBucketsOKBodyAO1 := swag.WriteJSON(dataGetFieldsExposedBucketsOKBodyAO1)
	if errGetFieldsExposedBucketsOKBodyAO1 != nil {
		return nil, errGetFieldsExposedBucketsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetFieldsExposedBucketsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get fields exposed buckets o k body
func (o *GetFieldsExposedBucketsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.APIResponse
	if err := o.APIResponse.Validate(formats); err != nil {
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

func (o *GetFieldsExposedBucketsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getFieldsExposedBucketsOK" + "." + "Data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getFieldsExposedBucketsOK" + "." + "Data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get fields exposed buckets o k body based on the context it is used
func (o *GetFieldsExposedBucketsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.APIResponse
	if err := o.APIResponse.ContextValidate(ctx, formats); err != nil {
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

func (o *GetFieldsExposedBucketsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getFieldsExposedBucketsOK" + "." + "Data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getFieldsExposedBucketsOK" + "." + "Data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFieldsExposedBucketsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFieldsExposedBucketsOKBody) UnmarshalBinary(b []byte) error {
	var res GetFieldsExposedBucketsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
