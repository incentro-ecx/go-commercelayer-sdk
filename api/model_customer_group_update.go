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

// checks if the CustomerGroupUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerGroupUpdate{}

// CustomerGroupUpdate struct for CustomerGroupUpdate
type CustomerGroupUpdate struct {
	Data CustomerGroupUpdateData `json:"data"`
}

// NewCustomerGroupUpdate instantiates a new CustomerGroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerGroupUpdate(data CustomerGroupUpdateData) *CustomerGroupUpdate {
	this := CustomerGroupUpdate{}
	this.Data = data
	return &this
}

// NewCustomerGroupUpdateWithDefaults instantiates a new CustomerGroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerGroupUpdateWithDefaults() *CustomerGroupUpdate {
	this := CustomerGroupUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerGroupUpdate) GetData() CustomerGroupUpdateData {
	if o == nil {
		var ret CustomerGroupUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerGroupUpdate) GetDataOk() (*CustomerGroupUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerGroupUpdate) SetData(v CustomerGroupUpdateData) {
	o.Data = v
}

func (o CustomerGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerGroupUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCustomerGroupUpdate struct {
	value *CustomerGroupUpdate
	isSet bool
}

func (v NullableCustomerGroupUpdate) Get() *CustomerGroupUpdate {
	return v.value
}

func (v *NullableCustomerGroupUpdate) Set(val *CustomerGroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerGroupUpdate(val *CustomerGroupUpdate) *NullableCustomerGroupUpdate {
	return &NullableCustomerGroupUpdate{value: val, isSet: true}
}

func (v NullableCustomerGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
