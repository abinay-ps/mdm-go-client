# \OrdinaryArticlesApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrdinaryArticlesBulkUpdateReturnsPost**](OrdinaryArticlesApi.md#OrdinaryArticlesBulkUpdateReturnsPost) | **Post** /ordinary-articles/bulk-update-returns | Unregistered Articles Returns
[**OrdinaryArticlesInvoicePost**](OrdinaryArticlesApi.md#OrdinaryArticlesInvoicePost) | **Post** /ordinary-articles/invoice | Create Invoice for Ordinary Articles
[**OrdinaryArticlesOfficeIdGet**](OrdinaryArticlesApi.md#OrdinaryArticlesOfficeIdGet) | **Get** /ordinary-articles/{office-id} | Get Unregistered Articles Details


# **OrdinaryArticlesBulkUpdateReturnsPost**
> string OrdinaryArticlesBulkUpdateReturnsPost(ctx, unregisteredState, body)
Unregistered Articles Returns

Submits returns for unregistered articles, handling both barcoded and unbarcoded cases.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unregisteredState** | **string**| State of Unregistered Articles Returns (unbarcoded or barcoded) | 
  **body** | [**HandlerCreateOrdinaryArticlesBulkUpdateReturnsRequest**](HandlerCreateOrdinaryArticlesBulkUpdateReturnsRequest.md)| Unregistered Articles Returns Request | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrdinaryArticlesInvoicePost**
> ResponseListOrdinaryArticlesResponse OrdinaryArticlesInvoicePost(ctx, unregisteredState, request)
Create Invoice for Ordinary Articles

Creates an invoice for ordinary articles based on the unregistered article state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unregisteredState** | **string**| Unregistered state(barcoded or unbarcoded) | 
  **request** | [**HandlerCreateOrdinaryArticlesInvoiceRequest**](HandlerCreateOrdinaryArticlesInvoiceRequest.md)| Create Ordinary Articles Invoice Request | 

### Return type

[**ResponseListOrdinaryArticlesResponse**](response.ListOrdinaryArticlesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrdinaryArticlesOfficeIdGet**
> ResponseListOrdinaryArticlesApiResponse OrdinaryArticlesOfficeIdGet(ctx, officeId, beatName, batchName, optional)
Get Unregistered Articles Details

Retrieves details of unregistered articles for a specific office and beat.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Reporting Office ID | 
  **beatName** | **string**| Beat Name | 
  **batchName** | **string**| Batch Name | 
 **optional** | ***OrdinaryArticlesApiOrdinaryArticlesOfficeIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdinaryArticlesApiOrdinaryArticlesOfficeIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListOrdinaryArticlesApiResponse**](response.ListOrdinaryArticlesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

