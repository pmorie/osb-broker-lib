package broker

import osb "github.com/pmorie/go-open-service-broker-client/v2"

// CatalogResponse is sent as the response to a catalog requests.
type CatalogResponse struct {
	osb.CatalogResponse
}

// ProvisionResponse is sent as the response to a provision call.
type ProvisionResponse struct {
	osb.ProvisionResponse

	// AlreadyReceived  is used to determine if the request was a duplicate
	// or not. should not be sent back in the respone.
	AlreadyReceived bool `json:"-"`
}

// UpdateInstanceResponse is sent as the response to a update call.
type UpdateInstanceResponse struct {
	osb.UpdateInstanceResponse
}

// DeprovisionResponse is sent as the response to a deprovision call.
type DeprovisionResponse struct {
	osb.DeprovisionResponse
}

// LastOperationResponse is sent as the response to a last operation call.
type LastOperationResponse struct {
	osb.LastOperationResponse
}

// BindResponse is sent as the response to a bind call.
type BindResponse struct {
	osb.BindResponse

	// AlreadyReceived  is used to determine if the request was a duplicate
	// or not. should not be sent back in the respone. This is needed for
	// async bind.
	AlreadyReceived bool `json:"-"`
}

// UnbindResponse is sent as the response to a bind call.
type UnbindResponse struct {
	osb.UnbindResponse
}
