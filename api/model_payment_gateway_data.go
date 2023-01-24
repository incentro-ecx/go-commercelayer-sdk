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

// PaymentGatewayData struct for PaymentGatewayData
type PaymentGatewayData struct {
	// The resource's type
	Type          string                                          `json:"type"`
	Attributes    GETKlarnaGateways200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *ManualGatewayDataRelationships                 `json:"relationships,omitempty"`
}

// NewPaymentGatewayData instantiates a new PaymentGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentGatewayData(type_ string, attributes GETKlarnaGateways200ResponseDataInnerAttributes) *PaymentGatewayData {
	this := PaymentGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPaymentGatewayDataWithDefaults instantiates a new PaymentGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentGatewayDataWithDefaults() *PaymentGatewayData {
	this := PaymentGatewayData{}
	return &this
}

// GetType returns the Type field value
func (o *PaymentGatewayData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentGatewayData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentGatewayData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PaymentGatewayData) GetAttributes() GETKlarnaGateways200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETKlarnaGateways200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PaymentGatewayData) GetAttributesOk() (*GETKlarnaGateways200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PaymentGatewayData) SetAttributes(v GETKlarnaGateways200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PaymentGatewayData) GetRelationships() ManualGatewayDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ManualGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayData) GetRelationshipsOk() (*ManualGatewayDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PaymentGatewayData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ManualGatewayDataRelationships and assigns it to the Relationships field.
func (o *PaymentGatewayData) SetRelationships(v ManualGatewayDataRelationships) {
	o.Relationships = &v
}

func (o PaymentGatewayData) MarshalJSON() ([]byte, error) {
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

type NullablePaymentGatewayData struct {
	value *PaymentGatewayData
	isSet bool
}

func (v NullablePaymentGatewayData) Get() *PaymentGatewayData {
	return v.value
}

func (v *NullablePaymentGatewayData) Set(val *PaymentGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentGatewayData(val *PaymentGatewayData) *NullablePaymentGatewayData {
	return &NullablePaymentGatewayData{value: val, isSet: true}
}

func (v NullablePaymentGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
