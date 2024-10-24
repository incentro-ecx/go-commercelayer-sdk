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

// checks if the GiftCardRecipientUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftCardRecipientUpdateData{}

// GiftCardRecipientUpdateData struct for GiftCardRecipientUpdateData
type GiftCardRecipientUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                     `json:"id"`
	Attributes    PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes `json:"attributes"`
	Relationships *CouponRecipientCreateDataRelationships                         `json:"relationships,omitempty"`
}

// NewGiftCardRecipientUpdateData instantiates a new GiftCardRecipientUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardRecipientUpdateData(type_ interface{}, id interface{}, attributes PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) *GiftCardRecipientUpdateData {
	this := GiftCardRecipientUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewGiftCardRecipientUpdateDataWithDefaults instantiates a new GiftCardRecipientUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardRecipientUpdateDataWithDefaults() *GiftCardRecipientUpdateData {
	this := GiftCardRecipientUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GiftCardRecipientUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardRecipientUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GiftCardRecipientUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GiftCardRecipientUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardRecipientUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GiftCardRecipientUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GiftCardRecipientUpdateData) GetAttributes() PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GiftCardRecipientUpdateData) GetAttributesOk() (*PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GiftCardRecipientUpdateData) SetAttributes(v PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GiftCardRecipientUpdateData) GetRelationships() CouponRecipientCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CouponRecipientCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardRecipientUpdateData) GetRelationshipsOk() (*CouponRecipientCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GiftCardRecipientUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CouponRecipientCreateDataRelationships and assigns it to the Relationships field.
func (o *GiftCardRecipientUpdateData) SetRelationships(v CouponRecipientCreateDataRelationships) {
	o.Relationships = &v
}

func (o GiftCardRecipientUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftCardRecipientUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableGiftCardRecipientUpdateData struct {
	value *GiftCardRecipientUpdateData
	isSet bool
}

func (v NullableGiftCardRecipientUpdateData) Get() *GiftCardRecipientUpdateData {
	return v.value
}

func (v *NullableGiftCardRecipientUpdateData) Set(val *GiftCardRecipientUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardRecipientUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardRecipientUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardRecipientUpdateData(val *GiftCardRecipientUpdateData) *NullableGiftCardRecipientUpdateData {
	return &NullableGiftCardRecipientUpdateData{value: val, isSet: true}
}

func (v NullableGiftCardRecipientUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardRecipientUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
