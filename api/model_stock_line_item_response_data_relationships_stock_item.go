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

// StockLineItemResponseDataRelationshipsStockItem struct for StockLineItemResponseDataRelationshipsStockItem
type StockLineItemResponseDataRelationshipsStockItem struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks       `json:"links,omitempty"`
	Data  *StockLineItemResponseDataRelationshipsStockItemData `json:"data,omitempty"`
}

// NewStockLineItemResponseDataRelationshipsStockItem instantiates a new StockLineItemResponseDataRelationshipsStockItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockLineItemResponseDataRelationshipsStockItem() *StockLineItemResponseDataRelationshipsStockItem {
	this := StockLineItemResponseDataRelationshipsStockItem{}
	return &this
}

// NewStockLineItemResponseDataRelationshipsStockItemWithDefaults instantiates a new StockLineItemResponseDataRelationshipsStockItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockLineItemResponseDataRelationshipsStockItemWithDefaults() *StockLineItemResponseDataRelationshipsStockItem {
	this := StockLineItemResponseDataRelationshipsStockItem{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *StockLineItemResponseDataRelationshipsStockItem) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLineItemResponseDataRelationshipsStockItem) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StockLineItemResponseDataRelationshipsStockItem) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *StockLineItemResponseDataRelationshipsStockItem) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StockLineItemResponseDataRelationshipsStockItem) GetData() StockLineItemResponseDataRelationshipsStockItemData {
	if o == nil || o.Data == nil {
		var ret StockLineItemResponseDataRelationshipsStockItemData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLineItemResponseDataRelationshipsStockItem) GetDataOk() (*StockLineItemResponseDataRelationshipsStockItemData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StockLineItemResponseDataRelationshipsStockItem) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given StockLineItemResponseDataRelationshipsStockItemData and assigns it to the Data field.
func (o *StockLineItemResponseDataRelationshipsStockItem) SetData(v StockLineItemResponseDataRelationshipsStockItemData) {
	o.Data = &v
}

func (o StockLineItemResponseDataRelationshipsStockItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStockLineItemResponseDataRelationshipsStockItem struct {
	value *StockLineItemResponseDataRelationshipsStockItem
	isSet bool
}

func (v NullableStockLineItemResponseDataRelationshipsStockItem) Get() *StockLineItemResponseDataRelationshipsStockItem {
	return v.value
}

func (v *NullableStockLineItemResponseDataRelationshipsStockItem) Set(val *StockLineItemResponseDataRelationshipsStockItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStockLineItemResponseDataRelationshipsStockItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStockLineItemResponseDataRelationshipsStockItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockLineItemResponseDataRelationshipsStockItem(val *StockLineItemResponseDataRelationshipsStockItem) *NullableStockLineItemResponseDataRelationshipsStockItem {
	return &NullableStockLineItemResponseDataRelationshipsStockItem{value: val, isSet: true}
}

func (v NullableStockLineItemResponseDataRelationshipsStockItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockLineItemResponseDataRelationshipsStockItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}