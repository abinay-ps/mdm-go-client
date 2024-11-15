# \SortingApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SortingSequencingsBulkUpdateSequencesPost**](SortingApi.md#SortingSequencingsBulkUpdateSequencesPost) | **Post** /sorting-sequencings/bulk-update-sequences | Update Postman Sorting Sequence Numbers
[**SortingSequencingsGet**](SortingApi.md#SortingSequencingsGet) | **Get** /sorting-sequencings | Retrieve Postman Sorting Sequence Data
[**SortingSequencingsPost**](SortingApi.md#SortingSequencingsPost) | **Post** /sorting-sequencings | Save Postman Sorting Data


# **SortingSequencingsBulkUpdateSequencesPost**
> string SortingSequencingsBulkUpdateSequencesPost(ctx, body)
Update Postman Sorting Sequence Numbers

This endpoint updates the sequence numbers for the provided postman sorting data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]DomainPostData**](domain.PostData.md)| Postman Sorting Data Array | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SortingSequencingsGet**
> []ResponseListSortingSequencingsApiResponse SortingSequencingsGet(ctx, batchName, beatName, officeId, optional)
Retrieve Postman Sorting Sequence Data

This endpoint fetches postman sorting sequence data using batch name, beat name, and office ID as query parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchName** | **string**| Batch Name | 
  **beatName** | **string**| Beat Name | 
  **officeId** | **string**| Office ID | 
 **optional** | ***SortingApiSortingSequencingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SortingApiSortingSequencingsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**[]ResponseListSortingSequencingsApiResponse**](response.ListSortingSequencingsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SortingSequencingsPost**
> string SortingSequencingsPost(ctx, request)
Save Postman Sorting Data

This endpoint allows the bulk saving of postman sorting data by receiving a JSON array of articles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**[]HandlerCreateSortingSequencingsRequest**](handler.CreateSortingSequencingsRequest.md)| Postman Sorting Data | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

