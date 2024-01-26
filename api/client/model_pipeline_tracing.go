/*
Merlin

API Guide for accessing Merlin's model management, deployment, and serving functionalities

API version: 0.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PipelineTracing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineTracing{}

// PipelineTracing struct for PipelineTracing
type PipelineTracing struct {
	OperationType *string                `json:"operation_type,omitempty"`
	Spec          map[string]interface{} `json:"spec,omitempty"`
	Input         map[string]interface{} `json:"input,omitempty"`
	Output        map[string]interface{} `json:"output,omitempty"`
}

// NewPipelineTracing instantiates a new PipelineTracing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineTracing() *PipelineTracing {
	this := PipelineTracing{}
	return &this
}

// NewPipelineTracingWithDefaults instantiates a new PipelineTracing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineTracingWithDefaults() *PipelineTracing {
	this := PipelineTracing{}
	return &this
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *PipelineTracing) GetOperationType() string {
	if o == nil || IsNil(o.OperationType) {
		var ret string
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineTracing) GetOperationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OperationType) {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *PipelineTracing) HasOperationType() bool {
	if o != nil && !IsNil(o.OperationType) {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given string and assigns it to the OperationType field.
func (o *PipelineTracing) SetOperationType(v string) {
	o.OperationType = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *PipelineTracing) GetSpec() map[string]interface{} {
	if o == nil || IsNil(o.Spec) {
		var ret map[string]interface{}
		return ret
	}
	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineTracing) GetSpecOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Spec) {
		return map[string]interface{}{}, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *PipelineTracing) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given map[string]interface{} and assigns it to the Spec field.
func (o *PipelineTracing) SetSpec(v map[string]interface{}) {
	o.Spec = v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *PipelineTracing) GetInput() map[string]interface{} {
	if o == nil || IsNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineTracing) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *PipelineTracing) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *PipelineTracing) SetInput(v map[string]interface{}) {
	o.Input = v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *PipelineTracing) GetOutput() map[string]interface{} {
	if o == nil || IsNil(o.Output) {
		var ret map[string]interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineTracing) GetOutputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return map[string]interface{}{}, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *PipelineTracing) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given map[string]interface{} and assigns it to the Output field.
func (o *PipelineTracing) SetOutput(v map[string]interface{}) {
	o.Output = v
}

func (o PipelineTracing) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineTracing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OperationType) {
		toSerialize["operation_type"] = o.OperationType
	}
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	return toSerialize, nil
}

type NullablePipelineTracing struct {
	value *PipelineTracing
	isSet bool
}

func (v NullablePipelineTracing) Get() *PipelineTracing {
	return v.value
}

func (v *NullablePipelineTracing) Set(val *PipelineTracing) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineTracing) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineTracing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineTracing(val *PipelineTracing) *NullablePipelineTracing {
	return &NullablePipelineTracing{value: val, isSet: true}
}

func (v NullablePipelineTracing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineTracing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}