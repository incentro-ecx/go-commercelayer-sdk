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

// checks if the AxervePayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AxervePayment{}

// AxervePayment struct for AxervePayment
type AxervePayment struct {
	Data *AxervePaymentData `json:"data,omitempty"`
}

// NewAxervePayment instantiates a new AxervePayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAxervePayment() *AxervePayment {
	this := AxervePayment{}
	return &this
}

// NewAxervePaymentWithDefaults instantiates a new AxervePayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAxervePaymentWithDefaults() *AxervePayment {
	this := AxervePayment{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AxervePayment) GetData() AxervePaymentData {
	if o == nil || IsNil(o.Data) {
		var ret AxervePaymentData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AxervePayment) GetDataOk() (*AxervePaymentData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AxervePayment) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AxervePaymentData and assigns it to the Data field.
func (o *AxervePayment) SetData(v AxervePaymentData) {
	o.Data = &v
}

func (o AxervePayment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AxervePayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAxervePayment struct {
	value *AxervePayment
	isSet bool
}

func (v NullableAxervePayment) Get() *AxervePayment {
	return v.value
}

func (v *NullableAxervePayment) Set(val *AxervePayment) {
	v.value = val
	v.isSet = true
}

func (v NullableAxervePayment) IsSet() bool {
	return v.isSet
}

func (v *NullableAxervePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAxervePayment(val *AxervePayment) *NullableAxervePayment {
	return &NullableAxervePayment{value: val, isSet: true}
}

func (v NullableAxervePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAxervePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
