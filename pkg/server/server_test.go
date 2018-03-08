package server

import (
	"github.com/pmorie/osb-broker-lib/pkg/broker"

	osb "github.com/pmorie/go-open-service-broker-client/v2"
)

// TODO: is this more of an integration test?

// FakeBroker provides an implementation of the broker.Interface.
type FakeBroker struct {
	validateAPIVersion func(string) error
	getCatalog         func(c *broker.RequestContext) (*osb.CatalogResponse, error)
	provision          func(pr *osb.ProvisionRequest, c *broker.RequestContext) (*osb.ProvisionResponse, error)
	deprovision        func(request *osb.DeprovisionRequest, c *broker.RequestContext) (*osb.DeprovisionResponse, error)
	lastOperation      func(request *osb.LastOperationRequest, c *broker.RequestContext) (*osb.LastOperationResponse, error)
	bind               func(request *osb.BindRequest, c *broker.RequestContext) (*osb.BindResponse, error)
	unbind             func(request *osb.UnbindRequest, c *broker.RequestContext) (*osb.UnbindResponse, error)
	update             func(request *osb.UpdateInstanceRequest, c *broker.RequestContext) (*osb.UpdateInstanceResponse, error)
}

var _ broker.Interface = &FakeBroker{}

func (b *FakeBroker) GetCatalog(c *broker.RequestContext) (*osb.CatalogResponse, error) {
	return b.getCatalog(c)
}

func (b *FakeBroker) Provision(pr *osb.ProvisionRequest, c *broker.RequestContext) (*osb.ProvisionResponse, error) {
	return b.provision(pr, c)
}

func (b *FakeBroker) Deprovision(request *osb.DeprovisionRequest, c *broker.RequestContext) (*osb.DeprovisionResponse, error) {
	return b.deprovision(request, c)
}

func (b *FakeBroker) LastOperation(request *osb.LastOperationRequest, c *broker.RequestContext) (*osb.LastOperationResponse, error) {
	return b.lastOperation(request, c)
}

func (b *FakeBroker) Bind(request *osb.BindRequest, c *broker.RequestContext) (*osb.BindResponse, error) {
	return b.bind(request, c)
}

func (b *FakeBroker) Unbind(request *osb.UnbindRequest, c *broker.RequestContext) (*osb.UnbindResponse, error) {
	return b.unbind(request, c)
}

func (b *FakeBroker) ValidateBrokerAPIVersion(version string) error {
	return b.validateAPIVersion(version)
}

func (b *FakeBroker) Update(request *osb.UpdateInstanceRequest, c *broker.RequestContext) (*osb.UpdateInstanceResponse, error) {
	return b.update(request, c)
}

func defaultValidateFunc(_ string) error {
	return nil
}

func strPtr(s string) *string {
	return &s
}

func defaultClientConfiguration() *osb.ClientConfiguration {
	conf := osb.DefaultClientConfiguration()
	conf.Verbose = true

	return conf
}

func originatingIdentity() *osb.OriginatingIdentity {
	return &osb.OriginatingIdentity{
		Platform: "kubernetes",
		Value:    `{"username":"test", "groups": [], "extra": {}}`,
	}
}
