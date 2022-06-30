/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OrderSubscriptionCreateData struct for OrderSubscriptionCreateData
type OrderSubscriptionCreateData struct {
	// The resource's type
	Type          string                                    `json:"type"`
	Attributes    OrderSubscriptionCreateDataAttributes     `json:"attributes"`
	Relationships *OrderSubscriptionCreateDataRelationships `json:"relationships,omitempty"`
}

// NewOrderSubscriptionCreateData instantiates a new OrderSubscriptionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionCreateData(type_ string, attributes OrderSubscriptionCreateDataAttributes) *OrderSubscriptionCreateData {
	this := OrderSubscriptionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOrderSubscriptionCreateDataWithDefaults instantiates a new OrderSubscriptionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionCreateDataWithDefaults() *OrderSubscriptionCreateData {
	this := OrderSubscriptionCreateData{}
	var type_ string = "order_subscriptions"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *OrderSubscriptionCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderSubscriptionCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *OrderSubscriptionCreateData) GetAttributes() OrderSubscriptionCreateDataAttributes {
	if o == nil {
		var ret OrderSubscriptionCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionCreateData) GetAttributesOk() (*OrderSubscriptionCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrderSubscriptionCreateData) SetAttributes(v OrderSubscriptionCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrderSubscriptionCreateData) GetRelationships() OrderSubscriptionCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret OrderSubscriptionCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionCreateData) GetRelationshipsOk() (*OrderSubscriptionCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrderSubscriptionCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrderSubscriptionCreateDataRelationships and assigns it to the Relationships field.
func (o *OrderSubscriptionCreateData) SetRelationships(v OrderSubscriptionCreateDataRelationships) {
	o.Relationships = &v
}

func (o OrderSubscriptionCreateData) MarshalJSON() ([]byte, error) {
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

type NullableOrderSubscriptionCreateData struct {
	value *OrderSubscriptionCreateData
	isSet bool
}

func (v NullableOrderSubscriptionCreateData) Get() *OrderSubscriptionCreateData {
	return v.value
}

func (v *NullableOrderSubscriptionCreateData) Set(val *OrderSubscriptionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionCreateData(val *OrderSubscriptionCreateData) *NullableOrderSubscriptionCreateData {
	return &NullableOrderSubscriptionCreateData{value: val, isSet: true}
}

func (v NullableOrderSubscriptionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
