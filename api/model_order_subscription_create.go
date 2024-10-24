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

// checks if the OrderSubscriptionCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSubscriptionCreate{}

// OrderSubscriptionCreate struct for OrderSubscriptionCreate
type OrderSubscriptionCreate struct {
	Data OrderSubscriptionCreateData `json:"data"`
}

// NewOrderSubscriptionCreate instantiates a new OrderSubscriptionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionCreate(data OrderSubscriptionCreateData) *OrderSubscriptionCreate {
	this := OrderSubscriptionCreate{}
	this.Data = data
	return &this
}

// NewOrderSubscriptionCreateWithDefaults instantiates a new OrderSubscriptionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionCreateWithDefaults() *OrderSubscriptionCreate {
	this := OrderSubscriptionCreate{}
	return &this
}

// GetData returns the Data field value
func (o *OrderSubscriptionCreate) GetData() OrderSubscriptionCreateData {
	if o == nil {
		var ret OrderSubscriptionCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionCreate) GetDataOk() (*OrderSubscriptionCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderSubscriptionCreate) SetData(v OrderSubscriptionCreateData) {
	o.Data = v
}

func (o OrderSubscriptionCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSubscriptionCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableOrderSubscriptionCreate struct {
	value *OrderSubscriptionCreate
	isSet bool
}

func (v NullableOrderSubscriptionCreate) Get() *OrderSubscriptionCreate {
	return v.value
}

func (v *NullableOrderSubscriptionCreate) Set(val *OrderSubscriptionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionCreate(val *OrderSubscriptionCreate) *NullableOrderSubscriptionCreate {
	return &NullableOrderSubscriptionCreate{value: val, isSet: true}
}

func (v NullableOrderSubscriptionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
