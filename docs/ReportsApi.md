# \ReportsApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportsOfficeIdArticlesAbstractGet**](ReportsApi.md#ReportsOfficeIdArticlesAbstractGet) | **Get** /reports/{office-id}/articles-abstract | Retrieve Articles Abstract
[**ReportsOfficeIdEmosAbstractGet**](ReportsApi.md#ReportsOfficeIdEmosAbstractGet) | **Get** /reports/{office-id}/emos-abstract | Retrieves Emo transaction reports
[**ReportsOfficeIdEmosPaidGet**](ReportsApi.md#ReportsOfficeIdEmosPaidGet) | **Get** /reports/{office-id}/emos-paid | Retrieve EMO Printed Data
[**ReportsOfficeIdUndelieveredArticlesGet**](ReportsApi.md#ReportsOfficeIdUndelieveredArticlesGet) | **Get** /reports/{office-id}/undelievered-articles | Retrieve Undelivered Articles
[**ReportsOfficeIdUnregisteredArticlesGet**](ReportsApi.md#ReportsOfficeIdUnregisteredArticlesGet) | **Get** /reports/{office-id}/unregistered-articles | Retrieve Unregistered Articles
[**ReportsPostmanPerformanceGet**](ReportsApi.md#ReportsPostmanPerformanceGet) | **Get** /reports/postman-performance | Generate Performance Reports


# **ReportsOfficeIdArticlesAbstractGet**
> ResponseListReportsArticlesAbstractApiResponse ReportsOfficeIdArticlesAbstractGet(ctx, officeId, optional)
Retrieve Articles Abstract

This endpoint fetches various statistics and details about articles based on the provided parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
 **optional** | ***ReportsApiReportsOfficeIdArticlesAbstractGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiReportsOfficeIdArticlesAbstractGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoicedDate** | **optional.String**| Invoice Date | 
 **deliveredDate** | **optional.String**| Delivered Date | 
 **articleRecivedDate** | **optional.String**| Article Received Date | 
 **beatName** | **optional.String**| Beat Name | 
 **batchName** | **optional.String**| Batch Name | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListReportsArticlesAbstractApiResponse**](response.ListReportsArticlesAbstractAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportsOfficeIdEmosAbstractGet**
> ResponseListReportsEmOsAbstractResponse ReportsOfficeIdEmosAbstractGet(ctx, officeId, optional)
Retrieves Emo transaction reports

Retrieves detailed Emo transaction data for a specific office based on office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
 **optional** | ***ReportsApiReportsOfficeIdEmosAbstractGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiReportsOfficeIdEmosAbstractGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoicedDate** | **optional.String**| Invoice Date | 
 **deliveredDate** | **optional.String**| Delivery Date | 
 **articleRecivedDate** | **optional.String**| Article Received Date | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListReportsEmOsAbstractResponse**](response.ListReportsEMOsAbstractResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportsOfficeIdEmosPaidGet**
> []ResponseListReportsEmOsPaidApiResponse ReportsOfficeIdEmosPaidGet(ctx, officeId, startDate, endDate, optional)
Retrieve EMO Printed Data

This endpoint fetches printed EMO data based on the provided start date, end date, and office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
  **startDate** | **string**| Start Date | 
  **endDate** | **string**| End Date | 
 **optional** | ***ReportsApiReportsOfficeIdEmosPaidGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiReportsOfficeIdEmosPaidGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**[]ResponseListReportsEmOsPaidApiResponse**](response.ListReportsEMOsPaidAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportsOfficeIdUndelieveredArticlesGet**
> []ResponseListReportsUndelieveredArticlesApiResponse ReportsOfficeIdUndelieveredArticlesGet(ctx, officeId, optional)
Retrieve Undelivered Articles

This endpoint fetches undelivered articles based on the provided invoice date and office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
 **optional** | ***ReportsApiReportsOfficeIdUndelieveredArticlesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiReportsOfficeIdUndelieveredArticlesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **optional.String**| Invoice Date | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**[]ResponseListReportsUndelieveredArticlesApiResponse**](response.ListReportsUndelieveredArticlesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportsOfficeIdUnregisteredArticlesGet**
> []ResponseListReportsUnregisteredArticlesResponse ReportsOfficeIdUnregisteredArticlesGet(ctx, officeId, optional)
Retrieve Unregistered Articles

This endpoint fetches unregistered articles based on the provided date and office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **string**| Office ID | 
 **optional** | ***ReportsApiReportsOfficeIdUnregisteredArticlesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiReportsOfficeIdUnregisteredArticlesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **optional.String**| Date | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**[]ResponseListReportsUnregisteredArticlesResponse**](response.ListReportsUnregisteredArticlesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportsPostmanPerformanceGet**
> []ResponseListReportsPostmanPerformanceResponse ReportsPostmanPerformanceGet(ctx, officeId, reportType, optional)
Generate Performance Reports

This endpoint generates performance reports based on report type, invoice date, and office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
  **reportType** | **string**| Report Type | 
 **optional** | ***ReportsApiReportsPostmanPerformanceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiReportsPostmanPerformanceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **invoicedDate** | **optional.String**| Invoice Date | 
 **deliveredDate** | **optional.String**| Delivered Date | 
 **articleReceivedDate** | **optional.String**| Article Received Date | 
 **beatName** | **optional.String**| Beat Name | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**[]ResponseListReportsPostmanPerformanceResponse**](response.ListReportsPostmanPerformanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

