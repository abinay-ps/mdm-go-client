# \CBSInterestApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InterestcalculationPosbPost**](CBSInterestApi.md#InterestcalculationPosbPost) | **Post** /interestcalculation/posb | Calculate Interest


# **InterestcalculationPosbPost**
> HandlerResponse InterestcalculationPosbPost(ctx, request)
Calculate Interest

Calculates CBS Interest charts for various Post Office Savings Schemes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**DomainInterestCalculationRequest**](DomainInterestCalculationRequest.md)| CBS Interest calculation request | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

