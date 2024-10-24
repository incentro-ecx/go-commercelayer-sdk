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

// checks if the POSTMarkets201ResponseDataRelationshipsSubscriptionModelData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarkets201ResponseDataRelationshipsSubscriptionModelData{}

// POSTMarkets201ResponseDataRelationshipsSubscriptionModelData struct for POSTMarkets201ResponseDataRelationshipsSubscriptionModelData
type POSTMarkets201ResponseDataRelationshipsSubscriptionModelData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTMarkets201ResponseDataRelationshipsSubscriptionModelData instantiates a new POSTMarkets201ResponseDataRelationshipsSubscriptionModelData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarkets201ResponseDataRelationshipsSubscriptionModelData() *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData {
	this := POSTMarkets201ResponseDataRelationshipsSubscriptionModelData{}
	return &this
}

// NewPOSTMarkets201ResponseDataRelationshipsSubscriptionModelDataWithDefaults instantiates a new POSTMarkets201ResponseDataRelationshipsSubscriptionModelData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarkets201ResponseDataRelationshipsSubscriptionModelDataWithDefaults() *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData {
	this := POSTMarkets201ResponseDataRelationshipsSubscriptionModelData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModelData struct {
	value *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData
	isSet bool
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModelData) Get() *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData {
	return v.value
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModelData) Set(val *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModelData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModelData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModelData(val *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) *NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModelData {
	return &NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModelData{value: val, isSet: true}
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModelData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModelData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
