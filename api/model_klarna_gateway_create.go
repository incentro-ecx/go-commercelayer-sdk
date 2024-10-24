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

// checks if the KlarnaGatewayCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KlarnaGatewayCreate{}

// KlarnaGatewayCreate struct for KlarnaGatewayCreate
type KlarnaGatewayCreate struct {
	Data KlarnaGatewayCreateData `json:"data"`
}

// NewKlarnaGatewayCreate instantiates a new KlarnaGatewayCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaGatewayCreate(data KlarnaGatewayCreateData) *KlarnaGatewayCreate {
	this := KlarnaGatewayCreate{}
	this.Data = data
	return &this
}

// NewKlarnaGatewayCreateWithDefaults instantiates a new KlarnaGatewayCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaGatewayCreateWithDefaults() *KlarnaGatewayCreate {
	this := KlarnaGatewayCreate{}
	return &this
}

// GetData returns the Data field value
func (o *KlarnaGatewayCreate) GetData() KlarnaGatewayCreateData {
	if o == nil {
		var ret KlarnaGatewayCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayCreate) GetDataOk() (*KlarnaGatewayCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *KlarnaGatewayCreate) SetData(v KlarnaGatewayCreateData) {
	o.Data = v
}

func (o KlarnaGatewayCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlarnaGatewayCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableKlarnaGatewayCreate struct {
	value *KlarnaGatewayCreate
	isSet bool
}

func (v NullableKlarnaGatewayCreate) Get() *KlarnaGatewayCreate {
	return v.value
}

func (v *NullableKlarnaGatewayCreate) Set(val *KlarnaGatewayCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaGatewayCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaGatewayCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaGatewayCreate(val *KlarnaGatewayCreate) *NullableKlarnaGatewayCreate {
	return &NullableKlarnaGatewayCreate{value: val, isSet: true}
}

func (v NullableKlarnaGatewayCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaGatewayCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
