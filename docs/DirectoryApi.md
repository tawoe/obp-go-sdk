# \DirectoryApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConsumer**](DirectoryApi.md#CreateConsumer) | **Post** /obp/v5.1.0/dynamic-registration/consumers | Create a Consumer
[**CreateRegulatedEntity**](DirectoryApi.md#CreateRegulatedEntity) | **Post** /obp/v5.1.0/regulated-entities | Create Regulated Entity
[**DeleteRegulatedEntity**](DirectoryApi.md#DeleteRegulatedEntity) | **Delete** /obp/v5.1.0/regulated-entities/REGULATED_ENTITY_ID | Delete Regulated Entity
[**GetRegulatedEntityById**](DirectoryApi.md#GetRegulatedEntityById) | **Get** /obp/v5.1.0/regulated-entities/REGULATED_ENTITY_ID | Get Regulated Entity
[**RegulatedEntities**](DirectoryApi.md#RegulatedEntities) | **Get** /obp/v5.1.0/regulated-entities | Get Regulated Entities


# **CreateConsumer**
> ConsumerJsonV510 CreateConsumer(ctx, body)
Create a Consumer

<p>Create a Consumer (mTLS access).</p><p>JWT payload:<br />- minimal<br />{ &quot;description&quot;:&quot;Description&quot; }<br />- full<br />{<br />&quot;description&quot;: &quot;Description&quot;,<br />&quot;app_name&quot;: &quot;Tesobe GmbH&quot;,<br />&quot;app_type&quot;: &quot;Sofit&quot;,<br />&quot;developer_email&quot;: &quot;<a href=\"m&#97;i&#x6c;&#x74;&#111;&#x3a;&#109;&#x61;&#x72;&#x6b;&#111;@&#116;&#x65;&#115;&#x6f;&#98;e&#x2e;&#x63;o&#109;\">&#109;&#x61;&#114;ko&#x40;te&#x73;&#x6f;&#98;&#101;.&#99;&#111;m</a>&quot;,<br />&quot;redirect_url&quot;: &quot;<a href=\"http://localhost:8082\">http://localhost:8082</a>&quot;<br />}<br />Please note that JWT must be signed with the counterpart private kew of the public key used to establish mTLS</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConsumerJwtPostJsonV510**](ConsumerJwtPostJsonV510.md)| ConsumerJwtPostJsonV510 object that needs to be added. | 

### Return type

[**ConsumerJsonV510**](ConsumerJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRegulatedEntity**
> RegulatedEntitiesJsonV510 CreateRegulatedEntity(ctx, body)
Create Regulated Entity

<p>Create Regulated Entity</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegulatedEntityPostJsonV510**](RegulatedEntityPostJsonV510.md)| RegulatedEntityPostJsonV510 object that needs to be added. | 

### Return type

[**RegulatedEntitiesJsonV510**](RegulatedEntitiesJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRegulatedEntity**
> DeleteRegulatedEntity(ctx, )
Delete Regulated Entity

<p>Delete Regulated Entity specified by REGULATED_ENTITY_ID</p><p>Authentication is Mandatory</p>

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

# **GetRegulatedEntityById**
> RegulatedEntitiesJsonV510 GetRegulatedEntityById(ctx, )
Get Regulated Entity

<p>Get Regulated Entity By REGULATED_ENTITY_ID</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RegulatedEntitiesJsonV510**](RegulatedEntitiesJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegulatedEntities**
> RegulatedEntitiesJsonV510 RegulatedEntities(ctx, )
Get Regulated Entities

<p>Returns information about:</p><ul><li>Regulated Entities</li></ul><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RegulatedEntitiesJsonV510**](RegulatedEntitiesJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

