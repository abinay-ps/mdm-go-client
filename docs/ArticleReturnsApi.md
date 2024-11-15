# \ArticleReturnsApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArticleReturnsBulkAlterReturnsPost**](ArticleReturnsApi.md#ArticleReturnsBulkAlterReturnsPost) | **Post** /article-returns/bulk-alter-returns | Modifies article returns in bulk
[**ArticleReturnsBulkMissentRedirectionWithoutInvoicingPost**](ArticleReturnsApi.md#ArticleReturnsBulkMissentRedirectionWithoutInvoicingPost) | **Post** /article-returns/bulk-missent-redirection-without-invoicing | Redirect missent articles without invoicing
[**ArticleReturnsBulkSuperApprovalForMissentPost**](ArticleReturnsApi.md#ArticleReturnsBulkSuperApprovalForMissentPost) | **Post** /article-returns/bulk-super-approval-for-missent | Approve missent articles redirection without invoicing
[**ArticleReturnsBulkSupervisorApprovalOfBoReturnsPost**](ArticleReturnsApi.md#ArticleReturnsBulkSupervisorApprovalOfBoReturnsPost) | **Post** /article-returns/bulk-supervisor-approval-of-bo-returns | Supervisor Approval for Branch Office Returns
[**ArticleReturnsBulkSupervisorApprovalOfBulkCustomerReturnsPost**](ArticleReturnsApi.md#ArticleReturnsBulkSupervisorApprovalOfBulkCustomerReturnsPost) | **Post** /article-returns/bulk-supervisor-approval-of-bulk-customer-returns | Supervisor Approval for Bulk Customer Returns
[**ArticleReturnsBulkSupervisorApprovalOfReturnsPost**](ArticleReturnsApi.md#ArticleReturnsBulkSupervisorApprovalOfReturnsPost) | **Post** /article-returns/bulk-supervisor-approval-of-returns | Approves bulk article returns
[**ArticleReturnsOfficeIdGet**](ArticleReturnsApi.md#ArticleReturnsOfficeIdGet) | **Get** /article-returns/{office-id} | Retrieves article returns
[**ArticleReturnsPost**](ArticleReturnsApi.md#ArticleReturnsPost) | **Post** /article-returns | Creates article returns based on the returns type


# **ArticleReturnsBulkAlterReturnsPost**
> ResponseCreateArticleReturnsBulkAlterReturnsApiResponse ArticleReturnsBulkAlterReturnsPost(ctx, returnsType, createArticleReturnsBulkAlterReturnsRequest)
Modifies article returns in bulk

Modifies article returns for either beat or branch office types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **returnsType** | **string**| Returns Type (beat, branch-office) | 
  **createArticleReturnsBulkAlterReturnsRequest** | [**HandlerCreateArticleReturnsBulkAlterReturnsRequest**](HandlerCreateArticleReturnsBulkAlterReturnsRequest.md)| Create Article Returns Bulk Alter Returns Request | 

### Return type

[**ResponseCreateArticleReturnsBulkAlterReturnsApiResponse**](response.CreateArticleReturnsBulkAlterReturnsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleReturnsBulkMissentRedirectionWithoutInvoicingPost**
> ResponseCreateArticleReturnsBulkMissentRedirectionWithoutInvoicingApiResponse ArticleReturnsBulkMissentRedirectionWithoutInvoicingPost(ctx, request)
Redirect missent articles without invoicing

Redirect a batch of missent articles to the correct office without generating an invoice.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateArticleReturnsBulkMissentRedirectionWithoutInvoicingRequest**](HandlerCreateArticleReturnsBulkMissentRedirectionWithoutInvoicingRequest.md)| Missent Article Redirection Without Invoicing Request | 

### Return type

[**ResponseCreateArticleReturnsBulkMissentRedirectionWithoutInvoicingApiResponse**](response.CreateArticleReturnsBulkMissentRedirectionWithoutInvoicingAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleReturnsBulkSuperApprovalForMissentPost**
> ResponseCreateArticleReturnsBulkSupervisorApprovalForMissentApiResponse ArticleReturnsBulkSuperApprovalForMissentPost(ctx, request)
Approve missent articles redirection without invoicing

Supervisor approval for redirecting missent articles to the correct office without invoicing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateArticleReturnsBulkSupervisorApprovalForMissentRequest**](HandlerCreateArticleReturnsBulkSupervisorApprovalForMissentRequest.md)| Supervisor Approval for Missent Article Redirection | 

### Return type

[**ResponseCreateArticleReturnsBulkSupervisorApprovalForMissentApiResponse**](response.CreateArticleReturnsBulkSupervisorApprovalForMissentAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleReturnsBulkSupervisorApprovalOfBoReturnsPost**
> ResponseCreateArticleReturnsBulkSupervisorApprovalOfBoReturnsApiResponse ArticleReturnsBulkSupervisorApprovalOfBoReturnsPost(ctx, request)
Supervisor Approval for Branch Office Returns

Approves the bulk returns of articles from branch office for supervisor approval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateArticleReturnsBulkSupervisorApprovalOfBoReturnsRequest**](HandlerCreateArticleReturnsBulkSupervisorApprovalOfBoReturnsRequest.md)| Supervisor Approval for Branch Office Returns | 

### Return type

[**ResponseCreateArticleReturnsBulkSupervisorApprovalOfBoReturnsApiResponse**](response.CreateArticleReturnsBulkSupervisorApprovalOfBOReturnsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleReturnsBulkSupervisorApprovalOfBulkCustomerReturnsPost**
> ResponseCreateArticleReturnsBulkSupervisorApprovalOfBoReturnsApiResponse ArticleReturnsBulkSupervisorApprovalOfBulkCustomerReturnsPost(ctx, request)
Supervisor Approval for Bulk Customer Returns

Approves the bulk returns of articles for customer supervision approval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateArticleReturnsBulkSupervisorApprovalOfBulkCustomerReturnsRequest**](HandlerCreateArticleReturnsBulkSupervisorApprovalOfBulkCustomerReturnsRequest.md)| Supervisor Approval for Bulk Customer Returns | 

### Return type

[**ResponseCreateArticleReturnsBulkSupervisorApprovalOfBoReturnsApiResponse**](response.CreateArticleReturnsBulkSupervisorApprovalOfBOReturnsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleReturnsBulkSupervisorApprovalOfReturnsPost**
> ResponseCreateArticleReturnsBulkSupervisorApprovalApiResponse ArticleReturnsBulkSupervisorApprovalOfReturnsPost(ctx, request)
Approves bulk article returns

Approves a list of article returns based on the provided article remarks, office ID, and beat name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateArticleReturnsBulkSupervisorApprovalRequest**](HandlerCreateArticleReturnsBulkSupervisorApprovalRequest.md)| Article Returns Approval Request | 

### Return type

[**ResponseCreateArticleReturnsBulkSupervisorApprovalApiResponse**](response.CreateArticleReturnsBulkSupervisorApprovalAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleReturnsOfficeIdGet**
> ResponseListArticleReturnsApiResponse ArticleReturnsOfficeIdGet(ctx, officeId, operationType, optional)
Retrieves article returns

Retrieves a list of article returns based on operation type, office ID, and other optional filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
  **operationType** | **string**| Operation Type (unprocessed, returns-not-taken, returned, window-delivered, returned-by-personnel-id) | 
 **optional** | ***ArticleReturnsApiArticleReturnsOfficeIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticleReturnsApiArticleReturnsOfficeIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchName** | **optional.String**| Batch Name | 
 **beatName** | **optional.String**| Beat Name | 
 **personnelId** | **optional.Int32**| Personnel ID (required if operation type is &#39;returned-by-personnel-id&#39;) | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListArticleReturnsApiResponse**](response.ListArticleReturnsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleReturnsPost**
> ResponseCreateArticleReturnsApiResponse ArticleReturnsPost(ctx, returnsType, createArticleReturnsRequest)
Creates article returns based on the returns type

Creates returns for articles based on their type (beat, branch-office, or direct)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **returnsType** | **string**| Returns Type (beat, branch-office, direct) | 
  **createArticleReturnsRequest** | [**HandlerCreateArticleReturnsRequest**](HandlerCreateArticleReturnsRequest.md)| Create Article Returns Request | 

### Return type

[**ResponseCreateArticleReturnsApiResponse**](response.CreateArticleReturnsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

