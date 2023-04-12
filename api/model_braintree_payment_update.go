/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the BraintreePaymentUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BraintreePaymentUpdate{}

// BraintreePaymentUpdate struct for BraintreePaymentUpdate
type BraintreePaymentUpdate struct {
	Data BraintreePaymentUpdateData `json:"data"`
}

// NewBraintreePaymentUpdate instantiates a new BraintreePaymentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreePaymentUpdate(data BraintreePaymentUpdateData) *BraintreePaymentUpdate {
	this := BraintreePaymentUpdate{}
	this.Data = data
	return &this
}

// NewBraintreePaymentUpdateWithDefaults instantiates a new BraintreePaymentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreePaymentUpdateWithDefaults() *BraintreePaymentUpdate {
	this := BraintreePaymentUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *BraintreePaymentUpdate) GetData() BraintreePaymentUpdateData {
	if o == nil {
		var ret BraintreePaymentUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BraintreePaymentUpdate) GetDataOk() (*BraintreePaymentUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BraintreePaymentUpdate) SetData(v BraintreePaymentUpdateData) {
	o.Data = v
}

func (o BraintreePaymentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BraintreePaymentUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBraintreePaymentUpdate struct {
	value *BraintreePaymentUpdate
	isSet bool
}

func (v NullableBraintreePaymentUpdate) Get() *BraintreePaymentUpdate {
	return v.value
}

func (v *NullableBraintreePaymentUpdate) Set(val *BraintreePaymentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreePaymentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreePaymentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreePaymentUpdate(val *BraintreePaymentUpdate) *NullableBraintreePaymentUpdate {
	return &NullableBraintreePaymentUpdate{value: val, isSet: true}
}

func (v NullableBraintreePaymentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreePaymentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
