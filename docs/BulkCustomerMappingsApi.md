# \BulkCustomerMappingsApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCustomerMappingsBulkCustomerIdPut**](BulkCustomerMappingsApi.md#BulkCustomerMappingsBulkCustomerIdPut) | **Put** /bulk-customer-mappings/{bulk-customer-id} | Update Bulk Customer Details
[**BulkCustomerMappingsGet**](BulkCustomerMappingsApi.md#BulkCustomerMappingsGet) | **Get** /bulk-customer-mappings | Retrieve Bulk Customers of Office
[**BulkCustomerMappingsPost**](BulkCustomerMappingsApi.md#BulkCustomerMappingsPost) | **Post** /bulk-customer-mappings | Create a new bulk customer


# **BulkCustomerMappingsBulkCustomerIdPut**
> HandlerBulkAddresseeDataResponse BulkCustomerMappingsBulkCustomerIdPut(ctx, bulkCustomerId, officeId, body)
Update Bulk Customer Details

Updates the details for a specified bulk customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bulkCustomerId** | **string**| Bulk Customer ID | 
  **officeId** | **int32**| Office ID | 
  **body** | [**HandlerUpdateBulkCustomerMappingsRequest**](HandlerUpdateBulkCustomerMappingsRequest.md)| Bulk Customer Details Request | 

### Return type

[**HandlerBulkAddresseeDataResponse**](handler.BulkAddresseeDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkCustomerMappingsGet**
> []ResponseListBulkCustomerMappingsRespone BulkCustomerMappingsGet(ctx, officeId, optional)
Retrieve Bulk Customers of Office

This endpoint retrieves bulk addressee details for a given office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| ID of the office | 
 **optional** | ***BulkCustomerMappingsApiBulkCustomerMappingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BulkCustomerMappingsApiBulkCustomerMappingsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**[]ResponseListBulkCustomerMappingsRespone**](response.ListBulkCustomerMappingsRespone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkCustomerMappingsPost**
> HandlerBulkAddresseeDataResponse BulkCustomerMappingsPost(ctx, createBulkCustomerMappingsRequest)
Create a new bulk customer

Adds a new bulk customer with the provided details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createBulkCustomerMappingsRequest** | [**HandlerCreateBulkCustomerMappingsRequest**](HandlerCreateBulkCustomerMappingsRequest.md)| Bulk Customer Data | 

### Return type

[**HandlerBulkAddresseeDataResponse**](handler.BulkAddresseeDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

