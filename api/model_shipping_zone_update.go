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

// checks if the ShippingZoneUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingZoneUpdate{}

// ShippingZoneUpdate struct for ShippingZoneUpdate
type ShippingZoneUpdate struct {
	Data ShippingZoneUpdateData `json:"data"`
}

// NewShippingZoneUpdate instantiates a new ShippingZoneUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingZoneUpdate(data ShippingZoneUpdateData) *ShippingZoneUpdate {
	this := ShippingZoneUpdate{}
	this.Data = data
	return &this
}

// NewShippingZoneUpdateWithDefaults instantiates a new ShippingZoneUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingZoneUpdateWithDefaults() *ShippingZoneUpdate {
	this := ShippingZoneUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ShippingZoneUpdate) GetData() ShippingZoneUpdateData {
	if o == nil {
		var ret ShippingZoneUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShippingZoneUpdate) GetDataOk() (*ShippingZoneUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShippingZoneUpdate) SetData(v ShippingZoneUpdateData) {
	o.Data = v
}

func (o ShippingZoneUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingZoneUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableShippingZoneUpdate struct {
	value *ShippingZoneUpdate
	isSet bool
}

func (v NullableShippingZoneUpdate) Get() *ShippingZoneUpdate {
	return v.value
}

func (v *NullableShippingZoneUpdate) Set(val *ShippingZoneUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingZoneUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingZoneUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingZoneUpdate(val *ShippingZoneUpdate) *NullableShippingZoneUpdate {
	return &NullableShippingZoneUpdate{value: val, isSet: true}
}

func (v NullableShippingZoneUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingZoneUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
