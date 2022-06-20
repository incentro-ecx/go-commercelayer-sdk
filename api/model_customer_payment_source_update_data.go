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

// CustomerPaymentSourceUpdateData struct for CustomerPaymentSourceUpdateData
type CustomerPaymentSourceUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                  `json:"id"`
	Attributes    AdyenPaymentCreateDataAttributes        `json:"attributes"`
	Relationships *CustomerPaymentSourceDataRelationships `json:"relationships,omitempty"`
}

// NewCustomerPaymentSourceUpdateData instantiates a new CustomerPaymentSourceUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSourceUpdateData(type_ string, id string, attributes AdyenPaymentCreateDataAttributes) *CustomerPaymentSourceUpdateData {
	this := CustomerPaymentSourceUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCustomerPaymentSourceUpdateDataWithDefaults instantiates a new CustomerPaymentSourceUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceUpdateDataWithDefaults() *CustomerPaymentSourceUpdateData {
	this := CustomerPaymentSourceUpdateData{}
	var type_ string = "customer_payment_sources"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CustomerPaymentSourceUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerPaymentSourceUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CustomerPaymentSourceUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerPaymentSourceUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerPaymentSourceUpdateData) GetAttributes() AdyenPaymentCreateDataAttributes {
	if o == nil {
		var ret AdyenPaymentCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceUpdateData) GetAttributesOk() (*AdyenPaymentCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerPaymentSourceUpdateData) SetAttributes(v AdyenPaymentCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerPaymentSourceUpdateData) GetRelationships() CustomerPaymentSourceDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CustomerPaymentSourceDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceUpdateData) GetRelationshipsOk() (*CustomerPaymentSourceDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerPaymentSourceUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerPaymentSourceDataRelationships and assigns it to the Relationships field.
func (o *CustomerPaymentSourceUpdateData) SetRelationships(v CustomerPaymentSourceDataRelationships) {
	o.Relationships = &v
}

func (o CustomerPaymentSourceUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableCustomerPaymentSourceUpdateData struct {
	value *CustomerPaymentSourceUpdateData
	isSet bool
}

func (v NullableCustomerPaymentSourceUpdateData) Get() *CustomerPaymentSourceUpdateData {
	return v.value
}

func (v *NullableCustomerPaymentSourceUpdateData) Set(val *CustomerPaymentSourceUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceUpdateData(val *CustomerPaymentSourceUpdateData) *NullableCustomerPaymentSourceUpdateData {
	return &NullableCustomerPaymentSourceUpdateData{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}