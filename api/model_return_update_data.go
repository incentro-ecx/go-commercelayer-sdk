/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReturnUpdateData struct for ReturnUpdateData
type ReturnUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                          `json:"id"`
	Attributes    ReturnUpdateDataAttributes      `json:"attributes"`
	Relationships *PackageUpdateDataRelationships `json:"relationships,omitempty"`
}

// NewReturnUpdateData instantiates a new ReturnUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnUpdateData(type_ string, id string, attributes ReturnUpdateDataAttributes) *ReturnUpdateData {
	this := ReturnUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewReturnUpdateDataWithDefaults instantiates a new ReturnUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnUpdateDataWithDefaults() *ReturnUpdateData {
	this := ReturnUpdateData{}
	var type_ string = "returns"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ReturnUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReturnUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReturnUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ReturnUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReturnUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReturnUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ReturnUpdateData) GetAttributes() ReturnUpdateDataAttributes {
	if o == nil {
		var ret ReturnUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ReturnUpdateData) GetAttributesOk() (*ReturnUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ReturnUpdateData) SetAttributes(v ReturnUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ReturnUpdateData) GetRelationships() PackageUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PackageUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnUpdateData) GetRelationshipsOk() (*PackageUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ReturnUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PackageUpdateDataRelationships and assigns it to the Relationships field.
func (o *ReturnUpdateData) SetRelationships(v PackageUpdateDataRelationships) {
	o.Relationships = &v
}

func (o ReturnUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableReturnUpdateData struct {
	value *ReturnUpdateData
	isSet bool
}

func (v NullableReturnUpdateData) Get() *ReturnUpdateData {
	return v.value
}

func (v *NullableReturnUpdateData) Set(val *ReturnUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnUpdateData(val *ReturnUpdateData) *NullableReturnUpdateData {
	return &NullableReturnUpdateData{value: val, isSet: true}
}

func (v NullableReturnUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
