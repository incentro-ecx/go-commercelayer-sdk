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

// checks if the POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData{}

// POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData struct for POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData
type POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData instantiates a new POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData() *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData {
	this := POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData{}
	return &this
}

// NewPOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleDataWithDefaults instantiates a new POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleDataWithDefaults() *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData {
	this := POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData struct {
	value *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData
	isSet bool
}

func (v NullablePOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) Get() *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData {
	return v.value
}

func (v *NullablePOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) Set(val *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData(val *POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) *NullablePOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData {
	return &NullablePOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData{value: val, isSet: true}
}

func (v NullablePOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
