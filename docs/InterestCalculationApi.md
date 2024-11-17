# \InterestCalculationApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PosbInterestCalculationsMisGet**](InterestCalculationApi.md#PosbInterestCalculationsMisGet) | **Get** /posb-interest-calculations/mis | Calculate interest for MIS accounts
[**PosbInterestCalculationsMsscGet**](InterestCalculationApi.md#PosbInterestCalculationsMsscGet) | **Get** /posb-interest-calculations/mssc | Calculate interest for MSSC accounts
[**PosbInterestCalculationsNscKvpGet**](InterestCalculationApi.md#PosbInterestCalculationsNscKvpGet) | **Get** /posb-interest-calculations/nsc-kvp | Calculate interest for NSC/KVP accounts
[**PosbInterestCalculationsPpfGet**](InterestCalculationApi.md#PosbInterestCalculationsPpfGet) | **Get** /posb-interest-calculations/ppf | Calculate interest for PPF accounts
[**PosbInterestCalculationsRdGet**](InterestCalculationApi.md#PosbInterestCalculationsRdGet) | **Get** /posb-interest-calculations/rd | Calculate interest for RD accounts
[**PosbInterestCalculationsSbGet**](InterestCalculationApi.md#PosbInterestCalculationsSbGet) | **Get** /posb-interest-calculations/sb | Calculate interest for SB accounts
[**PosbInterestCalculationsScssGet**](InterestCalculationApi.md#PosbInterestCalculationsScssGet) | **Get** /posb-interest-calculations/scss | Calculate interest for SCSS accounts
[**PosbInterestCalculationsSsaGet**](InterestCalculationApi.md#PosbInterestCalculationsSsaGet) | **Get** /posb-interest-calculations/ssa | Calculate interest for SSA accounts
[**PosbInterestCalculationsTdGet**](InterestCalculationApi.md#PosbInterestCalculationsTdGet) | **Get** /posb-interest-calculations/td | Calculate interest for TD accounts


# **PosbInterestCalculationsMisGet**
> ResponseGetMisInterestFinalApiResponse PosbInterestCalculationsMisGet(ctx, accountType, amount, productCode)
Calculate interest for MIS accounts

Calculates the interest for the Monthly Income Scheme (MIS) based on the provided amount and parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountType** | **string**| Type of account (Single or Joint) | 
  **amount** | **float32**| Amount for the interest calculation (must be non-negative) | 
  **productCode** | **string**| Product code for the Monthly Income Scheme account | 

### Return type

[**ResponseGetMisInterestFinalApiResponse**](response.GetMISInterestFinalAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PosbInterestCalculationsMsscGet**
> ResponseGetMsscInterestFinalApiResponse PosbInterestCalculationsMsscGet(ctx, amount, productCode)
Calculate interest for MSSC accounts

Calculates the interest for the Mahila Samman Savings Certificate (MSSC) based on the provided amount and parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amount** | **float32**| Amount for the interest calculation | 
  **productCode** | **string**| Product code for the Mahila Samman Savings Certificate | 

### Return type

[**ResponseGetMsscInterestFinalApiResponse**](response.GetMSSCInterestFinalAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PosbInterestCalculationsNscKvpGet**
> ResponseGetNsckvpInterestFinalApiResponse PosbInterestCalculationsNscKvpGet(ctx, amount, cbNscKvpType, productCode)
Calculate interest for NSC/KVP accounts

Calculates the interest for the National Saving Certificate (NSC) and Kisan Vikas Patra (KVP) based on the provided amount and parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amount** | **float32**| Amount for the interest calculation | 
  **cbNscKvpType** | **string**| Type of account (NSC or KVP) | 
  **productCode** | **string**| Product code for NSC/KVP accounts | 

### Return type

[**ResponseGetNsckvpInterestFinalApiResponse**](response.GetNSCKVPInterestFinalAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PosbInterestCalculationsPpfGet**
> ResponseGetPpfInterestFinalApiResponse PosbInterestCalculationsPpfGet(ctx, amount, productCode)
Calculate interest for PPF accounts

Calculates the interest for the Public Provident Fund (PPF) based on the provided amount and parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amount** | **float32**|  | 
  **productCode** | **string**|  | 

### Return type

[**ResponseGetPpfInterestFinalApiResponse**](response.GetPPFInterestFinalAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PosbInterestCalculationsRdGet**
> ResponseGetRdInterestFinalApiResponse PosbInterestCalculationsRdGet(ctx, amount, productCode)
Calculate interest for RD accounts

Calculates the interest for Recurring Deposit (RD) based on the provided amount and parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amount** | **float32**| Amount for the interest calculation | 
  **productCode** | **string**| Product code for the Recurring Deposit account | 

### Return type

[**ResponseGetRdInterestFinalApiResponse**](response.GetRDInterestFinalAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PosbInterestCalculationsSbGet**
> ResponseGetSbInterestApiResponse PosbInterestCalculationsSbGet(ctx, amount, productCode)
Calculate interest for SB accounts

Calculates the interest for Savings Bank (SB) based on the provided amount and parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amount** | **float32**| Amount for the interest calculation | 
  **productCode** | **string**| Product code for the Savings Bank account | 

### Return type

[**ResponseGetSbInterestApiResponse**](response.GetSBInterestAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PosbInterestCalculationsScssGet**
> ResponseGetScssInterestFinalApiResponse PosbInterestCalculationsScssGet(ctx, amount, productCode)
Calculate interest for SCSS accounts

Calculates the interest for the Senior Citizen Savings Scheme (SCSS) based on the provided amount and parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amount** | **float32**| Amount for the interest calculation | 
  **productCode** | **string**| Product code for the SCSS account | 

### Return type

[**ResponseGetScssInterestFinalApiResponse**](response.GetSCSSInterestFinalAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PosbInterestCalculationsSsaGet**
> ResponseGetSsaInterestFinalApiResponse PosbInterestCalculationsSsaGet(ctx, amount, girlAge, productCode)
Calculate interest for SSA accounts

Calculates the interest for the Sukanya Samriddhi Account (SSA) based on the provided amount and parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amount** | **float32**| Amount for the interest calculation | 
  **girlAge** | **int32**| Age of the girl (required, must be non-negative) | 
  **productCode** | **string**| Product code for the SSA account | 

### Return type

[**ResponseGetSsaInterestFinalApiResponse**](response.GetSSAInterestFinalAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PosbInterestCalculationsTdGet**
> ResponseGetTdInterestFinalApiResponse PosbInterestCalculationsTdGet(ctx, amount, productCode, optional)
Calculate interest for TD accounts

Calculates the interest for Time Deposit (TD) based on the provided amount and parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amount** | **float32**| Amount for the interest calculation | 
  **productCode** | **string**| Product code for the Time Deposit account | 
 **optional** | ***InterestCalculationApiPosbInterestCalculationsTdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InterestCalculationApiPosbInterestCalculationsTdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cbTdType** | **optional.Int32**| Optional: values: 1, 2, 3, 5 | 

### Return type

[**ResponseGetTdInterestFinalApiResponse**](response.GetTDInterestFinalAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

