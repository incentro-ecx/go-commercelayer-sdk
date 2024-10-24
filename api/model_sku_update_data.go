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

// checks if the SkuUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuUpdateData{}

// SkuUpdateData struct for SkuUpdateData
type SkuUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                             `json:"id"`
	Attributes    PATCHSkusSkuId200ResponseDataAttributes `json:"attributes"`
	Relationships *SkuUpdateDataRelationships             `json:"relationships,omitempty"`
}

// NewSkuUpdateData instantiates a new SkuUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuUpdateData(type_ interface{}, id interface{}, attributes PATCHSkusSkuId200ResponseDataAttributes) *SkuUpdateData {
	this := SkuUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewSkuUpdateDataWithDefaults instantiates a new SkuUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuUpdateDataWithDefaults() *SkuUpdateData {
	this := SkuUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SkuUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkuUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SkuUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkuUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SkuUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *SkuUpdateData) GetAttributes() PATCHSkusSkuId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHSkusSkuId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuUpdateData) GetAttributesOk() (*PATCHSkusSkuId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuUpdateData) SetAttributes(v PATCHSkusSkuId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuUpdateData) GetRelationships() SkuUpdateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SkuUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateData) GetRelationshipsOk() (*SkuUpdateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuUpdateDataRelationships and assigns it to the Relationships field.
func (o *SkuUpdateData) SetRelationships(v SkuUpdateDataRelationships) {
	o.Relationships = &v
}

func (o SkuUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableSkuUpdateData struct {
	value *SkuUpdateData
	isSet bool
}

func (v NullableSkuUpdateData) Get() *SkuUpdateData {
	return v.value
}

func (v *NullableSkuUpdateData) Set(val *SkuUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuUpdateData(val *SkuUpdateData) *NullableSkuUpdateData {
	return &NullableSkuUpdateData{value: val, isSet: true}
}

func (v NullableSkuUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
