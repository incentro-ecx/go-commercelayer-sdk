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

// GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments struct for GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments
type GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                           `json:"links,omitempty"`
	Data  []GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsDataInner `json:"data,omitempty"`
}

// NewGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments instantiates a new GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments() *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments {
	this := GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments{}
	return &this
}

// NewGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsWithDefaults instantiates a new GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsWithDefaults() *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments {
	this := GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) GetData() []GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsDataInner {
	if o == nil || o.Data == nil {
		var ret []GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) GetDataOk() ([]GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsDataInner and assigns it to the Data field.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) SetData(v []GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsDataInner) {
	o.Data = v
}

func (o GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments struct {
	value *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments
	isSet bool
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) Get() *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments {
	return v.value
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) Set(val *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments(val *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) *NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments {
	return &NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments{value: val, isSet: true}
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
