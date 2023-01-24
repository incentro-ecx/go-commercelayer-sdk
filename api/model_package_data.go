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

// PackageData struct for PackageData
type PackageData struct {
	// The resource's type
	Type          string                                    `json:"type"`
	Attributes    GETPackages200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *PackageDataRelationships                 `json:"relationships,omitempty"`
}

// NewPackageData instantiates a new PackageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageData(type_ string, attributes GETPackages200ResponseDataInnerAttributes) *PackageData {
	this := PackageData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPackageDataWithDefaults instantiates a new PackageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDataWithDefaults() *PackageData {
	this := PackageData{}
	return &this
}

// GetType returns the Type field value
func (o *PackageData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PackageData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PackageData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PackageData) GetAttributes() GETPackages200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETPackages200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PackageData) GetAttributesOk() (*GETPackages200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PackageData) SetAttributes(v GETPackages200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PackageData) GetRelationships() PackageDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PackageDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageData) GetRelationshipsOk() (*PackageDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PackageData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PackageDataRelationships and assigns it to the Relationships field.
func (o *PackageData) SetRelationships(v PackageDataRelationships) {
	o.Relationships = &v
}

func (o PackageData) MarshalJSON() ([]byte, error) {
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

type NullablePackageData struct {
	value *PackageData
	isSet bool
}

func (v NullablePackageData) Get() *PackageData {
	return v.value
}

func (v *NullablePackageData) Set(val *PackageData) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageData) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageData(val *PackageData) *NullablePackageData {
	return &NullablePackageData{value: val, isSet: true}
}

func (v NullablePackageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
