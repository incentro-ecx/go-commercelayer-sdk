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

// checks if the Customer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Customer{}

// Customer struct for Customer
type Customer struct {
	Data *CustomerData `json:"data,omitempty"`
}

// NewCustomer instantiates a new Customer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomer() *Customer {
	this := Customer{}
	return &this
}

// NewCustomerWithDefaults instantiates a new Customer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerWithDefaults() *Customer {
	this := Customer{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Customer) GetData() CustomerData {
	if o == nil || IsNil(o.Data) {
		var ret CustomerData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetDataOk() (*CustomerData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Customer) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CustomerData and assigns it to the Data field.
func (o *Customer) SetData(v CustomerData) {
	o.Data = &v
}

func (o Customer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Customer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCustomer struct {
	value *Customer
	isSet bool
}

func (v NullableCustomer) Get() *Customer {
	return v.value
}

func (v *NullableCustomer) Set(val *Customer) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomer(val *Customer) *NullableCustomer {
	return &NullableCustomer{value: val, isSet: true}
}

func (v NullableCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
