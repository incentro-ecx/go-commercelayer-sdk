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

// GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData struct for GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData
type GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData() *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData {
	this := GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataWithDefaults() *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData {
	this := GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) SetId(v string) {
	o.Id = &v
}

func (o GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData struct {
	value *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) Get() *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) Set(val *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData(val *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) *NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData {
	return &NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}