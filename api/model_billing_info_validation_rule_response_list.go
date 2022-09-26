/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BillingInfoValidationRuleResponseList struct for BillingInfoValidationRuleResponseList
type BillingInfoValidationRuleResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewBillingInfoValidationRuleResponseList instantiates a new BillingInfoValidationRuleResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoValidationRuleResponseList() *BillingInfoValidationRuleResponseList {
	this := BillingInfoValidationRuleResponseList{}
	return &this
}

// NewBillingInfoValidationRuleResponseListWithDefaults instantiates a new BillingInfoValidationRuleResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoValidationRuleResponseListWithDefaults() *BillingInfoValidationRuleResponseList {
	this := BillingInfoValidationRuleResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BillingInfoValidationRuleResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BillingInfoValidationRuleResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *BillingInfoValidationRuleResponseList) SetData(v []Data) {
	o.Data = v
}

func (o BillingInfoValidationRuleResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBillingInfoValidationRuleResponseList struct {
	value *BillingInfoValidationRuleResponseList
	isSet bool
}

func (v NullableBillingInfoValidationRuleResponseList) Get() *BillingInfoValidationRuleResponseList {
	return v.value
}

func (v *NullableBillingInfoValidationRuleResponseList) Set(val *BillingInfoValidationRuleResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoValidationRuleResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoValidationRuleResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoValidationRuleResponseList(val *BillingInfoValidationRuleResponseList) *NullableBillingInfoValidationRuleResponseList {
	return &NullableBillingInfoValidationRuleResponseList{value: val, isSet: true}
}

func (v NullableBillingInfoValidationRuleResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoValidationRuleResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}