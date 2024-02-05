# \DirectDebitApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDirectDebit**](DirectDebitApi.md#CreateDirectDebit) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/direct-debit | Create Direct Debit
[**CreateDirectDebitManagement**](DirectDebitApi.md#CreateDirectDebitManagement) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/direct-debit | Create Direct Debit (management)


# **CreateDirectDebit**
> DirectDebitJsonV400 CreateDirectDebit(ctx, body, vIEWID, aCCOUNTID, bANKID)
Create Direct Debit

<p>Create direct debit for an account.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostDirectDebitJsonV400**](PostDirectDebitJsonV400.md)| PostDirectDebitJsonV400 object that needs to be added. | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**DirectDebitJsonV400**](DirectDebitJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDirectDebitManagement**
> DirectDebitJsonV400 CreateDirectDebitManagement(ctx, body, aCCOUNTID, bANKID)
Create Direct Debit (management)

<p>Create direct debit for an account.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostDirectDebitJsonV400**](PostDirectDebitJsonV400.md)| PostDirectDebitJsonV400 object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**DirectDebitJsonV400**](DirectDebitJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

