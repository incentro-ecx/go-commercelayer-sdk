# ExportCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**ExportCreateDataAttributes**](ExportCreateDataAttributes.md) |  | 
**Relationships** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewExportCreateData

`func NewExportCreateData(type_ string, attributes ExportCreateDataAttributes, ) *ExportCreateData`

NewExportCreateData instantiates a new ExportCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportCreateDataWithDefaults

`func NewExportCreateDataWithDefaults() *ExportCreateData`

NewExportCreateDataWithDefaults instantiates a new ExportCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExportCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ExportCreateData) GetAttributes() ExportCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExportCreateData) GetAttributesOk() (*ExportCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExportCreateData) SetAttributes(v ExportCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ExportCreateData) GetRelationships() map[string]interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExportCreateData) GetRelationshipsOk() (*map[string]interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExportCreateData) SetRelationships(v map[string]interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExportCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

