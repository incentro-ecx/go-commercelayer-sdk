/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETAuthorizations200ResponseDataInnerRelationshipsCapturesData struct for GETAuthorizations200ResponseDataInnerRelationshipsCapturesData
type GETAuthorizations200ResponseDataInnerRelationshipsCapturesData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETAuthorizations200ResponseDataInnerRelationshipsCapturesData instantiates a new GETAuthorizations200ResponseDataInnerRelationshipsCapturesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAuthorizations200ResponseDataInnerRelationshipsCapturesData() *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData {
	this := GETAuthorizations200ResponseDataInnerRelationshipsCapturesData{}
	return &this
}

// NewGETAuthorizations200ResponseDataInnerRelationshipsCapturesDataWithDefaults instantiates a new GETAuthorizations200ResponseDataInnerRelationshipsCapturesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAuthorizations200ResponseDataInnerRelationshipsCapturesDataWithDefaults() *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData {
	this := GETAuthorizations200ResponseDataInnerRelationshipsCapturesData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData) SetId(v string) {
	o.Id = &v
}

func (o GETAuthorizations200ResponseDataInnerRelationshipsCapturesData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETAuthorizations200ResponseDataInnerRelationshipsCapturesData struct {
	value *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData
	isSet bool
}

func (v NullableGETAuthorizations200ResponseDataInnerRelationshipsCapturesData) Get() *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData {
	return v.value
}

func (v *NullableGETAuthorizations200ResponseDataInnerRelationshipsCapturesData) Set(val *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAuthorizations200ResponseDataInnerRelationshipsCapturesData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAuthorizations200ResponseDataInnerRelationshipsCapturesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAuthorizations200ResponseDataInnerRelationshipsCapturesData(val *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData) *NullableGETAuthorizations200ResponseDataInnerRelationshipsCapturesData {
	return &NullableGETAuthorizations200ResponseDataInnerRelationshipsCapturesData{value: val, isSet: true}
}

func (v NullableGETAuthorizations200ResponseDataInnerRelationshipsCapturesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAuthorizations200ResponseDataInnerRelationshipsCapturesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}