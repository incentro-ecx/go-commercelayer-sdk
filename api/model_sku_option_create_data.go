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

// checks if the SkuOptionCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuOptionCreateData{}

// SkuOptionCreateData struct for SkuOptionCreateData
type SkuOptionCreateData struct {
	// The resource's type
	Type          interface{}                             `json:"type"`
	Attributes    POSTSkuOptions201ResponseDataAttributes `json:"attributes"`
	Relationships *SkuOptionCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewSkuOptionCreateData instantiates a new SkuOptionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuOptionCreateData(type_ interface{}, attributes POSTSkuOptions201ResponseDataAttributes) *SkuOptionCreateData {
	this := SkuOptionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuOptionCreateDataWithDefaults instantiates a new SkuOptionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuOptionCreateDataWithDefaults() *SkuOptionCreateData {
	this := SkuOptionCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SkuOptionCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkuOptionCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuOptionCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuOptionCreateData) GetAttributes() POSTSkuOptions201ResponseDataAttributes {
	if o == nil {
		var ret POSTSkuOptions201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuOptionCreateData) GetAttributesOk() (*POSTSkuOptions201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuOptionCreateData) SetAttributes(v POSTSkuOptions201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuOptionCreateData) GetRelationships() SkuOptionCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SkuOptionCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionCreateData) GetRelationshipsOk() (*SkuOptionCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuOptionCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuOptionCreateDataRelationships and assigns it to the Relationships field.
func (o *SkuOptionCreateData) SetRelationships(v SkuOptionCreateDataRelationships) {
	o.Relationships = &v
}

func (o SkuOptionCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuOptionCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableSkuOptionCreateData struct {
	value *SkuOptionCreateData
	isSet bool
}

func (v NullableSkuOptionCreateData) Get() *SkuOptionCreateData {
	return v.value
}

func (v *NullableSkuOptionCreateData) Set(val *SkuOptionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuOptionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuOptionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuOptionCreateData(val *SkuOptionCreateData) *NullableSkuOptionCreateData {
	return &NullableSkuOptionCreateData{value: val, isSet: true}
}

func (v NullableSkuOptionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuOptionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
