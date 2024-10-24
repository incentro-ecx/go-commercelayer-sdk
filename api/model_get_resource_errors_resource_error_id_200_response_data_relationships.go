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

// checks if the GETResourceErrorsResourceErrorId200ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETResourceErrorsResourceErrorId200ResponseDataRelationships{}

// GETResourceErrorsResourceErrorId200ResponseDataRelationships struct for GETResourceErrorsResourceErrorId200ResponseDataRelationships
type GETResourceErrorsResourceErrorId200ResponseDataRelationships struct {
	Resource *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource `json:"resource,omitempty"`
}

// NewGETResourceErrorsResourceErrorId200ResponseDataRelationships instantiates a new GETResourceErrorsResourceErrorId200ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETResourceErrorsResourceErrorId200ResponseDataRelationships() *GETResourceErrorsResourceErrorId200ResponseDataRelationships {
	this := GETResourceErrorsResourceErrorId200ResponseDataRelationships{}
	return &this
}

// NewGETResourceErrorsResourceErrorId200ResponseDataRelationshipsWithDefaults instantiates a new GETResourceErrorsResourceErrorId200ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETResourceErrorsResourceErrorId200ResponseDataRelationshipsWithDefaults() *GETResourceErrorsResourceErrorId200ResponseDataRelationships {
	this := GETResourceErrorsResourceErrorId200ResponseDataRelationships{}
	return &this
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *GETResourceErrorsResourceErrorId200ResponseDataRelationships) GetResource() GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource {
	if o == nil || IsNil(o.Resource) {
		var ret GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETResourceErrorsResourceErrorId200ResponseDataRelationships) GetResourceOk() (*GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *GETResourceErrorsResourceErrorId200ResponseDataRelationships) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource and assigns it to the Resource field.
func (o *GETResourceErrorsResourceErrorId200ResponseDataRelationships) SetResource(v GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) {
	o.Resource = &v
}

func (o GETResourceErrorsResourceErrorId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETResourceErrorsResourceErrorId200ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return toSerialize, nil
}

type NullableGETResourceErrorsResourceErrorId200ResponseDataRelationships struct {
	value *GETResourceErrorsResourceErrorId200ResponseDataRelationships
	isSet bool
}

func (v NullableGETResourceErrorsResourceErrorId200ResponseDataRelationships) Get() *GETResourceErrorsResourceErrorId200ResponseDataRelationships {
	return v.value
}

func (v *NullableGETResourceErrorsResourceErrorId200ResponseDataRelationships) Set(val *GETResourceErrorsResourceErrorId200ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETResourceErrorsResourceErrorId200ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETResourceErrorsResourceErrorId200ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETResourceErrorsResourceErrorId200ResponseDataRelationships(val *GETResourceErrorsResourceErrorId200ResponseDataRelationships) *NullableGETResourceErrorsResourceErrorId200ResponseDataRelationships {
	return &NullableGETResourceErrorsResourceErrorId200ResponseDataRelationships{value: val, isSet: true}
}

func (v NullableGETResourceErrorsResourceErrorId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETResourceErrorsResourceErrorId200ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
