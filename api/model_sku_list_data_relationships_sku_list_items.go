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

// SkuListDataRelationshipsSkuListItems struct for SkuListDataRelationshipsSkuListItems
type SkuListDataRelationshipsSkuListItems struct {
	Data *SkuListDataRelationshipsSkuListItemsData `json:"data,omitempty"`
}

// NewSkuListDataRelationshipsSkuListItems instantiates a new SkuListDataRelationshipsSkuListItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListDataRelationshipsSkuListItems() *SkuListDataRelationshipsSkuListItems {
	this := SkuListDataRelationshipsSkuListItems{}
	return &this
}

// NewSkuListDataRelationshipsSkuListItemsWithDefaults instantiates a new SkuListDataRelationshipsSkuListItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListDataRelationshipsSkuListItemsWithDefaults() *SkuListDataRelationshipsSkuListItems {
	this := SkuListDataRelationshipsSkuListItems{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SkuListDataRelationshipsSkuListItems) GetData() SkuListDataRelationshipsSkuListItemsData {
	if o == nil || o.Data == nil {
		var ret SkuListDataRelationshipsSkuListItemsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListDataRelationshipsSkuListItems) GetDataOk() (*SkuListDataRelationshipsSkuListItemsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SkuListDataRelationshipsSkuListItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given SkuListDataRelationshipsSkuListItemsData and assigns it to the Data field.
func (o *SkuListDataRelationshipsSkuListItems) SetData(v SkuListDataRelationshipsSkuListItemsData) {
	o.Data = &v
}

func (o SkuListDataRelationshipsSkuListItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSkuListDataRelationshipsSkuListItems struct {
	value *SkuListDataRelationshipsSkuListItems
	isSet bool
}

func (v NullableSkuListDataRelationshipsSkuListItems) Get() *SkuListDataRelationshipsSkuListItems {
	return v.value
}

func (v *NullableSkuListDataRelationshipsSkuListItems) Set(val *SkuListDataRelationshipsSkuListItems) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListDataRelationshipsSkuListItems) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListDataRelationshipsSkuListItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListDataRelationshipsSkuListItems(val *SkuListDataRelationshipsSkuListItems) *NullableSkuListDataRelationshipsSkuListItems {
	return &NullableSkuListDataRelationshipsSkuListItems{value: val, isSet: true}
}

func (v NullableSkuListDataRelationshipsSkuListItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListDataRelationshipsSkuListItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
