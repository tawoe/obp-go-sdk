# \ConnectorMethodApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectorMethod**](ConnectorMethodApi.md#CreateConnectorMethod) | **Post** /obp/v5.1.0/management/connector-methods | Create Connector Method
[**GetAllConnectorMethods**](ConnectorMethodApi.md#GetAllConnectorMethods) | **Get** /obp/v5.1.0/management/connector-methods | Get all Connector Methods
[**GetConnectorMethod**](ConnectorMethodApi.md#GetConnectorMethod) | **Get** /obp/v5.1.0/management/connector-methods/CONNECTOR_METHOD_ID | Get Connector Method by Id
[**UpdateConnectorMethod**](ConnectorMethodApi.md#UpdateConnectorMethod) | **Put** /obp/v5.1.0/management/connector-methods/CONNECTOR_METHOD_ID | Update Connector Method


# **CreateConnectorMethod**
> JsonConnectorMethod CreateConnectorMethod(ctx, body)
Create Connector Method

<p>Create an internal connector.</p><p>The method_body is URL-encoded format String</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonConnectorMethod**](JsonConnectorMethod.md)| JsonConnectorMethod object that needs to be added. | 

### Return type

[**JsonConnectorMethod**](JsonConnectorMethod.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllConnectorMethods**
> InlineResponse2008 GetAllConnectorMethods(ctx, )
Get all Connector Methods

<p>Get all Connector Methods.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorMethod**
> JsonConnectorMethod GetConnectorMethod(ctx, )
Get Connector Method by Id

<p>Get an internal connector by CONNECTOR_METHOD_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JsonConnectorMethod**](JsonConnectorMethod.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnectorMethod**
> JsonConnectorMethod UpdateConnectorMethod(ctx, body)
Update Connector Method

<p>Update an internal connector.</p><p>The method_body is URL-encoded format String</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonConnectorMethodMethodBody**](JsonConnectorMethodMethodBody.md)| JsonConnectorMethodMethodBody object that needs to be added. | 

### Return type

[**JsonConnectorMethod**](JsonConnectorMethod.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

