# \PLIRPLIApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InsurancequotegenerationPliPost**](PLIRPLIApi.md#InsurancequotegenerationPliPost) | **Post** /insurancequotegeneration/pli | PLI Quote Generation
[**InsurancequotegenerationRpliPost**](PLIRPLIApi.md#InsurancequotegenerationRpliPost) | **Post** /insurancequotegeneration/rpli | RPLI Quote Generation


# **InsurancequotegenerationPliPost**
> HandlerResponse InsurancequotegenerationPliPost(ctx, pLIQuoteGenerationRequest)
PLI Quote Generation

Generates quote for six PLI policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pLIQuoteGenerationRequest** | [**DomainPliQuoteGenerationRequest**](DomainPliQuoteGenerationRequest.md)| PLI Quote Generation request particulars | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsurancequotegenerationRpliPost**
> HandlerResponse InsurancequotegenerationRpliPost(ctx, rPLIQuoteGenerationRequest)
RPLI Quote Generation

Generates quote for six RPLI policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rPLIQuoteGenerationRequest** | [**DomainRpliQuoteGenerationRequest**](DomainRpliQuoteGenerationRequest.md)| RPLI Quote Generation request particulars | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

