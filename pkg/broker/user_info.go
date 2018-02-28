package broker

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// Identity that is sent from the OSB spec in the Originating Identity Header
// https://github.com/openservicebrokerapi/servicebroker/blob/master/profile.md#originating-identity-header
type Identity interface {
	// Platform - Retrieve the platform for the identity.
	Platform() string
	// Value - Retrieve the value for the identity.
	Value() interface{}
}

// NewIdentity - will return a new identity
func NewIdentity(platform, value string) (Identity, error) {
	val, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil, fmt.Errorf("unable decode identity value")
	}
	switch strings.ToLower(platform) {
	case "kubernetes":
		u := UserInfo{}
		err = json.Unmarshal(val, &u)
		if err != nil {
			return nil, fmt.Errorf("unable to unmarshal json for value")
		}
		return K8SIdentity{userInfo: u, platform: platform}, nil
	case "cloudfoundry":
		m := map[string]interface{}{}
		err = json.Unmarshal(val, &m)
		if err != nil {
			return nil, fmt.Errorf("unable to unmarshal json for value")
		}
		return CloudFoundry{platform: platform, value: m}, nil
	}
	return nil, fmt.Errorf("unable to determine type of identity")
}

// K8SIdentity - Identity implementation based on the kubernetes user info object
type K8SIdentity struct {
	platform string
	userInfo UserInfo
}

// UserInfo - kubernetest user info object
type UserInfo struct {
	Username string
	UID      string
	Groups   []string
	Extra    map[string][]string
}

// Platform - Retrieve the platform for the identity.
func (k K8SIdentity) Platform() string {
	return k.platform
}

// Value - Retrieve the value for the identity.
func (k K8SIdentity) Value() interface{} {
	return k.userInfo
}

// CloudFoundry - Identity implementation based on cloud foundry
type CloudFoundry struct {
	platform string
	value    map[string]interface{}
}

// Platform - Retrieve the platform for the identity.
func (c CloudFoundry) Platform() string {
	return c.platform
}

// Value - Retrieve the value for the identity.
func (c CloudFoundry) Value() interface{} {
	return c.Value
}
