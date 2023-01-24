/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CustomerSubscription struct for CustomerSubscription
type CustomerSubscription struct {
	Data *CustomerSubscriptionData `json:"data,omitempty"`
}

// NewCustomerSubscription instantiates a new CustomerSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSubscription() *CustomerSubscription {
	this := CustomerSubscription{}
	return &this
}

// NewCustomerSubscriptionWithDefaults instantiates a new CustomerSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSubscriptionWithDefaults() *CustomerSubscription {
	this := CustomerSubscription{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerSubscription) GetData() CustomerSubscriptionData {
	if o == nil || o.Data == nil {
		var ret CustomerSubscriptionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscription) GetDataOk() (*CustomerSubscriptionData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerSubscription) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CustomerSubscriptionData and assigns it to the Data field.
func (o *CustomerSubscription) SetData(v CustomerSubscriptionData) {
	o.Data = &v
}

func (o CustomerSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerSubscription struct {
	value *CustomerSubscription
	isSet bool
}

func (v NullableCustomerSubscription) Get() *CustomerSubscription {
	return v.value
}

func (v *NullableCustomerSubscription) Set(val *CustomerSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSubscription(val *CustomerSubscription) *NullableCustomerSubscription {
	return &NullableCustomerSubscription{value: val, isSet: true}
}

func (v NullableCustomerSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
