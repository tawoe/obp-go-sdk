# \CustomerMessageApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomerMessage**](CustomerMessageApi.md#AddCustomerMessage) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customer/{CUSTOMER_ID}/messages | Create Customer Message
[**CreateCustomerMessage**](CustomerMessageApi.md#CreateCustomerMessage) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/messages | Create Customer Message
[**GetCustomerMessages**](CustomerMessageApi.md#GetCustomerMessages) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/messages | Get Customer Messages for a Customer
[**GetCustomersMessages**](CustomerMessageApi.md#GetCustomersMessages) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customer/messages | Get Customer Messages for all Customers


# **AddCustomerMessage**
> SuccessMessage AddCustomerMessage(ctx, body, cUSTOMERID, bANKID)
Create Customer Message

<p>Create a message for the customer specified by CUSTOMER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddCustomerMessageJson**](AddCustomerMessageJson.md)| AddCustomerMessageJson object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomerMessage**
> SuccessMessage CreateCustomerMessage(ctx, body, cUSTOMERID, bANKID)
Create Customer Message

<p>Create a message for the customer specified by CUSTOMER_ID<br />Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMessageJsonV400**](CreateMessageJsonV400.md)| CreateMessageJsonV400 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerMessages**
> CustomerMessagesJsonV400 GetCustomerMessages(ctx, cUSTOMERID, bANKID)
Get Customer Messages for a Customer

<p>Get messages for the customer specified by CUSTOMER_ID<br />Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerMessagesJsonV400**](CustomerMessagesJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersMessages**
> CustomerMessagesJson GetCustomersMessages(ctx, body, bANKID)
Get Customer Messages for all Customers

<p>Get messages for the logged in customer<br />Messages sent to the currently authenticated user.</p><p>Authentication via OAuth is required.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerMessagesJson**](CustomerMessagesJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

