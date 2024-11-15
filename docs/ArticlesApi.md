# \ArticlesApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArticlesArticleNumberAlterRemarksPut**](ArticlesApi.md#ArticlesArticleNumberAlterRemarksPut) | **Put** /articles/{article-number}/alter-remarks | Update Article Remarks
[**ArticlesArticleNumberDelete**](ArticlesApi.md#ArticlesArticleNumberDelete) | **Delete** /articles/{article-number} | Delete Article Deposited
[**ArticlesArticleNumberGet**](ArticlesApi.md#ArticlesArticleNumberGet) | **Get** /articles/{article-number} | Fetch Article by Article Number
[**ArticlesArticleNumberOfficeIdInternalTrackingGet**](ArticlesApi.md#ArticlesArticleNumberOfficeIdInternalTrackingGet) | **Get** /articles/{article-number}/{office-id}/internal-tracking | Fetch Article Internal Tracking
[**ArticlesArticleNumberPostageDuePut**](ArticlesApi.md#ArticlesArticleNumberPostageDuePut) | **Put** /articles/{article-number}/postage-due | Update Article Postage Due
[**ArticlesArticleNumberRemarksGet**](ArticlesApi.md#ArticlesArticleNumberRemarksGet) | **Get** /articles/{article-number}/remarks | Fetch Article Remarks
[**ArticlesBulkDataEntryPost**](ArticlesApi.md#ArticlesBulkDataEntryPost) | **Post** /articles/bulk-data-entry | Create Articles Bulk Data Entry
[**ArticlesBulkInvoicePost**](ArticlesApi.md#ArticlesBulkInvoicePost) | **Post** /articles/bulk-invoice | Create Bulk Invoice for Articles
[**ArticlesBulkRemoveMakerCheckerPost**](ArticlesApi.md#ArticlesBulkRemoveMakerCheckerPost) | **Post** /articles/bulk-remove-maker-checker | Remove Articles in Bulk with Maker-Checker
[**ArticlesBulkUpdateRemarksFromPmaPut**](ArticlesApi.md#ArticlesBulkUpdateRemarksFromPmaPut) | **Put** /articles/bulk-update-remarks-from-pma | Update Article Remarks from PMA
[**ArticlesBulkWindowDeliveryPost**](ArticlesApi.md#ArticlesBulkWindowDeliveryPost) | **Post** /articles/bulk-window-delivery | Create Bulk Window Delivery for Articles
[**ArticlesCountForAbstractGet**](ArticlesApi.md#ArticlesCountForAbstractGet) | **Get** /articles/count-for-abstract | Fetch Articles Count for Abstract
[**ArticlesDashboardStatisticsGet**](ArticlesApi.md#ArticlesDashboardStatisticsGet) | **Get** /articles/dashboard-statistics | Fetch Articles Dashboard Statistics
[**ArticlesDataEntriesGet**](ArticlesApi.md#ArticlesDataEntriesGet) | **Get** /articles/data-entries | List Articles Data Entries
[**ArticlesGet**](ArticlesApi.md#ArticlesGet) | **Get** /articles | List Articles
[**ArticlesInitialStatisticsGet**](ArticlesApi.md#ArticlesInitialStatisticsGet) | **Get** /articles/initial-statistics | Fetch Articles Initial Statistics
[**ArticlesSupervisorApprovalByPersonnelIdPost**](ArticlesApi.md#ArticlesSupervisorApprovalByPersonnelIdPost) | **Post** /articles/supervisor-approval-by-personnel-id | Create Articles Supervisor Approval by Personnel ID


# **ArticlesArticleNumberAlterRemarksPut**
> ResponseUpdateArticlesAlterRemarksApiResponse ArticlesArticleNumberAlterRemarksPut(ctx, articleNumber, request)
Update Article Remarks

Modifies the remarks for a specified article and its details based on the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleNumber** | **string**| Article Number | 
  **request** | [**HandlerUpdateArticlesAlterRemarksRequest**](HandlerUpdateArticlesAlterRemarksRequest.md)| Update Article Remarks Request | 

### Return type

[**ResponseUpdateArticlesAlterRemarksApiResponse**](response.UpdateArticlesAlterRemarksAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesArticleNumberDelete**
> ResponseDeleteArticleDepositedApiResponse ArticlesArticleNumberDelete(ctx, articleNumber)
Delete Article Deposited

This endpoint deletes a deposited article specified by the article number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleNumber** | **string**| Article Number | 

### Return type

[**ResponseDeleteArticleDepositedApiResponse**](response.DeleteArticleDepositedAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesArticleNumberGet**
> ResponseFetchArticleApiResponse ArticlesArticleNumberGet(ctx, articleNumber)
Fetch Article by Article Number

Retrieves article inward data based on the provided article number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleNumber** | **string**| Article Number | 

### Return type

[**ResponseFetchArticleApiResponse**](response.FetchArticleAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesArticleNumberOfficeIdInternalTrackingGet**
> ResponseFetchArticleInternalTrackingApiResponse ArticlesArticleNumberOfficeIdInternalTrackingGet(ctx, articleNumber, officeId)
Fetch Article Internal Tracking

This endpoint retrieves the internal tracking details of an article based on the provided article number and current office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleNumber** | **string**| Article Number | 
  **officeId** | **int32**| Current Office ID | 

### Return type

[**ResponseFetchArticleInternalTrackingApiResponse**](response.FetchArticleInternalTrackingAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesArticleNumberPostageDuePut**
> ResponseUpdateArticlePostageDueApiResponse ArticlesArticleNumberPostageDuePut(ctx, articleNumber, request)
Update Article Postage Due

This endpoint updates the postage due for a specified article number in the given office.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleNumber** | **string**| Article Number | 
  **request** | [**HandlerUpdateArticlePostageDueRequest**](HandlerUpdateArticlePostageDueRequest.md)| Postage due update request | 

### Return type

[**ResponseUpdateArticlePostageDueApiResponse**](response.UpdateArticlePostageDueAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesArticleNumberRemarksGet**
> ResponseFetchArticleRemarksApiResponse ArticlesArticleNumberRemarksGet(ctx, articleNumber, optional)
Fetch Article Remarks

Retrieves remarks for a specific article based on its article number and whether it is a single-hand office.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleNumber** | **string**| Article Number | 
 **optional** | ***ArticlesApiArticlesArticleNumberRemarksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticlesApiArticlesArticleNumberRemarksGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isSingleHandOffice** | **optional.Bool**| Is Single Hand Office | 

### Return type

[**ResponseFetchArticleRemarksApiResponse**](response.FetchArticleRemarksAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesBulkDataEntryPost**
> ResponseCreateArticlesBulkDataEntryApiResponse ArticlesBulkDataEntryPost(ctx, request)
Create Articles Bulk Data Entry

This endpoint allows for the bulk data entry of articles, including details like sender and receiver information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateArticlesBulkDataEntryRequest**](HandlerCreateArticlesBulkDataEntryRequest.md)| Create Articles Bulk Data Entry Request | 

### Return type

[**ResponseCreateArticlesBulkDataEntryApiResponse**](response.CreateArticlesBulkDataEntryAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesBulkInvoicePost**
> ResponseCreateArticlesBulkInvoiceApiResponse ArticlesBulkInvoicePost(ctx, invoiceType, request, optional)
Create Bulk Invoice for Articles

Creates invoices for articles based on the provided invoice type and delivery location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceType** | **string**| Invoice type for EMOs(invoice, reinvoice, transmit) | 
  **request** | [**HandlerCreateArticlesBulkInvoiceRequest**](HandlerCreateArticlesBulkInvoiceRequest.md)| Create Articles Bulk Invoice Request | 
 **optional** | ***ArticlesApiArticlesBulkInvoicePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticlesApiArticlesBulkInvoicePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deliveryLocation** | **optional.String**| Delivery location(office, branch-office, beat, bulk-customer) | 

### Return type

[**ResponseCreateArticlesBulkInvoiceApiResponse**](response.CreateArticlesBulkInvoiceAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesBulkRemoveMakerCheckerPost**
> ResponseCreateArticlesBulkRemoveMakerCheckerApiResponse ArticlesBulkRemoveMakerCheckerPost(ctx, request)
Remove Articles in Bulk with Maker-Checker

This endpoint removes articles in bulk using a maker-checker pattern based on the provided request data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateArticlesBulkRemoveMakerCheckerRequest**](HandlerCreateArticlesBulkRemoveMakerCheckerRequest.md)| Create Articles Bulk Remove Maker Checker Request | 

### Return type

[**ResponseCreateArticlesBulkRemoveMakerCheckerApiResponse**](response.CreateArticlesBulkRemoveMakerCheckerAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesBulkUpdateRemarksFromPmaPut**
> ResponseCreateArticlesBulkUpdateRemarksFromPmaapiResponse ArticlesBulkUpdateRemarksFromPmaPut(ctx, officeId)
Update Article Remarks from PMA

Fetches article remarks from PMA and updates them for the specified office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 

### Return type

[**ResponseCreateArticlesBulkUpdateRemarksFromPmaapiResponse**](response.CreateArticlesBulkUpdateRemarksFromPMAAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesBulkWindowDeliveryPost**
> ResponseCreateArticlesBulkWindowDeliveryApiResponse ArticlesBulkWindowDeliveryPost(ctx, request)
Create Bulk Window Delivery for Articles

Processes the bulk delivery of articles with the specified details and employee information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateArticlesBulkWindowDeliveryRequest**](HandlerCreateArticlesBulkWindowDeliveryRequest.md)| Create Articles Bulk Window Delivery Request | 

### Return type

[**ResponseCreateArticlesBulkWindowDeliveryApiResponse**](response.CreateArticlesBulkWindowDeliveryAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesCountForAbstractGet**
> ResponseFetchArticlesCountForAbstractApiResponse ArticlesCountForAbstractGet(ctx, officeId)
Fetch Articles Count for Abstract

This endpoint retrieves counts of deposited articles, articles received today, issued to postman, and more for a given office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 

### Return type

[**ResponseFetchArticlesCountForAbstractApiResponse**](response.FetchArticlesCountForAbstractAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesDashboardStatisticsGet**
> ResponseFetchArticlesDashboardStatisticsApiResponse ArticlesDashboardStatisticsGet(ctx, officeId)
Fetch Articles Dashboard Statistics

This endpoint retrieves dashboard statistics for articles based on the specified office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 

### Return type

[**ResponseFetchArticlesDashboardStatisticsApiResponse**](response.FetchArticlesDashboardStatisticsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesDataEntriesGet**
> ResponseListArticlesDataEntriesApiResponse ArticlesDataEntriesGet(ctx, startDate, endDate, officeId, optional)
List Articles Data Entries

This endpoint retrieves article data entries for a given date range and office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **startDate** | **string**| Start Date | 
  **endDate** | **string**| End Date | 
  **officeId** | **int32**| Office ID | 
 **optional** | ***ArticlesApiArticlesDataEntriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticlesApiArticlesDataEntriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListArticlesDataEntriesApiResponse**](response.ListArticlesDataEntriesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesGet**
> ResponseListArticlesApiResponse ArticlesGet(ctx, articlesState, deliveryLocation, officeId, optional)
List Articles

Retrieves a list of articles based on state, type, and delivery location

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articlesState** | **string**| State of the articles | 
  **deliveryLocation** | **string**| Delivery location | 
  **officeId** | **int32**| Office ID | 
 **optional** | ***ArticlesApiArticlesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticlesApiArticlesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **articleType** | **optional.String**| Type of the article | 
 **beatName** | **optional.String**| Beat Name | 
 **batchName** | **optional.String**| Batch Name | 
 **bulkCustomerId** | **optional.String**| Bulk Customer ID | 
 **boOfficeId** | **optional.Int32**| Branch Office ID | 
 **date** | **optional.String**| Date in DD-MM-YYYY format | 
 **personnelId** | **optional.Int32**| Personnel ID | 
 **sequencing** | **optional.Bool**| If true, orders by sequence | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListArticlesApiResponse**](response.ListArticlesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesInitialStatisticsGet**
> ResponseFetchArticlesInitialStatisticsApiResponse ArticlesInitialStatisticsGet(ctx, officeId)
Fetch Articles Initial Statistics

This endpoint retrieves initial statistics for articles based on the specified office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 

### Return type

[**ResponseFetchArticlesInitialStatisticsApiResponse**](response.FetchArticlesInitialStatisticsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesSupervisorApprovalByPersonnelIdPost**
> ResponseCreateArticlesSupervisorApprovalByPersonnelIdApiResponse ArticlesSupervisorApprovalByPersonnelIdPost(ctx, officeId, invoicedDate, optional)
Create Articles Supervisor Approval by Personnel ID

This endpoint retrieves a list of articles pending supervisor approval based on the provided office ID, invoiced date, and personnel ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
  **invoicedDate** | **string**| Invoiced Date | 
 **optional** | ***ArticlesApiArticlesSupervisorApprovalByPersonnelIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticlesApiArticlesSupervisorApprovalByPersonnelIdPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **personnelId** | **optional.Int32**| Personnel ID | 

### Return type

[**ResponseCreateArticlesSupervisorApprovalByPersonnelIdApiResponse**](response.CreateArticlesSupervisorApprovalByPersonnelIdAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

