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

// checks if the LinkData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkData{}

// LinkData struct for LinkData
type LinkData struct {
	// The resource's type
	Type          interface{}                             `json:"type"`
	Attributes    GETLinksLinkId200ResponseDataAttributes `json:"attributes"`
	Relationships *LinkDataRelationships                  `json:"relationships,omitempty"`
}

// NewLinkData instantiates a new LinkData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkData(type_ interface{}, attributes GETLinksLinkId200ResponseDataAttributes) *LinkData {
	this := LinkData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewLinkDataWithDefaults instantiates a new LinkData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDataWithDefaults() *LinkData {
	this := LinkData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *LinkData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LinkData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *LinkData) GetAttributes() GETLinksLinkId200ResponseDataAttributes {
	if o == nil {
		var ret GETLinksLinkId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LinkData) GetAttributesOk() (*GETLinksLinkId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *LinkData) SetAttributes(v GETLinksLinkId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *LinkData) GetRelationships() LinkDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret LinkDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkData) GetRelationshipsOk() (*LinkDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *LinkData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given LinkDataRelationships and assigns it to the Relationships field.
func (o *LinkData) SetRelationships(v LinkDataRelationships) {
	o.Relationships = &v
}

func (o LinkData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkData) ToMap() (map[string]interface{}, error) {
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

type NullableLinkData struct {
	value *LinkData
	isSet bool
}

func (v NullableLinkData) Get() *LinkData {
	return v.value
}

func (v *NullableLinkData) Set(val *LinkData) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkData) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkData(val *LinkData) *NullableLinkData {
	return &NullableLinkData{value: val, isSet: true}
}

func (v NullableLinkData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
