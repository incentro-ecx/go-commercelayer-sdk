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

// checks if the BundleCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleCreateData{}

// BundleCreateData struct for BundleCreateData
type BundleCreateData struct {
	// The resource's type
	Type          interface{}                          `json:"type"`
	Attributes    POSTBundles201ResponseDataAttributes `json:"attributes"`
	Relationships *BundleCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewBundleCreateData instantiates a new BundleCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleCreateData(type_ interface{}, attributes POSTBundles201ResponseDataAttributes) *BundleCreateData {
	this := BundleCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBundleCreateDataWithDefaults instantiates a new BundleCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleCreateDataWithDefaults() *BundleCreateData {
	this := BundleCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *BundleCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BundleCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BundleCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BundleCreateData) GetAttributes() POSTBundles201ResponseDataAttributes {
	if o == nil {
		var ret POSTBundles201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BundleCreateData) GetAttributesOk() (*POSTBundles201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BundleCreateData) SetAttributes(v POSTBundles201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BundleCreateData) GetRelationships() BundleCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BundleCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleCreateData) GetRelationshipsOk() (*BundleCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BundleCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BundleCreateDataRelationships and assigns it to the Relationships field.
func (o *BundleCreateData) SetRelationships(v BundleCreateDataRelationships) {
	o.Relationships = &v
}

func (o BundleCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableBundleCreateData struct {
	value *BundleCreateData
	isSet bool
}

func (v NullableBundleCreateData) Get() *BundleCreateData {
	return v.value
}

func (v *NullableBundleCreateData) Set(val *BundleCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleCreateData(val *BundleCreateData) *NullableBundleCreateData {
	return &NullableBundleCreateData{value: val, isSet: true}
}

func (v NullableBundleCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
