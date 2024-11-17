# \DemographicApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MiscDemographicsGet**](DemographicApi.md#MiscDemographicsGet) | **Get** /misc/demographics | List demographic data based on state, district, and taluk
[**MiscDistrictsGet**](DemographicApi.md#MiscDistrictsGet) | **Get** /misc/districts | List all districts under a specified state
[**MiscStatesGet**](DemographicApi.md#MiscStatesGet) | **Get** /misc/states | List all states


# **MiscDemographicsGet**
> ResponseListDemographicApiResponse MiscDemographicsGet(ctx, optional)
List demographic data based on state, district, and taluk

Retrieves demographic data including state names, district names, taluk names, and village names.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DemographicApiMiscDemographicsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DemographicApiMiscDemographicsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **districtName** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 
 **stateName** | **optional.String**|  | 
 **talukName** | **optional.String**|  | 

### Return type

[**ResponseListDemographicApiResponse**](response.ListDemographicAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MiscDistrictsGet**
> ResponseListDistrictNamesApiResponse MiscDistrictsGet(ctx, stateName, optional)
List all districts under a specified state

Retrieves a list of districts within a given state. The endpoint checks for cached district names in Redis and fetches them from the service if not available. Supports pagination.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stateName** | **string**|  | 
 **optional** | ***DemographicApiMiscDistrictsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DemographicApiMiscDistrictsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListDistrictNamesApiResponse**](response.ListDistrictNamesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MiscStatesGet**
> ResponseListStateNamesApiResponse MiscStatesGet(ctx, optional)
List all states

Retrieves a list of states. The endpoint checks for cached state names in Redis and fetches them from the service if not available. Supports pagination.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DemographicApiMiscStatesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DemographicApiMiscStatesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListStateNamesApiResponse**](response.ListStateNamesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

