# \OTPManagementApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OtpManagementArticleNumberGeneratePost**](OTPManagementApi.md#OtpManagementArticleNumberGeneratePost) | **Post** /otp-management/{article-number}/generate | generates OTP for delivery of articles
[**OtpManagementArticleNumberReSendPut**](OTPManagementApi.md#OtpManagementArticleNumberReSendPut) | **Put** /otp-management/{article-number}/re-send | Resend OTP for Article
[**OtpManagementArticleNumberValidateGet**](OTPManagementApi.md#OtpManagementArticleNumberValidateGet) | **Get** /otp-management/{article-number}/validate | Validate OTP for Article


# **OtpManagementArticleNumberGeneratePost**
> string OtpManagementArticleNumberGeneratePost(ctx, articleNumber)
generates OTP for delivery of articles

generates OTP for delivery of articles based on the article number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleNumber** | **string**| Article number | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OtpManagementArticleNumberReSendPut**
> string OtpManagementArticleNumberReSendPut(ctx, articleNumber)
Resend OTP for Article

Resends the OTP for the specified article number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleNumber** | **string**| Article Number | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OtpManagementArticleNumberValidateGet**
> ResponseFetchOtpManagementValidateApiResponse OtpManagementArticleNumberValidateGet(ctx, articleNumber, otp)
Validate OTP for Article

Validates the OTP for the specified article number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleNumber** | **string**| Article Number | 
  **otp** | **int32**| One-Time Password | 

### Return type

[**ResponseFetchOtpManagementValidateApiResponse**](response.FetchOTPManagementValidateAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

