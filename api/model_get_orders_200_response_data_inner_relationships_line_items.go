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

// GETOrders200ResponseDataInnerRelationshipsLineItems struct for GETOrders200ResponseDataInnerRelationshipsLineItems
type GETOrders200ResponseDataInnerRelationshipsLineItems struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks    `json:"links,omitempty"`
	Data  []GETOrders200ResponseDataInnerRelationshipsLineItemsDataInner `json:"data,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsLineItems instantiates a new GETOrders200ResponseDataInnerRelationshipsLineItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsLineItems() *GETOrders200ResponseDataInnerRelationshipsLineItems {
	this := GETOrders200ResponseDataInnerRelationshipsLineItems{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsLineItemsWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsLineItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsLineItemsWithDefaults() *GETOrders200ResponseDataInnerRelationshipsLineItems {
	this := GETOrders200ResponseDataInnerRelationshipsLineItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsLineItems) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsLineItems) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsLineItems) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETOrders200ResponseDataInnerRelationshipsLineItems) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsLineItems) GetData() []GETOrders200ResponseDataInnerRelationshipsLineItemsDataInner {
	if o == nil || o.Data == nil {
		var ret []GETOrders200ResponseDataInnerRelationshipsLineItemsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsLineItems) GetDataOk() ([]GETOrders200ResponseDataInnerRelationshipsLineItemsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsLineItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETOrders200ResponseDataInnerRelationshipsLineItemsDataInner and assigns it to the Data field.
func (o *GETOrders200ResponseDataInnerRelationshipsLineItems) SetData(v []GETOrders200ResponseDataInnerRelationshipsLineItemsDataInner) {
	o.Data = v
}

func (o GETOrders200ResponseDataInnerRelationshipsLineItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsLineItems struct {
	value *GETOrders200ResponseDataInnerRelationshipsLineItems
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsLineItems) Get() *GETOrders200ResponseDataInnerRelationshipsLineItems {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsLineItems) Set(val *GETOrders200ResponseDataInnerRelationshipsLineItems) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsLineItems) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsLineItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsLineItems(val *GETOrders200ResponseDataInnerRelationshipsLineItems) *NullableGETOrders200ResponseDataInnerRelationshipsLineItems {
	return &NullableGETOrders200ResponseDataInnerRelationshipsLineItems{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsLineItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsLineItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
