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

// PriceCreateData struct for PriceCreateData
type PriceCreateData struct {
	// The resource's type
	Type          string                        `json:"type"`
	Attributes    PriceCreateDataAttributes     `json:"attributes"`
	Relationships *PriceCreateDataRelationships `json:"relationships,omitempty"`
}

// NewPriceCreateData instantiates a new PriceCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceCreateData(type_ string, attributes PriceCreateDataAttributes) *PriceCreateData {
	this := PriceCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPriceCreateDataWithDefaults instantiates a new PriceCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceCreateDataWithDefaults() *PriceCreateData {
	this := PriceCreateData{}
	var type_ string = "prices"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *PriceCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PriceCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PriceCreateData) GetAttributes() PriceCreateDataAttributes {
	if o == nil {
		var ret PriceCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceCreateData) GetAttributesOk() (*PriceCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceCreateData) SetAttributes(v PriceCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceCreateData) GetRelationships() PriceCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PriceCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceCreateData) GetRelationshipsOk() (*PriceCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PriceCreateDataRelationships and assigns it to the Relationships field.
func (o *PriceCreateData) SetRelationships(v PriceCreateDataRelationships) {
	o.Relationships = &v
}

func (o PriceCreateData) MarshalJSON() ([]byte, error) {
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

type NullablePriceCreateData struct {
	value *PriceCreateData
	isSet bool
}

func (v NullablePriceCreateData) Get() *PriceCreateData {
	return v.value
}

func (v *NullablePriceCreateData) Set(val *PriceCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceCreateData(val *PriceCreateData) *NullablePriceCreateData {
	return &NullablePriceCreateData{value: val, isSet: true}
}

func (v NullablePriceCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
