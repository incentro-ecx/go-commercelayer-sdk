/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BraintreePaymentCreate struct for BraintreePaymentCreate
type BraintreePaymentCreate struct {
	Data BraintreePaymentCreateData `json:"data"`
}

// NewBraintreePaymentCreate instantiates a new BraintreePaymentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreePaymentCreate(data BraintreePaymentCreateData) *BraintreePaymentCreate {
	this := BraintreePaymentCreate{}
	this.Data = data
	return &this
}

// NewBraintreePaymentCreateWithDefaults instantiates a new BraintreePaymentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreePaymentCreateWithDefaults() *BraintreePaymentCreate {
	this := BraintreePaymentCreate{}
	return &this
}

// GetData returns the Data field value
func (o *BraintreePaymentCreate) GetData() BraintreePaymentCreateData {
	if o == nil {
		var ret BraintreePaymentCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BraintreePaymentCreate) GetDataOk() (*BraintreePaymentCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BraintreePaymentCreate) SetData(v BraintreePaymentCreateData) {
	o.Data = v
}

func (o BraintreePaymentCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBraintreePaymentCreate struct {
	value *BraintreePaymentCreate
	isSet bool
}

func (v NullableBraintreePaymentCreate) Get() *BraintreePaymentCreate {
	return v.value
}

func (v *NullableBraintreePaymentCreate) Set(val *BraintreePaymentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreePaymentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreePaymentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreePaymentCreate(val *BraintreePaymentCreate) *NullableBraintreePaymentCreate {
	return &NullableBraintreePaymentCreate{value: val, isSet: true}
}

func (v NullableBraintreePaymentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreePaymentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
