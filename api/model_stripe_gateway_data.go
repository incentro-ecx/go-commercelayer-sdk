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

// checks if the StripeGatewayData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripeGatewayData{}

// StripeGatewayData struct for StripeGatewayData
type StripeGatewayData struct {
	// The resource's type
	Type          interface{}                                               `json:"type"`
	Attributes    GETStripeGatewaysStripeGatewayId200ResponseDataAttributes `json:"attributes"`
	Relationships *StripeGatewayDataRelationships                           `json:"relationships,omitempty"`
}

// NewStripeGatewayData instantiates a new StripeGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeGatewayData(type_ interface{}, attributes GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) *StripeGatewayData {
	this := StripeGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewStripeGatewayDataWithDefaults instantiates a new StripeGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeGatewayDataWithDefaults() *StripeGatewayData {
	this := StripeGatewayData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *StripeGatewayData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StripeGatewayData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StripeGatewayData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *StripeGatewayData) GetAttributes() GETStripeGatewaysStripeGatewayId200ResponseDataAttributes {
	if o == nil {
		var ret GETStripeGatewaysStripeGatewayId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StripeGatewayData) GetAttributesOk() (*GETStripeGatewaysStripeGatewayId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StripeGatewayData) SetAttributes(v GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StripeGatewayData) GetRelationships() StripeGatewayDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret StripeGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeGatewayData) GetRelationshipsOk() (*StripeGatewayDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StripeGatewayData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given StripeGatewayDataRelationships and assigns it to the Relationships field.
func (o *StripeGatewayData) SetRelationships(v StripeGatewayDataRelationships) {
	o.Relationships = &v
}

func (o StripeGatewayData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripeGatewayData) ToMap() (map[string]interface{}, error) {
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

type NullableStripeGatewayData struct {
	value *StripeGatewayData
	isSet bool
}

func (v NullableStripeGatewayData) Get() *StripeGatewayData {
	return v.value
}

func (v *NullableStripeGatewayData) Set(val *StripeGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeGatewayData(val *StripeGatewayData) *NullableStripeGatewayData {
	return &NullableStripeGatewayData{value: val, isSet: true}
}

func (v NullableStripeGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
