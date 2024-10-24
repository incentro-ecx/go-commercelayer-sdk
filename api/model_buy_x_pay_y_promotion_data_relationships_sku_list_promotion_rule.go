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

// checks if the BuyXPayYPromotionDataRelationshipsSkuListPromotionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyXPayYPromotionDataRelationshipsSkuListPromotionRule{}

// BuyXPayYPromotionDataRelationshipsSkuListPromotionRule struct for BuyXPayYPromotionDataRelationshipsSkuListPromotionRule
type BuyXPayYPromotionDataRelationshipsSkuListPromotionRule struct {
	Data *BuyXPayYPromotionDataRelationshipsSkuListPromotionRuleData `json:"data,omitempty"`
}

// NewBuyXPayYPromotionDataRelationshipsSkuListPromotionRule instantiates a new BuyXPayYPromotionDataRelationshipsSkuListPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyXPayYPromotionDataRelationshipsSkuListPromotionRule() *BuyXPayYPromotionDataRelationshipsSkuListPromotionRule {
	this := BuyXPayYPromotionDataRelationshipsSkuListPromotionRule{}
	return &this
}

// NewBuyXPayYPromotionDataRelationshipsSkuListPromotionRuleWithDefaults instantiates a new BuyXPayYPromotionDataRelationshipsSkuListPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyXPayYPromotionDataRelationshipsSkuListPromotionRuleWithDefaults() *BuyXPayYPromotionDataRelationshipsSkuListPromotionRule {
	this := BuyXPayYPromotionDataRelationshipsSkuListPromotionRule{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BuyXPayYPromotionDataRelationshipsSkuListPromotionRule) GetData() BuyXPayYPromotionDataRelationshipsSkuListPromotionRuleData {
	if o == nil || IsNil(o.Data) {
		var ret BuyXPayYPromotionDataRelationshipsSkuListPromotionRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionDataRelationshipsSkuListPromotionRule) GetDataOk() (*BuyXPayYPromotionDataRelationshipsSkuListPromotionRuleData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BuyXPayYPromotionDataRelationshipsSkuListPromotionRule) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BuyXPayYPromotionDataRelationshipsSkuListPromotionRuleData and assigns it to the Data field.
func (o *BuyXPayYPromotionDataRelationshipsSkuListPromotionRule) SetData(v BuyXPayYPromotionDataRelationshipsSkuListPromotionRuleData) {
	o.Data = &v
}

func (o BuyXPayYPromotionDataRelationshipsSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyXPayYPromotionDataRelationshipsSkuListPromotionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBuyXPayYPromotionDataRelationshipsSkuListPromotionRule struct {
	value *BuyXPayYPromotionDataRelationshipsSkuListPromotionRule
	isSet bool
}

func (v NullableBuyXPayYPromotionDataRelationshipsSkuListPromotionRule) Get() *BuyXPayYPromotionDataRelationshipsSkuListPromotionRule {
	return v.value
}

func (v *NullableBuyXPayYPromotionDataRelationshipsSkuListPromotionRule) Set(val *BuyXPayYPromotionDataRelationshipsSkuListPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyXPayYPromotionDataRelationshipsSkuListPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyXPayYPromotionDataRelationshipsSkuListPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyXPayYPromotionDataRelationshipsSkuListPromotionRule(val *BuyXPayYPromotionDataRelationshipsSkuListPromotionRule) *NullableBuyXPayYPromotionDataRelationshipsSkuListPromotionRule {
	return &NullableBuyXPayYPromotionDataRelationshipsSkuListPromotionRule{value: val, isSet: true}
}

func (v NullableBuyXPayYPromotionDataRelationshipsSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyXPayYPromotionDataRelationshipsSkuListPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
