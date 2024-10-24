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

// checks if the CleanupCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CleanupCreate{}

// CleanupCreate struct for CleanupCreate
type CleanupCreate struct {
	Data CleanupCreateData `json:"data"`
}

// NewCleanupCreate instantiates a new CleanupCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCleanupCreate(data CleanupCreateData) *CleanupCreate {
	this := CleanupCreate{}
	this.Data = data
	return &this
}

// NewCleanupCreateWithDefaults instantiates a new CleanupCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCleanupCreateWithDefaults() *CleanupCreate {
	this := CleanupCreate{}
	return &this
}

// GetData returns the Data field value
func (o *CleanupCreate) GetData() CleanupCreateData {
	if o == nil {
		var ret CleanupCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CleanupCreate) GetDataOk() (*CleanupCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CleanupCreate) SetData(v CleanupCreateData) {
	o.Data = v
}

func (o CleanupCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CleanupCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCleanupCreate struct {
	value *CleanupCreate
	isSet bool
}

func (v NullableCleanupCreate) Get() *CleanupCreate {
	return v.value
}

func (v *NullableCleanupCreate) Set(val *CleanupCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableCleanupCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableCleanupCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCleanupCreate(val *CleanupCreate) *NullableCleanupCreate {
	return &NullableCleanupCreate{value: val, isSet: true}
}

func (v NullableCleanupCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCleanupCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
