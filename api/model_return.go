/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ModelReturn struct for ModelReturn
type ModelReturn struct {
	Data ReturnData `json:"data"`
}

// NewModelReturn instantiates a new ModelReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelReturn(data ReturnData) *ModelReturn {
	this := ModelReturn{}
	this.Data = data
	return &this
}

// NewModelReturnWithDefaults instantiates a new ModelReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelReturnWithDefaults() *ModelReturn {
	this := ModelReturn{}
	return &this
}

// GetData returns the Data field value
func (o *ModelReturn) GetData() ReturnData {
	if o == nil {
		var ret ReturnData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ModelReturn) GetDataOk() (*ReturnData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ModelReturn) SetData(v ReturnData) {
	o.Data = v
}

func (o ModelReturn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableModelReturn struct {
	value *ModelReturn
	isSet bool
}

func (v NullableModelReturn) Get() *ModelReturn {
	return v.value
}

func (v *NullableModelReturn) Set(val *ModelReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableModelReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableModelReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelReturn(val *ModelReturn) *NullableModelReturn {
	return &NullableModelReturn{value: val, isSet: true}
}

func (v NullableModelReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}