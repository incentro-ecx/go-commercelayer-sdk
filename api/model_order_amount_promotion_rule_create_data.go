/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OrderAmountPromotionRuleCreateData struct for OrderAmountPromotionRuleCreateData
type OrderAmountPromotionRuleCreateData struct {
	// The resource's type
	Type          string                                           `json:"type"`
	Attributes    OrderAmountPromotionRuleCreateDataAttributes     `json:"attributes"`
	Relationships *OrderAmountPromotionRuleCreateDataRelationships `json:"relationships,omitempty"`
}

// NewOrderAmountPromotionRuleCreateData instantiates a new OrderAmountPromotionRuleCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmountPromotionRuleCreateData(type_ string, attributes OrderAmountPromotionRuleCreateDataAttributes) *OrderAmountPromotionRuleCreateData {
	this := OrderAmountPromotionRuleCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOrderAmountPromotionRuleCreateDataWithDefaults instantiates a new OrderAmountPromotionRuleCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmountPromotionRuleCreateDataWithDefaults() *OrderAmountPromotionRuleCreateData {
	this := OrderAmountPromotionRuleCreateData{}
	var type_ string = "order_amount_promotion_rules"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *OrderAmountPromotionRuleCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderAmountPromotionRuleCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *OrderAmountPromotionRuleCreateData) GetAttributes() OrderAmountPromotionRuleCreateDataAttributes {
	if o == nil {
		var ret OrderAmountPromotionRuleCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleCreateData) GetAttributesOk() (*OrderAmountPromotionRuleCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrderAmountPromotionRuleCreateData) SetAttributes(v OrderAmountPromotionRuleCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleCreateData) GetRelationships() OrderAmountPromotionRuleCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret OrderAmountPromotionRuleCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleCreateData) GetRelationshipsOk() (*OrderAmountPromotionRuleCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrderAmountPromotionRuleCreateDataRelationships and assigns it to the Relationships field.
func (o *OrderAmountPromotionRuleCreateData) SetRelationships(v OrderAmountPromotionRuleCreateDataRelationships) {
	o.Relationships = &v
}

func (o OrderAmountPromotionRuleCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableOrderAmountPromotionRuleCreateData struct {
	value *OrderAmountPromotionRuleCreateData
	isSet bool
}

func (v NullableOrderAmountPromotionRuleCreateData) Get() *OrderAmountPromotionRuleCreateData {
	return v.value
}

func (v *NullableOrderAmountPromotionRuleCreateData) Set(val *OrderAmountPromotionRuleCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmountPromotionRuleCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmountPromotionRuleCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmountPromotionRuleCreateData(val *OrderAmountPromotionRuleCreateData) *NullableOrderAmountPromotionRuleCreateData {
	return &NullableOrderAmountPromotionRuleCreateData{value: val, isSet: true}
}

func (v NullableOrderAmountPromotionRuleCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmountPromotionRuleCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}