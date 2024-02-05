# \PublicDataApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrivateAccountsAtOneBank**](PublicDataApi.md#GetPrivateAccountsAtOneBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts | Get Accounts at Bank
[**PublicAccountsAllBanks**](PublicDataApi.md#PublicAccountsAllBanks) | **Get** /obp/v5.1.0/accounts/public | Get Public Accounts at all Banks
[**PublicAccountsAtOneBank**](PublicDataApi.md#PublicAccountsAtOneBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/public | Get Public Accounts at Bank


# **GetPrivateAccountsAtOneBank**
> BasicAccountsJson GetPrivateAccountsAtOneBank(ctx, bANKID)
Get Accounts at Bank

<p>Returns the list of accounts at BANK_ID that the user has access to.<br />For each account the API returns the account ID and the views available to the user..<br />Each account must have at least one private View.</p><p>optional request parameters for filter with attributes<br />URL params example: /banks/some-bank-id/accounts?manager=John&amp;count=8</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**BasicAccountsJson**](BasicAccountsJSON.md)

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

