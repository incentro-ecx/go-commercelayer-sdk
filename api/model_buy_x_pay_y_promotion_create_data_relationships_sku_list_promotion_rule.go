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

// checks if the BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule{}

// BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule struct for BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule
type BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule struct {
	Data BuyXPayYPromotionDataRelationshipsSkuListPromotionRuleData `json:"data"`
}

// NewBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule instantiates a new BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule(data BuyXPayYPromotionDataRelationshipsSkuListPromotionRuleData) *BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule {
	this := BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule{}
	this.Data = data
	return &this
}

// NewBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRuleWithDefaults instantiates a new BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRuleWithDefaults() *BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule {
	this := BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule{}
	return &this
}

// GetData returns the Data field value
func (o *BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) GetData() BuyXPayYPromotionDataRelationshipsSkuListPromotionRuleData {
	if o == nil {
		var ret BuyXPayYPromotionDataRelationshipsSkuListPromotionRuleData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) GetDataOk() (*BuyXPayYPromotionDataRelationshipsSkuListPromotionRuleData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) SetData(v BuyXPayYPromotionDataRelationshipsSkuListPromotionRuleData) {
	o.Data = v
}

func (o BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule struct {
	value *BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule
	isSet bool
}

func (v NullableBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) Get() *BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule {
	return v.value
}

func (v *NullableBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) Set(val *BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule(val *BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) *NullableBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule {
	return &NullableBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule{value: val, isSet: true}
}

func (v NullableBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
