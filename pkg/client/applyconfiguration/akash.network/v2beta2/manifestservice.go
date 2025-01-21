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

package v2beta2

// ManifestServiceApplyConfiguration represents a declarative configuration of the ManifestService type for use
// with apply.
type ManifestServiceApplyConfiguration struct {
	Name            *string                                       `json:"name,omitempty"`
	Image           *string                                       `json:"image,omitempty"`
	Command         []string                                      `json:"command,omitempty"`
	Args            []string                                      `json:"args,omitempty"`
	Env             []string                                      `json:"env,omitempty"`
	Resources       *ResourcesApplyConfiguration                  `json:"resources,omitempty"`
	Count           *uint32                                       `json:"count,omitempty"`
	Expose          []ManifestServiceExposeApplyConfiguration     `json:"expose,omitempty"`
	Params          *ManifestServiceParamsApplyConfiguration      `json:"params,omitempty"`
	SchedulerParams *SchedulerParamsApplyConfiguration            `json:"scheduler_params,omitempty"`
	Credentials     *ManifestServiceCredentialsApplyConfiguration `json:"credentials,omitempty"`
}

// ManifestServiceApplyConfiguration constructs a declarative configuration of the ManifestService type for use with
// apply.
func ManifestService() *ManifestServiceApplyConfiguration {
	return &ManifestServiceApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ManifestServiceApplyConfiguration) WithName(value string) *ManifestServiceApplyConfiguration {
	b.Name = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *ManifestServiceApplyConfiguration) WithImage(value string) *ManifestServiceApplyConfiguration {
	b.Image = &value
	return b
}

// WithCommand adds the given value to the Command field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Command field.
func (b *ManifestServiceApplyConfiguration) WithCommand(values ...string) *ManifestServiceApplyConfiguration {
	for i := range values {
		b.Command = append(b.Command, values[i])
	}
	return b
}

// WithArgs adds the given value to the Args field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Args field.
func (b *ManifestServiceApplyConfiguration) WithArgs(values ...string) *ManifestServiceApplyConfiguration {
	for i := range values {
		b.Args = append(b.Args, values[i])
	}
	return b
}

// WithEnv adds the given value to the Env field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Env field.
func (b *ManifestServiceApplyConfiguration) WithEnv(values ...string) *ManifestServiceApplyConfiguration {
	for i := range values {
		b.Env = append(b.Env, values[i])
	}
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *ManifestServiceApplyConfiguration) WithResources(value *ResourcesApplyConfiguration) *ManifestServiceApplyConfiguration {
	b.Resources = value
	return b
}

// WithCount sets the Count field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Count field is set to the value of the last call.
func (b *ManifestServiceApplyConfiguration) WithCount(value uint32) *ManifestServiceApplyConfiguration {
	b.Count = &value
	return b
}

// WithExpose adds the given value to the Expose field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Expose field.
func (b *ManifestServiceApplyConfiguration) WithExpose(values ...*ManifestServiceExposeApplyConfiguration) *ManifestServiceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExpose")
		}
		b.Expose = append(b.Expose, *values[i])
	}
	return b
}

// WithParams sets the Params field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Params field is set to the value of the last call.
func (b *ManifestServiceApplyConfiguration) WithParams(value *ManifestServiceParamsApplyConfiguration) *ManifestServiceApplyConfiguration {
	b.Params = value
	return b
}

// WithSchedulerParams sets the SchedulerParams field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SchedulerParams field is set to the value of the last call.
func (b *ManifestServiceApplyConfiguration) WithSchedulerParams(value *SchedulerParamsApplyConfiguration) *ManifestServiceApplyConfiguration {
	b.SchedulerParams = value
	return b
}

// WithCredentials sets the Credentials field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Credentials field is set to the value of the last call.
func (b *ManifestServiceApplyConfiguration) WithCredentials(value *ManifestServiceCredentialsApplyConfiguration) *ManifestServiceApplyConfiguration {
	b.Credentials = value
	return b
}
