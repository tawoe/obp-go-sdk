# \ConfirmationOfFundsServicePIISApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckFundsAvailable**](ConfirmationOfFundsServicePIISApi.md#CheckFundsAvailable) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/funds-available | Check Available Funds


# **CheckFundsAvailable**
> CheckFundsAvailableJson CheckFundsAvailable(ctx, vIEWID, aCCOUNTID, bANKID)
Check Available Funds

<p>Check Available Funds<br />Mandatory URL parameters:</p><ul><li>amount=NUMBER</li><li>currency=STRING</li></ul><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CheckFundsAvailableJson**](CheckFundsAvailableJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

