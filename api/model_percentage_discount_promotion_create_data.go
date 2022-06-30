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

// PercentageDiscountPromotionCreateData struct for PercentageDiscountPromotionCreateData
type PercentageDiscountPromotionCreateData struct {
	// The resource's type
	Type          string                                          `json:"type"`
	Attributes    PercentageDiscountPromotionCreateDataAttributes `json:"attributes"`
	Relationships *FixedPricePromotionUpdateDataRelationships     `json:"relationships,omitempty"`
}

// NewPercentageDiscountPromotionCreateData instantiates a new PercentageDiscountPromotionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPercentageDiscountPromotionCreateData(type_ string, attributes PercentageDiscountPromotionCreateDataAttributes) *PercentageDiscountPromotionCreateData {
	this := PercentageDiscountPromotionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPercentageDiscountPromotionCreateDataWithDefaults instantiates a new PercentageDiscountPromotionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPercentageDiscountPromotionCreateDataWithDefaults() *PercentageDiscountPromotionCreateData {
	this := PercentageDiscountPromotionCreateData{}
	var type_ string = "percentage_discount_promotions"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *PercentageDiscountPromotionCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotionCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PercentageDiscountPromotionCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PercentageDiscountPromotionCreateData) GetAttributes() PercentageDiscountPromotionCreateDataAttributes {
	if o == nil {
		var ret PercentageDiscountPromotionCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotionCreateData) GetAttributesOk() (*PercentageDiscountPromotionCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PercentageDiscountPromotionCreateData) SetAttributes(v PercentageDiscountPromotionCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PercentageDiscountPromotionCreateData) GetRelationships() FixedPricePromotionUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret FixedPricePromotionUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotionCreateData) GetRelationshipsOk() (*FixedPricePromotionUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PercentageDiscountPromotionCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given FixedPricePromotionUpdateDataRelationships and assigns it to the Relationships field.
func (o *PercentageDiscountPromotionCreateData) SetRelationships(v FixedPricePromotionUpdateDataRelationships) {
	o.Relationships = &v
}

func (o PercentageDiscountPromotionCreateData) MarshalJSON() ([]byte, error) {
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

type NullablePercentageDiscountPromotionCreateData struct {
	value *PercentageDiscountPromotionCreateData
	isSet bool
}

func (v NullablePercentageDiscountPromotionCreateData) Get() *PercentageDiscountPromotionCreateData {
	return v.value
}

func (v *NullablePercentageDiscountPromotionCreateData) Set(val *PercentageDiscountPromotionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePercentageDiscountPromotionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePercentageDiscountPromotionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePercentageDiscountPromotionCreateData(val *PercentageDiscountPromotionCreateData) *NullablePercentageDiscountPromotionCreateData {
	return &NullablePercentageDiscountPromotionCreateData{value: val, isSet: true}
}

func (v NullablePercentageDiscountPromotionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePercentageDiscountPromotionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
