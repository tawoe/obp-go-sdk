# \AccountApplicationApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountApplication**](AccountApplicationApi.md#CreateAccountApplication) | **Post** /obp/v5.1.0/banks/{BANK_ID}/account-applications | Create Account Application
[**GetAccountApplication**](AccountApplicationApi.md#GetAccountApplication) | **Get** /obp/v5.1.0/banks/{BANK_ID}/account-applications/{ACCOUNT_APPLICATION_ID} | Get Account Application by Id
[**GetAccountApplications**](AccountApplicationApi.md#GetAccountApplications) | **Get** /obp/v5.1.0/banks/{BANK_ID}/account-applications | Get Account Applications
[**UpdateAccountApplicationStatus**](AccountApplicationApi.md#UpdateAccountApplicationStatus) | **Put** /obp/v5.1.0/banks/{BANK_ID}/account-applications/{ACCOUNT_APPLICATION_ID} | Update Account Application Status


# **CreateAccountApplication**
> AccountApplicationResponseJson CreateAccountApplication(ctx, body, bANKID)
Create Account Application

<p>Create Account Application</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountApplicationJson**](AccountApplicationJson.md)| AccountApplicationJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountApplicationResponseJson**](AccountApplicationResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountApplication**
> AccountApplicationResponseJson GetAccountApplication(ctx, aCCOUNTAPPLICATIONID, bANKID)
Get Account Application by Id

<p>Get the Account Application.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aCCOUNTAPPLICATIONID** | **string**| the account application id  | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountApplicationResponseJson**](AccountApplicationResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountApplications**
> AccountApplicationsJsonV310 GetAccountApplications(ctx, bANKID)
Get Account Applications

<p>Get the Account Applications.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountApplicationsJsonV310**](AccountApplicationsJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountApplicationStatus**
> AccountApplicationResponseJson UpdateAccountApplicationStatus(ctx, body, aCCOUNTAPPLICATIONID, bANKID)
Update Account Application Status

<p>Update an Account Application status</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountApplicationUpdateStatusJson**](AccountApplicationUpdateStatusJson.md)| AccountApplicationUpdateStatusJson object that needs to be added. | 
  **aCCOUNTAPPLICATIONID** | **string**| the account application id  | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountApplicationResponseJson**](AccountApplicationResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

