# GETShippingZonesShippingZoneId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The shipping zone&#39;s internal name. | [optional] 
**CountryCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the shipping address country code, max size is 5000. | [optional] 
**NotCountryCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated as negative match for the shipping address country code, max size is 5000. | [optional] 
**StateCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the shipping address state code, max size is 5000. | [optional] 
**NotStateCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated as negative match for the shipping address state code, max size is 5000. | [optional] 
**ZipCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the shipping address zip code, max size is 5000. | [optional] 
**NotZipCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated as negative match for the shipping zip country code, max size is 5000. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETShippingZonesShippingZoneId200ResponseDataAttributes

`func NewGETShippingZonesShippingZoneId200ResponseDataAttributes() *GETShippingZonesShippingZoneId200ResponseDataAttributes`

NewGETShippingZonesShippingZoneId200ResponseDataAttributes instantiates a new GETShippingZonesShippingZoneId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETShippingZonesShippingZoneId200ResponseDataAttributesWithDefaults

`func NewGETShippingZonesShippingZoneId200ResponseDataAttributesWithDefaults() *GETShippingZonesShippingZoneId200ResponseDataAttributes`

NewGETShippingZonesShippingZoneId200ResponseDataAttributesWithDefaults instantiates a new GETShippingZonesShippingZoneId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCountryCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetCountryCodeRegex() interface{}`

GetCountryCodeRegex returns the CountryCodeRegex field if non-nil, zero value otherwise.

### GetCountryCodeRegexOk

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetCountryCodeRegexOk() (*interface{}, bool)`

GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetCountryCodeRegex(v interface{})`

SetCountryCodeRegex sets CountryCodeRegex field to given value.

### HasCountryCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) HasCountryCodeRegex() bool`

HasCountryCodeRegex returns a boolean if a field has been set.

### SetCountryCodeRegexNil

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetCountryCodeRegexNil(b bool)`

 SetCountryCodeRegexNil sets the value for CountryCodeRegex to be an explicit nil

### UnsetCountryCodeRegex
`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) UnsetCountryCodeRegex()`

UnsetCountryCodeRegex ensures that no value is present for CountryCodeRegex, not even an explicit nil
### GetNotCountryCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetNotCountryCodeRegex() interface{}`

GetNotCountryCodeRegex returns the NotCountryCodeRegex field if non-nil, zero value otherwise.

### GetNotCountryCodeRegexOk

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetNotCountryCodeRegexOk() (*interface{}, bool)`

GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCountryCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetNotCountryCodeRegex(v interface{})`

SetNotCountryCodeRegex sets NotCountryCodeRegex field to given value.

### HasNotCountryCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) HasNotCountryCodeRegex() bool`

HasNotCountryCodeRegex returns a boolean if a field has been set.

### SetNotCountryCodeRegexNil

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetNotCountryCodeRegexNil(b bool)`

 SetNotCountryCodeRegexNil sets the value for NotCountryCodeRegex to be an explicit nil

### UnsetNotCountryCodeRegex
`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) UnsetNotCountryCodeRegex()`

UnsetNotCountryCodeRegex ensures that no value is present for NotCountryCodeRegex, not even an explicit nil
### GetStateCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetStateCodeRegex() interface{}`

GetStateCodeRegex returns the StateCodeRegex field if non-nil, zero value otherwise.

### GetStateCodeRegexOk

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetStateCodeRegexOk() (*interface{}, bool)`

GetStateCodeRegexOk returns a tuple with the StateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetStateCodeRegex(v interface{})`

SetStateCodeRegex sets StateCodeRegex field to given value.

### HasStateCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) HasStateCodeRegex() bool`

HasStateCodeRegex returns a boolean if a field has been set.

### SetStateCodeRegexNil

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetStateCodeRegexNil(b bool)`

 SetStateCodeRegexNil sets the value for StateCodeRegex to be an explicit nil

### UnsetStateCodeRegex
`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) UnsetStateCodeRegex()`

UnsetStateCodeRegex ensures that no value is present for StateCodeRegex, not even an explicit nil
### GetNotStateCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetNotStateCodeRegex() interface{}`

GetNotStateCodeRegex returns the NotStateCodeRegex field if non-nil, zero value otherwise.

### GetNotStateCodeRegexOk

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetNotStateCodeRegexOk() (*interface{}, bool)`

GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotStateCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetNotStateCodeRegex(v interface{})`

SetNotStateCodeRegex sets NotStateCodeRegex field to given value.

### HasNotStateCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) HasNotStateCodeRegex() bool`

HasNotStateCodeRegex returns a boolean if a field has been set.

### SetNotStateCodeRegexNil

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetNotStateCodeRegexNil(b bool)`

 SetNotStateCodeRegexNil sets the value for NotStateCodeRegex to be an explicit nil

### UnsetNotStateCodeRegex
`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) UnsetNotStateCodeRegex()`

UnsetNotStateCodeRegex ensures that no value is present for NotStateCodeRegex, not even an explicit nil
### GetZipCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetZipCodeRegex() interface{}`

GetZipCodeRegex returns the ZipCodeRegex field if non-nil, zero value otherwise.

### GetZipCodeRegexOk

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetZipCodeRegexOk() (*interface{}, bool)`

GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetZipCodeRegex(v interface{})`

SetZipCodeRegex sets ZipCodeRegex field to given value.

### HasZipCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) HasZipCodeRegex() bool`

HasZipCodeRegex returns a boolean if a field has been set.

### SetZipCodeRegexNil

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetZipCodeRegexNil(b bool)`

 SetZipCodeRegexNil sets the value for ZipCodeRegex to be an explicit nil

### UnsetZipCodeRegex
`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) UnsetZipCodeRegex()`

UnsetZipCodeRegex ensures that no value is present for ZipCodeRegex, not even an explicit nil
### GetNotZipCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetNotZipCodeRegex() interface{}`

GetNotZipCodeRegex returns the NotZipCodeRegex field if non-nil, zero value otherwise.

### GetNotZipCodeRegexOk

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetNotZipCodeRegexOk() (*interface{}, bool)`

GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotZipCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetNotZipCodeRegex(v interface{})`

SetNotZipCodeRegex sets NotZipCodeRegex field to given value.

### HasNotZipCodeRegex

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) HasNotZipCodeRegex() bool`

HasNotZipCodeRegex returns a boolean if a field has been set.

### SetNotZipCodeRegexNil

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetNotZipCodeRegexNil(b bool)`

 SetNotZipCodeRegexNil sets the value for NotZipCodeRegex to be an explicit nil

### UnsetNotZipCodeRegex
`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) UnsetNotZipCodeRegex()`

UnsetNotZipCodeRegex ensures that no value is present for NotZipCodeRegex, not even an explicit nil
### GetCreatedAt

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETShippingZonesShippingZoneId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


