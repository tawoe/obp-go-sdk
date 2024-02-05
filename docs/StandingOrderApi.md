# \StandingOrderApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStandingOrder**](StandingOrderApi.md#CreateStandingOrder) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/standing-order | Create Standing Order
[**CreateStandingOrderManagement**](StandingOrderApi.md#CreateStandingOrderManagement) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/standing-order | Create Standing Order (management)


# **CreateStandingOrder**
> StandingOrderJsonV400 CreateStandingOrder(ctx, body, vIEWID, aCCOUNTID, bANKID)
Create Standing Order

<p>Create standing order for an account.</p><p>when -&gt; frequency = {‘YEARLY’,’MONTHLY, ‘WEEKLY’, ‘BI-WEEKLY’, DAILY’}<br />when -&gt; detail = { ‘FIRST_MONDAY’, ‘FIRST_DAY’, ‘LAST_DAY’}}</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostStandingOrderJsonV400**](PostStandingOrderJsonV400.md)| PostStandingOrderJsonV400 object that needs to be added. | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**StandingOrderJsonV400**](StandingOrderJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStandingOrderManagement**
> StandingOrderJsonV400 CreateStandingOrderManagement(ctx, body, aCCOUNTID, bANKID)
Create Standing Order (management)

<p>Create standing order for an account.</p><p>when -&gt; frequency = {‘YEARLY’,’MONTHLY, ‘WEEKLY’, ‘BI-WEEKLY’, DAILY’}<br />when -&gt; detail = { ‘FIRST_MONDAY’, ‘FIRST_DAY’, ‘LAST_DAY’}}</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostStandingOrderJsonV400**](PostStandingOrderJsonV400.md)| PostStandingOrderJsonV400 object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**StandingOrderJsonV400**](StandingOrderJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

