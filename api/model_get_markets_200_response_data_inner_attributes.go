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

// GETMarkets200ResponseDataInnerAttributes struct for GETMarkets200ResponseDataInnerAttributes
type GETMarkets200ResponseDataInnerAttributes struct {
	// Unique identifier for the market (numeric)
	Number *int32 `json:"number,omitempty"`
	// The market's internal name
	Name *string `json:"name,omitempty"`
	// The Facebook Pixed ID
	FacebookPixelId *string `json:"facebook_pixel_id,omitempty"`
	// The checkout URL for this market
	CheckoutUrl *string `json:"checkout_url,omitempty"`
	// The URL used to overwrite prices by an external source.
	ExternalPricesUrl *string `json:"external_prices_url,omitempty"`
	// The URL used to validate orders by an external source.
	ExternalOrderValidationUrl *string `json:"external_order_validation_url,omitempty"`
	// The shared secret used to sign the external requests payload.
	SharedSecret *string `json:"shared_secret,omitempty"`
	// Indicates if market belongs to a customer_group.
	Private *bool `json:"private,omitempty"`
	// Time at which the market was disabled.
	DisabledAt *string `json:"disabled_at,omitempty"`
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

// NewGETMarkets200ResponseDataInnerAttributes instantiates a new GETMarkets200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMarkets200ResponseDataInnerAttributes() *GETMarkets200ResponseDataInnerAttributes {
	this := GETMarkets200ResponseDataInnerAttributes{}
	return &this
}

// NewGETMarkets200ResponseDataInnerAttributesWithDefaults instantiates a new GETMarkets200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMarkets200ResponseDataInnerAttributesWithDefaults() *GETMarkets200ResponseDataInnerAttributes {
	this := GETMarkets200ResponseDataInnerAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetNumber(v int32) {
	o.Number = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetFacebookPixelId returns the FacebookPixelId field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetFacebookPixelId() string {
	if o == nil || o.FacebookPixelId == nil {
		var ret string
		return ret
	}
	return *o.FacebookPixelId
}

// GetFacebookPixelIdOk returns a tuple with the FacebookPixelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetFacebookPixelIdOk() (*string, bool) {
	if o == nil || o.FacebookPixelId == nil {
		return nil, false
	}
	return o.FacebookPixelId, true
}

// HasFacebookPixelId returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasFacebookPixelId() bool {
	if o != nil && o.FacebookPixelId != nil {
		return true
	}

	return false
}

// SetFacebookPixelId gets a reference to the given string and assigns it to the FacebookPixelId field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetFacebookPixelId(v string) {
	o.FacebookPixelId = &v
}

// GetCheckoutUrl returns the CheckoutUrl field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetCheckoutUrl() string {
	if o == nil || o.CheckoutUrl == nil {
		var ret string
		return ret
	}
	return *o.CheckoutUrl
}

// GetCheckoutUrlOk returns a tuple with the CheckoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetCheckoutUrlOk() (*string, bool) {
	if o == nil || o.CheckoutUrl == nil {
		return nil, false
	}
	return o.CheckoutUrl, true
}

// HasCheckoutUrl returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasCheckoutUrl() bool {
	if o != nil && o.CheckoutUrl != nil {
		return true
	}

	return false
}

// SetCheckoutUrl gets a reference to the given string and assigns it to the CheckoutUrl field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetCheckoutUrl(v string) {
	o.CheckoutUrl = &v
}

// GetExternalPricesUrl returns the ExternalPricesUrl field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetExternalPricesUrl() string {
	if o == nil || o.ExternalPricesUrl == nil {
		var ret string
		return ret
	}
	return *o.ExternalPricesUrl
}

// GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetExternalPricesUrlOk() (*string, bool) {
	if o == nil || o.ExternalPricesUrl == nil {
		return nil, false
	}
	return o.ExternalPricesUrl, true
}

// HasExternalPricesUrl returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasExternalPricesUrl() bool {
	if o != nil && o.ExternalPricesUrl != nil {
		return true
	}

	return false
}

// SetExternalPricesUrl gets a reference to the given string and assigns it to the ExternalPricesUrl field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetExternalPricesUrl(v string) {
	o.ExternalPricesUrl = &v
}

// GetExternalOrderValidationUrl returns the ExternalOrderValidationUrl field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetExternalOrderValidationUrl() string {
	if o == nil || o.ExternalOrderValidationUrl == nil {
		var ret string
		return ret
	}
	return *o.ExternalOrderValidationUrl
}

// GetExternalOrderValidationUrlOk returns a tuple with the ExternalOrderValidationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetExternalOrderValidationUrlOk() (*string, bool) {
	if o == nil || o.ExternalOrderValidationUrl == nil {
		return nil, false
	}
	return o.ExternalOrderValidationUrl, true
}

// HasExternalOrderValidationUrl returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasExternalOrderValidationUrl() bool {
	if o != nil && o.ExternalOrderValidationUrl != nil {
		return true
	}

	return false
}

// SetExternalOrderValidationUrl gets a reference to the given string and assigns it to the ExternalOrderValidationUrl field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetExternalOrderValidationUrl(v string) {
	o.ExternalOrderValidationUrl = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetPrivate(v bool) {
	o.Private = &v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetDisabledAt() string {
	if o == nil || o.DisabledAt == nil {
		var ret string
		return ret
	}
	return *o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetDisabledAtOk() (*string, bool) {
	if o == nil || o.DisabledAt == nil {
		return nil, false
	}
	return o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasDisabledAt() bool {
	if o != nil && o.DisabledAt != nil {
		return true
	}

	return false
}

// SetDisabledAt gets a reference to the given string and assigns it to the DisabledAt field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetDisabledAt(v string) {
	o.DisabledAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETMarkets200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETMarkets200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FacebookPixelId != nil {
		toSerialize["facebook_pixel_id"] = o.FacebookPixelId
	}
	if o.CheckoutUrl != nil {
		toSerialize["checkout_url"] = o.CheckoutUrl
	}
	if o.ExternalPricesUrl != nil {
		toSerialize["external_prices_url"] = o.ExternalPricesUrl
	}
	if o.ExternalOrderValidationUrl != nil {
		toSerialize["external_order_validation_url"] = o.ExternalOrderValidationUrl
	}
	if o.SharedSecret != nil {
		toSerialize["shared_secret"] = o.SharedSecret
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.DisabledAt != nil {
		toSerialize["disabled_at"] = o.DisabledAt
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

type NullableGETMarkets200ResponseDataInnerAttributes struct {
	value *GETMarkets200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETMarkets200ResponseDataInnerAttributes) Get() *GETMarkets200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETMarkets200ResponseDataInnerAttributes) Set(val *GETMarkets200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMarkets200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMarkets200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMarkets200ResponseDataInnerAttributes(val *GETMarkets200ResponseDataInnerAttributes) *NullableGETMarkets200ResponseDataInnerAttributes {
	return &NullableGETMarkets200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETMarkets200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMarkets200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
