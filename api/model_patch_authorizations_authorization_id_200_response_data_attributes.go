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

// checks if the PATCHAuthorizationsAuthorizationId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHAuthorizationsAuthorizationId200ResponseDataAttributes{}

// PATCHAuthorizationsAuthorizationId200ResponseDataAttributes struct for PATCHAuthorizationsAuthorizationId200ResponseDataAttributes
type PATCHAuthorizationsAuthorizationId200ResponseDataAttributes struct {
	// Indicates if the transaction is successful.
	Succeeded interface{} `json:"succeeded,omitempty"`
	// Send this attribute if you want to forward a stuck transaction to succeeded and update associated order states accordingly.
	Forward interface{} `json:"_forward,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// Send this attribute if you want to create a capture for this authorization.
	Capture interface{} `json:"_capture,omitempty"`
	// Send this attribute as a value in cents if you want to overwrite the amount to be captured.
	CaptureAmountCents interface{} `json:"_capture_amount_cents,omitempty"`
	// Send this attribute if you want to create a void for this authorization.
	Void interface{} `json:"_void,omitempty"`
	// Send this attribute if you want to void a succeeded authorization of a pending order (which is left unpaid).
	Cancel interface{} `json:"_cancel,omitempty"`
}

// NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributes instantiates a new PATCHAuthorizationsAuthorizationId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributes() *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes {
	this := PATCHAuthorizationsAuthorizationId200ResponseDataAttributes{}
	return &this
}

// NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributesWithDefaults instantiates a new PATCHAuthorizationsAuthorizationId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributesWithDefaults() *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes {
	this := PATCHAuthorizationsAuthorizationId200ResponseDataAttributes{}
	return &this
}

// GetSucceeded returns the Succeeded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetSucceeded() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Succeeded
}

// GetSucceededOk returns a tuple with the Succeeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetSucceededOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Succeeded) {
		return nil, false
	}
	return &o.Succeeded, true
}

// HasSucceeded returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasSucceeded() bool {
	if o != nil && IsNil(o.Succeeded) {
		return true
	}

	return false
}

// SetSucceeded gets a reference to the given interface{} and assigns it to the Succeeded field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetSucceeded(v interface{}) {
	o.Succeeded = v
}

// GetForward returns the Forward field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetForward() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Forward
}

// GetForwardOk returns a tuple with the Forward field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetForwardOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Forward) {
		return nil, false
	}
	return &o.Forward, true
}

// HasForward returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasForward() bool {
	if o != nil && IsNil(o.Forward) {
		return true
	}

	return false
}

// SetForward gets a reference to the given interface{} and assigns it to the Forward field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetForward(v interface{}) {
	o.Forward = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetCapture returns the Capture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCapture() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Capture
}

// GetCaptureOk returns a tuple with the Capture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Capture) {
		return nil, false
	}
	return &o.Capture, true
}

// HasCapture returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasCapture() bool {
	if o != nil && IsNil(o.Capture) {
		return true
	}

	return false
}

// SetCapture gets a reference to the given interface{} and assigns it to the Capture field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetCapture(v interface{}) {
	o.Capture = v
}

// GetCaptureAmountCents returns the CaptureAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CaptureAmountCents
}

// GetCaptureAmountCentsOk returns a tuple with the CaptureAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CaptureAmountCents) {
		return nil, false
	}
	return &o.CaptureAmountCents, true
}

// HasCaptureAmountCents returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasCaptureAmountCents() bool {
	if o != nil && IsNil(o.CaptureAmountCents) {
		return true
	}

	return false
}

// SetCaptureAmountCents gets a reference to the given interface{} and assigns it to the CaptureAmountCents field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetCaptureAmountCents(v interface{}) {
	o.CaptureAmountCents = v
}

// GetVoid returns the Void field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetVoid() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Void
}

// GetVoidOk returns a tuple with the Void field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetVoidOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Void) {
		return nil, false
	}
	return &o.Void, true
}

// HasVoid returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasVoid() bool {
	if o != nil && IsNil(o.Void) {
		return true
	}

	return false
}

// SetVoid gets a reference to the given interface{} and assigns it to the Void field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetVoid(v interface{}) {
	o.Void = v
}

// GetCancel returns the Cancel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCancel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCancelOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Cancel) {
		return nil, false
	}
	return &o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasCancel() bool {
	if o != nil && IsNil(o.Cancel) {
		return true
	}

	return false
}

// SetCancel gets a reference to the given interface{} and assigns it to the Cancel field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetCancel(v interface{}) {
	o.Cancel = v
}

func (o PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Succeeded != nil {
		toSerialize["succeeded"] = o.Succeeded
	}
	if o.Forward != nil {
		toSerialize["_forward"] = o.Forward
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
	if o.Capture != nil {
		toSerialize["_capture"] = o.Capture
	}
	if o.CaptureAmountCents != nil {
		toSerialize["_capture_amount_cents"] = o.CaptureAmountCents
	}
	if o.Void != nil {
		toSerialize["_void"] = o.Void
	}
	if o.Cancel != nil {
		toSerialize["_cancel"] = o.Cancel
	}
	return toSerialize, nil
}

type NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes struct {
	value *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes) Get() *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes) Set(val *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes(val *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) *NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes {
	return &NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
