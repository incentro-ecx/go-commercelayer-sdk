/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GiftCardRecipientData struct for GiftCardRecipientData
type GiftCardRecipientData struct {
	// The resource's type
	Type          string                            `json:"type"`
	Attributes    CouponRecipientDataAttributes     `json:"attributes"`
	Relationships *CouponRecipientDataRelationships `json:"relationships,omitempty"`
}

// NewGiftCardRecipientData instantiates a new GiftCardRecipientData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardRecipientData(type_ string, attributes CouponRecipientDataAttributes) *GiftCardRecipientData {
	this := GiftCardRecipientData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewGiftCardRecipientDataWithDefaults instantiates a new GiftCardRecipientData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardRecipientDataWithDefaults() *GiftCardRecipientData {
	this := GiftCardRecipientData{}
	var type_ string = "gift_card_recipients"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GiftCardRecipientData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GiftCardRecipientData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GiftCardRecipientData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GiftCardRecipientData) GetAttributes() CouponRecipientDataAttributes {
	if o == nil {
		var ret CouponRecipientDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GiftCardRecipientData) GetAttributesOk() (*CouponRecipientDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GiftCardRecipientData) SetAttributes(v CouponRecipientDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GiftCardRecipientData) GetRelationships() CouponRecipientDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CouponRecipientDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardRecipientData) GetRelationshipsOk() (*CouponRecipientDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GiftCardRecipientData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CouponRecipientDataRelationships and assigns it to the Relationships field.
func (o *GiftCardRecipientData) SetRelationships(v CouponRecipientDataRelationships) {
	o.Relationships = &v
}

func (o GiftCardRecipientData) MarshalJSON() ([]byte, error) {
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

type NullableGiftCardRecipientData struct {
	value *GiftCardRecipientData
	isSet bool
}

func (v NullableGiftCardRecipientData) Get() *GiftCardRecipientData {
	return v.value
}

func (v *NullableGiftCardRecipientData) Set(val *GiftCardRecipientData) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardRecipientData) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardRecipientData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardRecipientData(val *GiftCardRecipientData) *NullableGiftCardRecipientData {
	return &NullableGiftCardRecipientData{value: val, isSet: true}
}

func (v NullableGiftCardRecipientData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardRecipientData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}