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

// BillingInfoValidationRuleResponseDataRelationships struct for BillingInfoValidationRuleResponseDataRelationships
type BillingInfoValidationRuleResponseDataRelationships struct {
	Market *BillingInfoValidationRuleResponseDataRelationshipsMarket `json:"market,omitempty"`
}

// NewBillingInfoValidationRuleResponseDataRelationships instantiates a new BillingInfoValidationRuleResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoValidationRuleResponseDataRelationships() *BillingInfoValidationRuleResponseDataRelationships {
	this := BillingInfoValidationRuleResponseDataRelationships{}
	return &this
}

// NewBillingInfoValidationRuleResponseDataRelationshipsWithDefaults instantiates a new BillingInfoValidationRuleResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoValidationRuleResponseDataRelationshipsWithDefaults() *BillingInfoValidationRuleResponseDataRelationships {
	this := BillingInfoValidationRuleResponseDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *BillingInfoValidationRuleResponseDataRelationships) GetMarket() BillingInfoValidationRuleResponseDataRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret BillingInfoValidationRuleResponseDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleResponseDataRelationships) GetMarketOk() (*BillingInfoValidationRuleResponseDataRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *BillingInfoValidationRuleResponseDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BillingInfoValidationRuleResponseDataRelationshipsMarket and assigns it to the Market field.
func (o *BillingInfoValidationRuleResponseDataRelationships) SetMarket(v BillingInfoValidationRuleResponseDataRelationshipsMarket) {
	o.Market = &v
}

func (o BillingInfoValidationRuleResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	return json.Marshal(toSerialize)
}

type NullableBillingInfoValidationRuleResponseDataRelationships struct {
	value *BillingInfoValidationRuleResponseDataRelationships
	isSet bool
}

func (v NullableBillingInfoValidationRuleResponseDataRelationships) Get() *BillingInfoValidationRuleResponseDataRelationships {
	return v.value
}

func (v *NullableBillingInfoValidationRuleResponseDataRelationships) Set(val *BillingInfoValidationRuleResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoValidationRuleResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoValidationRuleResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoValidationRuleResponseDataRelationships(val *BillingInfoValidationRuleResponseDataRelationships) *NullableBillingInfoValidationRuleResponseDataRelationships {
	return &NullableBillingInfoValidationRuleResponseDataRelationships{value: val, isSet: true}
}

func (v NullableBillingInfoValidationRuleResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoValidationRuleResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}