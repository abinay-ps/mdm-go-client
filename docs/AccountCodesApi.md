# \AccountCodesApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlApproveGLPut**](AccountCodesApi.md#GlApproveGLPut) | **Put** /gl/approveGL | Approve Account Codes
[**GlCreateGLCodePost**](AccountCodesApi.md#GlCreateGLCodePost) | **Post** /gl/createGLCode/ | Create Account Code
[**GlGetActiveGLdetailsbyGLCodeAccountCodeGet**](AccountCodesApi.md#GlGetActiveGLdetailsbyGLCodeAccountCodeGet) | **Get** /gl/getActiveGLdetailsbyGLCode/{account_code} | Get Active Account Code details by Account Code
[**GlGetGLDetailsbyHOAHoaGet**](AccountCodesApi.md#GlGetGLDetailsbyHOAHoaGet) | **Get** /gl/getGLDetailsbyHOA/{hoa} | Get Account Code details by HOA
[**GlGetGLdetailsbyGLCodeAccountCodeGet**](AccountCodesApi.md#GlGetGLdetailsbyGLCodeAccountCodeGet) | **Get** /gl/getGLdetailsbyGLCode/{account_code} | Get Account Code details by Account Code
[**GlGetHOADetailsHoacodeGet**](AccountCodesApi.md#GlGetHOADetailsHoacodeGet) | **Get** /gl/getHOADetails/{hoacode} | Get HOA Details by HOA Code
[**GlGetallAccountCodesGet**](AccountCodesApi.md#GlGetallAccountCodesGet) | **Get** /gl/getallAccountCodes | Get All Account Code Details
[**GlGetallHOADetailsGet**](AccountCodesApi.md#GlGetallHOADetailsGet) | **Get** /gl/getallHOADetails | Get All HOA Details
[**GlGlverificationGet**](AccountCodesApi.md#GlGlverificationGet) | **Get** /gl/glverification | Get Account Code Verify
[**GlModifyGLActiveDeactivePut**](AccountCodesApi.md#GlModifyGLActiveDeactivePut) | **Put** /gl/modifyGLActiveDeactive | Modify an Account Code Activate/Deactivate
[**GlModifyGLchangeGLNamePut**](AccountCodesApi.md#GlModifyGLchangeGLNamePut) | **Put** /gl/modifyGLchangeGLName/ | Modify Account code Change Account Code Name
[**GlModifyGLchangeHOAPut**](AccountCodesApi.md#GlModifyGLchangeHOAPut) | **Put** /gl/modifyGLchangeHOA/ | Modify Account code change HOA


# **GlApproveGLPut**
> HandlerResponse GlApproveGLPut(ctx, verifyAccountCodeDetails)
Approve Account Codes

Approve Account Code Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verifyAccountCodeDetails** | [**DomainVerifyAccountCodeDetails**](DomainVerifyAccountCodeDetails.md)| Account Code Verification Request  | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlCreateGLCodePost**
> HandlerResponse GlCreateGLCodePost(ctx, createAccountCodeDetails)
Create Account Code

Create a new Account Code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createAccountCodeDetails** | [**DomainCreateAccountCodeDetails**](DomainCreateAccountCodeDetails.md)| Account Code Creation Request  | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlGetActiveGLdetailsbyGLCodeAccountCodeGet**
> HandlerResponse GlGetActiveGLdetailsbyGLCodeAccountCodeGet(ctx, accountCode)
Get Active Account Code details by Account Code

Fetch Active Account code details by Account code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountCode** | **int32**| Account Code | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlGetGLDetailsbyHOAHoaGet**
> HandlerResponse GlGetGLDetailsbyHOAHoaGet(ctx, hoa)
Get Account Code details by HOA

Fetch Account code details by HOA

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hoa** | **string**| HOA | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlGetGLdetailsbyGLCodeAccountCodeGet**
> HandlerResponse GlGetGLdetailsbyGLCodeAccountCodeGet(ctx, accountCode)
Get Account Code details by Account Code

Fetch Account code details by Account code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountCode** | **int32**| Account Code | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlGetHOADetailsHoacodeGet**
> HandlerResponse GlGetHOADetailsHoacodeGet(ctx, hoacode)
Get HOA Details by HOA Code

Gives Head of Account Details by HOA Code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hoacode** | **string**| HOA Code | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlGetallAccountCodesGet**
> HandlerResponse GlGetallAccountCodesGet(ctx, )
Get All Account Code Details

Fetches all Account Codes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlGetallHOADetailsGet**
> HandlerResponse GlGetallHOADetailsGet(ctx, )
Get All HOA Details

Fetches all HOA Codes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlGlverificationGet**
> HandlerResponse GlGlverificationGet(ctx, optional)
Get Account Code Verify

Fetch Account codes for verification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountCodesApiGlGlverificationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountCodesApiGlGlverificationGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glStatus** | **optional.String**| GL Status | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlModifyGLActiveDeactivePut**
> GlModifyGLActiveDeactivePut(ctx, modifyAccountCodeActiveDeactive)
Modify an Account Code Activate/Deactivate

Enable or Disable an Account Code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modifyAccountCodeActiveDeactive** | [**DomainModifyAccountCodeActiveDeactive**](DomainModifyAccountCodeActiveDeactive.md)| Modify Account Code modify request  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlModifyGLchangeGLNamePut**
> HandlerResponse GlModifyGLchangeGLNamePut(ctx, modifyAccountCodeName)
Modify Account code Change Account Code Name

Modify Account code name using Account code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modifyAccountCodeName** | [**DomainModifyAccountCodeName**](DomainModifyAccountCodeName.md)| Account Code HOA Modification Request | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlModifyGLchangeHOAPut**
> HandlerResponse GlModifyGLchangeHOAPut(ctx, modifyAccountCodeHOA)
Modify Account code change HOA

Modify Account code to change HOA

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modifyAccountCodeHOA** | [**DomainModifyAccountCodeHoa**](DomainModifyAccountCodeHoa.md)| Account Code HOA Modification Request | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

