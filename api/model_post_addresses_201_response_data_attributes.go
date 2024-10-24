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

// checks if the POSTAddresses201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAddresses201ResponseDataAttributes{}

// POSTAddresses201ResponseDataAttributes struct for POSTAddresses201ResponseDataAttributes
type POSTAddresses201ResponseDataAttributes struct {
	// Indicates if it's a business or a personal address.
	Business interface{} `json:"business,omitempty"`
	// Address first name (personal).
	FirstName interface{} `json:"first_name,omitempty"`
	// Address last name (personal).
	LastName interface{} `json:"last_name,omitempty"`
	// Address company name (business).
	Company interface{} `json:"company,omitempty"`
	// Address line 1, i.e. Street address, PO Box.
	Line1 interface{} `json:"line_1"`
	// Address line 2, i.e. Apartment, Suite, Building.
	Line2 interface{} `json:"line_2,omitempty"`
	// Address city.
	City interface{} `json:"city"`
	// ZIP or postal code.
	ZipCode interface{} `json:"zip_code,omitempty"`
	// State, province or region code.
	StateCode interface{} `json:"state_code"`
	// The international 2-letter country code as defined by the ISO 3166-1 standard.
	CountryCode interface{} `json:"country_code"`
	// Phone number (including extension).
	Phone interface{} `json:"phone"`
	// Email address.
	Email interface{} `json:"email,omitempty"`
	// A free notes attached to the address. When used as a shipping address, this can be useful to let the customers add specific delivery instructions.
	Notes interface{} `json:"notes,omitempty"`
	// The address geocoded latitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order's market.
	Lat interface{} `json:"lat,omitempty"`
	// The address geocoded longitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order's market.
	Lng interface{} `json:"lng,omitempty"`
	// Customer's billing information (i.e. VAT number, codice fiscale).
	BillingInfo interface{} `json:"billing_info,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTAddresses201ResponseDataAttributes instantiates a new POSTAddresses201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAddresses201ResponseDataAttributes(line1 interface{}, city interface{}, stateCode interface{}, countryCode interface{}, phone interface{}) *POSTAddresses201ResponseDataAttributes {
	this := POSTAddresses201ResponseDataAttributes{}
	this.Line1 = line1
	this.City = city
	this.StateCode = stateCode
	this.CountryCode = countryCode
	this.Phone = phone
	return &this
}

// NewPOSTAddresses201ResponseDataAttributesWithDefaults instantiates a new POSTAddresses201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAddresses201ResponseDataAttributesWithDefaults() *POSTAddresses201ResponseDataAttributes {
	this := POSTAddresses201ResponseDataAttributes{}
	return &this
}

// GetBusiness returns the Business field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetBusiness() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Business
}

// GetBusinessOk returns a tuple with the Business field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetBusinessOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Business) {
		return nil, false
	}
	return &o.Business, true
}

// HasBusiness returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasBusiness() bool {
	if o != nil && IsNil(o.Business) {
		return true
	}

	return false
}

// SetBusiness gets a reference to the given interface{} and assigns it to the Business field.
func (o *POSTAddresses201ResponseDataAttributes) SetBusiness(v interface{}) {
	o.Business = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetFirstName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetFirstNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return &o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasFirstName() bool {
	if o != nil && IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given interface{} and assigns it to the FirstName field.
func (o *POSTAddresses201ResponseDataAttributes) SetFirstName(v interface{}) {
	o.FirstName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetLastName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetLastNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return &o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasLastName() bool {
	if o != nil && IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given interface{} and assigns it to the LastName field.
func (o *POSTAddresses201ResponseDataAttributes) SetLastName(v interface{}) {
	o.LastName = v
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetCompany() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetCompanyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return &o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasCompany() bool {
	if o != nil && IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given interface{} and assigns it to the Company field.
func (o *POSTAddresses201ResponseDataAttributes) SetCompany(v interface{}) {
	o.Company = v
}

// GetLine1 returns the Line1 field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetLine1() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetLine1Ok() (*interface{}, bool) {
	if o == nil || IsNil(o.Line1) {
		return nil, false
	}
	return &o.Line1, true
}

// SetLine1 sets field value
func (o *POSTAddresses201ResponseDataAttributes) SetLine1(v interface{}) {
	o.Line1 = v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetLine2() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetLine2Ok() (*interface{}, bool) {
	if o == nil || IsNil(o.Line2) {
		return nil, false
	}
	return &o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasLine2() bool {
	if o != nil && IsNil(o.Line2) {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given interface{} and assigns it to the Line2 field.
func (o *POSTAddresses201ResponseDataAttributes) SetLine2(v interface{}) {
	o.Line2 = v
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetCity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetCityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *POSTAddresses201ResponseDataAttributes) SetCity(v interface{}) {
	o.City = v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetZipCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetZipCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return &o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasZipCode() bool {
	if o != nil && IsNil(o.ZipCode) {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given interface{} and assigns it to the ZipCode field.
func (o *POSTAddresses201ResponseDataAttributes) SetZipCode(v interface{}) {
	o.ZipCode = v
}

// GetStateCode returns the StateCode field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetStateCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.StateCode
}

// GetStateCodeOk returns a tuple with the StateCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetStateCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StateCode) {
		return nil, false
	}
	return &o.StateCode, true
}

// SetStateCode sets field value
func (o *POSTAddresses201ResponseDataAttributes) SetStateCode(v interface{}) {
	o.StateCode = v
}

// GetCountryCode returns the CountryCode field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetCountryCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetCountryCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *POSTAddresses201ResponseDataAttributes) SetCountryCode(v interface{}) {
	o.CountryCode = v
}

// GetPhone returns the Phone field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetPhone() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetPhoneOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *POSTAddresses201ResponseDataAttributes) SetPhone(v interface{}) {
	o.Phone = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return &o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasEmail() bool {
	if o != nil && IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given interface{} and assigns it to the Email field.
func (o *POSTAddresses201ResponseDataAttributes) SetEmail(v interface{}) {
	o.Email = v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetNotes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetNotesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return &o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasNotes() bool {
	if o != nil && IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given interface{} and assigns it to the Notes field.
func (o *POSTAddresses201ResponseDataAttributes) SetNotes(v interface{}) {
	o.Notes = v
}

// GetLat returns the Lat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetLat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetLatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return &o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasLat() bool {
	if o != nil && IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given interface{} and assigns it to the Lat field.
func (o *POSTAddresses201ResponseDataAttributes) SetLat(v interface{}) {
	o.Lat = v
}

// GetLng returns the Lng field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetLng() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetLngOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Lng) {
		return nil, false
	}
	return &o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasLng() bool {
	if o != nil && IsNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given interface{} and assigns it to the Lng field.
func (o *POSTAddresses201ResponseDataAttributes) SetLng(v interface{}) {
	o.Lng = v
}

// GetBillingInfo returns the BillingInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetBillingInfo() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BillingInfo
}

// GetBillingInfoOk returns a tuple with the BillingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetBillingInfoOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BillingInfo) {
		return nil, false
	}
	return &o.BillingInfo, true
}

// HasBillingInfo returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasBillingInfo() bool {
	if o != nil && IsNil(o.BillingInfo) {
		return true
	}

	return false
}

// SetBillingInfo gets a reference to the given interface{} and assigns it to the BillingInfo field.
func (o *POSTAddresses201ResponseDataAttributes) SetBillingInfo(v interface{}) {
	o.BillingInfo = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTAddresses201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTAddresses201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTAddresses201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTAddresses201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAddresses201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Business != nil {
		toSerialize["business"] = o.Business
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Line1 != nil {
		toSerialize["line_1"] = o.Line1
	}
	if o.Line2 != nil {
		toSerialize["line_2"] = o.Line2
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.ZipCode != nil {
		toSerialize["zip_code"] = o.ZipCode
	}
	if o.StateCode != nil {
		toSerialize["state_code"] = o.StateCode
	}
	if o.CountryCode != nil {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Lng != nil {
		toSerialize["lng"] = o.Lng
	}
	if o.BillingInfo != nil {
		toSerialize["billing_info"] = o.BillingInfo
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

type NullablePOSTAddresses201ResponseDataAttributes struct {
	value *POSTAddresses201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTAddresses201ResponseDataAttributes) Get() *POSTAddresses201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTAddresses201ResponseDataAttributes) Set(val *POSTAddresses201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAddresses201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAddresses201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAddresses201ResponseDataAttributes(val *POSTAddresses201ResponseDataAttributes) *NullablePOSTAddresses201ResponseDataAttributes {
	return &NullablePOSTAddresses201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTAddresses201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAddresses201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
