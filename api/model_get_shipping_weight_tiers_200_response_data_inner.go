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

// GETShippingWeightTiers200ResponseDataInner struct for GETShippingWeightTiers200ResponseDataInner
type GETShippingWeightTiers200ResponseDataInner struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                                  `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks                   `json:"links,omitempty"`
	Attributes    *GETShippingMethodTiers200ResponseDataInnerAttributes    `json:"attributes,omitempty"`
	Relationships *GETShippingMethodTiers200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewGETShippingWeightTiers200ResponseDataInner instantiates a new GETShippingWeightTiers200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingWeightTiers200ResponseDataInner() *GETShippingWeightTiers200ResponseDataInner {
	this := GETShippingWeightTiers200ResponseDataInner{}
	return &this
}

// NewGETShippingWeightTiers200ResponseDataInnerWithDefaults instantiates a new GETShippingWeightTiers200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingWeightTiers200ResponseDataInnerWithDefaults() *GETShippingWeightTiers200ResponseDataInner {
	this := GETShippingWeightTiers200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETShippingWeightTiers200ResponseDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingWeightTiers200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETShippingWeightTiers200ResponseDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETShippingWeightTiers200ResponseDataInner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETShippingWeightTiers200ResponseDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingWeightTiers200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETShippingWeightTiers200ResponseDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETShippingWeightTiers200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETShippingWeightTiers200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingWeightTiers200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETShippingWeightTiers200ResponseDataInner) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *GETShippingWeightTiers200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GETShippingWeightTiers200ResponseDataInner) GetAttributes() GETShippingMethodTiers200ResponseDataInnerAttributes {
	if o == nil || o.Attributes == nil {
		var ret GETShippingMethodTiers200ResponseDataInnerAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingWeightTiers200ResponseDataInner) GetAttributesOk() (*GETShippingMethodTiers200ResponseDataInnerAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GETShippingWeightTiers200ResponseDataInner) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GETShippingMethodTiers200ResponseDataInnerAttributes and assigns it to the Attributes field.
func (o *GETShippingWeightTiers200ResponseDataInner) SetAttributes(v GETShippingMethodTiers200ResponseDataInnerAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GETShippingWeightTiers200ResponseDataInner) GetRelationships() GETShippingMethodTiers200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETShippingMethodTiers200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingWeightTiers200ResponseDataInner) GetRelationshipsOk() (*GETShippingMethodTiers200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GETShippingWeightTiers200ResponseDataInner) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETShippingMethodTiers200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *GETShippingWeightTiers200ResponseDataInner) SetRelationships(v GETShippingMethodTiers200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o GETShippingWeightTiers200ResponseDataInner) MarshalJSON() ([]byte, error) {
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

type NullableGETShippingWeightTiers200ResponseDataInner struct {
	value *GETShippingWeightTiers200ResponseDataInner
	isSet bool
}

func (v NullableGETShippingWeightTiers200ResponseDataInner) Get() *GETShippingWeightTiers200ResponseDataInner {
	return v.value
}

func (v *NullableGETShippingWeightTiers200ResponseDataInner) Set(val *GETShippingWeightTiers200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingWeightTiers200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingWeightTiers200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingWeightTiers200ResponseDataInner(val *GETShippingWeightTiers200ResponseDataInner) *NullableGETShippingWeightTiers200ResponseDataInner {
	return &NullableGETShippingWeightTiers200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGETShippingWeightTiers200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingWeightTiers200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
