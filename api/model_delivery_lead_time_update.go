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

// checks if the DeliveryLeadTimeUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryLeadTimeUpdate{}

// DeliveryLeadTimeUpdate struct for DeliveryLeadTimeUpdate
type DeliveryLeadTimeUpdate struct {
	Data DeliveryLeadTimeUpdateData `json:"data"`
}

// NewDeliveryLeadTimeUpdate instantiates a new DeliveryLeadTimeUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeUpdate(data DeliveryLeadTimeUpdateData) *DeliveryLeadTimeUpdate {
	this := DeliveryLeadTimeUpdate{}
	this.Data = data
	return &this
}

// NewDeliveryLeadTimeUpdateWithDefaults instantiates a new DeliveryLeadTimeUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeUpdateWithDefaults() *DeliveryLeadTimeUpdate {
	this := DeliveryLeadTimeUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *DeliveryLeadTimeUpdate) GetData() DeliveryLeadTimeUpdateData {
	if o == nil {
		var ret DeliveryLeadTimeUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeUpdate) GetDataOk() (*DeliveryLeadTimeUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeliveryLeadTimeUpdate) SetData(v DeliveryLeadTimeUpdateData) {
	o.Data = v
}

func (o DeliveryLeadTimeUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryLeadTimeUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableDeliveryLeadTimeUpdate struct {
	value *DeliveryLeadTimeUpdate
	isSet bool
}

func (v NullableDeliveryLeadTimeUpdate) Get() *DeliveryLeadTimeUpdate {
	return v.value
}

func (v *NullableDeliveryLeadTimeUpdate) Set(val *DeliveryLeadTimeUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeUpdate(val *DeliveryLeadTimeUpdate) *NullableDeliveryLeadTimeUpdate {
	return &NullableDeliveryLeadTimeUpdate{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
