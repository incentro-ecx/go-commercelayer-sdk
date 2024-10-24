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

// checks if the StockLineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockLineItem{}

// StockLineItem struct for StockLineItem
type StockLineItem struct {
	Data *StockLineItemData `json:"data,omitempty"`
}

// NewStockLineItem instantiates a new StockLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockLineItem() *StockLineItem {
	this := StockLineItem{}
	return &this
}

// NewStockLineItemWithDefaults instantiates a new StockLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockLineItemWithDefaults() *StockLineItem {
	this := StockLineItem{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StockLineItem) GetData() StockLineItemData {
	if o == nil || IsNil(o.Data) {
		var ret StockLineItemData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLineItem) GetDataOk() (*StockLineItemData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StockLineItem) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given StockLineItemData and assigns it to the Data field.
func (o *StockLineItem) SetData(v StockLineItemData) {
	o.Data = &v
}

func (o StockLineItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockLineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableStockLineItem struct {
	value *StockLineItem
	isSet bool
}

func (v NullableStockLineItem) Get() *StockLineItem {
	return v.value
}

func (v *NullableStockLineItem) Set(val *StockLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStockLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStockLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockLineItem(val *StockLineItem) *NullableStockLineItem {
	return &NullableStockLineItem{value: val, isSet: true}
}

func (v NullableStockLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
