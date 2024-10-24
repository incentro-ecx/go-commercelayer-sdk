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

// checks if the POSTOrders201ResponseDataRelationshipsAuthorizationsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsAuthorizationsData{}

// POSTOrders201ResponseDataRelationshipsAuthorizationsData struct for POSTOrders201ResponseDataRelationshipsAuthorizationsData
type POSTOrders201ResponseDataRelationshipsAuthorizationsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsAuthorizationsData instantiates a new POSTOrders201ResponseDataRelationshipsAuthorizationsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsAuthorizationsData() *POSTOrders201ResponseDataRelationshipsAuthorizationsData {
	this := POSTOrders201ResponseDataRelationshipsAuthorizationsData{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsAuthorizationsDataWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsAuthorizationsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsAuthorizationsDataWithDefaults() *POSTOrders201ResponseDataRelationshipsAuthorizationsData {
	this := POSTOrders201ResponseDataRelationshipsAuthorizationsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataRelationshipsAuthorizationsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataRelationshipsAuthorizationsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsAuthorizationsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTOrders201ResponseDataRelationshipsAuthorizationsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataRelationshipsAuthorizationsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataRelationshipsAuthorizationsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsAuthorizationsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTOrders201ResponseDataRelationshipsAuthorizationsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTOrders201ResponseDataRelationshipsAuthorizationsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsAuthorizationsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsAuthorizationsData struct {
	value *POSTOrders201ResponseDataRelationshipsAuthorizationsData
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsAuthorizationsData) Get() *POSTOrders201ResponseDataRelationshipsAuthorizationsData {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsAuthorizationsData) Set(val *POSTOrders201ResponseDataRelationshipsAuthorizationsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsAuthorizationsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsAuthorizationsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsAuthorizationsData(val *POSTOrders201ResponseDataRelationshipsAuthorizationsData) *NullablePOSTOrders201ResponseDataRelationshipsAuthorizationsData {
	return &NullablePOSTOrders201ResponseDataRelationshipsAuthorizationsData{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsAuthorizationsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsAuthorizationsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
