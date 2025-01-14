# PATCHCustomersCustomerId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **interface{}** | The customer&#39;s email address. | [optional] 
**Password** | Pointer to **interface{}** | The customer&#39;s password. Initiate a customer password reset flow if you need to change it. | [optional] 
**ShopperReference** | Pointer to **interface{}** | A reference to uniquely identify the shopper during payment sessions. | [optional] 
**ProfileId** | Pointer to **interface{}** | A reference to uniquely identify the customer on any connected external services. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHCustomersCustomerId200ResponseDataAttributes

`func NewPATCHCustomersCustomerId200ResponseDataAttributes() *PATCHCustomersCustomerId200ResponseDataAttributes`

NewPATCHCustomersCustomerId200ResponseDataAttributes instantiates a new PATCHCustomersCustomerId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHCustomersCustomerId200ResponseDataAttributesWithDefaults

`func NewPATCHCustomersCustomerId200ResponseDataAttributesWithDefaults() *PATCHCustomersCustomerId200ResponseDataAttributes`

NewPATCHCustomersCustomerId200ResponseDataAttributesWithDefaults instantiates a new PATCHCustomersCustomerId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPassword

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetPassword() interface{}`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetPasswordOk() (*interface{}, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetPassword(v interface{})`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetShopperReference

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetShopperReference() interface{}`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetShopperReferenceOk() (*interface{}, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetShopperReference(v interface{})`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### SetShopperReferenceNil

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetShopperReferenceNil(b bool)`

 SetShopperReferenceNil sets the value for ShopperReference to be an explicit nil

### UnsetShopperReference
`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) UnsetShopperReference()`

UnsetShopperReference ensures that no value is present for ShopperReference, not even an explicit nil
### GetProfileId

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetReference

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHCustomersCustomerId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


