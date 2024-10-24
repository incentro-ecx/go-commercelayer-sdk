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

// checks if the AdjustmentDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdjustmentDataRelationships{}

// AdjustmentDataRelationships struct for AdjustmentDataRelationships
type AdjustmentDataRelationships struct {
	Versions *AddressDataRelationshipsVersions `json:"versions,omitempty"`
}

// NewAdjustmentDataRelationships instantiates a new AdjustmentDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustmentDataRelationships() *AdjustmentDataRelationships {
	this := AdjustmentDataRelationships{}
	return &this
}

// NewAdjustmentDataRelationshipsWithDefaults instantiates a new AdjustmentDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustmentDataRelationshipsWithDefaults() *AdjustmentDataRelationships {
	this := AdjustmentDataRelationships{}
	return &this
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *AdjustmentDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *AdjustmentDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *AdjustmentDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o AdjustmentDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdjustmentDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableAdjustmentDataRelationships struct {
	value *AdjustmentDataRelationships
	isSet bool
}

func (v NullableAdjustmentDataRelationships) Get() *AdjustmentDataRelationships {
	return v.value
}

func (v *NullableAdjustmentDataRelationships) Set(val *AdjustmentDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustmentDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustmentDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustmentDataRelationships(val *AdjustmentDataRelationships) *NullableAdjustmentDataRelationships {
	return &NullableAdjustmentDataRelationships{value: val, isSet: true}
}

func (v NullableAdjustmentDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdjustmentDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}