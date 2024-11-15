# \SubmitAccountsApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubmitAccountsEmoCashRequestForWindowDeliveryPost**](SubmitAccountsApi.md#SubmitAccountsEmoCashRequestForWindowDeliveryPost) | **Post** /submit-accounts/emo-cash-request-for-window-delivery | Submits a cash request for EMO window delivery
[**SubmitAccountsGet**](SubmitAccountsApi.md#SubmitAccountsGet) | **Get** /submit-accounts | List submit accounts
[**SubmitAccountsOfficeIdAmountStatisticsForDashboardGet**](SubmitAccountsApi.md#SubmitAccountsOfficeIdAmountStatisticsForDashboardGet) | **Get** /submit-accounts/{office-id}/amount-statistics-for-dashboard | Retrieves amount statistics for the specified office and date
[**SubmitAccountsOfficeIdStatusGet**](SubmitAccountsApi.md#SubmitAccountsOfficeIdStatusGet) | **Get** /submit-accounts/{office-id}/status | Retrieves the submission status of accounts for a specific office, beat, batch, and date
[**SubmitAccountsOfficeIdTotalAmountOfWindowDeliveryGet**](SubmitAccountsApi.md#SubmitAccountsOfficeIdTotalAmountOfWindowDeliveryGet) | **Get** /submit-accounts/{office-id}/total-amount-of-window-delivery | Calculates total amount for window delivery
[**SubmitAccountsPost**](SubmitAccountsApi.md#SubmitAccountsPost) | **Post** /submit-accounts | Submits account details for specified office types
[**SubmitAccountsTransactionsAmountCollectedGet**](SubmitAccountsApi.md#SubmitAccountsTransactionsAmountCollectedGet) | **Get** /submit-accounts/transactions/amount-collected | Submit account amount collected


# **SubmitAccountsEmoCashRequestForWindowDeliveryPost**
> string SubmitAccountsEmoCashRequestForWindowDeliveryPost(ctx, body)
Submits a cash request for EMO window delivery

Submits a request to the treasury for cash needed for EMO window delivery based on the provided details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainSubmitAccount**](DomainSubmitAccount.md)| Submit Account Request | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitAccountsGet**
> ResponseListSubmitAccountsApiResponse SubmitAccountsGet(ctx, submitAccountType, officeId, optional)
List submit accounts

Retrieves a list of submit accounts for a specified type (offices, branchOffices, or windowDelivery).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **submitAccountType** | **string**| The type of submit account | 
  **officeId** | **int32**| Office ID | 
 **optional** | ***SubmitAccountsApiSubmitAccountsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubmitAccountsApiSubmitAccountsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **destinationPincode** | **optional.String**| Destination pincode | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListSubmitAccountsApiResponse**](response.ListSubmitAccountsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitAccountsOfficeIdAmountStatisticsForDashboardGet**
> ResponseAmountStatisticsDashboardApiResponse SubmitAccountsOfficeIdAmountStatisticsForDashboardGet(ctx, officeId, date)
Retrieves amount statistics for the specified office and date

Fetches requested and received amounts, VPP and COD counts, and remitted amounts for the given office and date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
  **date** | **string**| Date | 

### Return type

[**ResponseAmountStatisticsDashboardApiResponse**](response.AmountStatisticsDashboardAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitAccountsOfficeIdStatusGet**
> string SubmitAccountsOfficeIdStatusGet(ctx, officeId, beatName, batchName, submitDate)
Retrieves the submission status of accounts for a specific office, beat, batch, and date

Fetches the account submission status based on the provided office ID, beat name, batch name, and submission date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
  **beatName** | **string**| Beat Name | 
  **batchName** | **string**| Batch Name | 
  **submitDate** | **string**| Submission Date | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitAccountsOfficeIdTotalAmountOfWindowDeliveryGet**
> ResponseFetchSubmitAccountsTotalAmountForWindowDeliveryApiResponse SubmitAccountsOfficeIdTotalAmountOfWindowDeliveryGet(ctx, officeId, optional)
Calculates total amount for window delivery

Calculates the total amount based on the current office ID, created by user, and date provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
 **optional** | ***SubmitAccountsApiSubmitAccountsOfficeIdTotalAmountOfWindowDeliveryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubmitAccountsApiSubmitAccountsOfficeIdTotalAmountOfWindowDeliveryGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdBy** | **optional.String**| User who created the request | 
 **date** | **optional.String**| Date for which total amount is calculated | 

### Return type

[**ResponseFetchSubmitAccountsTotalAmountForWindowDeliveryApiResponse**](response.FetchSubmitAccountsTotalAmountForWindowDeliveryAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitAccountsPost**
> HandlerCombinedResponse SubmitAccountsPost(ctx, body, submitAccountType)
Submits account details for specified office types

This endpoint allows the submission of account data based on the specified office type (offices, branch offices, or window delivery).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainSubmitAccount**](DomainSubmitAccount.md)| Submit Account Request | 
  **submitAccountType** | **string**| Type of the account submission | 

### Return type

[**HandlerCombinedResponse**](handler.CombinedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitAccountsTransactionsAmountCollectedGet**
> ResponseListSubmitAccountsAmountCollectedApiResponse SubmitAccountsTransactionsAmountCollectedGet(ctx, submitAccountOfficeType, date, officeId, employeeId, optional)
Submit account amount collected

Retrieves the amount collected for a specific office type (single hand or others).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **submitAccountOfficeType** | **string**| The type of office submitting the account | 
  **date** | **string**| Date for the collection | 
  **officeId** | **int32**| Office ID | 
  **employeeId** | **int32**| Office ID | 
 **optional** | ***SubmitAccountsApiSubmitAccountsTransactionsAmountCollectedGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubmitAccountsApiSubmitAccountsTransactionsAmountCollectedGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **batchName** | **optional.String**| Batch Name | 
 **beatName** | **optional.String**| Beat Name | 

### Return type

[**ResponseListSubmitAccountsAmountCollectedApiResponse**](response.ListSubmitAccountsAmountCollectedAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

