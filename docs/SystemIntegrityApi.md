# \SystemIntegrityApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountAccessUniqueIndexCheck**](SystemIntegrityApi.md#AccountAccessUniqueIndexCheck) | **Get** /obp/v5.1.0/management/system/integrity/account-access-unique-index-1-check | Check Unique Index at Account Access
[**AccountCurrencyCheck**](SystemIntegrityApi.md#AccountCurrencyCheck) | **Get** /obp/v5.1.0/management/system/integrity/banks/{BANK_ID}/account-currency-check | Check for Sensible Currencies
[**CustomViewNamesCheck**](SystemIntegrityApi.md#CustomViewNamesCheck) | **Get** /obp/v5.1.0/management/system/integrity/custom-view-names-check | Check Custom View Names
[**OrphanedAccountCheck**](SystemIntegrityApi.md#OrphanedAccountCheck) | **Get** /obp/v5.1.0/management/system/integrity/banks/{BANK_ID}/orphaned-account-check | Check for Orphaned Accounts
[**SystemViewNamesCheck**](SystemIntegrityApi.md#SystemViewNamesCheck) | **Get** /obp/v5.1.0/management/system/integrity/system-view-names-check | Check System View Names


# **AccountAccessUniqueIndexCheck**
> CheckSystemIntegrityJsonV510 AccountAccessUniqueIndexCheck(ctx, )
Check Unique Index at Account Access

<p>Check unique index at account access table.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CheckSystemIntegrityJsonV510**](CheckSystemIntegrityJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountCurrencyCheck**
> CheckSystemIntegrityJsonV510 AccountCurrencyCheck(ctx, bANKID)
Check for Sensible Currencies

<p>Check for sensible currencies at Bank Account model</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**CheckSystemIntegrityJsonV510**](CheckSystemIntegrityJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomViewNamesCheck**
> CheckSystemIntegrityJsonV510 CustomViewNamesCheck(ctx, )
Check Custom View Names

<p>Check custom view names.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CheckSystemIntegrityJsonV510**](CheckSystemIntegrityJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrphanedAccountCheck**
> CheckSystemIntegrityJsonV510 OrphanedAccountCheck(ctx, bANKID)
Check for Orphaned Accounts

<p>Check for orphaned accounts at Bank Account model</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**CheckSystemIntegrityJsonV510**](CheckSystemIntegrityJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemViewNamesCheck**
> CheckSystemIntegrityJsonV510 SystemViewNamesCheck(ctx, )
Check System View Names

<p>Check system view names.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CheckSystemIntegrityJsonV510**](CheckSystemIntegrityJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

