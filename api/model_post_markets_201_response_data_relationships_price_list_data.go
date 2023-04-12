/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTMarkets201ResponseDataRelationshipsPriceListData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarkets201ResponseDataRelationshipsPriceListData{}

// POSTMarkets201ResponseDataRelationshipsPriceListData struct for POSTMarkets201ResponseDataRelationshipsPriceListData
type POSTMarkets201ResponseDataRelationshipsPriceListData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTMarkets201ResponseDataRelationshipsPriceListData instantiates a new POSTMarkets201ResponseDataRelationshipsPriceListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarkets201ResponseDataRelationshipsPriceListData() *POSTMarkets201ResponseDataRelationshipsPriceListData {
	this := POSTMarkets201ResponseDataRelationshipsPriceListData{}
	return &this
}

// NewPOSTMarkets201ResponseDataRelationshipsPriceListDataWithDefaults instantiates a new POSTMarkets201ResponseDataRelationshipsPriceListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarkets201ResponseDataRelationshipsPriceListDataWithDefaults() *POSTMarkets201ResponseDataRelationshipsPriceListData {
	this := POSTMarkets201ResponseDataRelationshipsPriceListData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataRelationshipsPriceListData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataRelationshipsPriceListData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsPriceListData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTMarkets201ResponseDataRelationshipsPriceListData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataRelationshipsPriceListData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataRelationshipsPriceListData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsPriceListData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTMarkets201ResponseDataRelationshipsPriceListData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTMarkets201ResponseDataRelationshipsPriceListData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarkets201ResponseDataRelationshipsPriceListData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTMarkets201ResponseDataRelationshipsPriceListData struct {
	value *POSTMarkets201ResponseDataRelationshipsPriceListData
	isSet bool
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsPriceListData) Get() *POSTMarkets201ResponseDataRelationshipsPriceListData {
	return v.value
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsPriceListData) Set(val *POSTMarkets201ResponseDataRelationshipsPriceListData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsPriceListData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsPriceListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarkets201ResponseDataRelationshipsPriceListData(val *POSTMarkets201ResponseDataRelationshipsPriceListData) *NullablePOSTMarkets201ResponseDataRelationshipsPriceListData {
	return &NullablePOSTMarkets201ResponseDataRelationshipsPriceListData{value: val, isSet: true}
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsPriceListData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsPriceListData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}