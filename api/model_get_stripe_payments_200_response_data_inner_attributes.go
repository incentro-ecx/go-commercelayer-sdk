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

// GETStripePayments200ResponseDataInnerAttributes struct for GETStripePayments200ResponseDataInnerAttributes
type GETStripePayments200ResponseDataInnerAttributes struct {
	// The Stripe payment intent client secret. Required to create a charge through Stripe.js.
	ClientSecret *string `json:"client_secret,omitempty"`
	// The Stripe publishable API key.
	PublishableKey *string `json:"publishable_key,omitempty"`
	// Stripe payment options: 'setup_future_usage', 'customer', and 'payment_method'
	Options map[string]interface{} `json:"options,omitempty"`
	// Stripe 'payment_method', set by webhook.
	PaymentMethod map[string]interface{} `json:"payment_method,omitempty"`
	// Indicates if the order current amount differs form the one of the created payment intent.
	MismatchedAmounts *bool `json:"mismatched_amounts,omitempty"`
	// The amount of the associated payment intent, in cents.
	IntentAmountCents *int32 `json:"intent_amount_cents,omitempty"`
	// The amount of the associated payment intent, float.
	IntentAmountFloat *float32 `json:"intent_amount_float,omitempty"`
	// The amount of the associated payment intent, formatted.
	FormattedIntentAmount *string `json:"formatted_intent_amount,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewGETStripePayments200ResponseDataInnerAttributes instantiates a new GETStripePayments200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStripePayments200ResponseDataInnerAttributes() *GETStripePayments200ResponseDataInnerAttributes {
	this := GETStripePayments200ResponseDataInnerAttributes{}
	return &this
}

// NewGETStripePayments200ResponseDataInnerAttributesWithDefaults instantiates a new GETStripePayments200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStripePayments200ResponseDataInnerAttributesWithDefaults() *GETStripePayments200ResponseDataInnerAttributes {
	this := GETStripePayments200ResponseDataInnerAttributes{}
	return &this
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetPublishableKey returns the PublishableKey field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetPublishableKey() string {
	if o == nil || o.PublishableKey == nil {
		var ret string
		return ret
	}
	return *o.PublishableKey
}

// GetPublishableKeyOk returns a tuple with the PublishableKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetPublishableKeyOk() (*string, bool) {
	if o == nil || o.PublishableKey == nil {
		return nil, false
	}
	return o.PublishableKey, true
}

// HasPublishableKey returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasPublishableKey() bool {
	if o != nil && o.PublishableKey != nil {
		return true
	}

	return false
}

// SetPublishableKey gets a reference to the given string and assigns it to the PublishableKey field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetPublishableKey(v string) {
	o.PublishableKey = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetPaymentMethod() map[string]interface{} {
	if o == nil || o.PaymentMethod == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetPaymentMethodOk() (map[string]interface{}, bool) {
	if o == nil || o.PaymentMethod == nil {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod != nil {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given map[string]interface{} and assigns it to the PaymentMethod field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetPaymentMethod(v map[string]interface{}) {
	o.PaymentMethod = v
}

// GetMismatchedAmounts returns the MismatchedAmounts field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetMismatchedAmounts() bool {
	if o == nil || o.MismatchedAmounts == nil {
		var ret bool
		return ret
	}
	return *o.MismatchedAmounts
}

// GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetMismatchedAmountsOk() (*bool, bool) {
	if o == nil || o.MismatchedAmounts == nil {
		return nil, false
	}
	return o.MismatchedAmounts, true
}

// HasMismatchedAmounts returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasMismatchedAmounts() bool {
	if o != nil && o.MismatchedAmounts != nil {
		return true
	}

	return false
}

// SetMismatchedAmounts gets a reference to the given bool and assigns it to the MismatchedAmounts field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetMismatchedAmounts(v bool) {
	o.MismatchedAmounts = &v
}

// GetIntentAmountCents returns the IntentAmountCents field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetIntentAmountCents() int32 {
	if o == nil || o.IntentAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.IntentAmountCents
}

// GetIntentAmountCentsOk returns a tuple with the IntentAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetIntentAmountCentsOk() (*int32, bool) {
	if o == nil || o.IntentAmountCents == nil {
		return nil, false
	}
	return o.IntentAmountCents, true
}

// HasIntentAmountCents returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasIntentAmountCents() bool {
	if o != nil && o.IntentAmountCents != nil {
		return true
	}

	return false
}

// SetIntentAmountCents gets a reference to the given int32 and assigns it to the IntentAmountCents field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetIntentAmountCents(v int32) {
	o.IntentAmountCents = &v
}

// GetIntentAmountFloat returns the IntentAmountFloat field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetIntentAmountFloat() float32 {
	if o == nil || o.IntentAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.IntentAmountFloat
}

// GetIntentAmountFloatOk returns a tuple with the IntentAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetIntentAmountFloatOk() (*float32, bool) {
	if o == nil || o.IntentAmountFloat == nil {
		return nil, false
	}
	return o.IntentAmountFloat, true
}

// HasIntentAmountFloat returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasIntentAmountFloat() bool {
	if o != nil && o.IntentAmountFloat != nil {
		return true
	}

	return false
}

// SetIntentAmountFloat gets a reference to the given float32 and assigns it to the IntentAmountFloat field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetIntentAmountFloat(v float32) {
	o.IntentAmountFloat = &v
}

// GetFormattedIntentAmount returns the FormattedIntentAmount field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetFormattedIntentAmount() string {
	if o == nil || o.FormattedIntentAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedIntentAmount
}

// GetFormattedIntentAmountOk returns a tuple with the FormattedIntentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetFormattedIntentAmountOk() (*string, bool) {
	if o == nil || o.FormattedIntentAmount == nil {
		return nil, false
	}
	return o.FormattedIntentAmount, true
}

// HasFormattedIntentAmount returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasFormattedIntentAmount() bool {
	if o != nil && o.FormattedIntentAmount != nil {
		return true
	}

	return false
}

// SetFormattedIntentAmount gets a reference to the given string and assigns it to the FormattedIntentAmount field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetFormattedIntentAmount(v string) {
	o.FormattedIntentAmount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETStripePayments200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETStripePayments200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETStripePayments200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.PublishableKey != nil {
		toSerialize["publishable_key"] = o.PublishableKey
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.PaymentMethod != nil {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if o.MismatchedAmounts != nil {
		toSerialize["mismatched_amounts"] = o.MismatchedAmounts
	}
	if o.IntentAmountCents != nil {
		toSerialize["intent_amount_cents"] = o.IntentAmountCents
	}
	if o.IntentAmountFloat != nil {
		toSerialize["intent_amount_float"] = o.IntentAmountFloat
	}
	if o.FormattedIntentAmount != nil {
		toSerialize["formatted_intent_amount"] = o.FormattedIntentAmount
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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

type NullableGETStripePayments200ResponseDataInnerAttributes struct {
	value *GETStripePayments200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETStripePayments200ResponseDataInnerAttributes) Get() *GETStripePayments200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETStripePayments200ResponseDataInnerAttributes) Set(val *GETStripePayments200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStripePayments200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStripePayments200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStripePayments200ResponseDataInnerAttributes(val *GETStripePayments200ResponseDataInnerAttributes) *NullableGETStripePayments200ResponseDataInnerAttributes {
	return &NullableGETStripePayments200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETStripePayments200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStripePayments200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
