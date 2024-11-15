# \EstablishmentApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EstablishmentmasterDisplayOfficeListGet**](EstablishmentApi.md#EstablishmentmasterDisplayOfficeListGet) | **Get** /establishmentmaster/DisplayOfficeList | Get List Of Offices Within Establishment Register
[**EstablishmentmasterOfficeIDGet**](EstablishmentApi.md#EstablishmentmasterOfficeIDGet) | **Get** /establishmentmaster/OfficeID/ | Get Establishment Register Details by Office ID
[**EstablishmentmasterOfficeNameOfficenameGet**](EstablishmentApi.md#EstablishmentmasterOfficeNameOfficenameGet) | **Get** /establishmentmaster/OfficeName/{officename} | Get Establishment Register Details By Office Name
[**EstablishmentmasterPost**](EstablishmentApi.md#EstablishmentmasterPost) | **Post** /establishmentmaster/ | Create a new Establishment Register
[**EstablishmentmasterRegisterIDEridGet**](EstablishmentApi.md#EstablishmentmasterRegisterIDEridGet) | **Get** /establishmentmaster/RegisterID/{erid} | Get Establishment Register Details by ER ID
[**EstablishmentmasterRegisterNameErnameGet**](EstablishmentApi.md#EstablishmentmasterRegisterNameErnameGet) | **Get** /establishmentmaster/RegisterName/{ername} | Get Establishment Register Details By ER Name
[**EstablishmentmasterUpdateEridPut**](EstablishmentApi.md#EstablishmentmasterUpdateEridPut) | **Put** /establishmentmaster/update/{erid} | Update an Establishment Register
[**EstablishmentmasterVerificationGet**](EstablishmentApi.md#EstablishmentmasterVerificationGet) | **Get** /establishmentmaster/verification | Get all new entries of Establishment Register pending for Verification
[**EstablishmentmasterVerificationrecordPut**](EstablishmentApi.md#EstablishmentmasterVerificationrecordPut) | **Put** /establishmentmaster/verificationrecord | Verify a  New or Existing Establishment Register


# **EstablishmentmasterDisplayOfficeListGet**
> HandlerResponse EstablishmentmasterDisplayOfficeListGet(ctx, optional)
Get List Of Offices Within Establishment Register

Fetches List Of Offices Within an Establishment Register

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EstablishmentApiEstablishmentmasterDisplayOfficeListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EstablishmentApiEstablishmentmasterDisplayOfficeListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **circleName** | **optional.String**| Circle Name | 
 **regionName** | **optional.String**| Region Name | 
 **divisionName** | **optional.String**| Division Name | 
 **subdivisionName** | **optional.String**| Sub Division Name | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstablishmentmasterOfficeIDGet**
> HandlerResponse EstablishmentmasterOfficeIDGet(ctx, optional)
Get Establishment Register Details by Office ID

Fetches Establishment Register Details by Office ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EstablishmentApiEstablishmentmasterOfficeIDGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EstablishmentApiEstablishmentmasterOfficeIDGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **optional.Int32**| Office ID | 
 **userOfficeId** | **optional.Int32**| User Office ID | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstablishmentmasterOfficeNameOfficenameGet**
> HandlerResponse EstablishmentmasterOfficeNameOfficenameGet(ctx, officename)
Get Establishment Register Details By Office Name

Fetches Establishment Register Details by Office Name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officename** | **string**| Office Name | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstablishmentmasterPost**
> HandlerResponse EstablishmentmasterPost(ctx, request)
Create a new Establishment Register

Create a new Establishment Register for Office By giving Required Details ER_Name,Office_Name,Office_ID,OrderMemoNo,Remarks,Supporting_Doc_Path,Total_Sanctioned_Strength

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**HandlerCreateEstablishmentRegisterRequest**](HandlerCreateEstablishmentRegisterRequest.md)| Establishment Register Input Field  | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstablishmentmasterRegisterIDEridGet**
> HandlerResponse EstablishmentmasterRegisterIDEridGet(ctx, erid)
Get Establishment Register Details by ER ID

Fetches Establishment Register details by giving ER ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **erid** | **string**| ER ID | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstablishmentmasterRegisterNameErnameGet**
> HandlerResponse EstablishmentmasterRegisterNameErnameGet(ctx, ername)
Get Establishment Register Details By ER Name

Fetches Establishment Register Details by giving Establishment Register Name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ername** | **string**| ER Name | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstablishmentmasterUpdateEridPut**
> HandlerResponse EstablishmentmasterUpdateEridPut(ctx, erid)
Update an Establishment Register

Updates an existing Establishment Register of an Office By giving Required Details ER_Name,Office_Name,Office_ID,OrderMemoNo,Remarks,Supporting_Doc_Path,Total_Sanctioned_Strength

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **erid** | **string**| ER ID | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstablishmentmasterVerificationGet**
> HandlerResponse EstablishmentmasterVerificationGet(ctx, status)
Get all new entries of Establishment Register pending for Verification

Lists all new entries of Establishment Register that are Pending for Verification by giving status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **status** | **string**| Status  | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstablishmentmasterVerificationrecordPut**
> HandlerResponse EstablishmentmasterVerificationrecordPut(ctx, erID)
Verify a  New or Existing Establishment Register

Verify a  New or Existing Establishment Register By Giving Erid and status at query parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **erID** | **string**| ER ID | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

