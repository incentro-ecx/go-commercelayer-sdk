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

// InventoryStockLocationData struct for InventoryStockLocationData
type InventoryStockLocationData struct {
	// The resource's type
	Type          string                                    `json:"type"`
	Attributes    InventoryStockLocationDataAttributes      `json:"attributes"`
	Relationships *InventoryReturnLocationDataRelationships `json:"relationships,omitempty"`
}

// NewInventoryStockLocationData instantiates a new InventoryStockLocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryStockLocationData(type_ string, attributes InventoryStockLocationDataAttributes) *InventoryStockLocationData {
	this := InventoryStockLocationData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInventoryStockLocationDataWithDefaults instantiates a new InventoryStockLocationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryStockLocationDataWithDefaults() *InventoryStockLocationData {
	this := InventoryStockLocationData{}
	var type_ string = "inventory_stock_locations"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InventoryStockLocationData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryStockLocationData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InventoryStockLocationData) GetAttributes() InventoryStockLocationDataAttributes {
	if o == nil {
		var ret InventoryStockLocationDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationData) GetAttributesOk() (*InventoryStockLocationDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InventoryStockLocationData) SetAttributes(v InventoryStockLocationDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InventoryStockLocationData) GetRelationships() InventoryReturnLocationDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret InventoryReturnLocationDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationData) GetRelationshipsOk() (*InventoryReturnLocationDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InventoryStockLocationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InventoryReturnLocationDataRelationships and assigns it to the Relationships field.
func (o *InventoryStockLocationData) SetRelationships(v InventoryReturnLocationDataRelationships) {
	o.Relationships = &v
}

func (o InventoryStockLocationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryStockLocationData struct {
	value *InventoryStockLocationData
	isSet bool
}

func (v NullableInventoryStockLocationData) Get() *InventoryStockLocationData {
	return v.value
}

func (v *NullableInventoryStockLocationData) Set(val *InventoryStockLocationData) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryStockLocationData) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryStockLocationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryStockLocationData(val *InventoryStockLocationData) *NullableInventoryStockLocationData {
	return &NullableInventoryStockLocationData{value: val, isSet: true}
}

func (v NullableInventoryStockLocationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryStockLocationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
