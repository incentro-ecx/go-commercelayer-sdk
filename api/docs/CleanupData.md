# CleanupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**GETCleanups200ResponseDataInnerAttributes**](GETCleanups200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**CleanupDataRelationships**](CleanupDataRelationships.md) |  | [optional] 

## Methods

### NewCleanupData

`func NewCleanupData(type_ string, attributes GETCleanups200ResponseDataInnerAttributes, ) *CleanupData`

NewCleanupData instantiates a new CleanupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanupDataWithDefaults

`func NewCleanupDataWithDefaults() *CleanupData`

NewCleanupDataWithDefaults instantiates a new CleanupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CleanupData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CleanupData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CleanupData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CleanupData) GetAttributes() GETCleanups200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CleanupData) GetAttributesOk() (*GETCleanups200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CleanupData) SetAttributes(v GETCleanups200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CleanupData) GetRelationships() CleanupDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CleanupData) GetRelationshipsOk() (*CleanupDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CleanupData) SetRelationships(v CleanupDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CleanupData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

