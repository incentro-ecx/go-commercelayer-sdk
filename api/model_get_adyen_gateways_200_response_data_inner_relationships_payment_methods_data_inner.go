/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner struct for GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner
type GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner instantiates a new GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner() *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner {
	this := GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner{}
	return &this
}

// NewGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInnerWithDefaults instantiates a new GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInnerWithDefaults() *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner {
	this := GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) SetId(v string) {
	o.Id = &v
}

func (o GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner struct {
	value *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner
	isSet bool
}

func (v NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) Get() *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner {
	return v.value
}

func (v *NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) Set(val *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner(val *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) *NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner {
	return &NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner{value: val, isSet: true}
}

func (v NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}