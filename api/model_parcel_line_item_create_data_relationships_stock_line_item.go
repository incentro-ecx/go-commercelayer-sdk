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

// ParcelLineItemCreateDataRelationshipsStockLineItem struct for ParcelLineItemCreateDataRelationshipsStockLineItem
type ParcelLineItemCreateDataRelationshipsStockLineItem struct {
	Data LineItemDataRelationshipsStockLineItemsData `json:"data"`
}

// NewParcelLineItemCreateDataRelationshipsStockLineItem instantiates a new ParcelLineItemCreateDataRelationshipsStockLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelLineItemCreateDataRelationshipsStockLineItem(data LineItemDataRelationshipsStockLineItemsData) *ParcelLineItemCreateDataRelationshipsStockLineItem {
	this := ParcelLineItemCreateDataRelationshipsStockLineItem{}
	this.Data = data
	return &this
}

// NewParcelLineItemCreateDataRelationshipsStockLineItemWithDefaults instantiates a new ParcelLineItemCreateDataRelationshipsStockLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelLineItemCreateDataRelationshipsStockLineItemWithDefaults() *ParcelLineItemCreateDataRelationshipsStockLineItem {
	this := ParcelLineItemCreateDataRelationshipsStockLineItem{}
	return &this
}

// GetData returns the Data field value
func (o *ParcelLineItemCreateDataRelationshipsStockLineItem) GetData() LineItemDataRelationshipsStockLineItemsData {
	if o == nil {
		var ret LineItemDataRelationshipsStockLineItemsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ParcelLineItemCreateDataRelationshipsStockLineItem) GetDataOk() (*LineItemDataRelationshipsStockLineItemsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ParcelLineItemCreateDataRelationshipsStockLineItem) SetData(v LineItemDataRelationshipsStockLineItemsData) {
	o.Data = v
}

func (o ParcelLineItemCreateDataRelationshipsStockLineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableParcelLineItemCreateDataRelationshipsStockLineItem struct {
	value *ParcelLineItemCreateDataRelationshipsStockLineItem
	isSet bool
}

func (v NullableParcelLineItemCreateDataRelationshipsStockLineItem) Get() *ParcelLineItemCreateDataRelationshipsStockLineItem {
	return v.value
}

func (v *NullableParcelLineItemCreateDataRelationshipsStockLineItem) Set(val *ParcelLineItemCreateDataRelationshipsStockLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelLineItemCreateDataRelationshipsStockLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelLineItemCreateDataRelationshipsStockLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelLineItemCreateDataRelationshipsStockLineItem(val *ParcelLineItemCreateDataRelationshipsStockLineItem) *NullableParcelLineItemCreateDataRelationshipsStockLineItem {
	return &NullableParcelLineItemCreateDataRelationshipsStockLineItem{value: val, isSet: true}
}

func (v NullableParcelLineItemCreateDataRelationshipsStockLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelLineItemCreateDataRelationshipsStockLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
