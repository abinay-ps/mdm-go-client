# \InsuranceApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InsuranceQuotesPliGet**](InsuranceApi.md#InsuranceQuotesPliGet) | **Get** /insurance-quotes/pli | Generate PLI (Postal Life Insurance) Quote
[**InsuranceQuotesRpliGet**](InsuranceApi.md#InsuranceQuotesRpliGet) | **Get** /insurance-quotes/rpli | Generate RPLI (Rural Postal Life Insurance) Quote


# **InsuranceQuotesPliGet**
> ResponseGetPliQuoteFinalResponseOffApiResponse InsuranceQuotesPliGet(ctx, sumAssured, ageOnNextBirthday, periodicity, isNewPolicy, optional)
Generate PLI (Postal Life Insurance) Quote

Generates a quote for Postal Life Insurance based on the provided details such as sum assured, age, periodicity, and policy type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sumAssured** | **float32**| Sum assured for the policy | 
  **ageOnNextBirthday** | **float32**| Age of the policyholder on next birthday | 
  **periodicity** | **int32**| Periodicity of premium payment (1: Monthly, 2: Quarterly, 3: Half-Yearly, 4: Yearly) | 
  **isNewPolicy** | **int32**| Indicates if it is a new policy (1: Yes, 2: No) | 
 **optional** | ***InsuranceApiInsuranceQuotesPliGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsuranceApiInsuranceQuotesPliGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **spouseAge** | **optional.Float32**| Age of the spouse (if applicable) | 
 **childAge** | **optional.Float32**| Age of the child (if applicable) | 
 **sumAssuredChildPolicy** | **optional.Float32**| Sum assured for child policy (if child policy is opted) | 

### Return type

[**ResponseGetPliQuoteFinalResponseOffApiResponse**](response.GetPLIQuoteFinalResponseOffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsuranceQuotesRpliGet**
> ResponseGetRpliQuoteFinalResponseOffApiResponse InsuranceQuotesRpliGet(ctx, sumAssured, ageOnNextBirthday, periodicity, isNewPolicy, isNonStdAge, optional)
Generate RPLI (Rural Postal Life Insurance) Quote

Generates a quote for Rural Postal Life Insurance based on the provided details such as sum assured, age, periodicity, and policy type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sumAssured** | **float32**| Sum assured for the policy | 
  **ageOnNextBirthday** | **float32**| Age of the policyholder on next birthday | 
  **periodicity** | **int32**| Periodicity of premium payment (1: Monthly, 2: Quarterly, 3: Half-Yearly, 4: Yearly) | 
  **isNewPolicy** | **int32**| Indicates if it is a new policy (1: New, 2: Old) | 
  **isNonStdAge** | **int32**| Indicates if the age is non-standard (1: Standard, 2: Non-Standard) | 
 **optional** | ***InsuranceApiInsuranceQuotesRpliGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsuranceApiInsuranceQuotesRpliGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **childAge** | **optional.Float32**| Age of the child (if applicable) | 
 **sumAssuredChildPolicy** | **optional.Float32**| Sum assured for child policy (if child policy is opted) | 

### Return type

[**ResponseGetRpliQuoteFinalResponseOffApiResponse**](response.GetRPLIQuoteFinalResponseOffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

