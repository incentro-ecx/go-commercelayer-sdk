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

// checks if the KlarnaGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KlarnaGateway{}

// KlarnaGateway struct for KlarnaGateway
type KlarnaGateway struct {
	Data *KlarnaGatewayData `json:"data,omitempty"`
}

// NewKlarnaGateway instantiates a new KlarnaGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaGateway() *KlarnaGateway {
	this := KlarnaGateway{}
	return &this
}

// NewKlarnaGatewayWithDefaults instantiates a new KlarnaGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaGatewayWithDefaults() *KlarnaGateway {
	this := KlarnaGateway{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *KlarnaGateway) GetData() KlarnaGatewayData {
	if o == nil || IsNil(o.Data) {
		var ret KlarnaGatewayData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaGateway) GetDataOk() (*KlarnaGatewayData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KlarnaGateway) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given KlarnaGatewayData and assigns it to the Data field.
func (o *KlarnaGateway) SetData(v KlarnaGatewayData) {
	o.Data = &v
}

func (o KlarnaGateway) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlarnaGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableKlarnaGateway struct {
	value *KlarnaGateway
	isSet bool
}

func (v NullableKlarnaGateway) Get() *KlarnaGateway {
	return v.value
}

func (v *NullableKlarnaGateway) Set(val *KlarnaGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaGateway(val *KlarnaGateway) *NullableKlarnaGateway {
	return &NullableKlarnaGateway{value: val, isSet: true}
}

func (v NullableKlarnaGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
