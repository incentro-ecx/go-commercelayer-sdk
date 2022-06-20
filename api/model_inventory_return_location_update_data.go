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

// InventoryReturnLocationUpdateData struct for InventoryReturnLocationUpdateData
type InventoryReturnLocationUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                      `json:"id"`
	Attributes    InventoryReturnLocationUpdateDataAttributes `json:"attributes"`
	Relationships *InventoryReturnLocationDataRelationships   `json:"relationships,omitempty"`
}

// NewInventoryReturnLocationUpdateData instantiates a new InventoryReturnLocationUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocationUpdateData(type_ string, id string, attributes InventoryReturnLocationUpdateDataAttributes) *InventoryReturnLocationUpdateData {
	this := InventoryReturnLocationUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewInventoryReturnLocationUpdateDataWithDefaults instantiates a new InventoryReturnLocationUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReturnLocationUpdateDataWithDefaults() *InventoryReturnLocationUpdateData {
	this := InventoryReturnLocationUpdateData{}
	var type_ string = "inventory_return_locations"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InventoryReturnLocationUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryReturnLocationUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InventoryReturnLocationUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InventoryReturnLocationUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *InventoryReturnLocationUpdateData) GetAttributes() InventoryReturnLocationUpdateDataAttributes {
	if o == nil {
		var ret InventoryReturnLocationUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationUpdateData) GetAttributesOk() (*InventoryReturnLocationUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InventoryReturnLocationUpdateData) SetAttributes(v InventoryReturnLocationUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InventoryReturnLocationUpdateData) GetRelationships() InventoryReturnLocationDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret InventoryReturnLocationDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationUpdateData) GetRelationshipsOk() (*InventoryReturnLocationDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InventoryReturnLocationUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InventoryReturnLocationDataRelationships and assigns it to the Relationships field.
func (o *InventoryReturnLocationUpdateData) SetRelationships(v InventoryReturnLocationDataRelationships) {
	o.Relationships = &v
}

func (o InventoryReturnLocationUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryReturnLocationUpdateData struct {
	value *InventoryReturnLocationUpdateData
	isSet bool
}

func (v NullableInventoryReturnLocationUpdateData) Get() *InventoryReturnLocationUpdateData {
	return v.value
}

func (v *NullableInventoryReturnLocationUpdateData) Set(val *InventoryReturnLocationUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReturnLocationUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReturnLocationUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReturnLocationUpdateData(val *InventoryReturnLocationUpdateData) *NullableInventoryReturnLocationUpdateData {
	return &NullableInventoryReturnLocationUpdateData{value: val, isSet: true}
}

func (v NullableInventoryReturnLocationUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReturnLocationUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}