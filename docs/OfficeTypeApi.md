# \OfficeTypeApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficeTypesGet**](OfficeTypeApi.md#OfficeTypesGet) | **Get** /office-types | List all office types
[**OfficeTypesOfficeTypeIdGet**](OfficeTypeApi.md#OfficeTypesOfficeTypeIdGet) | **Get** /office-types/{office-type-id} | Retrieve office type by ID
[**OfficeTypesOfficeTypeIdPut**](OfficeTypeApi.md#OfficeTypesOfficeTypeIdPut) | **Put** /office-types/{office-type-id} | Update details of a specific office type
[**OfficeTypesPost**](OfficeTypeApi.md#OfficeTypesPost) | **Post** /office-types | Creates a new office type


# **OfficeTypesGet**
> ResponseListOfficeTypeApiResponse OfficeTypesGet(ctx, optional)
List all office types

Retrieves a list of office types with optional pagination

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OfficeTypeApiOfficeTypesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeTypeApiOfficeTypesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListOfficeTypeApiResponse**](response.ListOfficeTypeAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficeTypesOfficeTypeIdGet**
> ResponseGetOfficeTypeApiResponse OfficeTypesOfficeTypeIdGet(ctx, officeTypeId)
Retrieve office type by ID

Gets the office type details based on the specified office type ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeTypeId** | **int32**| The ID of the office type | 

### Return type

[**ResponseGetOfficeTypeApiResponse**](response.GetOfficeTypeAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficeTypesOfficeTypeIdPut**
> ResponseUpdateOfficeTypesApiResponse OfficeTypesOfficeTypeIdPut(ctx, officeTypeId, body)
Update details of a specific office type

Updates office type details identified by the office type ID. Requires the office type ID as a URI parameter and the updated details in the JSON body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeTypeId** | **int32**| ID of the office type to update | 
  **body** | [**HandlerUpdateOfficeTypesRequest**](HandlerUpdateOfficeTypesRequest.md)| Updated details for the office type | 

### Return type

[**ResponseUpdateOfficeTypesApiResponse**](response.UpdateOfficeTypesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficeTypesPost**
> ResponseCreateOfficeTypeApiResponse OfficeTypesPost(ctx, body)
Creates a new office type

Creates a new office type based on the provided JSON request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HandlerCreateOfficeTypeRequest**](HandlerCreateOfficeTypeRequest.md)| Request body containing details of the new office type | 

### Return type

[**ResponseCreateOfficeTypeApiResponse**](response.CreateOfficeTypeAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

