# \PaymentsApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentsQrCodesPost**](PaymentsApi.md#PaymentsQrCodesPost) | **Post** /payments/qr-codes | Generate a QR Code
[**PaymentsTransactionIdGet**](PaymentsApi.md#PaymentsTransactionIdGet) | **Get** /payments/{transaction-id} | Fetch Payment Transaction Details


# **PaymentsQrCodesPost**
> string PaymentsQrCodesPost(ctx, body)
Generate a QR Code

This endpoint generates a QR code based on the provided transaction details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HandlerCreatePaymentsRequest**](HandlerCreatePaymentsRequest.md)| QR Code Generation Request | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PaymentsTransactionIdGet**
> ResponseFetchPaymentsApiResponse PaymentsTransactionIdGet(ctx, transactionId, request)
Fetch Payment Transaction Details

Fetches payment transaction details based on the provided transaction ID and updates the internal transaction data if successful.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transactionId** | **string**| Transaction ID | 
  **request** | [**HandlerFetchPaymentsRequest**](HandlerFetchPaymentsRequest.md)| Payment Fetch Request | 

### Return type

[**ResponseFetchPaymentsApiResponse**](response.FetchPaymentsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

