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

// checks if the ShippingMethodTierData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingMethodTierData{}

// ShippingMethodTierData struct for ShippingMethodTierData
type ShippingMethodTierData struct {
	// The resource's type
	Type          interface{}                                                         `json:"type"`
	Attributes    GETShippingMethodTiersShippingMethodTierId200ResponseDataAttributes `json:"attributes"`
	Relationships *ShippingMethodTierDataRelationships                                `json:"relationships,omitempty"`
}

// NewShippingMethodTierData instantiates a new ShippingMethodTierData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodTierData(type_ interface{}, attributes GETShippingMethodTiersShippingMethodTierId200ResponseDataAttributes) *ShippingMethodTierData {
	this := ShippingMethodTierData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewShippingMethodTierDataWithDefaults instantiates a new ShippingMethodTierData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodTierDataWithDefaults() *ShippingMethodTierData {
	this := ShippingMethodTierData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ShippingMethodTierData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingMethodTierData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShippingMethodTierData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ShippingMethodTierData) GetAttributes() GETShippingMethodTiersShippingMethodTierId200ResponseDataAttributes {
	if o == nil {
		var ret GETShippingMethodTiersShippingMethodTierId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierData) GetAttributesOk() (*GETShippingMethodTiersShippingMethodTierId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShippingMethodTierData) SetAttributes(v GETShippingMethodTiersShippingMethodTierId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShippingMethodTierData) GetRelationships() ShippingMethodTierDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ShippingMethodTierDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierData) GetRelationshipsOk() (*ShippingMethodTierDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShippingMethodTierData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ShippingMethodTierDataRelationships and assigns it to the Relationships field.
func (o *ShippingMethodTierData) SetRelationships(v ShippingMethodTierDataRelationships) {
	o.Relationships = &v
}

func (o ShippingMethodTierData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingMethodTierData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableShippingMethodTierData struct {
	value *ShippingMethodTierData
	isSet bool
}

func (v NullableShippingMethodTierData) Get() *ShippingMethodTierData {
	return v.value
}

func (v *NullableShippingMethodTierData) Set(val *ShippingMethodTierData) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodTierData) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodTierData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodTierData(val *ShippingMethodTierData) *NullableShippingMethodTierData {
	return &NullableShippingMethodTierData{value: val, isSet: true}
}

func (v NullableShippingMethodTierData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodTierData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
