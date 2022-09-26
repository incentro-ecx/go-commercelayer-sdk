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

// FixedPricePromotionResponseData struct for FixedPricePromotionResponseData
type FixedPricePromotionResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                       `json:"type,omitempty"`
	Links         *AddressResponseDataLinks                     `json:"links,omitempty"`
	Attributes    *Attributes                                   `json:"attributes,omitempty"`
	Relationships *FixedPricePromotionResponseDataRelationships `json:"relationships,omitempty"`
}

// NewFixedPricePromotionResponseData instantiates a new FixedPricePromotionResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedPricePromotionResponseData() *FixedPricePromotionResponseData {
	this := FixedPricePromotionResponseData{}
	return &this
}

// NewFixedPricePromotionResponseDataWithDefaults instantiates a new FixedPricePromotionResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedPricePromotionResponseDataWithDefaults() *FixedPricePromotionResponseData {
	this := FixedPricePromotionResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FixedPricePromotionResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FixedPricePromotionResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseData) GetLinks() AddressResponseDataLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseData) GetLinksOk() (*AddressResponseDataLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataLinks and assigns it to the Links field.
func (o *FixedPricePromotionResponseData) SetLinks(v AddressResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseData) GetAttributes() Attributes {
	if o == nil || o.Attributes == nil {
		var ret Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseData) GetAttributesOk() (*Attributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Attributes and assigns it to the Attributes field.
func (o *FixedPricePromotionResponseData) SetAttributes(v Attributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseData) GetRelationships() FixedPricePromotionResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret FixedPricePromotionResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseData) GetRelationshipsOk() (*FixedPricePromotionResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given FixedPricePromotionResponseDataRelationships and assigns it to the Relationships field.
func (o *FixedPricePromotionResponseData) SetRelationships(v FixedPricePromotionResponseDataRelationships) {
	o.Relationships = &v
}

func (o FixedPricePromotionResponseData) MarshalJSON() ([]byte, error) {
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

type NullableFixedPricePromotionResponseData struct {
	value *FixedPricePromotionResponseData
	isSet bool
}

func (v NullableFixedPricePromotionResponseData) Get() *FixedPricePromotionResponseData {
	return v.value
}

func (v *NullableFixedPricePromotionResponseData) Set(val *FixedPricePromotionResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedPricePromotionResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedPricePromotionResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedPricePromotionResponseData(val *FixedPricePromotionResponseData) *NullableFixedPricePromotionResponseData {
	return &NullableFixedPricePromotionResponseData{value: val, isSet: true}
}

func (v NullableFixedPricePromotionResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedPricePromotionResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}