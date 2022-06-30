/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ShipmentDataRelationshipsShippingCategory struct for ShipmentDataRelationshipsShippingCategory
type ShipmentDataRelationshipsShippingCategory struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewShipmentDataRelationshipsShippingCategory instantiates a new ShipmentDataRelationshipsShippingCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDataRelationshipsShippingCategory(type_ string, id string) *ShipmentDataRelationshipsShippingCategory {
	this := ShipmentDataRelationshipsShippingCategory{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewShipmentDataRelationshipsShippingCategoryWithDefaults instantiates a new ShipmentDataRelationshipsShippingCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDataRelationshipsShippingCategoryWithDefaults() *ShipmentDataRelationshipsShippingCategory {
	this := ShipmentDataRelationshipsShippingCategory{}
	var type_ string = "shipping_categories"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ShipmentDataRelationshipsShippingCategory) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationshipsShippingCategory) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShipmentDataRelationshipsShippingCategory) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ShipmentDataRelationshipsShippingCategory) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationshipsShippingCategory) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ShipmentDataRelationshipsShippingCategory) SetId(v string) {
	o.Id = v
}

func (o ShipmentDataRelationshipsShippingCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableShipmentDataRelationshipsShippingCategory struct {
	value *ShipmentDataRelationshipsShippingCategory
	isSet bool
}

func (v NullableShipmentDataRelationshipsShippingCategory) Get() *ShipmentDataRelationshipsShippingCategory {
	return v.value
}

func (v *NullableShipmentDataRelationshipsShippingCategory) Set(val *ShipmentDataRelationshipsShippingCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDataRelationshipsShippingCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDataRelationshipsShippingCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDataRelationshipsShippingCategory(val *ShipmentDataRelationshipsShippingCategory) *NullableShipmentDataRelationshipsShippingCategory {
	return &NullableShipmentDataRelationshipsShippingCategory{value: val, isSet: true}
}

func (v NullableShipmentDataRelationshipsShippingCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDataRelationshipsShippingCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
