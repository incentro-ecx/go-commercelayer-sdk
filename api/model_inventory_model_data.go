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

// checks if the InventoryModelData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryModelData{}

// InventoryModelData struct for InventoryModelData
type InventoryModelData struct {
	// The resource's type
	Type          interface{}                                                 `json:"type"`
	Attributes    GETInventoryModelsInventoryModelId200ResponseDataAttributes `json:"attributes"`
	Relationships *InventoryModelDataRelationships                            `json:"relationships,omitempty"`
}

// NewInventoryModelData instantiates a new InventoryModelData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryModelData(type_ interface{}, attributes GETInventoryModelsInventoryModelId200ResponseDataAttributes) *InventoryModelData {
	this := InventoryModelData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInventoryModelDataWithDefaults instantiates a new InventoryModelData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryModelDataWithDefaults() *InventoryModelData {
	this := InventoryModelData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InventoryModelData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryModelData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryModelData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InventoryModelData) GetAttributes() GETInventoryModelsInventoryModelId200ResponseDataAttributes {
	if o == nil {
		var ret GETInventoryModelsInventoryModelId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InventoryModelData) GetAttributesOk() (*GETInventoryModelsInventoryModelId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InventoryModelData) SetAttributes(v GETInventoryModelsInventoryModelId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InventoryModelData) GetRelationships() InventoryModelDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret InventoryModelDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryModelData) GetRelationshipsOk() (*InventoryModelDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InventoryModelData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InventoryModelDataRelationships and assigns it to the Relationships field.
func (o *InventoryModelData) SetRelationships(v InventoryModelDataRelationships) {
	o.Relationships = &v
}

func (o InventoryModelData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryModelData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableInventoryModelData struct {
	value *InventoryModelData
	isSet bool
}

func (v NullableInventoryModelData) Get() *InventoryModelData {
	return v.value
}

func (v *NullableInventoryModelData) Set(val *InventoryModelData) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryModelData) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryModelData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryModelData(val *InventoryModelData) *NullableInventoryModelData {
	return &NullableInventoryModelData{value: val, isSet: true}
}

func (v NullableInventoryModelData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryModelData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
