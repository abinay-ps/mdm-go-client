# \EMOsApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmosBoDeliverySlipGenerationPost**](EMOsApi.md#EmosBoDeliverySlipGenerationPost) | **Post** /emos/bo-delivery-slip-generation | Generate EMO Delivery Slips
[**EmosBulkPrintPost**](EMOsApi.md#EmosBulkPrintPost) | **Post** /emos/bulk-print | Bulk Print EMOs
[**EmosBulkReceiptPost**](EMOsApi.md#EmosBulkReceiptPost) | **Post** /emos/bulk-receipt | Bulk Receipt of EMOs
[**EmosBulkSupervisorApprovalForCashPost**](EMOsApi.md#EmosBulkSupervisorApprovalForCashPost) | **Post** /emos/bulk-supervisor-approval-for-cash | Bulk Supervisor Approval for EMOs Cash
[**EmosBulkUpdateRemarksFromPmaPut**](EMOsApi.md#EmosBulkUpdateRemarksFromPmaPut) | **Put** /emos/bulk-update-remarks-from-pma | Update EMOs remarks from PMA
[**EmosBulkWindowDeliveryPost**](EMOsApi.md#EmosBulkWindowDeliveryPost) | **Post** /emos/bulk-window-delivery | Bulk Delivery of EMOs through Window
[**EmosCountForAbstractGet**](EMOsApi.md#EmosCountForAbstractGet) | **Get** /emos/count-for-abstract | Fetch EMOs Count for Abstract
[**EmosEmoNumberOfficeIdInternalTrackingGet**](EMOsApi.md#EmosEmoNumberOfficeIdInternalTrackingGet) | **Get** /emos/{emo-number}/{office-id}/internal-tracking | Fetch EMO Internal Tracking Details
[**EmosEmoNumberRemarksGet**](EMOsApi.md#EmosEmoNumberRemarksGet) | **Get** /emos/{emo-number}/remarks | Retrieve EMO Remarks
[**EmosGet**](EMOsApi.md#EmosGet) | **Get** /emos | List EMOs
[**EmosInvoicePost**](EMOsApi.md#EmosInvoicePost) | **Post** /emos/invoice | Create EMOs Invoice
[**EmosSupervisorApprovalGet**](EMOsApi.md#EmosSupervisorApprovalGet) | **Get** /emos/supervisor-approval | List EMOs for Supervisor Approval


# **EmosBoDeliverySlipGenerationPost**
> ResponseCreateEmOsBoDeliverySlipGenerationApiResponse EmosBoDeliverySlipGenerationPost(ctx, request)
Generate EMO Delivery Slips

Creates delivery slips for EMOs invoiced to the specified BO on the provided invoice date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateEmOsBoDeliverySlipGenerationRequest**](HandlerCreateEmOsBoDeliverySlipGenerationRequest.md)| Create EMOs BO Delivery Slip Generation Request | 

### Return type

[**ResponseCreateEmOsBoDeliverySlipGenerationApiResponse**](response.CreateEMOsBODeliverySlipGenerationAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmosBulkPrintPost**
> ResponseCreateEmOsBulkReceiptApiResponse EmosBulkPrintPost(ctx, request)
Bulk Print EMOs

Initiates the printing process for a batch of EMOs based on the provided parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateEmOsBulkPrintRequest**](HandlerCreateEmOsBulkPrintRequest.md)| Create EMOs Bulk Print Request | 

### Return type

[**ResponseCreateEmOsBulkReceiptApiResponse**](response.CreateEMOsBulkReceiptAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmosBulkReceiptPost**
> ResponseCreateEmOsBulkReceiptApiResponse EmosBulkReceiptPost(ctx, request)
Bulk Receipt of EMOs

Initiates the process for bulk receipt of EMOs based on the provided office ID and current user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateEmOsBulkReceiptRequest**](HandlerCreateEmOsBulkReceiptRequest.md)| Create EMOs Bulk Receipt Request | 

### Return type

[**ResponseCreateEmOsBulkReceiptApiResponse**](response.CreateEMOsBulkReceiptAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmosBulkSupervisorApprovalForCashPost**
> ResponseCreateEmOsBulkSupervisorApprovalForCashApiResponse EmosBulkSupervisorApprovalForCashPost(ctx, request)
Bulk Supervisor Approval for EMOs Cash

Processes the approval of multiple EMOs by the supervisor for cash transactions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateEmOsBulkSupervisorApprovalForCashRequest**](HandlerCreateEmOsBulkSupervisorApprovalForCashRequest.md)| Create EMOs Bulk Supervisor Approval Request | 

### Return type

[**ResponseCreateEmOsBulkSupervisorApprovalForCashApiResponse**](response.CreateEMOsBulkSupervisorApprovalForCashAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmosBulkUpdateRemarksFromPmaPut**
> ResponseCreateEmOsBulkUpdateRemarksFromPmaapiResponse EmosBulkUpdateRemarksFromPmaPut(ctx, officeId)
Update EMOs remarks from PMA

Updates the articles remarks taken by postman in the article_transaction table based on the office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 

### Return type

[**ResponseCreateEmOsBulkUpdateRemarksFromPmaapiResponse**](response.CreateEMOsBulkUpdateRemarksFromPMAAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmosBulkWindowDeliveryPost**
> ResponseCreateEmOsBulkWindowDeliveryApiResponse EmosBulkWindowDeliveryPost(ctx, request)
Bulk Delivery of EMOs through Window

Processes the delivery of multiple EMOs through the window for the specified office.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateEmOsBulkWindowDeliveryRequest**](HandlerCreateEmOsBulkWindowDeliveryRequest.md)| Create EMOs Bulk Window Delivery Request | 

### Return type

[**ResponseCreateEmOsBulkWindowDeliveryApiResponse**](response.CreateEMOsBulkWindowDeliveryAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmosCountForAbstractGet**
> ResponseFetchEmOsCountForAbstractApiResponse EmosCountForAbstractGet(ctx, officeId)
Fetch EMOs Count for Abstract

Fetch counts of various EMOs based on the office ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 

### Return type

[**ResponseFetchEmOsCountForAbstractApiResponse**](response.FetchEMOsCountForAbstractAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmosEmoNumberOfficeIdInternalTrackingGet**
> ResponseFetchEmoInternalTrackingApiResponse EmosEmoNumberOfficeIdInternalTrackingGet(ctx, emoNumber, officeId)
Fetch EMO Internal Tracking Details

Retrieve internal tracking details of an EMO based on the EMO number and office ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emoNumber** | **string**| EMO Number | 
  **officeId** | **int32**| Current Office ID | 

### Return type

[**ResponseFetchEmoInternalTrackingApiResponse**](response.FetchEMOInternalTrackingAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmosEmoNumberRemarksGet**
> ResponseFetchEmoRemarksApiResponse EmosEmoNumberRemarksGet(ctx, emoNumber, optional)
Retrieve EMO Remarks

Fetches the remarks for the given EMO number and office type (single-hand or multi-hand).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emoNumber** | **string**| EMO Number | 
 **optional** | ***EMOsApiEmosEmoNumberRemarksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EMOsApiEmosEmoNumberRemarksGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isSingleHandOffice** | **optional.Bool**| Is Single Hand Office | 

### Return type

[**ResponseFetchEmoRemarksApiResponse**](response.FetchEMORemarksAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmosGet**
> ResponseListEmOsApiResponse EmosGet(ctx, emoState, officeId, optional)
List EMOs

Returns a filtered list of EMOs based on their state, receipt type, and delivery location

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emoState** | **string**| Article state(received, deposited, invoiced, transmitted, pending) | 
  **officeId** | **int32**| ID of the office | 
 **optional** | ***EMOsApiEmosGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EMOsApiEmosGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **receiptType** | **optional.String**| Receipt type for received EMOs(pending-emos, all-emos, received-emos) | 
 **pendingType** | **optional.String**| Pending type for pending EMOs(not-printed, above-seven-days) | 
 **deliveryLocation** | **optional.String**| Delivery location(office, branch-office, beat, bulk-customer) | 
 **beatName** | **optional.String**| Name of the beat | 
 **batchName** | **optional.String**| Name of the batch | 
 **bulkCustomerId** | **optional.String**| ID of the bulk customer | 
 **boOfficeId** | **optional.Int32**| ID of the branch office | 
 **date** | **optional.String**| Date for invoiced or transmitted EMOs | 
 **personnelId** | **optional.Int32**| ID of the personnel for branch office | 
 **sequencing** | **optional.Bool**| Flag for sequencing EMOs | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListEmOsApiResponse**](response.ListEMOsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmosInvoicePost**
> ResponseCreateEmOsInvoiceApiResponse EmosInvoicePost(ctx, invoiceType, request, optional)
Create EMOs Invoice

Generates invoices for EMOs based on the provided parameters, including delivery location and invoice type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceType** | **string**| Invoice type for EMOs(invoice,transmit) | 
  **request** | [**HandlerCreateEmOsInvoiceRequest**](HandlerCreateEmOsInvoiceRequest.md)| Create EMOs Invoice Request | 
 **optional** | ***EMOsApiEmosInvoicePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EMOsApiEmosInvoicePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deliveryLocation** | **optional.String**| Delivery location(office, branch-office, beat, bulk-customer) | 
 **isSingleHandOffice** | **optional.Bool**| true/false | 

### Return type

[**ResponseCreateEmOsInvoiceApiResponse**](response.CreateEMOsInvoiceAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmosSupervisorApprovalGet**
> ResponseListEmOsSupervisorApprovalApiResponse EmosSupervisorApprovalGet(ctx, supervisorApprovalType, officeId, optional)
List EMOs for Supervisor Approval

This endpoint retrieves EMOs pending for supervisor approval based on the specified approval type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supervisorApprovalType** | **string**| Supervisor Approval Type(cash-for-postman, returns, bulk-delivered) | 
  **officeId** | **int32**| Office ID | 
 **optional** | ***EMOsApiEmosSupervisorApprovalGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EMOsApiEmosSupervisorApprovalGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **beatName** | **optional.String**| Beat Name | 
 **batchName** | **optional.String**| Batch Name | 
 **bulkAddresseeId** | **optional.String**| Bulk Customer ID | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListEmOsSupervisorApprovalApiResponse**](response.ListEMOsSupervisorApprovalAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

