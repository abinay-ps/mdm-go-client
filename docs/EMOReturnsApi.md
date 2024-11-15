# \EMOReturnsApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmoReturnsBulkAlterReturnsPost**](EMOReturnsApi.md#EmoReturnsBulkAlterReturnsPost) | **Post** /emo-returns/bulk-alter-returns | Modify EMO Returns in Bulk
[**EmoReturnsBulkMissentRedirectionWithoutInvoicingPost**](EMOReturnsApi.md#EmoReturnsBulkMissentRedirectionWithoutInvoicingPost) | **Post** /emo-returns/bulk-missent-redirection-without-invoicing | Redirect Missent EMO Returns Without Invoicing
[**EmoReturnsBulkSupervisorApprovalOfBulkCustomersDeliveredPost**](EMOReturnsApi.md#EmoReturnsBulkSupervisorApprovalOfBulkCustomersDeliveredPost) | **Post** /emo-returns/bulk-supervisor-approval-of-bulk-customers-delivered | Supervisor Approval of Bulk EMO Customer Deliveries
[**EmoReturnsBulkSupervisorApprovalOfReturnsPost**](EMOReturnsApi.md#EmoReturnsBulkSupervisorApprovalOfReturnsPost) | **Post** /emo-returns/bulk-supervisor-approval-of-returns | Approve Bulk EMO Returns
[**EmoReturnsOfficeIdGet**](EMOReturnsApi.md#EmoReturnsOfficeIdGet) | **Get** /emo-returns/{office-id} | List EMO Returns
[**EmoReturnsPost**](EMOReturnsApi.md#EmoReturnsPost) | **Post** /emo-returns | Create EMO Returns


# **EmoReturnsBulkAlterReturnsPost**
> ResponseCreateEmoReturnsBulkAlterReturnsApiResponse EmoReturnsBulkAlterReturnsPost(ctx, returnsType, createEMOReturnsBulkAlterReturnsRequest)
Modify EMO Returns in Bulk

This endpoint modifies EMO returns in bulk based on the specified returns type (beat or branch-office).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **returnsType** | **string**| Modify Returns Type (beat, branch-office) | 
  **createEMOReturnsBulkAlterReturnsRequest** | [**HandlerCreateEmoReturnsBulkAlterReturnsRequest**](HandlerCreateEmoReturnsBulkAlterReturnsRequest.md)| Bulk Modify EMO Returns Request Body | 

### Return type

[**ResponseCreateEmoReturnsBulkAlterReturnsApiResponse**](response.CreateEMOReturnsBulkAlterReturnsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmoReturnsBulkMissentRedirectionWithoutInvoicingPost**
> ResponseCreateEmoReturnsBulkMissentRedirectionWithoutInvoicingApiResponse EmoReturnsBulkMissentRedirectionWithoutInvoicingPost(ctx, createEMOReturnsBulkMissentRedirectionWithoutInvoicingRequest)
Redirect Missent EMO Returns Without Invoicing

This endpoint handles the redirection of bulk missent EMO returns without generating an invoice.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createEMOReturnsBulkMissentRedirectionWithoutInvoicingRequest** | [**HandlerCreateEmoReturnsBulkMissentRedirectionWithoutInvoicingRequest**](HandlerCreateEmoReturnsBulkMissentRedirectionWithoutInvoicingRequest.md)| Bulk Missent EMO Redirection Request Body | 

### Return type

[**ResponseCreateEmoReturnsBulkMissentRedirectionWithoutInvoicingApiResponse**](response.CreateEMOReturnsBulkMissentRedirectionWithoutInvoicingAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmoReturnsBulkSupervisorApprovalOfBulkCustomersDeliveredPost**
> ResponseCreateEmoReturnsBulkSupervisorApprovalOfBulkCustomersDeliveryApiResponse EmoReturnsBulkSupervisorApprovalOfBulkCustomersDeliveredPost(ctx, createEMOReturnsBulkSupervisorApprovalOfBulkCustomersDeliveryRequest)
Supervisor Approval of Bulk EMO Customer Deliveries

This endpoint handles the supervisor's approval for the bulk delivery of EMOs to customers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createEMOReturnsBulkSupervisorApprovalOfBulkCustomersDeliveryRequest** | [**HandlerCreateEmoReturnsBulkSupervisorApprovalOfBulkCustomersDeliveryRequest**](HandlerCreateEmoReturnsBulkSupervisorApprovalOfBulkCustomersDeliveryRequest.md)| Bulk EMO Customer Delivery Approval Request Body | 

### Return type

[**ResponseCreateEmoReturnsBulkSupervisorApprovalOfBulkCustomersDeliveryApiResponse**](response.CreateEMOReturnsBulkSupervisorApprovalOfBulkCustomersDeliveryAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmoReturnsBulkSupervisorApprovalOfReturnsPost**
> ResponseCreateEmoReturnsBulkSupervisorApprovalApiResponse EmoReturnsBulkSupervisorApprovalOfReturnsPost(ctx, createEMOReturnsBulkSupervisorApprovalRequest)
Approve Bulk EMO Returns

This endpoint is used by supervisors to confirm and approve bulk EMO returns.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createEMOReturnsBulkSupervisorApprovalRequest** | [**HandlerCreateEmoReturnsBulkSupervisorApprovalRequest**](HandlerCreateEmoReturnsBulkSupervisorApprovalRequest.md)| Bulk Supervisor Approval of EMO Returns Request Body | 

### Return type

[**ResponseCreateEmoReturnsBulkSupervisorApprovalApiResponse**](response.CreateEMOReturnsBulkSupervisorApprovalAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmoReturnsOfficeIdGet**
> ResponseListEmoReturnsApiResponse EmoReturnsOfficeIdGet(ctx, officeId, operationType, optional)
List EMO Returns

Lists EMO returns based on the specified operation type, including unprocessed, returns not taken, returned, or returned by personnel ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
  **operationType** | **string**| Operation type(&#39;unprocessed&#39;, &#39;returns-not-taken&#39;, &#39;returned&#39;, &#39;returned-by-personnel-id&#39;) | 
 **optional** | ***EMOReturnsApiEmoReturnsOfficeIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EMOReturnsApiEmoReturnsOfficeIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchName** | **optional.String**| Batch name filter | 
 **beatName** | **optional.String**| Beat name filter | 
 **personnelId** | **optional.Int32**| Personnel ID (required for &#39;returned-by-personnel-id&#39; operation type) | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListEmoReturnsApiResponse**](response.ListEMOReturnsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmoReturnsPost**
> ResponseCreateEmoReturnsApiResponse EmoReturnsPost(ctx, returnsType, createEMOReturnsRequest)
Create EMO Returns

This endpoint creates EMO returns depending on the specified returns type (beat, branch-office, or direct).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **returnsType** | **string**| Returns Type (beat, branch-office, direct) | 
  **createEMOReturnsRequest** | [**HandlerCreateEmoReturnsRequest**](HandlerCreateEmoReturnsRequest.md)| Bulk EMO Returns Request Body | 

### Return type

[**ResponseCreateEmoReturnsApiResponse**](response.CreateEMOReturnsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

