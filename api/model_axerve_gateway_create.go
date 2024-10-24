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

// checks if the AxerveGatewayCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AxerveGatewayCreate{}

// AxerveGatewayCreate struct for AxerveGatewayCreate
type AxerveGatewayCreate struct {
	Data AxerveGatewayCreateData `json:"data"`
}

// NewAxerveGatewayCreate instantiates a new AxerveGatewayCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAxerveGatewayCreate(data AxerveGatewayCreateData) *AxerveGatewayCreate {
	this := AxerveGatewayCreate{}
	this.Data = data
	return &this
}

// NewAxerveGatewayCreateWithDefaults instantiates a new AxerveGatewayCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAxerveGatewayCreateWithDefaults() *AxerveGatewayCreate {
	this := AxerveGatewayCreate{}
	return &this
}

// GetData returns the Data field value
func (o *AxerveGatewayCreate) GetData() AxerveGatewayCreateData {
	if o == nil {
		var ret AxerveGatewayCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AxerveGatewayCreate) GetDataOk() (*AxerveGatewayCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AxerveGatewayCreate) SetData(v AxerveGatewayCreateData) {
	o.Data = v
}

func (o AxerveGatewayCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AxerveGatewayCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAxerveGatewayCreate struct {
	value *AxerveGatewayCreate
	isSet bool
}

func (v NullableAxerveGatewayCreate) Get() *AxerveGatewayCreate {
	return v.value
}

func (v *NullableAxerveGatewayCreate) Set(val *AxerveGatewayCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAxerveGatewayCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAxerveGatewayCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAxerveGatewayCreate(val *AxerveGatewayCreate) *NullableAxerveGatewayCreate {
	return &NullableAxerveGatewayCreate{value: val, isSet: true}
}

func (v NullableAxerveGatewayCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAxerveGatewayCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
