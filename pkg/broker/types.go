package broker

import osb "github.com/pmorie/go-open-service-broker-client/v2"

// CatalogResponse is sent as the response to a catalog requests.
type CatalogResponse struct {
	osb.CatalogResponse
}

// ProvisionResponse is sent as the response to a provision call.
type ProvisionResponse struct {
	osb.ProvisionResponse
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
}

// UnbindResponse is sent as the response to a bind call.
type UnbindResponse struct {
	osb.UnbindResponse
}
