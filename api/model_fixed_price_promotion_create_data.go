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

// FixedPricePromotionCreateData struct for FixedPricePromotionCreateData
type FixedPricePromotionCreateData struct {
	// The resource's type
	Type          string                                      `json:"type"`
	Attributes    FixedPricePromotionCreateDataAttributes     `json:"attributes"`
	Relationships *FixedPricePromotionCreateDataRelationships `json:"relationships,omitempty"`
}

// NewFixedPricePromotionCreateData instantiates a new FixedPricePromotionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedPricePromotionCreateData(type_ string, attributes FixedPricePromotionCreateDataAttributes) *FixedPricePromotionCreateData {
	this := FixedPricePromotionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewFixedPricePromotionCreateDataWithDefaults instantiates a new FixedPricePromotionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedPricePromotionCreateDataWithDefaults() *FixedPricePromotionCreateData {
	this := FixedPricePromotionCreateData{}
	var type_ string = "fixed_price_promotions"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *FixedPricePromotionCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FixedPricePromotionCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *FixedPricePromotionCreateData) GetAttributes() FixedPricePromotionCreateDataAttributes {
	if o == nil {
		var ret FixedPricePromotionCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionCreateData) GetAttributesOk() (*FixedPricePromotionCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *FixedPricePromotionCreateData) SetAttributes(v FixedPricePromotionCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FixedPricePromotionCreateData) GetRelationships() FixedPricePromotionCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret FixedPricePromotionCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionCreateData) GetRelationshipsOk() (*FixedPricePromotionCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FixedPricePromotionCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given FixedPricePromotionCreateDataRelationships and assigns it to the Relationships field.
func (o *FixedPricePromotionCreateData) SetRelationships(v FixedPricePromotionCreateDataRelationships) {
	o.Relationships = &v
}

func (o FixedPricePromotionCreateData) MarshalJSON() ([]byte, error) {
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

type NullableFixedPricePromotionCreateData struct {
	value *FixedPricePromotionCreateData
	isSet bool
}

func (v NullableFixedPricePromotionCreateData) Get() *FixedPricePromotionCreateData {
	return v.value
}

func (v *NullableFixedPricePromotionCreateData) Set(val *FixedPricePromotionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedPricePromotionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedPricePromotionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedPricePromotionCreateData(val *FixedPricePromotionCreateData) *NullableFixedPricePromotionCreateData {
	return &NullableFixedPricePromotionCreateData{value: val, isSet: true}
}

func (v NullableFixedPricePromotionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedPricePromotionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
