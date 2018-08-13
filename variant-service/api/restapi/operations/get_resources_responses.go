// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/CanDIG/go-model-service/variant-service/api/models"
)

// GetResourcesOKCode is the HTTP code returned for type GetResourcesOK
const GetResourcesOKCode int = 200

/*GetResourcesOK Return resources

swagger:response getResourcesOK
*/
type GetResourcesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Resource `json:"body,omitempty"`
}

// NewGetResourcesOK creates GetResourcesOK with default headers values
func NewGetResourcesOK() *GetResourcesOK {

	return &GetResourcesOK{}
}

// WithPayload adds the payload to the get resources o k response
func (o *GetResourcesOK) WithPayload(payload []*models.Resource) *GetResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get resources o k response
func (o *GetResourcesOK) SetPayload(payload []*models.Resource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Resource, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetResourcesInternalServerErrorCode is the HTTP code returned for type GetResourcesInternalServerError
const GetResourcesInternalServerErrorCode int = 500

/*GetResourcesInternalServerError Internal error

swagger:response getResourcesInternalServerError
*/
type GetResourcesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetResourcesInternalServerError creates GetResourcesInternalServerError with default headers values
func NewGetResourcesInternalServerError() *GetResourcesInternalServerError {

	return &GetResourcesInternalServerError{}
}

// WithPayload adds the payload to the get resources internal server error response
func (o *GetResourcesInternalServerError) WithPayload(payload *models.Error) *GetResourcesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get resources internal server error response
func (o *GetResourcesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetResourcesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
