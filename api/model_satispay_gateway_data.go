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

// checks if the SatispayGatewayData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SatispayGatewayData{}

// SatispayGatewayData struct for SatispayGatewayData
type SatispayGatewayData struct {
	// The resource's type
	Type          interface{}                                                   `json:"type"`
	Attributes    GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes `json:"attributes"`
	Relationships *SatispayGatewayDataRelationships                             `json:"relationships,omitempty"`
}

// NewSatispayGatewayData instantiates a new SatispayGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSatispayGatewayData(type_ interface{}, attributes GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) *SatispayGatewayData {
	this := SatispayGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSatispayGatewayDataWithDefaults instantiates a new SatispayGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSatispayGatewayDataWithDefaults() *SatispayGatewayData {
	this := SatispayGatewayData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SatispayGatewayData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SatispayGatewayData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SatispayGatewayData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SatispayGatewayData) GetAttributes() GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes {
	if o == nil {
		var ret GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SatispayGatewayData) GetAttributesOk() (*GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SatispayGatewayData) SetAttributes(v GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SatispayGatewayData) GetRelationships() SatispayGatewayDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SatispayGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SatispayGatewayData) GetRelationshipsOk() (*SatispayGatewayDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SatispayGatewayData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SatispayGatewayDataRelationships and assigns it to the Relationships field.
func (o *SatispayGatewayData) SetRelationships(v SatispayGatewayDataRelationships) {
	o.Relationships = &v
}

func (o SatispayGatewayData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SatispayGatewayData) ToMap() (map[string]interface{}, error) {
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

type NullableSatispayGatewayData struct {
	value *SatispayGatewayData
	isSet bool
}

func (v NullableSatispayGatewayData) Get() *SatispayGatewayData {
	return v.value
}

func (v *NullableSatispayGatewayData) Set(val *SatispayGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableSatispayGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableSatispayGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSatispayGatewayData(val *SatispayGatewayData) *NullableSatispayGatewayData {
	return &NullableSatispayGatewayData{value: val, isSet: true}
}

func (v NullableSatispayGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSatispayGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
