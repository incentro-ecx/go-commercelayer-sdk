# LinkData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETLinksLinkId200ResponseDataAttributes**](GETLinksLinkId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**LinkDataRelationships**](LinkDataRelationships.md) |  | [optional] 

## Methods

### NewLinkData

`func NewLinkData(type_ interface{}, attributes GETLinksLinkId200ResponseDataAttributes, ) *LinkData`

NewLinkData instantiates a new LinkData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkDataWithDefaults

`func NewLinkDataWithDefaults() *LinkData`

NewLinkDataWithDefaults instantiates a new LinkData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LinkData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *LinkData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LinkData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *LinkData) GetAttributes() GETLinksLinkId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LinkData) GetAttributesOk() (*GETLinksLinkId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LinkData) SetAttributes(v GETLinksLinkId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *LinkData) GetRelationships() LinkDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *LinkData) GetRelationshipsOk() (*LinkDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *LinkData) SetRelationships(v LinkDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *LinkData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


