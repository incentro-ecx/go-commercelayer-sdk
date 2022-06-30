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

// CustomerPaymentSourceDataRelationships struct for CustomerPaymentSourceDataRelationships
type CustomerPaymentSourceDataRelationships struct {
	Customer      *CouponRecipientDataRelationshipsCustomer            `json:"customer,omitempty"`
	PaymentSource *CustomerPaymentSourceDataRelationshipsPaymentSource `json:"payment_source,omitempty"`
}

// NewCustomerPaymentSourceDataRelationships instantiates a new CustomerPaymentSourceDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSourceDataRelationships() *CustomerPaymentSourceDataRelationships {
	this := CustomerPaymentSourceDataRelationships{}
	return &this
}

// NewCustomerPaymentSourceDataRelationshipsWithDefaults instantiates a new CustomerPaymentSourceDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceDataRelationshipsWithDefaults() *CustomerPaymentSourceDataRelationships {
	this := CustomerPaymentSourceDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CustomerPaymentSourceDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *CustomerPaymentSourceDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customer field.
func (o *CustomerPaymentSourceDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *CustomerPaymentSourceDataRelationships) GetPaymentSource() CustomerPaymentSourceDataRelationshipsPaymentSource {
	if o == nil || o.PaymentSource == nil {
		var ret CustomerPaymentSourceDataRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceDataRelationshipsPaymentSource, bool) {
	if o == nil || o.PaymentSource == nil {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *CustomerPaymentSourceDataRelationships) HasPaymentSource() bool {
	if o != nil && o.PaymentSource != nil {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given CustomerPaymentSourceDataRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *CustomerPaymentSourceDataRelationships) SetPaymentSource(v CustomerPaymentSourceDataRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

func (o CustomerPaymentSourceDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.PaymentSource != nil {
		toSerialize["payment_source"] = o.PaymentSource
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPaymentSourceDataRelationships struct {
	value *CustomerPaymentSourceDataRelationships
	isSet bool
}

func (v NullableCustomerPaymentSourceDataRelationships) Get() *CustomerPaymentSourceDataRelationships {
	return v.value
}

func (v *NullableCustomerPaymentSourceDataRelationships) Set(val *CustomerPaymentSourceDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceDataRelationships(val *CustomerPaymentSourceDataRelationships) *NullableCustomerPaymentSourceDataRelationships {
	return &NullableCustomerPaymentSourceDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
