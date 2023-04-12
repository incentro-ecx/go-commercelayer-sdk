/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the LineItemDataRelationshipsShipmentLineItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemDataRelationshipsShipmentLineItems{}

// LineItemDataRelationshipsShipmentLineItems struct for LineItemDataRelationshipsShipmentLineItems
type LineItemDataRelationshipsShipmentLineItems struct {
	Data *LineItemDataRelationshipsShipmentLineItemsData `json:"data,omitempty"`
}

// NewLineItemDataRelationshipsShipmentLineItems instantiates a new LineItemDataRelationshipsShipmentLineItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemDataRelationshipsShipmentLineItems() *LineItemDataRelationshipsShipmentLineItems {
	this := LineItemDataRelationshipsShipmentLineItems{}
	return &this
}

// NewLineItemDataRelationshipsShipmentLineItemsWithDefaults instantiates a new LineItemDataRelationshipsShipmentLineItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemDataRelationshipsShipmentLineItemsWithDefaults() *LineItemDataRelationshipsShipmentLineItems {
	this := LineItemDataRelationshipsShipmentLineItems{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LineItemDataRelationshipsShipmentLineItems) GetData() LineItemDataRelationshipsShipmentLineItemsData {
	if o == nil || IsNil(o.Data) {
		var ret LineItemDataRelationshipsShipmentLineItemsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemDataRelationshipsShipmentLineItems) GetDataOk() (*LineItemDataRelationshipsShipmentLineItemsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LineItemDataRelationshipsShipmentLineItems) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LineItemDataRelationshipsShipmentLineItemsData and assigns it to the Data field.
func (o *LineItemDataRelationshipsShipmentLineItems) SetData(v LineItemDataRelationshipsShipmentLineItemsData) {
	o.Data = &v
}

func (o LineItemDataRelationshipsShipmentLineItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemDataRelationshipsShipmentLineItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLineItemDataRelationshipsShipmentLineItems struct {
	value *LineItemDataRelationshipsShipmentLineItems
	isSet bool
}

func (v NullableLineItemDataRelationshipsShipmentLineItems) Get() *LineItemDataRelationshipsShipmentLineItems {
	return v.value
}

func (v *NullableLineItemDataRelationshipsShipmentLineItems) Set(val *LineItemDataRelationshipsShipmentLineItems) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemDataRelationshipsShipmentLineItems) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemDataRelationshipsShipmentLineItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemDataRelationshipsShipmentLineItems(val *LineItemDataRelationshipsShipmentLineItems) *NullableLineItemDataRelationshipsShipmentLineItems {
	return &NullableLineItemDataRelationshipsShipmentLineItems{value: val, isSet: true}
}

func (v NullableLineItemDataRelationshipsShipmentLineItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemDataRelationshipsShipmentLineItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
