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

// LineItemOptionCreateDataRelationships struct for LineItemOptionCreateDataRelationships
type LineItemOptionCreateDataRelationships struct {
	LineItem  LineItemOptionCreateDataRelationshipsLineItem  `json:"line_item"`
	SkuOption LineItemOptionCreateDataRelationshipsSkuOption `json:"sku_option"`
}

// NewLineItemOptionCreateDataRelationships instantiates a new LineItemOptionCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionCreateDataRelationships(lineItem LineItemOptionCreateDataRelationshipsLineItem, skuOption LineItemOptionCreateDataRelationshipsSkuOption) *LineItemOptionCreateDataRelationships {
	this := LineItemOptionCreateDataRelationships{}
	this.LineItem = lineItem
	this.SkuOption = skuOption
	return &this
}

// NewLineItemOptionCreateDataRelationshipsWithDefaults instantiates a new LineItemOptionCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionCreateDataRelationshipsWithDefaults() *LineItemOptionCreateDataRelationships {
	this := LineItemOptionCreateDataRelationships{}
	return &this
}

// GetLineItem returns the LineItem field value
func (o *LineItemOptionCreateDataRelationships) GetLineItem() LineItemOptionCreateDataRelationshipsLineItem {
	if o == nil {
		var ret LineItemOptionCreateDataRelationshipsLineItem
		return ret
	}

	return o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreateDataRelationships) GetLineItemOk() (*LineItemOptionCreateDataRelationshipsLineItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LineItem, true
}

// SetLineItem sets field value
func (o *LineItemOptionCreateDataRelationships) SetLineItem(v LineItemOptionCreateDataRelationshipsLineItem) {
	o.LineItem = v
}

// GetSkuOption returns the SkuOption field value
func (o *LineItemOptionCreateDataRelationships) GetSkuOption() LineItemOptionCreateDataRelationshipsSkuOption {
	if o == nil {
		var ret LineItemOptionCreateDataRelationshipsSkuOption
		return ret
	}

	return o.SkuOption
}

// GetSkuOptionOk returns a tuple with the SkuOption field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreateDataRelationships) GetSkuOptionOk() (*LineItemOptionCreateDataRelationshipsSkuOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkuOption, true
}

// SetSkuOption sets field value
func (o *LineItemOptionCreateDataRelationships) SetSkuOption(v LineItemOptionCreateDataRelationshipsSkuOption) {
	o.SkuOption = v
}

func (o LineItemOptionCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["line_item"] = o.LineItem
	}
	if true {
		toSerialize["sku_option"] = o.SkuOption
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemOptionCreateDataRelationships struct {
	value *LineItemOptionCreateDataRelationships
	isSet bool
}

func (v NullableLineItemOptionCreateDataRelationships) Get() *LineItemOptionCreateDataRelationships {
	return v.value
}

func (v *NullableLineItemOptionCreateDataRelationships) Set(val *LineItemOptionCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionCreateDataRelationships(val *LineItemOptionCreateDataRelationships) *NullableLineItemOptionCreateDataRelationships {
	return &NullableLineItemOptionCreateDataRelationships{value: val, isSet: true}
}

func (v NullableLineItemOptionCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
