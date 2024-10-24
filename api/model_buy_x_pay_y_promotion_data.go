/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the BuyXPayYPromotionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyXPayYPromotionData{}

// BuyXPayYPromotionData struct for BuyXPayYPromotionData
type BuyXPayYPromotionData struct {
	// The resource's type
	Type          interface{}                                                       `json:"type"`
	Attributes    GETBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes `json:"attributes"`
	Relationships *BuyXPayYPromotionDataRelationships                               `json:"relationships,omitempty"`
}

// NewBuyXPayYPromotionData instantiates a new BuyXPayYPromotionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyXPayYPromotionData(type_ interface{}, attributes GETBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) *BuyXPayYPromotionData {
	this := BuyXPayYPromotionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBuyXPayYPromotionDataWithDefaults instantiates a new BuyXPayYPromotionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyXPayYPromotionDataWithDefaults() *BuyXPayYPromotionData {
	this := BuyXPayYPromotionData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *BuyXPayYPromotionData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuyXPayYPromotionData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BuyXPayYPromotionData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BuyXPayYPromotionData) GetAttributes() GETBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes {
	if o == nil {
		var ret GETBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionData) GetAttributesOk() (*GETBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BuyXPayYPromotionData) SetAttributes(v GETBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BuyXPayYPromotionData) GetRelationships() BuyXPayYPromotionDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BuyXPayYPromotionDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionData) GetRelationshipsOk() (*BuyXPayYPromotionDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BuyXPayYPromotionData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BuyXPayYPromotionDataRelationships and assigns it to the Relationships field.
func (o *BuyXPayYPromotionData) SetRelationships(v BuyXPayYPromotionDataRelationships) {
	o.Relationships = &v
}

func (o BuyXPayYPromotionData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyXPayYPromotionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableBuyXPayYPromotionData struct {
	value *BuyXPayYPromotionData
	isSet bool
}

func (v NullableBuyXPayYPromotionData) Get() *BuyXPayYPromotionData {
	return v.value
}

func (v *NullableBuyXPayYPromotionData) Set(val *BuyXPayYPromotionData) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyXPayYPromotionData) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyXPayYPromotionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyXPayYPromotionData(val *BuyXPayYPromotionData) *NullableBuyXPayYPromotionData {
	return &NullableBuyXPayYPromotionData{value: val, isSet: true}
}

func (v NullableBuyXPayYPromotionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyXPayYPromotionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
