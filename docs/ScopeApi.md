# \ScopeApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScope**](ScopeApi.md#AddScope) | **Post** /obp/v5.1.0/consumers/{CONSUMER_ID}/scopes | Create Scope for a Consumer
[**DeleteScope**](ScopeApi.md#DeleteScope) | **Delete** /obp/v5.1.0/consumers/{CONSUMER_ID}/scope/{SCOPE_ID} | Delete Consumer Scope
[**GetScopes**](ScopeApi.md#GetScopes) | **Get** /obp/v5.1.0/consumers/{CONSUMER_ID}/scopes | Get Scopes for Consumer


# **AddScope**
> ScopeJson AddScope(ctx, body, cONSUMERID)
Create Scope for a Consumer

<p>Create Scope. Grant Role to Consumer.</p><p>Scopes are used to grant System or Bank level roles to the Consumer (App). (For Account level privileges, see Views)</p><p>For a System level Role (.e.g CanGetAnyUser), set bank_id to an empty string i.e. &quot;bank_id&quot;:&quot;&quot;</p><p>For a Bank level Role (e.g. CanCreateAccount), set bank_id to a valid value e.g. &quot;bank_id&quot;:&quot;my-bank-id&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateScopeJson**](CreateScopeJson.md)| CreateScopeJson object that needs to be added. | 
  **cONSUMERID** | **string**| new consumer id | 

### Return type

[**ScopeJson**](ScopeJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScope**
> DeleteScope(ctx, sCOPEID, cONSUMERID)
Delete Consumer Scope

<p>Delete Consumer Scope specified by SCOPE_ID for an consumer specified by CONSUMER_ID</p><p>Authentication is required and the user needs to be a Super Admin.<br />Super Admins are listed in the Props file.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sCOPEID** | **string**| the scope id | 
  **cONSUMERID** | **string**| new consumer id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScopes**
> ScopeJsons GetScopes(ctx, cONSUMERID)
Get Scopes for Consumer

<p>Get all the scopes for an consumer specified by CONSUMER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cONSUMERID** | **string**| new consumer id | 

### Return type

[**ScopeJsons**](ScopeJsons.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

