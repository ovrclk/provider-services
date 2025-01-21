/*
Copyright The Akash Network Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v2beta1

// ProviderLeasedIPSpecApplyConfiguration represents a declarative configuration of the ProviderLeasedIPSpec type for use
// with apply.
type ProviderLeasedIPSpecApplyConfiguration struct {
	LeaseID      *LeaseIDApplyConfiguration `json:"lease_id,omitempty"`
	ServiceName  *string                    `json:"service_name,omitempty"`
	Port         *uint32                    `json:"port,omitempty"`
	ExternalPort *uint32                    `json:"external_port,omitempty"`
	SharingKey   *string                    `json:"sharing_key,omitempty"`
	Protocol     *string                    `json:"protocol,omitempty"`
}

// ProviderLeasedIPSpecApplyConfiguration constructs a declarative configuration of the ProviderLeasedIPSpec type for use with
// apply.
func ProviderLeasedIPSpec() *ProviderLeasedIPSpecApplyConfiguration {
	return &ProviderLeasedIPSpecApplyConfiguration{}
}

// WithLeaseID sets the LeaseID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LeaseID field is set to the value of the last call.
func (b *ProviderLeasedIPSpecApplyConfiguration) WithLeaseID(value *LeaseIDApplyConfiguration) *ProviderLeasedIPSpecApplyConfiguration {
	b.LeaseID = value
	return b
}

// WithServiceName sets the ServiceName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceName field is set to the value of the last call.
func (b *ProviderLeasedIPSpecApplyConfiguration) WithServiceName(value string) *ProviderLeasedIPSpecApplyConfiguration {
	b.ServiceName = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *ProviderLeasedIPSpecApplyConfiguration) WithPort(value uint32) *ProviderLeasedIPSpecApplyConfiguration {
	b.Port = &value
	return b
}

// WithExternalPort sets the ExternalPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExternalPort field is set to the value of the last call.
func (b *ProviderLeasedIPSpecApplyConfiguration) WithExternalPort(value uint32) *ProviderLeasedIPSpecApplyConfiguration {
	b.ExternalPort = &value
	return b
}

// WithSharingKey sets the SharingKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SharingKey field is set to the value of the last call.
func (b *ProviderLeasedIPSpecApplyConfiguration) WithSharingKey(value string) *ProviderLeasedIPSpecApplyConfiguration {
	b.SharingKey = &value
	return b
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *ProviderLeasedIPSpecApplyConfiguration) WithProtocol(value string) *ProviderLeasedIPSpecApplyConfiguration {
	b.Protocol = &value
	return b
}
