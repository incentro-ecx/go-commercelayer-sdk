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

// GETLineItemOptions200ResponseDataInnerRelationships struct for GETLineItemOptions200ResponseDataInnerRelationships
type GETLineItemOptions200ResponseDataInnerRelationships struct {
	LineItem  *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem  `json:"line_item,omitempty"`
	SkuOption *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption `json:"sku_option,omitempty"`
}

// NewGETLineItemOptions200ResponseDataInnerRelationships instantiates a new GETLineItemOptions200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItemOptions200ResponseDataInnerRelationships() *GETLineItemOptions200ResponseDataInnerRelationships {
	this := GETLineItemOptions200ResponseDataInnerRelationships{}
	return &this
}

// NewGETLineItemOptions200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETLineItemOptions200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItemOptions200ResponseDataInnerRelationshipsWithDefaults() *GETLineItemOptions200ResponseDataInnerRelationships {
	this := GETLineItemOptions200ResponseDataInnerRelationships{}
	return &this
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *GETLineItemOptions200ResponseDataInnerRelationships) GetLineItem() GETLineItemOptions200ResponseDataInnerRelationshipsLineItem {
	if o == nil || o.LineItem == nil {
		var ret GETLineItemOptions200ResponseDataInnerRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationships) GetLineItemOk() (*GETLineItemOptions200ResponseDataInnerRelationshipsLineItem, bool) {
	if o == nil || o.LineItem == nil {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationships) HasLineItem() bool {
	if o != nil && o.LineItem != nil {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given GETLineItemOptions200ResponseDataInnerRelationshipsLineItem and assigns it to the LineItem field.
func (o *GETLineItemOptions200ResponseDataInnerRelationships) SetLineItem(v GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) {
	o.LineItem = &v
}

// GetSkuOption returns the SkuOption field value if set, zero value otherwise.
func (o *GETLineItemOptions200ResponseDataInnerRelationships) GetSkuOption() GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption {
	if o == nil || o.SkuOption == nil {
		var ret GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption
		return ret
	}
	return *o.SkuOption
}

// GetSkuOptionOk returns a tuple with the SkuOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationships) GetSkuOptionOk() (*GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption, bool) {
	if o == nil || o.SkuOption == nil {
		return nil, false
	}
	return o.SkuOption, true
}

// HasSkuOption returns a boolean if a field has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationships) HasSkuOption() bool {
	if o != nil && o.SkuOption != nil {
		return true
	}

	return false
}

// SetSkuOption gets a reference to the given GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption and assigns it to the SkuOption field.
func (o *GETLineItemOptions200ResponseDataInnerRelationships) SetSkuOption(v GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption) {
	o.SkuOption = &v
}

func (o GETLineItemOptions200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LineItem != nil {
		toSerialize["line_item"] = o.LineItem
	}
	if o.SkuOption != nil {
		toSerialize["sku_option"] = o.SkuOption
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItemOptions200ResponseDataInnerRelationships struct {
	value *GETLineItemOptions200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETLineItemOptions200ResponseDataInnerRelationships) Get() *GETLineItemOptions200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETLineItemOptions200ResponseDataInnerRelationships) Set(val *GETLineItemOptions200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItemOptions200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItemOptions200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItemOptions200ResponseDataInnerRelationships(val *GETLineItemOptions200ResponseDataInnerRelationships) *NullableGETLineItemOptions200ResponseDataInnerRelationships {
	return &NullableGETLineItemOptions200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETLineItemOptions200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItemOptions200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
