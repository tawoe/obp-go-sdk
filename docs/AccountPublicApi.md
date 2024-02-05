# \AccountPublicApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublicAccountById**](AccountPublicApi.md#GetPublicAccountById) | **Get** /obp/v5.1.0/banks/{BANK_ID}/public/accounts/{ACCOUNT_ID}/{VIEW_ID}/account | Get Public Account by Id
[**PublicAccountsAllBanks**](AccountPublicApi.md#PublicAccountsAllBanks) | **Get** /obp/v5.1.0/accounts/public | Get Public Accounts at all Banks
[**PublicAccountsAtOneBank**](AccountPublicApi.md#PublicAccountsAtOneBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/public | Get Public Accounts at Bank


# **GetPublicAccountById**
> ModeratedCoreAccountJsonV300 GetPublicAccountById(ctx, vIEWID, aCCOUNTID, bANKID)
Get Public Account by Id

<p>Returns information about an account that has a public view.</p><p>The account is specified by ACCOUNT_ID. The information is moderated by the view specified by VIEW_ID.</p><ul><li>Number</li><li>Owners</li><li>Type</li><li>Balance</li><li>Routing</li></ul><p>PSD2 Context: PSD2 requires customers to have access to their account information via third party applications.<br />This call provides balance and other account information via delegated authentication using OAuth.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ModeratedCoreAccountJsonV300**](ModeratedCoreAccountJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicAccountsAllBanks**
> BasicAccountsJson PublicAccountsAllBanks(ctx, body)
Get Public Accounts at all Banks

<p>Get public accounts at all banks (Anonymous access).<br />Returns accounts that contain at least one public view (a view where is_public is true)<br />For each account the API returns the ID and the available views.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 

### Return type

[**BasicAccountsJson**](BasicAccountsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicAccountsAtOneBank**
> BasicAccountsJson PublicAccountsAtOneBank(ctx, body, bANKID)
Get Public Accounts at Bank

<p>Returns a list of the public accounts (Anonymous access) at BANK_ID. For each account the API returns the ID and the available views.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**BasicAccountsJson**](BasicAccountsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

