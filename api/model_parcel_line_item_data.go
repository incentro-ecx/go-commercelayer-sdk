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

// ParcelLineItemData struct for ParcelLineItemData
type ParcelLineItemData struct {
	// The resource's type
	Type          string                           `json:"type"`
	Attributes    ParcelLineItemDataAttributes     `json:"attributes"`
	Relationships *ParcelLineItemDataRelationships `json:"relationships,omitempty"`
}

// NewParcelLineItemData instantiates a new ParcelLineItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelLineItemData(type_ string, attributes ParcelLineItemDataAttributes) *ParcelLineItemData {
	this := ParcelLineItemData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewParcelLineItemDataWithDefaults instantiates a new ParcelLineItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelLineItemDataWithDefaults() *ParcelLineItemData {
	this := ParcelLineItemData{}
	var type_ string = "parcel_line_items"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ParcelLineItemData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ParcelLineItemData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ParcelLineItemData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ParcelLineItemData) GetAttributes() ParcelLineItemDataAttributes {
	if o == nil {
		var ret ParcelLineItemDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ParcelLineItemData) GetAttributesOk() (*ParcelLineItemDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ParcelLineItemData) SetAttributes(v ParcelLineItemDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ParcelLineItemData) GetRelationships() ParcelLineItemDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ParcelLineItemDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelLineItemData) GetRelationshipsOk() (*ParcelLineItemDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ParcelLineItemData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ParcelLineItemDataRelationships and assigns it to the Relationships field.
func (o *ParcelLineItemData) SetRelationships(v ParcelLineItemDataRelationships) {
	o.Relationships = &v
}

func (o ParcelLineItemData) MarshalJSON() ([]byte, error) {
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

type NullableParcelLineItemData struct {
	value *ParcelLineItemData
	isSet bool
}

func (v NullableParcelLineItemData) Get() *ParcelLineItemData {
	return v.value
}

func (v *NullableParcelLineItemData) Set(val *ParcelLineItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelLineItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelLineItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelLineItemData(val *ParcelLineItemData) *NullableParcelLineItemData {
	return &NullableParcelLineItemData{value: val, isSet: true}
}

func (v NullableParcelLineItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelLineItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
