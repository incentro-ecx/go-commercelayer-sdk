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

// GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments struct for GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments
type GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks         `json:"links,omitempty"`
	Data  *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPaymentsData `json:"data,omitempty"`
}

// NewGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments instantiates a new GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments() *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments {
	this := GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments{}
	return &this
}

// NewGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPaymentsWithDefaults instantiates a new GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPaymentsWithDefaults() *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments {
	this := GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) GetData() GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPaymentsData {
	if o == nil || o.Data == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) GetDataOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPaymentsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPaymentsData and assigns it to the Data field.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) SetData(v GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPaymentsData) {
	o.Data = &v
}

func (o GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments struct {
	value *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments
	isSet bool
}

func (v NullableGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) Get() *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments {
	return v.value
}

func (v *NullableGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) Set(val *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments(val *GETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) *NullableGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments {
	return &NullableGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments{value: val, isSet: true}
}

func (v NullableGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdyenGateways200ResponseDataInnerRelationshipsAdyenPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
