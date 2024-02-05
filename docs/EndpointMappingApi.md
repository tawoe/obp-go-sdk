# \EndpointMappingApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBankLevelEndpointMapping**](EndpointMappingApi.md#CreateBankLevelEndpointMapping) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/endpoint-mappings | Create Bank Level Endpoint Mapping
[**CreateEndpointMapping**](EndpointMappingApi.md#CreateEndpointMapping) | **Post** /obp/v5.1.0/management/endpoint-mappings | Create Endpoint Mapping
[**DeleteBankLevelEndpointMapping**](EndpointMappingApi.md#DeleteBankLevelEndpointMapping) | **Delete** /obp/v5.1.0/management/banks/{BANK_ID}/endpoint-mappings/ENDPOINT_MAPPING_ID | Delete Bank Level Endpoint Mapping
[**DeleteEndpointMapping**](EndpointMappingApi.md#DeleteEndpointMapping) | **Delete** /obp/v5.1.0/management/endpoint-mappings/ENDPOINT_MAPPING_ID | Delete Endpoint Mapping
[**GetAllBankLevelEndpointMappings**](EndpointMappingApi.md#GetAllBankLevelEndpointMappings) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/endpoint-mappings | Get all Bank Level Endpoint Mappings
[**GetAllEndpointMappings**](EndpointMappingApi.md#GetAllEndpointMappings) | **Get** /obp/v5.1.0/management/endpoint-mappings | Get all Endpoint Mappings
[**GetBankLevelEndpointMapping**](EndpointMappingApi.md#GetBankLevelEndpointMapping) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/endpoint-mappings/ENDPOINT_MAPPING_ID | Get Bank Level Endpoint Mapping
[**GetEndpointMapping**](EndpointMappingApi.md#GetEndpointMapping) | **Get** /obp/v5.1.0/management/endpoint-mappings/ENDPOINT_MAPPING_ID | Get Endpoint Mapping by Id
[**UpdateBankLevelEndpointMapping**](EndpointMappingApi.md#UpdateBankLevelEndpointMapping) | **Put** /obp/v5.1.0/management/banks/{BANK_ID}/endpoint-mappings/ENDPOINT_MAPPING_ID | Update Bank Level Endpoint Mapping
[**UpdateEndpointMapping**](EndpointMappingApi.md#UpdateEndpointMapping) | **Put** /obp/v5.1.0/management/endpoint-mappings/ENDPOINT_MAPPING_ID | Update Endpoint Mapping


# **CreateBankLevelEndpointMapping**
> Body1 CreateBankLevelEndpointMapping(ctx, body, bANKID)
Create Bank Level Endpoint Mapping

<p>Create an Bank Level Endpoint Mapping.</p><p>Note: at moment only support the dynamic endpoints</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body1**](Body1.md)| JObject object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**Body1**](body_1.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEndpointMapping**
> Body4 CreateEndpointMapping(ctx, body)
Create Endpoint Mapping

<p>Create an Endpoint Mapping.</p><p>Note: at moment only support the dynamic endpoints</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body4**](Body4.md)| JObject object that needs to be added. | 

### Return type

[**Body4**](body_4.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBankLevelEndpointMapping**
> bool DeleteBankLevelEndpointMapping(ctx, bANKID)
Delete Bank Level Endpoint Mapping

<p>Delete a Bank Level Endpoint Mapping.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

**bool**

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEndpointMapping**
> bool DeleteEndpointMapping(ctx, )
Delete Endpoint Mapping

<p>Delete a Endpoint Mapping.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

**bool**

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllBankLevelEndpointMappings**
> InlineResponse2007 GetAllBankLevelEndpointMappings(ctx, bANKID)
Get all Bank Level Endpoint Mappings

<p>Get all Bank Level Endpoint Mappings.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllEndpointMappings**
> InlineResponse2007 GetAllEndpointMappings(ctx, )
Get all Endpoint Mappings

<p>Get all Endpoint Mappings.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankLevelEndpointMapping**
> InlineResponse2007Endpointmappings GetBankLevelEndpointMapping(ctx, bANKID)
Get Bank Level Endpoint Mapping

<p>Get an Bank Level Endpoint Mapping by ENDPOINT_MAPPING_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**InlineResponse2007Endpointmappings**](inline_response_200_7_endpointmappings.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEndpointMapping**
> InlineResponse2007Endpointmappings GetEndpointMapping(ctx, )
Get Endpoint Mapping by Id

<p>Get an Endpoint Mapping by ENDPOINT_MAPPING_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2007Endpointmappings**](inline_response_200_7_endpointmappings.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBankLevelEndpointMapping**
> InlineResponse2007Endpointmappings UpdateBankLevelEndpointMapping(ctx, body, bANKID)
Update Bank Level Endpoint Mapping

<p>Update an Bank Level Endpoint Mapping.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body2**](Body2.md)| JObject object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**InlineResponse2007Endpointmappings**](inline_response_200_7_endpointmappings.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEndpointMapping**
> InlineResponse2007Endpointmappings UpdateEndpointMapping(ctx, body)
Update Endpoint Mapping

<p>Update an Endpoint Mapping.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body5**](Body5.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse2007Endpointmappings**](inline_response_200_7_endpointmappings.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

