# \DynamicEndpointManageApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBankLevelDynamicEndpoint**](DynamicEndpointManageApi.md#CreateBankLevelDynamicEndpoint) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-endpoints | Create Bank Level Dynamic Endpoint
[**CreateDynamicEndpoint**](DynamicEndpointManageApi.md#CreateDynamicEndpoint) | **Post** /obp/v5.1.0/management/dynamic-endpoints | Create Dynamic Endpoint
[**DeleteBankLevelDynamicEndpoint**](DynamicEndpointManageApi.md#DeleteBankLevelDynamicEndpoint) | **Delete** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-endpoints/DYNAMIC_ENDPOINT_ID |  Delete Bank Level Dynamic Endpoint
[**DeleteDynamicEndpoint**](DynamicEndpointManageApi.md#DeleteDynamicEndpoint) | **Delete** /obp/v5.1.0/management/dynamic-endpoints/DYNAMIC_ENDPOINT_ID |  Delete Dynamic Endpoint
[**DeleteMyDynamicEndpoint**](DynamicEndpointManageApi.md#DeleteMyDynamicEndpoint) | **Delete** /obp/v5.1.0/my/dynamic-endpoints/DYNAMIC_ENDPOINT_ID | Delete My Dynamic Endpoint
[**GetBankLevelDynamicEndpoint**](DynamicEndpointManageApi.md#GetBankLevelDynamicEndpoint) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-endpoints/DYNAMIC_ENDPOINT_ID |  Get Bank Level Dynamic Endpoint
[**GetBankLevelDynamicEndpoints**](DynamicEndpointManageApi.md#GetBankLevelDynamicEndpoints) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-endpoints | Get Bank Level Dynamic Endpoints
[**GetDynamicEndpoint**](DynamicEndpointManageApi.md#GetDynamicEndpoint) | **Get** /obp/v5.1.0/management/dynamic-endpoints/DYNAMIC_ENDPOINT_ID | Get Dynamic Endpoint
[**GetDynamicEndpoints**](DynamicEndpointManageApi.md#GetDynamicEndpoints) | **Get** /obp/v5.1.0/management/dynamic-endpoints |  Get Dynamic Endpoints
[**GetMyDynamicEndpoints**](DynamicEndpointManageApi.md#GetMyDynamicEndpoints) | **Get** /obp/v5.1.0/my/dynamic-endpoints | Get My Dynamic Endpoints
[**UpdateBankLevelDynamicEndpointHost**](DynamicEndpointManageApi.md#UpdateBankLevelDynamicEndpointHost) | **Put** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-endpoints/DYNAMIC_ENDPOINT_ID/host |  Update Bank Level Dynamic Endpoint Host
[**UpdateDynamicEndpointHost**](DynamicEndpointManageApi.md#UpdateDynamicEndpointHost) | **Put** /obp/v5.1.0/management/dynamic-endpoints/DYNAMIC_ENDPOINT_ID/host |  Update Dynamic Endpoint Host


# **CreateBankLevelDynamicEndpoint**
> InlineResponse201 CreateBankLevelDynamicEndpoint(ctx, body, bANKID)
Create Bank Level Dynamic Endpoint

<p>Create dynamic endpoints.</p><p>Create dynamic endpoints with one json format swagger content.</p><p>If the host of swagger is <code>dynamic_entity</code>, then you need link the swagger fields to the dynamic entity fields,<br />please check <code>Endpoint Mapping</code> endpoints.</p><p>If the host of swagger is <code>obp_mock</code>, every dynamic endpoint will return example response of swagger,</p><p>when create MethodRouting for given dynamic endpoint, it will be routed to given url.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body**](Body.md)| JObject object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDynamicEndpoint**
> InlineResponse2011 CreateDynamicEndpoint(ctx, body)
Create Dynamic Endpoint

<p>Create dynamic endpoints.</p><p>Create dynamic endpoints with one json format swagger content.</p><p>If the host of swagger is <code>dynamic_entity</code>, then you need link the swagger fields to the dynamic entity fields,<br />please check <code>Endpoint Mapping</code> endpoints.</p><p>If the host of swagger is <code>obp_mock</code>, every dynamic endpoint will return example response of swagger,</p><p>when create MethodRouting for given dynamic endpoint, it will be routed to given url.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body3**](Body3.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBankLevelDynamicEndpoint**
> DeleteBankLevelDynamicEndpoint(ctx, bANKID)
 Delete Bank Level Dynamic Endpoint

<p>Delete a Bank Level DynamicEndpoint specified by DYNAMIC_ENDPOINT_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDynamicEndpoint**
> DeleteDynamicEndpoint(ctx, )
 Delete Dynamic Endpoint

<p>Delete a DynamicEndpoint specified by DYNAMIC_ENDPOINT_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMyDynamicEndpoint**
> DeleteMyDynamicEndpoint(ctx, )
Delete My Dynamic Endpoint

<p>Delete a DynamicEndpoint specified by DYNAMIC_ENDPOINT_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankLevelDynamicEndpoint**
> InlineResponse201 GetBankLevelDynamicEndpoint(ctx, bANKID)
 Get Bank Level Dynamic Endpoint

<p>Get a Bank Level Dynamic Endpoint.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankLevelDynamicEndpoints**
> InlineResponse2003 GetBankLevelDynamicEndpoints(ctx, bANKID)
Get Bank Level Dynamic Endpoints

<p>Get Bank Level Dynamic Endpoints.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDynamicEndpoint**
> InlineResponse2011 GetDynamicEndpoint(ctx, )
Get Dynamic Endpoint

<p>Get a Dynamic Endpoint.</p><p>Get one DynamicEndpoint,</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDynamicEndpoints**
> InlineResponse2009 GetDynamicEndpoints(ctx, )
 Get Dynamic Endpoints

<p>Get Dynamic Endpoints.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyDynamicEndpoints**
> InlineResponse20012 GetMyDynamicEndpoints(ctx, )
Get My Dynamic Endpoints

<p>Get My Dynamic Endpoints.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBankLevelDynamicEndpointHost**
> DynamicEndpointHostJson400 UpdateBankLevelDynamicEndpointHost(ctx, body, bANKID)
 Update Bank Level Dynamic Endpoint Host

<p>Update Bank Level  dynamic endpoint Host.<br />The value can be obp_mock, dynamic_entity, or some service url.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DynamicEndpointHostJson400**](DynamicEndpointHostJson400.md)| DynamicEndpointHostJson400 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**DynamicEndpointHostJson400**](DynamicEndpointHostJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDynamicEndpointHost**
> DynamicEndpointHostJson400 UpdateDynamicEndpointHost(ctx, body)
 Update Dynamic Endpoint Host

<p>Update dynamic endpoint Host.<br />The value can be obp_mock, dynamic_entity, or some service url.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DynamicEndpointHostJson400**](DynamicEndpointHostJson400.md)| DynamicEndpointHostJson400 object that needs to be added. | 

### Return type

[**DynamicEndpointHostJson400**](DynamicEndpointHostJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

