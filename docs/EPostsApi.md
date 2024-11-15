# \EPostsApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EpostsGet**](EPostsApi.md#EpostsGet) | **Get** /eposts | Performs operations on epost based on the given state.


# **EpostsGet**
> ResponseListEpostApiResponse EpostsGet(ctx, epostState, officeId, optional)
Performs operations on epost based on the given state.

Retrieves epost details or performs updates based on epost state, office ID, and optionally epost ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **epostState** | **string**| Epost State (receipt, download, printed) | 
  **officeId** | **int32**| Office ID | 
 **optional** | ***EPostsApiEpostsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EPostsApiEpostsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **epostId** | **optional.Int32**| Epost ID (required if epost-state is receipt) | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListEpostApiResponse**](response.ListEpostAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

