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

// ShipmentDataRelationshipsShippingCategoryData struct for ShipmentDataRelationshipsShippingCategoryData
type ShipmentDataRelationshipsShippingCategoryData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource's id
	Id *string `json:"id,omitempty"`
}

// NewShipmentDataRelationshipsShippingCategoryData instantiates a new ShipmentDataRelationshipsShippingCategoryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDataRelationshipsShippingCategoryData() *ShipmentDataRelationshipsShippingCategoryData {
	this := ShipmentDataRelationshipsShippingCategoryData{}
	return &this
}

// NewShipmentDataRelationshipsShippingCategoryDataWithDefaults instantiates a new ShipmentDataRelationshipsShippingCategoryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDataRelationshipsShippingCategoryDataWithDefaults() *ShipmentDataRelationshipsShippingCategoryData {
	this := ShipmentDataRelationshipsShippingCategoryData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ShipmentDataRelationshipsShippingCategoryData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationshipsShippingCategoryData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ShipmentDataRelationshipsShippingCategoryData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ShipmentDataRelationshipsShippingCategoryData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShipmentDataRelationshipsShippingCategoryData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationshipsShippingCategoryData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShipmentDataRelationshipsShippingCategoryData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ShipmentDataRelationshipsShippingCategoryData) SetId(v string) {
	o.Id = &v
}

func (o ShipmentDataRelationshipsShippingCategoryData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableShipmentDataRelationshipsShippingCategoryData struct {
	value *ShipmentDataRelationshipsShippingCategoryData
	isSet bool
}

func (v NullableShipmentDataRelationshipsShippingCategoryData) Get() *ShipmentDataRelationshipsShippingCategoryData {
	return v.value
}

func (v *NullableShipmentDataRelationshipsShippingCategoryData) Set(val *ShipmentDataRelationshipsShippingCategoryData) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDataRelationshipsShippingCategoryData) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDataRelationshipsShippingCategoryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDataRelationshipsShippingCategoryData(val *ShipmentDataRelationshipsShippingCategoryData) *NullableShipmentDataRelationshipsShippingCategoryData {
	return &NullableShipmentDataRelationshipsShippingCategoryData{value: val, isSet: true}
}

func (v NullableShipmentDataRelationshipsShippingCategoryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDataRelationshipsShippingCategoryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}