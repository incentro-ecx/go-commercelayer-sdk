/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GiftCardResponseData struct for GiftCardResponseData
type GiftCardResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                            `json:"type,omitempty"`
	Links         *AddressResponseDataLinks          `json:"links,omitempty"`
	Attributes    *Attributes                        `json:"attributes,omitempty"`
	Relationships *GiftCardResponseDataRelationships `json:"relationships,omitempty"`
}

// NewGiftCardResponseData instantiates a new GiftCardResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardResponseData() *GiftCardResponseData {
	this := GiftCardResponseData{}
	return &this
}

// NewGiftCardResponseDataWithDefaults instantiates a new GiftCardResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardResponseDataWithDefaults() *GiftCardResponseData {
	this := GiftCardResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GiftCardResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GiftCardResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GiftCardResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GiftCardResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GiftCardResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GiftCardResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GiftCardResponseData) GetLinks() AddressResponseDataLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardResponseData) GetLinksOk() (*AddressResponseDataLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GiftCardResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataLinks and assigns it to the Links field.
func (o *GiftCardResponseData) SetLinks(v AddressResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GiftCardResponseData) GetAttributes() Attributes {
	if o == nil || o.Attributes == nil {
		var ret Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardResponseData) GetAttributesOk() (*Attributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GiftCardResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Attributes and assigns it to the Attributes field.
func (o *GiftCardResponseData) SetAttributes(v Attributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GiftCardResponseData) GetRelationships() GiftCardResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret GiftCardResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardResponseData) GetRelationshipsOk() (*GiftCardResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GiftCardResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GiftCardResponseDataRelationships and assigns it to the Relationships field.
func (o *GiftCardResponseData) SetRelationships(v GiftCardResponseDataRelationships) {
	o.Relationships = &v
}

func (o GiftCardResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCardResponseData struct {
	value *GiftCardResponseData
	isSet bool
}

func (v NullableGiftCardResponseData) Get() *GiftCardResponseData {
	return v.value
}

func (v *NullableGiftCardResponseData) Set(val *GiftCardResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardResponseData(val *GiftCardResponseData) *NullableGiftCardResponseData {
	return &NullableGiftCardResponseData{value: val, isSet: true}
}

func (v NullableGiftCardResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}