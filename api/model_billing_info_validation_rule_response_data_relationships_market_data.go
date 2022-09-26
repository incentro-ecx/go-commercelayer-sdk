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

// BillingInfoValidationRuleResponseDataRelationshipsMarketData struct for BillingInfoValidationRuleResponseDataRelationshipsMarketData
type BillingInfoValidationRuleResponseDataRelationshipsMarketData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewBillingInfoValidationRuleResponseDataRelationshipsMarketData instantiates a new BillingInfoValidationRuleResponseDataRelationshipsMarketData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoValidationRuleResponseDataRelationshipsMarketData() *BillingInfoValidationRuleResponseDataRelationshipsMarketData {
	this := BillingInfoValidationRuleResponseDataRelationshipsMarketData{}
	return &this
}

// NewBillingInfoValidationRuleResponseDataRelationshipsMarketDataWithDefaults instantiates a new BillingInfoValidationRuleResponseDataRelationshipsMarketData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoValidationRuleResponseDataRelationshipsMarketDataWithDefaults() *BillingInfoValidationRuleResponseDataRelationshipsMarketData {
	this := BillingInfoValidationRuleResponseDataRelationshipsMarketData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingInfoValidationRuleResponseDataRelationshipsMarketData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleResponseDataRelationshipsMarketData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingInfoValidationRuleResponseDataRelationshipsMarketData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BillingInfoValidationRuleResponseDataRelationshipsMarketData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingInfoValidationRuleResponseDataRelationshipsMarketData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleResponseDataRelationshipsMarketData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingInfoValidationRuleResponseDataRelationshipsMarketData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BillingInfoValidationRuleResponseDataRelationshipsMarketData) SetId(v string) {
	o.Id = &v
}

func (o BillingInfoValidationRuleResponseDataRelationshipsMarketData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBillingInfoValidationRuleResponseDataRelationshipsMarketData struct {
	value *BillingInfoValidationRuleResponseDataRelationshipsMarketData
	isSet bool
}

func (v NullableBillingInfoValidationRuleResponseDataRelationshipsMarketData) Get() *BillingInfoValidationRuleResponseDataRelationshipsMarketData {
	return v.value
}

func (v *NullableBillingInfoValidationRuleResponseDataRelationshipsMarketData) Set(val *BillingInfoValidationRuleResponseDataRelationshipsMarketData) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoValidationRuleResponseDataRelationshipsMarketData) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoValidationRuleResponseDataRelationshipsMarketData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoValidationRuleResponseDataRelationshipsMarketData(val *BillingInfoValidationRuleResponseDataRelationshipsMarketData) *NullableBillingInfoValidationRuleResponseDataRelationshipsMarketData {
	return &NullableBillingInfoValidationRuleResponseDataRelationshipsMarketData{value: val, isSet: true}
}

func (v NullableBillingInfoValidationRuleResponseDataRelationshipsMarketData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoValidationRuleResponseDataRelationshipsMarketData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}