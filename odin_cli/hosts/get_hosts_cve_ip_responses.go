// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cybledev/odin-cli/models"
)

// GetHostsCveIPReader is a Reader for the GetHostsCveIP structure.
type GetHostsCveIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostsCveIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostsCveIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400,402,401:
		result := NewGetHostsCveIPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetHostsCveIPRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHostsCveIPInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /hosts/cve/{ip}/] GetHostsCveIP", response, response.Code())
	}
}

// NewGetHostsCveIPOK creates a GetHostsCveIPOK with default headers values
func NewGetHostsCveIPOK() *GetHostsCveIPOK {
	return &GetHostsCveIPOK{}
}

/*
GetHostsCveIPOK describes a response with status code 200, with default header values.

OK
*/
type GetHostsCveIPOK struct {
	Payload *models.IpservicesAPIResponse
}

// IsSuccess returns true when this get hosts cve Ip o k response has a 2xx status code
func (o *GetHostsCveIPOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get hosts cve Ip o k response has a 3xx status code
func (o *GetHostsCveIPOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hosts cve Ip o k response has a 4xx status code
func (o *GetHostsCveIPOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get hosts cve Ip o k response has a 5xx status code
func (o *GetHostsCveIPOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get hosts cve Ip o k response a status code equal to that given
func (o *GetHostsCveIPOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get hosts cve Ip o k response
func (o *GetHostsCveIPOK) Code() int {
	return 200
}

func (o *GetHostsCveIPOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /hosts/cve/{ip}/][%d] getHostsCveIpOK %s", 200, payload)
}

func (o *GetHostsCveIPOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /hosts/cve/{ip}/][%d] getHostsCveIpOK %s", 200, payload)
}

func (o *GetHostsCveIPOK) GetPayload() *models.IpservicesAPIResponse {
	return o.Payload
}

func (o *GetHostsCveIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpservicesAPIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostsCveIPBadRequest creates a GetHostsCveIPBadRequest with default headers values
func NewGetHostsCveIPBadRequest() *GetHostsCveIPBadRequest {
	return &GetHostsCveIPBadRequest{}
}

/*
GetHostsCveIPBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetHostsCveIPBadRequest struct {
	Payload *models.IpservicesErrorResponse
}

// IsSuccess returns true when this get hosts cve Ip bad request response has a 2xx status code
func (o *GetHostsCveIPBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hosts cve Ip bad request response has a 3xx status code
func (o *GetHostsCveIPBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hosts cve Ip bad request response has a 4xx status code
func (o *GetHostsCveIPBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get hosts cve Ip bad request response has a 5xx status code
func (o *GetHostsCveIPBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get hosts cve Ip bad request response a status code equal to that given
func (o *GetHostsCveIPBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get hosts cve Ip bad request response
func (o *GetHostsCveIPBadRequest) Code() int {
	return 400
}

func (o *GetHostsCveIPBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /hosts/cve/{ip}/][%d] getHostsCveIpBadRequest %s", 400, payload)
}

func (o *GetHostsCveIPBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /hosts/cve/{ip}/][%d] getHostsCveIpBadRequest %s", 400, payload)
}

func (o *GetHostsCveIPBadRequest) GetPayload() *models.IpservicesErrorResponse {
	return o.Payload
}

func (o *GetHostsCveIPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpservicesErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostsCveIPRequestTimeout creates a GetHostsCveIPRequestTimeout with default headers values
func NewGetHostsCveIPRequestTimeout() *GetHostsCveIPRequestTimeout {
	return &GetHostsCveIPRequestTimeout{}
}

/*
GetHostsCveIPRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetHostsCveIPRequestTimeout struct {
	Payload *models.IpservicesErrorResponse
}

// IsSuccess returns true when this get hosts cve Ip request timeout response has a 2xx status code
func (o *GetHostsCveIPRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hosts cve Ip request timeout response has a 3xx status code
func (o *GetHostsCveIPRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hosts cve Ip request timeout response has a 4xx status code
func (o *GetHostsCveIPRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get hosts cve Ip request timeout response has a 5xx status code
func (o *GetHostsCveIPRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get hosts cve Ip request timeout response a status code equal to that given
func (o *GetHostsCveIPRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get hosts cve Ip request timeout response
func (o *GetHostsCveIPRequestTimeout) Code() int {
	return 408
}

func (o *GetHostsCveIPRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /hosts/cve/{ip}/][%d] getHostsCveIpRequestTimeout %s", 408, payload)
}

func (o *GetHostsCveIPRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /hosts/cve/{ip}/][%d] getHostsCveIpRequestTimeout %s", 408, payload)
}

func (o *GetHostsCveIPRequestTimeout) GetPayload() *models.IpservicesErrorResponse {
	return o.Payload
}

func (o *GetHostsCveIPRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpservicesErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostsCveIPInternalServerError creates a GetHostsCveIPInternalServerError with default headers values
func NewGetHostsCveIPInternalServerError() *GetHostsCveIPInternalServerError {
	return &GetHostsCveIPInternalServerError{}
}

/*
GetHostsCveIPInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetHostsCveIPInternalServerError struct {
	Payload *models.IpservicesErrorResponse
}

// IsSuccess returns true when this get hosts cve Ip internal server error response has a 2xx status code
func (o *GetHostsCveIPInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hosts cve Ip internal server error response has a 3xx status code
func (o *GetHostsCveIPInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hosts cve Ip internal server error response has a 4xx status code
func (o *GetHostsCveIPInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get hosts cve Ip internal server error response has a 5xx status code
func (o *GetHostsCveIPInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get hosts cve Ip internal server error response a status code equal to that given
func (o *GetHostsCveIPInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get hosts cve Ip internal server error response
func (o *GetHostsCveIPInternalServerError) Code() int {
	return 500
}

func (o *GetHostsCveIPInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /hosts/cve/{ip}/][%d] getHostsCveIpInternalServerError %s", 500, payload)
}

func (o *GetHostsCveIPInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /hosts/cve/{ip}/][%d] getHostsCveIpInternalServerError %s", 500, payload)
}

func (o *GetHostsCveIPInternalServerError) GetPayload() *models.IpservicesErrorResponse {
	return o.Payload
}

func (o *GetHostsCveIPInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpservicesErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
