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

// checks if the TaxRuleUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxRuleUpdate{}

// TaxRuleUpdate struct for TaxRuleUpdate
type TaxRuleUpdate struct {
	Data TaxRuleUpdateData `json:"data"`
}

// NewTaxRuleUpdate instantiates a new TaxRuleUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleUpdate(data TaxRuleUpdateData) *TaxRuleUpdate {
	this := TaxRuleUpdate{}
	this.Data = data
	return &this
}

// NewTaxRuleUpdateWithDefaults instantiates a new TaxRuleUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleUpdateWithDefaults() *TaxRuleUpdate {
	this := TaxRuleUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *TaxRuleUpdate) GetData() TaxRuleUpdateData {
	if o == nil {
		var ret TaxRuleUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TaxRuleUpdate) GetDataOk() (*TaxRuleUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TaxRuleUpdate) SetData(v TaxRuleUpdateData) {
	o.Data = v
}

func (o TaxRuleUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxRuleUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableTaxRuleUpdate struct {
	value *TaxRuleUpdate
	isSet bool
}

func (v NullableTaxRuleUpdate) Get() *TaxRuleUpdate {
	return v.value
}

func (v *NullableTaxRuleUpdate) Set(val *TaxRuleUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleUpdate(val *TaxRuleUpdate) *NullableTaxRuleUpdate {
	return &NullableTaxRuleUpdate{value: val, isSet: true}
}

func (v NullableTaxRuleUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
