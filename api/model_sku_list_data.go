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

// checks if the SkuListData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuListData{}

// SkuListData struct for SkuListData
type SkuListData struct {
	// The resource's type
	Type          interface{}                                   `json:"type"`
	Attributes    GETSkuListsSkuListId200ResponseDataAttributes `json:"attributes"`
	Relationships *SkuListDataRelationships                     `json:"relationships,omitempty"`
}

// NewSkuListData instantiates a new SkuListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListData(type_ interface{}, attributes GETSkuListsSkuListId200ResponseDataAttributes) *SkuListData {
	this := SkuListData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuListDataWithDefaults instantiates a new SkuListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListDataWithDefaults() *SkuListData {
	this := SkuListData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SkuListData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkuListData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuListData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuListData) GetAttributes() GETSkuListsSkuListId200ResponseDataAttributes {
	if o == nil {
		var ret GETSkuListsSkuListId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuListData) GetAttributesOk() (*GETSkuListsSkuListId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuListData) SetAttributes(v GETSkuListsSkuListId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuListData) GetRelationships() SkuListDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SkuListDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListData) GetRelationshipsOk() (*SkuListDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuListData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuListDataRelationships and assigns it to the Relationships field.
func (o *SkuListData) SetRelationships(v SkuListDataRelationships) {
	o.Relationships = &v
}

func (o SkuListData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuListData) ToMap() (map[string]interface{}, error) {
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

type NullableSkuListData struct {
	value *SkuListData
	isSet bool
}

func (v NullableSkuListData) Get() *SkuListData {
	return v.value
}

func (v *NullableSkuListData) Set(val *SkuListData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListData(val *SkuListData) *NullableSkuListData {
	return &NullableSkuListData{value: val, isSet: true}
}

func (v NullableSkuListData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
