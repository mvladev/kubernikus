package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sapcc/kubernikus/pkg/api/models"
)

// PostAPIV1ClustersOKCode is the HTTP code returned for type PostAPIV1ClustersOK
const PostAPIV1ClustersOKCode int = 200

/*PostAPIV1ClustersOK OK

swagger:response postApiV1ClustersOK
*/
type PostAPIV1ClustersOK struct {

	/*
	  In: Body
	*/
	Payload *models.Cluster `json:"body,omitempty"`
}

// NewPostAPIV1ClustersOK creates PostAPIV1ClustersOK with default headers values
func NewPostAPIV1ClustersOK() *PostAPIV1ClustersOK {
	return &PostAPIV1ClustersOK{}
}

// WithPayload adds the payload to the post Api v1 clusters o k response
func (o *PostAPIV1ClustersOK) WithPayload(payload *models.Cluster) *PostAPIV1ClustersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Api v1 clusters o k response
func (o *PostAPIV1ClustersOK) SetPayload(payload *models.Cluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAPIV1ClustersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}