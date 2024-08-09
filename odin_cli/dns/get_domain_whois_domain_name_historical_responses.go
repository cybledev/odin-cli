// Code generated by go-swagger; DO NOT EDIT.

package dns

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

// GetDomainWhoisDomainNameHistoricalReader is a Reader for the GetDomainWhoisDomainNameHistorical structure.
type GetDomainWhoisDomainNameHistoricalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainWhoisDomainNameHistoricalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDomainWhoisDomainNameHistoricalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400,401:
		result := NewGetDomainWhoisDomainNameHistoricalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewGetDomainWhoisDomainNameHistoricalPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetDomainWhoisDomainNameHistoricalRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDomainWhoisDomainNameHistoricalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /domain/whois/{domain-name}/historical] GetDomainWhoisDomainNameHistorical", response, response.Code())
	}
}

// NewGetDomainWhoisDomainNameHistoricalOK creates a GetDomainWhoisDomainNameHistoricalOK with default headers values
func NewGetDomainWhoisDomainNameHistoricalOK() *GetDomainWhoisDomainNameHistoricalOK {
	return &GetDomainWhoisDomainNameHistoricalOK{}
}

/*
GetDomainWhoisDomainNameHistoricalOK describes a response with status code 200, with default header values.

OK
*/
type GetDomainWhoisDomainNameHistoricalOK struct {
	Payload *GetDomainWhoisDomainNameHistoricalOKBody
}

// IsSuccess returns true when this get domain whois domain name historical o k response has a 2xx status code
func (o *GetDomainWhoisDomainNameHistoricalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get domain whois domain name historical o k response has a 3xx status code
func (o *GetDomainWhoisDomainNameHistoricalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain whois domain name historical o k response has a 4xx status code
func (o *GetDomainWhoisDomainNameHistoricalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domain whois domain name historical o k response has a 5xx status code
func (o *GetDomainWhoisDomainNameHistoricalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain whois domain name historical o k response a status code equal to that given
func (o *GetDomainWhoisDomainNameHistoricalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get domain whois domain name historical o k response
func (o *GetDomainWhoisDomainNameHistoricalOK) Code() int {
	return 200
}

func (o *GetDomainWhoisDomainNameHistoricalOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domain/whois/{domain-name}/historical][%d] getDomainWhoisDomainNameHistoricalOK %s", 200, payload)
}

func (o *GetDomainWhoisDomainNameHistoricalOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domain/whois/{domain-name}/historical][%d] getDomainWhoisDomainNameHistoricalOK %s", 200, payload)
}

func (o *GetDomainWhoisDomainNameHistoricalOK) GetPayload() *GetDomainWhoisDomainNameHistoricalOKBody {
	return o.Payload
}

func (o *GetDomainWhoisDomainNameHistoricalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDomainWhoisDomainNameHistoricalOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainWhoisDomainNameHistoricalBadRequest creates a GetDomainWhoisDomainNameHistoricalBadRequest with default headers values
func NewGetDomainWhoisDomainNameHistoricalBadRequest() *GetDomainWhoisDomainNameHistoricalBadRequest {
	return &GetDomainWhoisDomainNameHistoricalBadRequest{}
}

/*
GetDomainWhoisDomainNameHistoricalBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDomainWhoisDomainNameHistoricalBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this get domain whois domain name historical bad request response has a 2xx status code
func (o *GetDomainWhoisDomainNameHistoricalBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain whois domain name historical bad request response has a 3xx status code
func (o *GetDomainWhoisDomainNameHistoricalBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain whois domain name historical bad request response has a 4xx status code
func (o *GetDomainWhoisDomainNameHistoricalBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get domain whois domain name historical bad request response has a 5xx status code
func (o *GetDomainWhoisDomainNameHistoricalBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain whois domain name historical bad request response a status code equal to that given
func (o *GetDomainWhoisDomainNameHistoricalBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get domain whois domain name historical bad request response
func (o *GetDomainWhoisDomainNameHistoricalBadRequest) Code() int {
	return 400
}

func (o *GetDomainWhoisDomainNameHistoricalBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domain/whois/{domain-name}/historical][%d] getDomainWhoisDomainNameHistoricalBadRequest %s", 400, payload)
}

func (o *GetDomainWhoisDomainNameHistoricalBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domain/whois/{domain-name}/historical][%d] getDomainWhoisDomainNameHistoricalBadRequest %s", 400, payload)
}

func (o *GetDomainWhoisDomainNameHistoricalBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDomainWhoisDomainNameHistoricalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainWhoisDomainNameHistoricalPaymentRequired creates a GetDomainWhoisDomainNameHistoricalPaymentRequired with default headers values
func NewGetDomainWhoisDomainNameHistoricalPaymentRequired() *GetDomainWhoisDomainNameHistoricalPaymentRequired {
	return &GetDomainWhoisDomainNameHistoricalPaymentRequired{}
}

/*
GetDomainWhoisDomainNameHistoricalPaymentRequired describes a response with status code 402, with default header values.

Payment Required
*/
type GetDomainWhoisDomainNameHistoricalPaymentRequired struct {
	Payload interface{}
}

// IsSuccess returns true when this get domain whois domain name historical payment required response has a 2xx status code
func (o *GetDomainWhoisDomainNameHistoricalPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain whois domain name historical payment required response has a 3xx status code
func (o *GetDomainWhoisDomainNameHistoricalPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain whois domain name historical payment required response has a 4xx status code
func (o *GetDomainWhoisDomainNameHistoricalPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this get domain whois domain name historical payment required response has a 5xx status code
func (o *GetDomainWhoisDomainNameHistoricalPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain whois domain name historical payment required response a status code equal to that given
func (o *GetDomainWhoisDomainNameHistoricalPaymentRequired) IsCode(code int) bool {
	return code == 402
}

// Code gets the status code for the get domain whois domain name historical payment required response
func (o *GetDomainWhoisDomainNameHistoricalPaymentRequired) Code() int {
	return 402
}

func (o *GetDomainWhoisDomainNameHistoricalPaymentRequired) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domain/whois/{domain-name}/historical][%d] getDomainWhoisDomainNameHistoricalPaymentRequired %s", 402, payload)
}

func (o *GetDomainWhoisDomainNameHistoricalPaymentRequired) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domain/whois/{domain-name}/historical][%d] getDomainWhoisDomainNameHistoricalPaymentRequired %s", 402, payload)
}

func (o *GetDomainWhoisDomainNameHistoricalPaymentRequired) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDomainWhoisDomainNameHistoricalPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainWhoisDomainNameHistoricalRequestTimeout creates a GetDomainWhoisDomainNameHistoricalRequestTimeout with default headers values
func NewGetDomainWhoisDomainNameHistoricalRequestTimeout() *GetDomainWhoisDomainNameHistoricalRequestTimeout {
	return &GetDomainWhoisDomainNameHistoricalRequestTimeout{}
}

/*
GetDomainWhoisDomainNameHistoricalRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetDomainWhoisDomainNameHistoricalRequestTimeout struct {
	Payload interface{}
}

// IsSuccess returns true when this get domain whois domain name historical request timeout response has a 2xx status code
func (o *GetDomainWhoisDomainNameHistoricalRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain whois domain name historical request timeout response has a 3xx status code
func (o *GetDomainWhoisDomainNameHistoricalRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain whois domain name historical request timeout response has a 4xx status code
func (o *GetDomainWhoisDomainNameHistoricalRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get domain whois domain name historical request timeout response has a 5xx status code
func (o *GetDomainWhoisDomainNameHistoricalRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain whois domain name historical request timeout response a status code equal to that given
func (o *GetDomainWhoisDomainNameHistoricalRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get domain whois domain name historical request timeout response
func (o *GetDomainWhoisDomainNameHistoricalRequestTimeout) Code() int {
	return 408
}

func (o *GetDomainWhoisDomainNameHistoricalRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domain/whois/{domain-name}/historical][%d] getDomainWhoisDomainNameHistoricalRequestTimeout %s", 408, payload)
}

func (o *GetDomainWhoisDomainNameHistoricalRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domain/whois/{domain-name}/historical][%d] getDomainWhoisDomainNameHistoricalRequestTimeout %s", 408, payload)
}

func (o *GetDomainWhoisDomainNameHistoricalRequestTimeout) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDomainWhoisDomainNameHistoricalRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainWhoisDomainNameHistoricalInternalServerError creates a GetDomainWhoisDomainNameHistoricalInternalServerError with default headers values
func NewGetDomainWhoisDomainNameHistoricalInternalServerError() *GetDomainWhoisDomainNameHistoricalInternalServerError {
	return &GetDomainWhoisDomainNameHistoricalInternalServerError{}
}

/*
GetDomainWhoisDomainNameHistoricalInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDomainWhoisDomainNameHistoricalInternalServerError struct {
	Payload interface{}
}

// IsSuccess returns true when this get domain whois domain name historical internal server error response has a 2xx status code
func (o *GetDomainWhoisDomainNameHistoricalInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain whois domain name historical internal server error response has a 3xx status code
func (o *GetDomainWhoisDomainNameHistoricalInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain whois domain name historical internal server error response has a 4xx status code
func (o *GetDomainWhoisDomainNameHistoricalInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domain whois domain name historical internal server error response has a 5xx status code
func (o *GetDomainWhoisDomainNameHistoricalInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get domain whois domain name historical internal server error response a status code equal to that given
func (o *GetDomainWhoisDomainNameHistoricalInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get domain whois domain name historical internal server error response
func (o *GetDomainWhoisDomainNameHistoricalInternalServerError) Code() int {
	return 500
}

func (o *GetDomainWhoisDomainNameHistoricalInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domain/whois/{domain-name}/historical][%d] getDomainWhoisDomainNameHistoricalInternalServerError %s", 500, payload)
}

func (o *GetDomainWhoisDomainNameHistoricalInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domain/whois/{domain-name}/historical][%d] getDomainWhoisDomainNameHistoricalInternalServerError %s", 500, payload)
}

func (o *GetDomainWhoisDomainNameHistoricalInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDomainWhoisDomainNameHistoricalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetDomainWhoisDomainNameHistoricalOKBody get domain whois domain name historical o k body
swagger:model GetDomainWhoisDomainNameHistoricalOKBody
*/
type GetDomainWhoisDomainNameHistoricalOKBody struct {
	models.SchemaAPIResponse

	// meta
	Meta *models.SchemaPaginationMeta `json:" meta,omitempty"`

	// data
	Data []*models.SchemaDomainWhoisResponse `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetDomainWhoisDomainNameHistoricalOKBody) UnmarshalJSON(raw []byte) error {
	// GetDomainWhoisDomainNameHistoricalOKBodyAO0
	var getDomainWhoisDomainNameHistoricalOKBodyAO0 models.SchemaAPIResponse
	if err := swag.ReadJSON(raw, &getDomainWhoisDomainNameHistoricalOKBodyAO0); err != nil {
		return err
	}
	o.SchemaAPIResponse = getDomainWhoisDomainNameHistoricalOKBodyAO0

	// GetDomainWhoisDomainNameHistoricalOKBodyAO1
	var dataGetDomainWhoisDomainNameHistoricalOKBodyAO1 struct {
		Meta *models.SchemaPaginationMeta `json:" meta,omitempty"`

		Data []*models.SchemaDomainWhoisResponse `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataGetDomainWhoisDomainNameHistoricalOKBodyAO1); err != nil {
		return err
	}

	o.Meta = dataGetDomainWhoisDomainNameHistoricalOKBodyAO1.Meta

	o.Data = dataGetDomainWhoisDomainNameHistoricalOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetDomainWhoisDomainNameHistoricalOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getDomainWhoisDomainNameHistoricalOKBodyAO0, err := swag.WriteJSON(o.SchemaAPIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getDomainWhoisDomainNameHistoricalOKBodyAO0)
	var dataGetDomainWhoisDomainNameHistoricalOKBodyAO1 struct {
		Meta *models.SchemaPaginationMeta `json:" meta,omitempty"`

		Data []*models.SchemaDomainWhoisResponse `json:"data"`
	}

	dataGetDomainWhoisDomainNameHistoricalOKBodyAO1.Meta = o.Meta

	dataGetDomainWhoisDomainNameHistoricalOKBodyAO1.Data = o.Data

	jsonDataGetDomainWhoisDomainNameHistoricalOKBodyAO1, errGetDomainWhoisDomainNameHistoricalOKBodyAO1 := swag.WriteJSON(dataGetDomainWhoisDomainNameHistoricalOKBodyAO1)
	if errGetDomainWhoisDomainNameHistoricalOKBodyAO1 != nil {
		return nil, errGetDomainWhoisDomainNameHistoricalOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetDomainWhoisDomainNameHistoricalOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get domain whois domain name historical o k body
func (o *GetDomainWhoisDomainNameHistoricalOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.SchemaAPIResponse
	if err := o.SchemaAPIResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMeta(formats); err != nil {
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

func (o *GetDomainWhoisDomainNameHistoricalOKBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDomainWhoisDomainNameHistoricalOK" + "." + " meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getDomainWhoisDomainNameHistoricalOK" + "." + " meta")
			}
			return err
		}
	}

	return nil
}

func (o *GetDomainWhoisDomainNameHistoricalOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("getDomainWhoisDomainNameHistoricalOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getDomainWhoisDomainNameHistoricalOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get domain whois domain name historical o k body based on the context it is used
func (o *GetDomainWhoisDomainNameHistoricalOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.SchemaAPIResponse
	if err := o.SchemaAPIResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMeta(ctx, formats); err != nil {
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

func (o *GetDomainWhoisDomainNameHistoricalOKBody) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if o.Meta != nil {

		if swag.IsZero(o.Meta) { // not required
			return nil
		}

		if err := o.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDomainWhoisDomainNameHistoricalOK" + "." + " meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getDomainWhoisDomainNameHistoricalOK" + "." + " meta")
			}
			return err
		}
	}

	return nil
}

func (o *GetDomainWhoisDomainNameHistoricalOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDomainWhoisDomainNameHistoricalOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getDomainWhoisDomainNameHistoricalOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDomainWhoisDomainNameHistoricalOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDomainWhoisDomainNameHistoricalOKBody) UnmarshalBinary(b []byte) error {
	var res GetDomainWhoisDomainNameHistoricalOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
