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

// checks if the POSTAdjustments201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAdjustments201ResponseDataRelationships{}

// POSTAdjustments201ResponseDataRelationships struct for POSTAdjustments201ResponseDataRelationships
type POSTAdjustments201ResponseDataRelationships struct {
	Versions *POSTAddresses201ResponseDataRelationshipsVersions `json:"versions,omitempty"`
}

// NewPOSTAdjustments201ResponseDataRelationships instantiates a new POSTAdjustments201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdjustments201ResponseDataRelationships() *POSTAdjustments201ResponseDataRelationships {
	this := POSTAdjustments201ResponseDataRelationships{}
	return &this
}

// NewPOSTAdjustments201ResponseDataRelationshipsWithDefaults instantiates a new POSTAdjustments201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdjustments201ResponseDataRelationshipsWithDefaults() *POSTAdjustments201ResponseDataRelationships {
	this := POSTAdjustments201ResponseDataRelationships{}
	return &this
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTAdjustments201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdjustments201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTAdjustments201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTAdjustments201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTAdjustments201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAdjustments201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTAdjustments201ResponseDataRelationships struct {
	value *POSTAdjustments201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTAdjustments201ResponseDataRelationships) Get() *POSTAdjustments201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTAdjustments201ResponseDataRelationships) Set(val *POSTAdjustments201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdjustments201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdjustments201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdjustments201ResponseDataRelationships(val *POSTAdjustments201ResponseDataRelationships) *NullablePOSTAdjustments201ResponseDataRelationships {
	return &NullablePOSTAdjustments201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTAdjustments201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdjustments201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}