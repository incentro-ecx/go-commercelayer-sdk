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

// StripePaymentUpdateData struct for StripePaymentUpdateData
type StripePaymentUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                      `json:"id"`
	Attributes    PATCHStripePaymentsStripePaymentId200ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentUpdateDataRelationships                        `json:"relationships,omitempty"`
}

// NewStripePaymentUpdateData instantiates a new StripePaymentUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripePaymentUpdateData(type_ string, id string, attributes PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) *StripePaymentUpdateData {
	this := StripePaymentUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewStripePaymentUpdateDataWithDefaults instantiates a new StripePaymentUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripePaymentUpdateDataWithDefaults() *StripePaymentUpdateData {
	this := StripePaymentUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *StripePaymentUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StripePaymentUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StripePaymentUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *StripePaymentUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StripePaymentUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StripePaymentUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *StripePaymentUpdateData) GetAttributes() PATCHStripePaymentsStripePaymentId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHStripePaymentsStripePaymentId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StripePaymentUpdateData) GetAttributesOk() (*PATCHStripePaymentsStripePaymentId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StripePaymentUpdateData) SetAttributes(v PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StripePaymentUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StripePaymentUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentUpdateDataRelationships and assigns it to the Relationships field.
func (o *StripePaymentUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships) {
	o.Relationships = &v
}

func (o StripePaymentUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableStripePaymentUpdateData struct {
	value *StripePaymentUpdateData
	isSet bool
}

func (v NullableStripePaymentUpdateData) Get() *StripePaymentUpdateData {
	return v.value
}

func (v *NullableStripePaymentUpdateData) Set(val *StripePaymentUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableStripePaymentUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableStripePaymentUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripePaymentUpdateData(val *StripePaymentUpdateData) *NullableStripePaymentUpdateData {
	return &NullableStripePaymentUpdateData{value: val, isSet: true}
}

func (v NullableStripePaymentUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripePaymentUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
