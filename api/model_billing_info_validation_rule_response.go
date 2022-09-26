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

// BillingInfoValidationRuleResponse struct for BillingInfoValidationRuleResponse
type BillingInfoValidationRuleResponse struct {
	Data *BillingInfoValidationRuleResponseData `json:"data,omitempty"`
}

// NewBillingInfoValidationRuleResponse instantiates a new BillingInfoValidationRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoValidationRuleResponse() *BillingInfoValidationRuleResponse {
	this := BillingInfoValidationRuleResponse{}
	return &this
}

// NewBillingInfoValidationRuleResponseWithDefaults instantiates a new BillingInfoValidationRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoValidationRuleResponseWithDefaults() *BillingInfoValidationRuleResponse {
	this := BillingInfoValidationRuleResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BillingInfoValidationRuleResponse) GetData() BillingInfoValidationRuleResponseData {
	if o == nil || o.Data == nil {
		var ret BillingInfoValidationRuleResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleResponse) GetDataOk() (*BillingInfoValidationRuleResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BillingInfoValidationRuleResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given BillingInfoValidationRuleResponseData and assigns it to the Data field.
func (o *BillingInfoValidationRuleResponse) SetData(v BillingInfoValidationRuleResponseData) {
	o.Data = &v
}

func (o BillingInfoValidationRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBillingInfoValidationRuleResponse struct {
	value *BillingInfoValidationRuleResponse
	isSet bool
}

func (v NullableBillingInfoValidationRuleResponse) Get() *BillingInfoValidationRuleResponse {
	return v.value
}

func (v *NullableBillingInfoValidationRuleResponse) Set(val *BillingInfoValidationRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoValidationRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoValidationRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoValidationRuleResponse(val *BillingInfoValidationRuleResponse) *NullableBillingInfoValidationRuleResponse {
	return &NullableBillingInfoValidationRuleResponse{value: val, isSet: true}
}

func (v NullableBillingInfoValidationRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoValidationRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}