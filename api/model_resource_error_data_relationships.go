/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ResourceErrorDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceErrorDataRelationships{}

// ResourceErrorDataRelationships struct for ResourceErrorDataRelationships
type ResourceErrorDataRelationships struct {
	Resource *ResourceErrorDataRelationshipsResource `json:"resource,omitempty"`
}

// NewResourceErrorDataRelationships instantiates a new ResourceErrorDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceErrorDataRelationships() *ResourceErrorDataRelationships {
	this := ResourceErrorDataRelationships{}
	return &this
}

// NewResourceErrorDataRelationshipsWithDefaults instantiates a new ResourceErrorDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceErrorDataRelationshipsWithDefaults() *ResourceErrorDataRelationships {
	this := ResourceErrorDataRelationships{}
	return &this
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ResourceErrorDataRelationships) GetResource() ResourceErrorDataRelationshipsResource {
	if o == nil || IsNil(o.Resource) {
		var ret ResourceErrorDataRelationshipsResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceErrorDataRelationships) GetResourceOk() (*ResourceErrorDataRelationshipsResource, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ResourceErrorDataRelationships) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given ResourceErrorDataRelationshipsResource and assigns it to the Resource field.
func (o *ResourceErrorDataRelationships) SetResource(v ResourceErrorDataRelationshipsResource) {
	o.Resource = &v
}

func (o ResourceErrorDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceErrorDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return toSerialize, nil
}

type NullableResourceErrorDataRelationships struct {
	value *ResourceErrorDataRelationships
	isSet bool
}

func (v NullableResourceErrorDataRelationships) Get() *ResourceErrorDataRelationships {
	return v.value
}

func (v *NullableResourceErrorDataRelationships) Set(val *ResourceErrorDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceErrorDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceErrorDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceErrorDataRelationships(val *ResourceErrorDataRelationships) *NullableResourceErrorDataRelationships {
	return &NullableResourceErrorDataRelationships{value: val, isSet: true}
}

func (v NullableResourceErrorDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceErrorDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
