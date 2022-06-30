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

// CaptureDataRelationshipsRefunds struct for CaptureDataRelationshipsRefunds
type CaptureDataRelationshipsRefunds struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewCaptureDataRelationshipsRefunds instantiates a new CaptureDataRelationshipsRefunds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureDataRelationshipsRefunds(type_ string, id string) *CaptureDataRelationshipsRefunds {
	this := CaptureDataRelationshipsRefunds{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCaptureDataRelationshipsRefundsWithDefaults instantiates a new CaptureDataRelationshipsRefunds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureDataRelationshipsRefundsWithDefaults() *CaptureDataRelationshipsRefunds {
	this := CaptureDataRelationshipsRefunds{}
	var type_ string = "refunds"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CaptureDataRelationshipsRefunds) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaptureDataRelationshipsRefunds) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CaptureDataRelationshipsRefunds) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CaptureDataRelationshipsRefunds) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CaptureDataRelationshipsRefunds) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CaptureDataRelationshipsRefunds) SetId(v string) {
	o.Id = v
}

func (o CaptureDataRelationshipsRefunds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCaptureDataRelationshipsRefunds struct {
	value *CaptureDataRelationshipsRefunds
	isSet bool
}

func (v NullableCaptureDataRelationshipsRefunds) Get() *CaptureDataRelationshipsRefunds {
	return v.value
}

func (v *NullableCaptureDataRelationshipsRefunds) Set(val *CaptureDataRelationshipsRefunds) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureDataRelationshipsRefunds) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureDataRelationshipsRefunds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureDataRelationshipsRefunds(val *CaptureDataRelationshipsRefunds) *NullableCaptureDataRelationshipsRefunds {
	return &NullableCaptureDataRelationshipsRefunds{value: val, isSet: true}
}

func (v NullableCaptureDataRelationshipsRefunds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureDataRelationshipsRefunds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
