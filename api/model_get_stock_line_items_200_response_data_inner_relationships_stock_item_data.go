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

// GETStockLineItems200ResponseDataInnerRelationshipsStockItemData struct for GETStockLineItems200ResponseDataInnerRelationshipsStockItemData
type GETStockLineItems200ResponseDataInnerRelationshipsStockItemData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETStockLineItems200ResponseDataInnerRelationshipsStockItemData instantiates a new GETStockLineItems200ResponseDataInnerRelationshipsStockItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockLineItems200ResponseDataInnerRelationshipsStockItemData() *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData {
	this := GETStockLineItems200ResponseDataInnerRelationshipsStockItemData{}
	return &this
}

// NewGETStockLineItems200ResponseDataInnerRelationshipsStockItemDataWithDefaults instantiates a new GETStockLineItems200ResponseDataInnerRelationshipsStockItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockLineItems200ResponseDataInnerRelationshipsStockItemDataWithDefaults() *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData {
	this := GETStockLineItems200ResponseDataInnerRelationshipsStockItemData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData) SetId(v string) {
	o.Id = &v
}

func (o GETStockLineItems200ResponseDataInnerRelationshipsStockItemData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETStockLineItems200ResponseDataInnerRelationshipsStockItemData struct {
	value *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData
	isSet bool
}

func (v NullableGETStockLineItems200ResponseDataInnerRelationshipsStockItemData) Get() *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData {
	return v.value
}

func (v *NullableGETStockLineItems200ResponseDataInnerRelationshipsStockItemData) Set(val *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockLineItems200ResponseDataInnerRelationshipsStockItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockLineItems200ResponseDataInnerRelationshipsStockItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockLineItems200ResponseDataInnerRelationshipsStockItemData(val *GETStockLineItems200ResponseDataInnerRelationshipsStockItemData) *NullableGETStockLineItems200ResponseDataInnerRelationshipsStockItemData {
	return &NullableGETStockLineItems200ResponseDataInnerRelationshipsStockItemData{value: val, isSet: true}
}

func (v NullableGETStockLineItems200ResponseDataInnerRelationshipsStockItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockLineItems200ResponseDataInnerRelationshipsStockItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
