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

// CustomerPaymentSourceResponseDataRelationships struct for CustomerPaymentSourceResponseDataRelationships
type CustomerPaymentSourceResponseDataRelationships struct {
	Customer      *CouponRecipientResponseDataRelationshipsCustomer            `json:"customer,omitempty"`
	PaymentSource *CustomerPaymentSourceResponseDataRelationshipsPaymentSource `json:"payment_source,omitempty"`
}

// NewCustomerPaymentSourceResponseDataRelationships instantiates a new CustomerPaymentSourceResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSourceResponseDataRelationships() *CustomerPaymentSourceResponseDataRelationships {
	this := CustomerPaymentSourceResponseDataRelationships{}
	return &this
}

// NewCustomerPaymentSourceResponseDataRelationshipsWithDefaults instantiates a new CustomerPaymentSourceResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceResponseDataRelationshipsWithDefaults() *CustomerPaymentSourceResponseDataRelationships {
	this := CustomerPaymentSourceResponseDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CustomerPaymentSourceResponseDataRelationships) GetCustomer() CouponRecipientResponseDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientResponseDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceResponseDataRelationships) GetCustomerOk() (*CouponRecipientResponseDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *CustomerPaymentSourceResponseDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientResponseDataRelationshipsCustomer and assigns it to the Customer field.
func (o *CustomerPaymentSourceResponseDataRelationships) SetCustomer(v CouponRecipientResponseDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *CustomerPaymentSourceResponseDataRelationships) GetPaymentSource() CustomerPaymentSourceResponseDataRelationshipsPaymentSource {
	if o == nil || o.PaymentSource == nil {
		var ret CustomerPaymentSourceResponseDataRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceResponseDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceResponseDataRelationshipsPaymentSource, bool) {
	if o == nil || o.PaymentSource == nil {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *CustomerPaymentSourceResponseDataRelationships) HasPaymentSource() bool {
	if o != nil && o.PaymentSource != nil {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given CustomerPaymentSourceResponseDataRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *CustomerPaymentSourceResponseDataRelationships) SetPaymentSource(v CustomerPaymentSourceResponseDataRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

func (o CustomerPaymentSourceResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.PaymentSource != nil {
		toSerialize["payment_source"] = o.PaymentSource
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPaymentSourceResponseDataRelationships struct {
	value *CustomerPaymentSourceResponseDataRelationships
	isSet bool
}

func (v NullableCustomerPaymentSourceResponseDataRelationships) Get() *CustomerPaymentSourceResponseDataRelationships {
	return v.value
}

func (v *NullableCustomerPaymentSourceResponseDataRelationships) Set(val *CustomerPaymentSourceResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceResponseDataRelationships(val *CustomerPaymentSourceResponseDataRelationships) *NullableCustomerPaymentSourceResponseDataRelationships {
	return &NullableCustomerPaymentSourceResponseDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}