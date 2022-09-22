/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods struct for GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods
type GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                  `json:"links,omitempty"`
	Data  []GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner `json:"data,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods() *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods {
	this := GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsWithDefaults() *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods {
	this := GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) GetData() []GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner {
	if o == nil || o.Data == nil {
		var ret []GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) GetDataOk() ([]GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner and assigns it to the Data field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) SetData(v []GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) {
	o.Data = v
}

func (o GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods struct {
	value *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) Get() *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) Set(val *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods(val *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) *NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods {
	return &NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
