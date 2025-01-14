# POSTExports201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **interface{}** | The type of resource being exported. | 
**Format** | Pointer to **interface{}** | The format of the export one of &#39;json&#39; (default) or &#39;csv&#39;. | [optional] 
**Includes** | Pointer to **interface{}** | List of related resources that should be included in the export (redundant when &#39;fields&#39; are specified). | [optional] 
**Fields** | Pointer to **interface{}** | List of fields to export for the main and related resources (automatically included). Pass the asterisk &#39;*&#39; to include all exportable fields for the main and related resources. | [optional] 
**Filters** | Pointer to **interface{}** | The filters used to select the records to be exported. | [optional] 
**DryData** | Pointer to **interface{}** | Send this attribute if you want to skip exporting redundant attributes (IDs, timestamps, blanks, etc.), useful when combining export and import to duplicate your dataset. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTExports201ResponseDataAttributes

`func NewPOSTExports201ResponseDataAttributes(resourceType interface{}, ) *POSTExports201ResponseDataAttributes`

NewPOSTExports201ResponseDataAttributes instantiates a new POSTExports201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTExports201ResponseDataAttributesWithDefaults

`func NewPOSTExports201ResponseDataAttributesWithDefaults() *POSTExports201ResponseDataAttributes`

NewPOSTExports201ResponseDataAttributesWithDefaults instantiates a new POSTExports201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *POSTExports201ResponseDataAttributes) GetResourceType() interface{}`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *POSTExports201ResponseDataAttributes) GetResourceTypeOk() (*interface{}, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *POSTExports201ResponseDataAttributes) SetResourceType(v interface{})`

SetResourceType sets ResourceType field to given value.


### SetResourceTypeNil

`func (o *POSTExports201ResponseDataAttributes) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *POSTExports201ResponseDataAttributes) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetFormat

`func (o *POSTExports201ResponseDataAttributes) GetFormat() interface{}`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *POSTExports201ResponseDataAttributes) GetFormatOk() (*interface{}, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *POSTExports201ResponseDataAttributes) SetFormat(v interface{})`

SetFormat sets Format field to given value.

### HasFormat

`func (o *POSTExports201ResponseDataAttributes) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *POSTExports201ResponseDataAttributes) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *POSTExports201ResponseDataAttributes) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetIncludes

`func (o *POSTExports201ResponseDataAttributes) GetIncludes() interface{}`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *POSTExports201ResponseDataAttributes) GetIncludesOk() (*interface{}, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *POSTExports201ResponseDataAttributes) SetIncludes(v interface{})`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *POSTExports201ResponseDataAttributes) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### SetIncludesNil

`func (o *POSTExports201ResponseDataAttributes) SetIncludesNil(b bool)`

 SetIncludesNil sets the value for Includes to be an explicit nil

### UnsetIncludes
`func (o *POSTExports201ResponseDataAttributes) UnsetIncludes()`

UnsetIncludes ensures that no value is present for Includes, not even an explicit nil
### GetFields

`func (o *POSTExports201ResponseDataAttributes) GetFields() interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *POSTExports201ResponseDataAttributes) GetFieldsOk() (*interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *POSTExports201ResponseDataAttributes) SetFields(v interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *POSTExports201ResponseDataAttributes) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *POSTExports201ResponseDataAttributes) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *POSTExports201ResponseDataAttributes) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetFilters

`func (o *POSTExports201ResponseDataAttributes) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *POSTExports201ResponseDataAttributes) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *POSTExports201ResponseDataAttributes) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *POSTExports201ResponseDataAttributes) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *POSTExports201ResponseDataAttributes) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *POSTExports201ResponseDataAttributes) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetDryData

`func (o *POSTExports201ResponseDataAttributes) GetDryData() interface{}`

GetDryData returns the DryData field if non-nil, zero value otherwise.

### GetDryDataOk

`func (o *POSTExports201ResponseDataAttributes) GetDryDataOk() (*interface{}, bool)`

GetDryDataOk returns a tuple with the DryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryData

`func (o *POSTExports201ResponseDataAttributes) SetDryData(v interface{})`

SetDryData sets DryData field to given value.

### HasDryData

`func (o *POSTExports201ResponseDataAttributes) HasDryData() bool`

HasDryData returns a boolean if a field has been set.

### SetDryDataNil

`func (o *POSTExports201ResponseDataAttributes) SetDryDataNil(b bool)`

 SetDryDataNil sets the value for DryData to be an explicit nil

### UnsetDryData
`func (o *POSTExports201ResponseDataAttributes) UnsetDryData()`

UnsetDryData ensures that no value is present for DryData, not even an explicit nil
### GetReference

`func (o *POSTExports201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTExports201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTExports201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTExports201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTExports201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTExports201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTExports201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTExports201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTExports201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTExports201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTExports201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTExports201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTExports201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTExports201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTExports201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTExports201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTExports201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTExports201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


