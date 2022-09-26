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

// MarketResponseDataRelationshipsPriceList struct for MarketResponseDataRelationshipsPriceList
type MarketResponseDataRelationshipsPriceList struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *MarketResponseDataRelationshipsPriceListData  `json:"data,omitempty"`
}

// NewMarketResponseDataRelationshipsPriceList instantiates a new MarketResponseDataRelationshipsPriceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketResponseDataRelationshipsPriceList() *MarketResponseDataRelationshipsPriceList {
	this := MarketResponseDataRelationshipsPriceList{}
	return &this
}

// NewMarketResponseDataRelationshipsPriceListWithDefaults instantiates a new MarketResponseDataRelationshipsPriceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketResponseDataRelationshipsPriceListWithDefaults() *MarketResponseDataRelationshipsPriceList {
	this := MarketResponseDataRelationshipsPriceList{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MarketResponseDataRelationshipsPriceList) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketResponseDataRelationshipsPriceList) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MarketResponseDataRelationshipsPriceList) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *MarketResponseDataRelationshipsPriceList) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MarketResponseDataRelationshipsPriceList) GetData() MarketResponseDataRelationshipsPriceListData {
	if o == nil || o.Data == nil {
		var ret MarketResponseDataRelationshipsPriceListData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketResponseDataRelationshipsPriceList) GetDataOk() (*MarketResponseDataRelationshipsPriceListData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MarketResponseDataRelationshipsPriceList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given MarketResponseDataRelationshipsPriceListData and assigns it to the Data field.
func (o *MarketResponseDataRelationshipsPriceList) SetData(v MarketResponseDataRelationshipsPriceListData) {
	o.Data = &v
}

func (o MarketResponseDataRelationshipsPriceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMarketResponseDataRelationshipsPriceList struct {
	value *MarketResponseDataRelationshipsPriceList
	isSet bool
}

func (v NullableMarketResponseDataRelationshipsPriceList) Get() *MarketResponseDataRelationshipsPriceList {
	return v.value
}

func (v *NullableMarketResponseDataRelationshipsPriceList) Set(val *MarketResponseDataRelationshipsPriceList) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketResponseDataRelationshipsPriceList) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketResponseDataRelationshipsPriceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketResponseDataRelationshipsPriceList(val *MarketResponseDataRelationshipsPriceList) *NullableMarketResponseDataRelationshipsPriceList {
	return &NullableMarketResponseDataRelationshipsPriceList{value: val, isSet: true}
}

func (v NullableMarketResponseDataRelationshipsPriceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketResponseDataRelationshipsPriceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}