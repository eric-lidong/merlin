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

// checks if the BuildImageOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildImageOptions{}

// BuildImageOptions struct for BuildImageOptions
type BuildImageOptions struct {
	BackoffLimit    *int32           `json:"backoff_limit,omitempty"`
	ResourceRequest *ResourceRequest `json:"resource_request,omitempty"`
}

// NewBuildImageOptions instantiates a new BuildImageOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildImageOptions() *BuildImageOptions {
	this := BuildImageOptions{}
	return &this
}

// NewBuildImageOptionsWithDefaults instantiates a new BuildImageOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildImageOptionsWithDefaults() *BuildImageOptions {
	this := BuildImageOptions{}
	return &this
}

// GetBackoffLimit returns the BackoffLimit field value if set, zero value otherwise.
func (o *BuildImageOptions) GetBackoffLimit() int32 {
	if o == nil || IsNil(o.BackoffLimit) {
		var ret int32
		return ret
	}
	return *o.BackoffLimit
}

// GetBackoffLimitOk returns a tuple with the BackoffLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildImageOptions) GetBackoffLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.BackoffLimit) {
		return nil, false
	}
	return o.BackoffLimit, true
}

// HasBackoffLimit returns a boolean if a field has been set.
func (o *BuildImageOptions) HasBackoffLimit() bool {
	if o != nil && !IsNil(o.BackoffLimit) {
		return true
	}

	return false
}

// SetBackoffLimit gets a reference to the given int32 and assigns it to the BackoffLimit field.
func (o *BuildImageOptions) SetBackoffLimit(v int32) {
	o.BackoffLimit = &v
}

// GetResourceRequest returns the ResourceRequest field value if set, zero value otherwise.
func (o *BuildImageOptions) GetResourceRequest() ResourceRequest {
	if o == nil || IsNil(o.ResourceRequest) {
		var ret ResourceRequest
		return ret
	}
	return *o.ResourceRequest
}

// GetResourceRequestOk returns a tuple with the ResourceRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildImageOptions) GetResourceRequestOk() (*ResourceRequest, bool) {
	if o == nil || IsNil(o.ResourceRequest) {
		return nil, false
	}
	return o.ResourceRequest, true
}

// HasResourceRequest returns a boolean if a field has been set.
func (o *BuildImageOptions) HasResourceRequest() bool {
	if o != nil && !IsNil(o.ResourceRequest) {
		return true
	}

	return false
}

// SetResourceRequest gets a reference to the given ResourceRequest and assigns it to the ResourceRequest field.
func (o *BuildImageOptions) SetResourceRequest(v ResourceRequest) {
	o.ResourceRequest = &v
}

func (o BuildImageOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildImageOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackoffLimit) {
		toSerialize["backoff_limit"] = o.BackoffLimit
	}
	if !IsNil(o.ResourceRequest) {
		toSerialize["resource_request"] = o.ResourceRequest
	}
	return toSerialize, nil
}

type NullableBuildImageOptions struct {
	value *BuildImageOptions
	isSet bool
}

func (v NullableBuildImageOptions) Get() *BuildImageOptions {
	return v.value
}

func (v *NullableBuildImageOptions) Set(val *BuildImageOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildImageOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildImageOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildImageOptions(val *BuildImageOptions) *NullableBuildImageOptions {
	return &NullableBuildImageOptions{value: val, isSet: true}
}

func (v NullableBuildImageOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildImageOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
