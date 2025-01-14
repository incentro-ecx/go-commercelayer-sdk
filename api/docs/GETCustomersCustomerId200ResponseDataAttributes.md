# GETCustomersCustomerId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **interface{}** | The customer&#39;s email address. | [optional] 
**Status** | Pointer to **interface{}** | The customer&#39;s status. One of &#39;prospect&#39; (default), &#39;acquired&#39;, or &#39;repeat&#39;. | [optional] 
**HasPassword** | Pointer to **interface{}** | Indicates if the customer has a password. | [optional] 
**TotalOrdersCount** | Pointer to **interface{}** | The total number of orders for the customer. | [optional] 
**ShopperReference** | Pointer to **interface{}** | A reference to uniquely identify the shopper during payment sessions. | [optional] 
**ProfileId** | Pointer to **interface{}** | A reference to uniquely identify the customer on any connected external services. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETCustomersCustomerId200ResponseDataAttributes

`func NewGETCustomersCustomerId200ResponseDataAttributes() *GETCustomersCustomerId200ResponseDataAttributes`

NewGETCustomersCustomerId200ResponseDataAttributes instantiates a new GETCustomersCustomerId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCustomersCustomerId200ResponseDataAttributesWithDefaults

`func NewGETCustomersCustomerId200ResponseDataAttributesWithDefaults() *GETCustomersCustomerId200ResponseDataAttributes`

NewGETCustomersCustomerId200ResponseDataAttributesWithDefaults instantiates a new GETCustomersCustomerId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GETCustomersCustomerId200ResponseDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GETCustomersCustomerId200ResponseDataAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetStatus

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETCustomersCustomerId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETCustomersCustomerId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetHasPassword

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetHasPassword() interface{}`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetHasPasswordOk() (*interface{}, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetHasPassword(v interface{})`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *GETCustomersCustomerId200ResponseDataAttributes) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### SetHasPasswordNil

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetHasPasswordNil(b bool)`

 SetHasPasswordNil sets the value for HasPassword to be an explicit nil

### UnsetHasPassword
`func (o *GETCustomersCustomerId200ResponseDataAttributes) UnsetHasPassword()`

UnsetHasPassword ensures that no value is present for HasPassword, not even an explicit nil
### GetTotalOrdersCount

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetTotalOrdersCount() interface{}`

GetTotalOrdersCount returns the TotalOrdersCount field if non-nil, zero value otherwise.

### GetTotalOrdersCountOk

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetTotalOrdersCountOk() (*interface{}, bool)`

GetTotalOrdersCountOk returns a tuple with the TotalOrdersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOrdersCount

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetTotalOrdersCount(v interface{})`

SetTotalOrdersCount sets TotalOrdersCount field to given value.

### HasTotalOrdersCount

`func (o *GETCustomersCustomerId200ResponseDataAttributes) HasTotalOrdersCount() bool`

HasTotalOrdersCount returns a boolean if a field has been set.

### SetTotalOrdersCountNil

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetTotalOrdersCountNil(b bool)`

 SetTotalOrdersCountNil sets the value for TotalOrdersCount to be an explicit nil

### UnsetTotalOrdersCount
`func (o *GETCustomersCustomerId200ResponseDataAttributes) UnsetTotalOrdersCount()`

UnsetTotalOrdersCount ensures that no value is present for TotalOrdersCount, not even an explicit nil
### GetShopperReference

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetShopperReference() interface{}`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetShopperReferenceOk() (*interface{}, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetShopperReference(v interface{})`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *GETCustomersCustomerId200ResponseDataAttributes) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### SetShopperReferenceNil

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetShopperReferenceNil(b bool)`

 SetShopperReferenceNil sets the value for ShopperReference to be an explicit nil

### UnsetShopperReference
`func (o *GETCustomersCustomerId200ResponseDataAttributes) UnsetShopperReference()`

UnsetShopperReference ensures that no value is present for ShopperReference, not even an explicit nil
### GetProfileId

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *GETCustomersCustomerId200ResponseDataAttributes) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *GETCustomersCustomerId200ResponseDataAttributes) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetCreatedAt

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETCustomersCustomerId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETCustomersCustomerId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETCustomersCustomerId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETCustomersCustomerId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETCustomersCustomerId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETCustomersCustomerId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETCustomersCustomerId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETCustomersCustomerId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETCustomersCustomerId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETCustomersCustomerId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETCustomersCustomerId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETCustomersCustomerId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


