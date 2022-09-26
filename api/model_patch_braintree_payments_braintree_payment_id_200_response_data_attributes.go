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

// PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes struct for PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes
type PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes struct {
	// The Braintree payment method nonce. Sent by the Braintree JS SDK.
	PaymentMethodNonce *string `json:"payment_method_nonce,omitempty"`
	// The Braintree payment ID used by local payment and sent by the Braintree JS SDK.
	PaymentId *string `json:"payment_id,omitempty"`
	// Indicates if the payment is local, in such case Braintree will trigger a webhook call passing the \"payment_id\" and \"payment_method_nonce\" in order to complete the transaction.
	Local *bool `json:"local,omitempty"`
	// Braintree payment options: 'customer_id' and 'payment_method_token'
	Options map[string]interface{} `json:"options,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes instantiates a new PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes() *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes {
	this := PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes{}
	return &this
}

// NewPATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributesWithDefaults() *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes {
	this := PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes{}
	return &this
}

// GetPaymentMethodNonce returns the PaymentMethodNonce field value if set, zero value otherwise.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentMethodNonce() string {
	if o == nil || o.PaymentMethodNonce == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodNonce
}

// GetPaymentMethodNonceOk returns a tuple with the PaymentMethodNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentMethodNonceOk() (*string, bool) {
	if o == nil || o.PaymentMethodNonce == nil {
		return nil, false
	}
	return o.PaymentMethodNonce, true
}

// HasPaymentMethodNonce returns a boolean if a field has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasPaymentMethodNonce() bool {
	if o != nil && o.PaymentMethodNonce != nil {
		return true
	}

	return false
}

// SetPaymentMethodNonce gets a reference to the given string and assigns it to the PaymentMethodNonce field.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentMethodNonce(v string) {
	o.PaymentMethodNonce = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentId() string {
	if o == nil || o.PaymentId == nil {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentIdOk() (*string, bool) {
	if o == nil || o.PaymentId == nil {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasPaymentId() bool {
	if o != nil && o.PaymentId != nil {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetLocal() bool {
	if o == nil || o.Local == nil {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetLocalOk() (*bool, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetLocal(v bool) {
	o.Local = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethodNonce != nil {
		toSerialize["payment_method_nonce"] = o.PaymentMethodNonce
	}
	if o.PaymentId != nil {
		toSerialize["payment_id"] = o.PaymentId
	}
	if o.Local != nil {
		toSerialize["local"] = o.Local
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes struct {
	value *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) Get() *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) Set(val *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes(val *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) *NullablePATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes {
	return &NullablePATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}