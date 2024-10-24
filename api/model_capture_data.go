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

// checks if the CaptureData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptureData{}

// CaptureData struct for CaptureData
type CaptureData struct {
	// The resource's type
	Type          interface{}                                   `json:"type"`
	Attributes    GETCapturesCaptureId200ResponseDataAttributes `json:"attributes"`
	Relationships *CaptureDataRelationships                     `json:"relationships,omitempty"`
}

// NewCaptureData instantiates a new CaptureData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureData(type_ interface{}, attributes GETCapturesCaptureId200ResponseDataAttributes) *CaptureData {
	this := CaptureData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCaptureDataWithDefaults instantiates a new CaptureData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureDataWithDefaults() *CaptureData {
	this := CaptureData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CaptureData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CaptureData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CaptureData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CaptureData) GetAttributes() GETCapturesCaptureId200ResponseDataAttributes {
	if o == nil {
		var ret GETCapturesCaptureId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CaptureData) GetAttributesOk() (*GETCapturesCaptureId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CaptureData) SetAttributes(v GETCapturesCaptureId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CaptureData) GetRelationships() CaptureDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CaptureDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureData) GetRelationshipsOk() (*CaptureDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CaptureData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CaptureDataRelationships and assigns it to the Relationships field.
func (o *CaptureData) SetRelationships(v CaptureDataRelationships) {
	o.Relationships = &v
}

func (o CaptureData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptureData) ToMap() (map[string]interface{}, error) {
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

type NullableCaptureData struct {
	value *CaptureData
	isSet bool
}

func (v NullableCaptureData) Get() *CaptureData {
	return v.value
}

func (v *NullableCaptureData) Set(val *CaptureData) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureData) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureData(val *CaptureData) *NullableCaptureData {
	return &NullableCaptureData{value: val, isSet: true}
}

func (v NullableCaptureData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
