/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LineItemOptionCreateData struct for LineItemOptionCreateData
type LineItemOptionCreateData struct {
	// The resource's type
	Type          string                                       `json:"type"`
	Attributes    POSTLineItemOptions201ResponseDataAttributes `json:"attributes"`
	Relationships *LineItemOptionCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewLineItemOptionCreateData instantiates a new LineItemOptionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionCreateData(type_ string, attributes POSTLineItemOptions201ResponseDataAttributes) *LineItemOptionCreateData {
	this := LineItemOptionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewLineItemOptionCreateDataWithDefaults instantiates a new LineItemOptionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionCreateDataWithDefaults() *LineItemOptionCreateData {
	this := LineItemOptionCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *LineItemOptionCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LineItemOptionCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *LineItemOptionCreateData) GetAttributes() POSTLineItemOptions201ResponseDataAttributes {
	if o == nil {
		var ret POSTLineItemOptions201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreateData) GetAttributesOk() (*POSTLineItemOptions201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *LineItemOptionCreateData) SetAttributes(v POSTLineItemOptions201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *LineItemOptionCreateData) GetRelationships() LineItemOptionCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret LineItemOptionCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreateData) GetRelationshipsOk() (*LineItemOptionCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *LineItemOptionCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given LineItemOptionCreateDataRelationships and assigns it to the Relationships field.
func (o *LineItemOptionCreateData) SetRelationships(v LineItemOptionCreateDataRelationships) {
	o.Relationships = &v
}

func (o LineItemOptionCreateData) MarshalJSON() ([]byte, error) {
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

type NullableLineItemOptionCreateData struct {
	value *LineItemOptionCreateData
	isSet bool
}

func (v NullableLineItemOptionCreateData) Get() *LineItemOptionCreateData {
	return v.value
}

func (v *NullableLineItemOptionCreateData) Set(val *LineItemOptionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionCreateData(val *LineItemOptionCreateData) *NullableLineItemOptionCreateData {
	return &NullableLineItemOptionCreateData{value: val, isSet: true}
}

func (v NullableLineItemOptionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
