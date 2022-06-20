/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LineItemOptionDataRelationshipsLineItem struct for LineItemOptionDataRelationshipsLineItem
type LineItemOptionDataRelationshipsLineItem struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewLineItemOptionDataRelationshipsLineItem instantiates a new LineItemOptionDataRelationshipsLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionDataRelationshipsLineItem(type_ string, id string) *LineItemOptionDataRelationshipsLineItem {
	this := LineItemOptionDataRelationshipsLineItem{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewLineItemOptionDataRelationshipsLineItemWithDefaults instantiates a new LineItemOptionDataRelationshipsLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionDataRelationshipsLineItemWithDefaults() *LineItemOptionDataRelationshipsLineItem {
	this := LineItemOptionDataRelationshipsLineItem{}
	var type_ string = "line_items"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *LineItemOptionDataRelationshipsLineItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataRelationshipsLineItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LineItemOptionDataRelationshipsLineItem) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *LineItemOptionDataRelationshipsLineItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataRelationshipsLineItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LineItemOptionDataRelationshipsLineItem) SetId(v string) {
	o.Id = v
}

func (o LineItemOptionDataRelationshipsLineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemOptionDataRelationshipsLineItem struct {
	value *LineItemOptionDataRelationshipsLineItem
	isSet bool
}

func (v NullableLineItemOptionDataRelationshipsLineItem) Get() *LineItemOptionDataRelationshipsLineItem {
	return v.value
}

func (v *NullableLineItemOptionDataRelationshipsLineItem) Set(val *LineItemOptionDataRelationshipsLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionDataRelationshipsLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionDataRelationshipsLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionDataRelationshipsLineItem(val *LineItemOptionDataRelationshipsLineItem) *NullableLineItemOptionDataRelationshipsLineItem {
	return &NullableLineItemOptionDataRelationshipsLineItem{value: val, isSet: true}
}

func (v NullableLineItemOptionDataRelationshipsLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionDataRelationshipsLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}