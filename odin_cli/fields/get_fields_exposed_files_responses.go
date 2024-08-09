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

// GetFieldsExposedFilesReader is a Reader for the GetFieldsExposedFiles structure.
type GetFieldsExposedFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFieldsExposedFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFieldsExposedFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400,401:
		result := NewGetFieldsExposedFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewGetFieldsExposedFilesPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFieldsExposedFilesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFieldsExposedFilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /fields/exposed/files/] GetFieldsExposedFiles", response, response.Code())
	}
}

// NewGetFieldsExposedFilesOK creates a GetFieldsExposedFilesOK with default headers values
func NewGetFieldsExposedFilesOK() *GetFieldsExposedFilesOK {
	return &GetFieldsExposedFilesOK{}
}

/*
GetFieldsExposedFilesOK describes a response with status code 200, with default header values.

OK
*/
type GetFieldsExposedFilesOK struct {
	Payload *GetFieldsExposedFilesOKBody
}

// IsSuccess returns true when this get fields exposed files o k response has a 2xx status code
func (o *GetFieldsExposedFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get fields exposed files o k response has a 3xx status code
func (o *GetFieldsExposedFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fields exposed files o k response has a 4xx status code
func (o *GetFieldsExposedFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fields exposed files o k response has a 5xx status code
func (o *GetFieldsExposedFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get fields exposed files o k response a status code equal to that given
func (o *GetFieldsExposedFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get fields exposed files o k response
func (o *GetFieldsExposedFilesOK) Code() int {
	return 200
}

func (o *GetFieldsExposedFilesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/files/][%d] getFieldsExposedFilesOK %s", 200, payload)
}

func (o *GetFieldsExposedFilesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/files/][%d] getFieldsExposedFilesOK %s", 200, payload)
}

func (o *GetFieldsExposedFilesOK) GetPayload() *GetFieldsExposedFilesOKBody {
	return o.Payload
}

func (o *GetFieldsExposedFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetFieldsExposedFilesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFieldsExposedFilesBadRequest creates a GetFieldsExposedFilesBadRequest with default headers values
func NewGetFieldsExposedFilesBadRequest() *GetFieldsExposedFilesBadRequest {
	return &GetFieldsExposedFilesBadRequest{}
}

/*
GetFieldsExposedFilesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetFieldsExposedFilesBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get fields exposed files bad request response has a 2xx status code
func (o *GetFieldsExposedFilesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fields exposed files bad request response has a 3xx status code
func (o *GetFieldsExposedFilesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fields exposed files bad request response has a 4xx status code
func (o *GetFieldsExposedFilesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fields exposed files bad request response has a 5xx status code
func (o *GetFieldsExposedFilesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get fields exposed files bad request response a status code equal to that given
func (o *GetFieldsExposedFilesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get fields exposed files bad request response
func (o *GetFieldsExposedFilesBadRequest) Code() int {
	return 400
}

func (o *GetFieldsExposedFilesBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/files/][%d] getFieldsExposedFilesBadRequest %s", 400, payload)
}

func (o *GetFieldsExposedFilesBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/files/][%d] getFieldsExposedFilesBadRequest %s", 400, payload)
}

func (o *GetFieldsExposedFilesBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFieldsExposedFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFieldsExposedFilesPaymentRequired creates a GetFieldsExposedFilesPaymentRequired with default headers values
func NewGetFieldsExposedFilesPaymentRequired() *GetFieldsExposedFilesPaymentRequired {
	return &GetFieldsExposedFilesPaymentRequired{}
}

/*
GetFieldsExposedFilesPaymentRequired describes a response with status code 402, with default header values.

Payment Required
*/
type GetFieldsExposedFilesPaymentRequired struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get fields exposed files payment required response has a 2xx status code
func (o *GetFieldsExposedFilesPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fields exposed files payment required response has a 3xx status code
func (o *GetFieldsExposedFilesPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fields exposed files payment required response has a 4xx status code
func (o *GetFieldsExposedFilesPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fields exposed files payment required response has a 5xx status code
func (o *GetFieldsExposedFilesPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this get fields exposed files payment required response a status code equal to that given
func (o *GetFieldsExposedFilesPaymentRequired) IsCode(code int) bool {
	return code == 402
}

// Code gets the status code for the get fields exposed files payment required response
func (o *GetFieldsExposedFilesPaymentRequired) Code() int {
	return 402
}

func (o *GetFieldsExposedFilesPaymentRequired) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/files/][%d] getFieldsExposedFilesPaymentRequired %s", 402, payload)
}

func (o *GetFieldsExposedFilesPaymentRequired) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/files/][%d] getFieldsExposedFilesPaymentRequired %s", 402, payload)
}

func (o *GetFieldsExposedFilesPaymentRequired) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFieldsExposedFilesPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFieldsExposedFilesRequestTimeout creates a GetFieldsExposedFilesRequestTimeout with default headers values
func NewGetFieldsExposedFilesRequestTimeout() *GetFieldsExposedFilesRequestTimeout {
	return &GetFieldsExposedFilesRequestTimeout{}
}

/*
GetFieldsExposedFilesRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetFieldsExposedFilesRequestTimeout struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get fields exposed files request timeout response has a 2xx status code
func (o *GetFieldsExposedFilesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fields exposed files request timeout response has a 3xx status code
func (o *GetFieldsExposedFilesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fields exposed files request timeout response has a 4xx status code
func (o *GetFieldsExposedFilesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fields exposed files request timeout response has a 5xx status code
func (o *GetFieldsExposedFilesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get fields exposed files request timeout response a status code equal to that given
func (o *GetFieldsExposedFilesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get fields exposed files request timeout response
func (o *GetFieldsExposedFilesRequestTimeout) Code() int {
	return 408
}

func (o *GetFieldsExposedFilesRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/files/][%d] getFieldsExposedFilesRequestTimeout %s", 408, payload)
}

func (o *GetFieldsExposedFilesRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/files/][%d] getFieldsExposedFilesRequestTimeout %s", 408, payload)
}

func (o *GetFieldsExposedFilesRequestTimeout) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFieldsExposedFilesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFieldsExposedFilesInternalServerError creates a GetFieldsExposedFilesInternalServerError with default headers values
func NewGetFieldsExposedFilesInternalServerError() *GetFieldsExposedFilesInternalServerError {
	return &GetFieldsExposedFilesInternalServerError{}
}

/*
GetFieldsExposedFilesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetFieldsExposedFilesInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get fields exposed files internal server error response has a 2xx status code
func (o *GetFieldsExposedFilesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fields exposed files internal server error response has a 3xx status code
func (o *GetFieldsExposedFilesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fields exposed files internal server error response has a 4xx status code
func (o *GetFieldsExposedFilesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fields exposed files internal server error response has a 5xx status code
func (o *GetFieldsExposedFilesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get fields exposed files internal server error response a status code equal to that given
func (o *GetFieldsExposedFilesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get fields exposed files internal server error response
func (o *GetFieldsExposedFilesInternalServerError) Code() int {
	return 500
}

func (o *GetFieldsExposedFilesInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/files/][%d] getFieldsExposedFilesInternalServerError %s", 500, payload)
}

func (o *GetFieldsExposedFilesInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fields/exposed/files/][%d] getFieldsExposedFilesInternalServerError %s", 500, payload)
}

func (o *GetFieldsExposedFilesInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFieldsExposedFilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetFieldsExposedFilesOKBody get fields exposed files o k body
swagger:model GetFieldsExposedFilesOKBody
*/
type GetFieldsExposedFilesOKBody struct {
	models.APIResponse

	// data
	Data []*models.Field `json:"Data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetFieldsExposedFilesOKBody) UnmarshalJSON(raw []byte) error {
	// GetFieldsExposedFilesOKBodyAO0
	var getFieldsExposedFilesOKBodyAO0 models.APIResponse
	if err := swag.ReadJSON(raw, &getFieldsExposedFilesOKBodyAO0); err != nil {
		return err
	}
	o.APIResponse = getFieldsExposedFilesOKBodyAO0

	// GetFieldsExposedFilesOKBodyAO1
	var dataGetFieldsExposedFilesOKBodyAO1 struct {
		Data []*models.Field `json:"Data"`
	}
	if err := swag.ReadJSON(raw, &dataGetFieldsExposedFilesOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetFieldsExposedFilesOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetFieldsExposedFilesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getFieldsExposedFilesOKBodyAO0, err := swag.WriteJSON(o.APIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getFieldsExposedFilesOKBodyAO0)
	var dataGetFieldsExposedFilesOKBodyAO1 struct {
		Data []*models.Field `json:"Data"`
	}

	dataGetFieldsExposedFilesOKBodyAO1.Data = o.Data

	jsonDataGetFieldsExposedFilesOKBodyAO1, errGetFieldsExposedFilesOKBodyAO1 := swag.WriteJSON(dataGetFieldsExposedFilesOKBodyAO1)
	if errGetFieldsExposedFilesOKBodyAO1 != nil {
		return nil, errGetFieldsExposedFilesOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetFieldsExposedFilesOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get fields exposed files o k body
func (o *GetFieldsExposedFilesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetFieldsExposedFilesOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("getFieldsExposedFilesOK" + "." + "Data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getFieldsExposedFilesOK" + "." + "Data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get fields exposed files o k body based on the context it is used
func (o *GetFieldsExposedFilesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetFieldsExposedFilesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getFieldsExposedFilesOK" + "." + "Data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getFieldsExposedFilesOK" + "." + "Data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFieldsExposedFilesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFieldsExposedFilesOKBody) UnmarshalBinary(b []byte) error {
	var res GetFieldsExposedFilesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
