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

// checks if the KlarnaGatewayData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KlarnaGatewayData{}

// KlarnaGatewayData struct for KlarnaGatewayData
type KlarnaGatewayData struct {
	// The resource's type
	Type          interface{}                                               `json:"type"`
	Attributes    GETKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes `json:"attributes"`
	Relationships *KlarnaGatewayDataRelationships                           `json:"relationships,omitempty"`
}

// NewKlarnaGatewayData instantiates a new KlarnaGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaGatewayData(type_ interface{}, attributes GETKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) *KlarnaGatewayData {
	this := KlarnaGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewKlarnaGatewayDataWithDefaults instantiates a new KlarnaGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaGatewayDataWithDefaults() *KlarnaGatewayData {
	this := KlarnaGatewayData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *KlarnaGatewayData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KlarnaGatewayData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KlarnaGatewayData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *KlarnaGatewayData) GetAttributes() GETKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes {
	if o == nil {
		var ret GETKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayData) GetAttributesOk() (*GETKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *KlarnaGatewayData) SetAttributes(v GETKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *KlarnaGatewayData) GetRelationships() KlarnaGatewayDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret KlarnaGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayData) GetRelationshipsOk() (*KlarnaGatewayDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *KlarnaGatewayData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given KlarnaGatewayDataRelationships and assigns it to the Relationships field.
func (o *KlarnaGatewayData) SetRelationships(v KlarnaGatewayDataRelationships) {
	o.Relationships = &v
}

func (o KlarnaGatewayData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlarnaGatewayData) ToMap() (map[string]interface{}, error) {
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

type NullableKlarnaGatewayData struct {
	value *KlarnaGatewayData
	isSet bool
}

func (v NullableKlarnaGatewayData) Get() *KlarnaGatewayData {
	return v.value
}

func (v *NullableKlarnaGatewayData) Set(val *KlarnaGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaGatewayData(val *KlarnaGatewayData) *NullableKlarnaGatewayData {
	return &NullableKlarnaGatewayData{value: val, isSet: true}
}

func (v NullableKlarnaGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
