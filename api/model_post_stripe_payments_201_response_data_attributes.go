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

// checks if the POSTStripePayments201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStripePayments201ResponseDataAttributes{}

// POSTStripePayments201ResponseDataAttributes struct for POSTStripePayments201ResponseDataAttributes
type POSTStripePayments201ResponseDataAttributes struct {
	// The Stripe payment intent ID. Required to identify a payment session on stripe.
	StripeId interface{} `json:"stripe_id,omitempty"`
	// The Stripe payment intent client secret. Required to create a charge through Stripe.js.
	ClientSecret interface{} `json:"client_secret,omitempty"`
	// Stripe payment options: 'customer', 'payment_method', 'return_url', etc. Check Stripe payment intent API for more details.
	Options interface{} `json:"options,omitempty"`
	// The URL to return to when a redirect payment is completed.
	ReturnUrl interface{} `json:"return_url,omitempty"`
	// The email address to send the receipt to.
	ReceiptEmail interface{} `json:"receipt_email,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTStripePayments201ResponseDataAttributes instantiates a new POSTStripePayments201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStripePayments201ResponseDataAttributes() *POSTStripePayments201ResponseDataAttributes {
	this := POSTStripePayments201ResponseDataAttributes{}
	return &this
}

// NewPOSTStripePayments201ResponseDataAttributesWithDefaults instantiates a new POSTStripePayments201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStripePayments201ResponseDataAttributesWithDefaults() *POSTStripePayments201ResponseDataAttributes {
	this := POSTStripePayments201ResponseDataAttributes{}
	return &this
}

// GetStripeId returns the StripeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStripePayments201ResponseDataAttributes) GetStripeId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StripeId
}

// GetStripeIdOk returns a tuple with the StripeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStripePayments201ResponseDataAttributes) GetStripeIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StripeId) {
		return nil, false
	}
	return &o.StripeId, true
}

// HasStripeId returns a boolean if a field has been set.
func (o *POSTStripePayments201ResponseDataAttributes) HasStripeId() bool {
	if o != nil && IsNil(o.StripeId) {
		return true
	}

	return false
}

// SetStripeId gets a reference to the given interface{} and assigns it to the StripeId field.
func (o *POSTStripePayments201ResponseDataAttributes) SetStripeId(v interface{}) {
	o.StripeId = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStripePayments201ResponseDataAttributes) GetClientSecret() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStripePayments201ResponseDataAttributes) GetClientSecretOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return &o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *POSTStripePayments201ResponseDataAttributes) HasClientSecret() bool {
	if o != nil && IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given interface{} and assigns it to the ClientSecret field.
func (o *POSTStripePayments201ResponseDataAttributes) SetClientSecret(v interface{}) {
	o.ClientSecret = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStripePayments201ResponseDataAttributes) GetOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStripePayments201ResponseDataAttributes) GetOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *POSTStripePayments201ResponseDataAttributes) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given interface{} and assigns it to the Options field.
func (o *POSTStripePayments201ResponseDataAttributes) SetOptions(v interface{}) {
	o.Options = v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStripePayments201ResponseDataAttributes) GetReturnUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStripePayments201ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return &o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *POSTStripePayments201ResponseDataAttributes) HasReturnUrl() bool {
	if o != nil && IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given interface{} and assigns it to the ReturnUrl field.
func (o *POSTStripePayments201ResponseDataAttributes) SetReturnUrl(v interface{}) {
	o.ReturnUrl = v
}

// GetReceiptEmail returns the ReceiptEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStripePayments201ResponseDataAttributes) GetReceiptEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReceiptEmail
}

// GetReceiptEmailOk returns a tuple with the ReceiptEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStripePayments201ResponseDataAttributes) GetReceiptEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReceiptEmail) {
		return nil, false
	}
	return &o.ReceiptEmail, true
}

// HasReceiptEmail returns a boolean if a field has been set.
func (o *POSTStripePayments201ResponseDataAttributes) HasReceiptEmail() bool {
	if o != nil && IsNil(o.ReceiptEmail) {
		return true
	}

	return false
}

// SetReceiptEmail gets a reference to the given interface{} and assigns it to the ReceiptEmail field.
func (o *POSTStripePayments201ResponseDataAttributes) SetReceiptEmail(v interface{}) {
	o.ReceiptEmail = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStripePayments201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStripePayments201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTStripePayments201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTStripePayments201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStripePayments201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStripePayments201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTStripePayments201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTStripePayments201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStripePayments201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStripePayments201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTStripePayments201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTStripePayments201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTStripePayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStripePayments201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.StripeId != nil {
		toSerialize["stripe_id"] = o.StripeId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.ReturnUrl != nil {
		toSerialize["return_url"] = o.ReturnUrl
	}
	if o.ReceiptEmail != nil {
		toSerialize["receipt_email"] = o.ReceiptEmail
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
	return toSerialize, nil
}

type NullablePOSTStripePayments201ResponseDataAttributes struct {
	value *POSTStripePayments201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTStripePayments201ResponseDataAttributes) Get() *POSTStripePayments201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTStripePayments201ResponseDataAttributes) Set(val *POSTStripePayments201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStripePayments201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStripePayments201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStripePayments201ResponseDataAttributes(val *POSTStripePayments201ResponseDataAttributes) *NullablePOSTStripePayments201ResponseDataAttributes {
	return &NullablePOSTStripePayments201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTStripePayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStripePayments201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
