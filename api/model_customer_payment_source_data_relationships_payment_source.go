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
	"fmt"
)

// CustomerPaymentSourceDataRelationshipsPaymentSource - struct for CustomerPaymentSourceDataRelationshipsPaymentSource
type CustomerPaymentSourceDataRelationshipsPaymentSource struct {
	AdyenPayment       *AdyenPayment
	BraintreePayment   *BraintreePayment
	CheckoutComPayment *CheckoutComPayment
	ExternalPayment    *ExternalPayment
	KlarnaPayment      *KlarnaPayment
	PaypalPayment      *PaypalPayment
	StripePayment      *StripePayment
	WireTransfer       *WireTransfer
}

// AdyenPaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource is a convenience function that returns AdyenPayment wrapped in CustomerPaymentSourceDataRelationshipsPaymentSource
func AdyenPaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource(v *AdyenPayment) CustomerPaymentSourceDataRelationshipsPaymentSource {
	return CustomerPaymentSourceDataRelationshipsPaymentSource{
		AdyenPayment: v,
	}
}

// BraintreePaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource is a convenience function that returns BraintreePayment wrapped in CustomerPaymentSourceDataRelationshipsPaymentSource
func BraintreePaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource(v *BraintreePayment) CustomerPaymentSourceDataRelationshipsPaymentSource {
	return CustomerPaymentSourceDataRelationshipsPaymentSource{
		BraintreePayment: v,
	}
}

// CheckoutComPaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource is a convenience function that returns CheckoutComPayment wrapped in CustomerPaymentSourceDataRelationshipsPaymentSource
func CheckoutComPaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource(v *CheckoutComPayment) CustomerPaymentSourceDataRelationshipsPaymentSource {
	return CustomerPaymentSourceDataRelationshipsPaymentSource{
		CheckoutComPayment: v,
	}
}

// ExternalPaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource is a convenience function that returns ExternalPayment wrapped in CustomerPaymentSourceDataRelationshipsPaymentSource
func ExternalPaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource(v *ExternalPayment) CustomerPaymentSourceDataRelationshipsPaymentSource {
	return CustomerPaymentSourceDataRelationshipsPaymentSource{
		ExternalPayment: v,
	}
}

// KlarnaPaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource is a convenience function that returns KlarnaPayment wrapped in CustomerPaymentSourceDataRelationshipsPaymentSource
func KlarnaPaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource(v *KlarnaPayment) CustomerPaymentSourceDataRelationshipsPaymentSource {
	return CustomerPaymentSourceDataRelationshipsPaymentSource{
		KlarnaPayment: v,
	}
}

// PaypalPaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource is a convenience function that returns PaypalPayment wrapped in CustomerPaymentSourceDataRelationshipsPaymentSource
func PaypalPaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource(v *PaypalPayment) CustomerPaymentSourceDataRelationshipsPaymentSource {
	return CustomerPaymentSourceDataRelationshipsPaymentSource{
		PaypalPayment: v,
	}
}

// StripePaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource is a convenience function that returns StripePayment wrapped in CustomerPaymentSourceDataRelationshipsPaymentSource
func StripePaymentAsCustomerPaymentSourceDataRelationshipsPaymentSource(v *StripePayment) CustomerPaymentSourceDataRelationshipsPaymentSource {
	return CustomerPaymentSourceDataRelationshipsPaymentSource{
		StripePayment: v,
	}
}

// WireTransferAsCustomerPaymentSourceDataRelationshipsPaymentSource is a convenience function that returns WireTransfer wrapped in CustomerPaymentSourceDataRelationshipsPaymentSource
func WireTransferAsCustomerPaymentSourceDataRelationshipsPaymentSource(v *WireTransfer) CustomerPaymentSourceDataRelationshipsPaymentSource {
	return CustomerPaymentSourceDataRelationshipsPaymentSource{
		WireTransfer: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CustomerPaymentSourceDataRelationshipsPaymentSource) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AdyenPayment
	err = newStrictDecoder(data).Decode(&dst.AdyenPayment)
	if err == nil {
		jsonAdyenPayment, _ := json.Marshal(dst.AdyenPayment)
		if string(jsonAdyenPayment) == "{}" { // empty struct
			dst.AdyenPayment = nil
		} else {
			match++
		}
	} else {
		dst.AdyenPayment = nil
	}

	// try to unmarshal data into BraintreePayment
	err = newStrictDecoder(data).Decode(&dst.BraintreePayment)
	if err == nil {
		jsonBraintreePayment, _ := json.Marshal(dst.BraintreePayment)
		if string(jsonBraintreePayment) == "{}" { // empty struct
			dst.BraintreePayment = nil
		} else {
			match++
		}
	} else {
		dst.BraintreePayment = nil
	}

	// try to unmarshal data into CheckoutComPayment
	err = newStrictDecoder(data).Decode(&dst.CheckoutComPayment)
	if err == nil {
		jsonCheckoutComPayment, _ := json.Marshal(dst.CheckoutComPayment)
		if string(jsonCheckoutComPayment) == "{}" { // empty struct
			dst.CheckoutComPayment = nil
		} else {
			match++
		}
	} else {
		dst.CheckoutComPayment = nil
	}

	// try to unmarshal data into ExternalPayment
	err = newStrictDecoder(data).Decode(&dst.ExternalPayment)
	if err == nil {
		jsonExternalPayment, _ := json.Marshal(dst.ExternalPayment)
		if string(jsonExternalPayment) == "{}" { // empty struct
			dst.ExternalPayment = nil
		} else {
			match++
		}
	} else {
		dst.ExternalPayment = nil
	}

	// try to unmarshal data into KlarnaPayment
	err = newStrictDecoder(data).Decode(&dst.KlarnaPayment)
	if err == nil {
		jsonKlarnaPayment, _ := json.Marshal(dst.KlarnaPayment)
		if string(jsonKlarnaPayment) == "{}" { // empty struct
			dst.KlarnaPayment = nil
		} else {
			match++
		}
	} else {
		dst.KlarnaPayment = nil
	}

	// try to unmarshal data into PaypalPayment
	err = newStrictDecoder(data).Decode(&dst.PaypalPayment)
	if err == nil {
		jsonPaypalPayment, _ := json.Marshal(dst.PaypalPayment)
		if string(jsonPaypalPayment) == "{}" { // empty struct
			dst.PaypalPayment = nil
		} else {
			match++
		}
	} else {
		dst.PaypalPayment = nil
	}

	// try to unmarshal data into StripePayment
	err = newStrictDecoder(data).Decode(&dst.StripePayment)
	if err == nil {
		jsonStripePayment, _ := json.Marshal(dst.StripePayment)
		if string(jsonStripePayment) == "{}" { // empty struct
			dst.StripePayment = nil
		} else {
			match++
		}
	} else {
		dst.StripePayment = nil
	}

	// try to unmarshal data into WireTransfer
	err = newStrictDecoder(data).Decode(&dst.WireTransfer)
	if err == nil {
		jsonWireTransfer, _ := json.Marshal(dst.WireTransfer)
		if string(jsonWireTransfer) == "{}" { // empty struct
			dst.WireTransfer = nil
		} else {
			match++
		}
	} else {
		dst.WireTransfer = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AdyenPayment = nil
		dst.BraintreePayment = nil
		dst.CheckoutComPayment = nil
		dst.ExternalPayment = nil
		dst.KlarnaPayment = nil
		dst.PaypalPayment = nil
		dst.StripePayment = nil
		dst.WireTransfer = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CustomerPaymentSourceDataRelationshipsPaymentSource)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CustomerPaymentSourceDataRelationshipsPaymentSource)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CustomerPaymentSourceDataRelationshipsPaymentSource) MarshalJSON() ([]byte, error) {
	if src.AdyenPayment != nil {
		return json.Marshal(&src.AdyenPayment)
	}

	if src.BraintreePayment != nil {
		return json.Marshal(&src.BraintreePayment)
	}

	if src.CheckoutComPayment != nil {
		return json.Marshal(&src.CheckoutComPayment)
	}

	if src.ExternalPayment != nil {
		return json.Marshal(&src.ExternalPayment)
	}

	if src.KlarnaPayment != nil {
		return json.Marshal(&src.KlarnaPayment)
	}

	if src.PaypalPayment != nil {
		return json.Marshal(&src.PaypalPayment)
	}

	if src.StripePayment != nil {
		return json.Marshal(&src.StripePayment)
	}

	if src.WireTransfer != nil {
		return json.Marshal(&src.WireTransfer)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CustomerPaymentSourceDataRelationshipsPaymentSource) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AdyenPayment != nil {
		return obj.AdyenPayment
	}

	if obj.BraintreePayment != nil {
		return obj.BraintreePayment
	}

	if obj.CheckoutComPayment != nil {
		return obj.CheckoutComPayment
	}

	if obj.ExternalPayment != nil {
		return obj.ExternalPayment
	}

	if obj.KlarnaPayment != nil {
		return obj.KlarnaPayment
	}

	if obj.PaypalPayment != nil {
		return obj.PaypalPayment
	}

	if obj.StripePayment != nil {
		return obj.StripePayment
	}

	if obj.WireTransfer != nil {
		return obj.WireTransfer
	}

	// all schemas are nil
	return nil
}

type NullableCustomerPaymentSourceDataRelationshipsPaymentSource struct {
	value *CustomerPaymentSourceDataRelationshipsPaymentSource
	isSet bool
}

func (v NullableCustomerPaymentSourceDataRelationshipsPaymentSource) Get() *CustomerPaymentSourceDataRelationshipsPaymentSource {
	return v.value
}

func (v *NullableCustomerPaymentSourceDataRelationshipsPaymentSource) Set(val *CustomerPaymentSourceDataRelationshipsPaymentSource) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceDataRelationshipsPaymentSource) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceDataRelationshipsPaymentSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceDataRelationshipsPaymentSource(val *CustomerPaymentSourceDataRelationshipsPaymentSource) *NullableCustomerPaymentSourceDataRelationshipsPaymentSource {
	return &NullableCustomerPaymentSourceDataRelationshipsPaymentSource{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceDataRelationshipsPaymentSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceDataRelationshipsPaymentSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
