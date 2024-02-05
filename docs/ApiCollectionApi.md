# \ApiCollectionApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMyApiCollection**](ApiCollectionApi.md#CreateMyApiCollection) | **Post** /obp/v5.1.0/my/api-collections | Create My Api Collection
[**CreateMyApiCollectionEndpoint**](ApiCollectionApi.md#CreateMyApiCollectionEndpoint) | **Post** /obp/v5.1.0/my/api-collections/API_COLLECTION_NAME/api-collection-endpoints | Create My Api Collection Endpoint
[**CreateMyApiCollectionEndpointById**](ApiCollectionApi.md#CreateMyApiCollectionEndpointById) | **Post** /obp/v5.1.0/my/api-collection-ids/API_COLLECTION_ID/api-collection-endpoints | Create My Api Collection Endpoint By Id
[**DeleteMyApiCollection**](ApiCollectionApi.md#DeleteMyApiCollection) | **Delete** /obp/v5.1.0/my/api-collections/API_COLLECTION_ID | Delete My Api Collection
[**DeleteMyApiCollectionEndpoint**](ApiCollectionApi.md#DeleteMyApiCollectionEndpoint) | **Delete** /obp/v5.1.0/my/api-collections/API_COLLECTION_NAME/api-collection-endpoints/OPERATION_ID | Delete My Api Collection Endpoint
[**DeleteMyApiCollectionEndpointById**](ApiCollectionApi.md#DeleteMyApiCollectionEndpointById) | **Delete** /obp/v5.1.0/my/api-collection-ids/API_COLLECTION_ID/api-collection-endpoint-ids/API_COLLECTION_ENDPOINT_ID | Delete My Api Collection Endpoint By Id
[**DeleteMyApiCollectionEndpointByOperationId**](ApiCollectionApi.md#DeleteMyApiCollectionEndpointByOperationId) | **Delete** /obp/v5.1.0/my/api-collection-ids/API_COLLECTION_ID/api-collection-endpoints/OPERATION_ID | Delete My Api Collection Endpoint By Id
[**GetAllApiCollections**](ApiCollectionApi.md#GetAllApiCollections) | **Get** /obp/v5.1.0/management/api-collections | Get All API Collections
[**GetApiCollectionEndpoints**](ApiCollectionApi.md#GetApiCollectionEndpoints) | **Get** /obp/v5.1.0/api-collections/API_COLLECTION_ID/api-collection-endpoints | Get Api Collection Endpoints
[**GetApiCollectionsForUser**](ApiCollectionApi.md#GetApiCollectionsForUser) | **Get** /obp/v5.1.0/users/{USER_ID}/api-collections | Get Api Collections for User
[**GetFeaturedApiCollections**](ApiCollectionApi.md#GetFeaturedApiCollections) | **Get** /obp/v5.1.0/api-collections/featured | Get Featured Api Collections
[**GetMyApiCollectionById**](ApiCollectionApi.md#GetMyApiCollectionById) | **Get** /obp/v5.1.0/my/api-collections/API_COLLECTION_ID | Get My Api Collection By Id
[**GetMyApiCollectionByName**](ApiCollectionApi.md#GetMyApiCollectionByName) | **Get** /obp/v5.1.0/my/api-collections/name/API_COLLECTION_NAME | Get My Api Collection By Name
[**GetMyApiCollectionEndpoint**](ApiCollectionApi.md#GetMyApiCollectionEndpoint) | **Get** /obp/v5.1.0/my/api-collections/API_COLLECTION_NAME/api-collection-endpoints/OPERATION_ID | Get My Api Collection Endpoint
[**GetMyApiCollectionEndpoints**](ApiCollectionApi.md#GetMyApiCollectionEndpoints) | **Get** /obp/v5.1.0/my/api-collections/API_COLLECTION_NAME/api-collection-endpoints | Get My Api Collection Endpoints
[**GetMyApiCollectionEndpointsById**](ApiCollectionApi.md#GetMyApiCollectionEndpointsById) | **Get** /obp/v5.1.0/my/api-collection-ids/API_COLLECTION_ID/api-collection-endpoints | Get My Api Collection Endpoints By Id
[**GetMyApiCollections**](ApiCollectionApi.md#GetMyApiCollections) | **Get** /obp/v5.1.0/my/api-collections | Get My Api Collections
[**GetSharableApiCollectionById**](ApiCollectionApi.md#GetSharableApiCollectionById) | **Get** /obp/v5.1.0/api-collections/sharable/API_COLLECTION_ID | Get Sharable Api Collection By Id
[**UpdateMyApiCollection**](ApiCollectionApi.md#UpdateMyApiCollection) | **Put** /obp/v5.1.0/my/api-collections/API_COLLECTION_ID | Update My Api Collection By API_COLLECTION_ID


# **CreateMyApiCollection**
> ApiCollectionJson400 CreateMyApiCollection(ctx, body)
Create My Api Collection

<p>Create Api Collection for logged in user.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostApiCollectionJson400**](PostApiCollectionJson400.md)| PostApiCollectionJson400 object that needs to be added. | 

### Return type

[**ApiCollectionJson400**](ApiCollectionJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMyApiCollectionEndpoint**
> ApiCollectionEndpointJson400 CreateMyApiCollectionEndpoint(ctx, body)
Create My Api Collection Endpoint

<p>Create Api Collection Endpoint.</p><p>glossary-item-not-found</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostApiCollectionEndpointJson400**](PostApiCollectionEndpointJson400.md)| PostApiCollectionEndpointJson400 object that needs to be added. | 

### Return type

[**ApiCollectionEndpointJson400**](ApiCollectionEndpointJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMyApiCollectionEndpointById**
> ApiCollectionEndpointJson400 CreateMyApiCollectionEndpointById(ctx, body)
Create My Api Collection Endpoint By Id

<p>Create Api Collection Endpoint By Id.</p><p>glossary-item-not-found</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostApiCollectionEndpointJson400**](PostApiCollectionEndpointJson400.md)| PostApiCollectionEndpointJson400 object that needs to be added. | 

### Return type

[**ApiCollectionEndpointJson400**](ApiCollectionEndpointJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMyApiCollection**
> Full DeleteMyApiCollection(ctx, )
Delete My Api Collection

<p>Delete Api Collection By API_COLLECTION_ID</p><p>glossary-item-not-found</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Full**](Full.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMyApiCollectionEndpoint**
> Full DeleteMyApiCollectionEndpoint(ctx, )
Delete My Api Collection Endpoint

<p>glossary-item-not-found</p><p>Delete Api Collection Endpoint By OPERATION_ID</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Full**](Full.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMyApiCollectionEndpointById**
> Full DeleteMyApiCollectionEndpointById(ctx, )
Delete My Api Collection Endpoint By Id

<p>glossary-item-not-found<br />Delete Api Collection Endpoint<br />Delete Api Collection Endpoint By Id</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Full**](Full.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMyApiCollectionEndpointByOperationId**
> Full DeleteMyApiCollectionEndpointByOperationId(ctx, )
Delete My Api Collection Endpoint By Id

<p>glossary-item-not-found</p><p>Delete Api Collection Endpoint By OPERATION_ID</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Full**](Full.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllApiCollections**
> ApiCollectionsJson400 GetAllApiCollections(ctx, )
Get All API Collections

<p>Get All API Collections.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCollectionsJson400**](ApiCollectionsJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiCollectionEndpoints**
> ApiCollectionEndpointsJson400 GetApiCollectionEndpoints(ctx, )
Get Api Collection Endpoints

<p>Get Api Collection Endpoints By API_COLLECTION_ID.</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCollectionEndpointsJson400**](ApiCollectionEndpointsJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiCollectionsForUser**
> ApiCollectionsJson400 GetApiCollectionsForUser(ctx, uSERID)
Get Api Collections for User

<p>Get Api Collections for User.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 

### Return type

[**ApiCollectionsJson400**](ApiCollectionsJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeaturedApiCollections**
> ApiCollectionsJson400 GetFeaturedApiCollections(ctx, )
Get Featured Api Collections

<p>Get Featured Api Collections.</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCollectionsJson400**](ApiCollectionsJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyApiCollectionById**
> ApiCollectionJson400 GetMyApiCollectionById(ctx, )
Get My Api Collection By Id

<p>Get Api Collection By API_COLLECTION_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCollectionJson400**](ApiCollectionJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyApiCollectionByName**
> ApiCollectionJson400 GetMyApiCollectionByName(ctx, )
Get My Api Collection By Name

<p>Get Api Collection By API_COLLECTION_NAME.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCollectionJson400**](ApiCollectionJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyApiCollectionEndpoint**
> ApiCollectionEndpointJson400 GetMyApiCollectionEndpoint(ctx, )
Get My Api Collection Endpoint

<p>Get Api Collection Endpoint By API_COLLECTION_NAME and OPERATION_ID.</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCollectionEndpointJson400**](ApiCollectionEndpointJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyApiCollectionEndpoints**
> ApiCollectionEndpointsJson400 GetMyApiCollectionEndpoints(ctx, )
Get My Api Collection Endpoints

<p>Get Api Collection Endpoints By API_COLLECTION_NAME.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCollectionEndpointsJson400**](ApiCollectionEndpointsJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyApiCollectionEndpointsById**
> ApiCollectionEndpointsJson400 GetMyApiCollectionEndpointsById(ctx, )
Get My Api Collection Endpoints By Id

<p>Get Api Collection Endpoints By API_COLLECTION_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCollectionEndpointsJson400**](ApiCollectionEndpointsJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyApiCollections**
> ApiCollectionsJson400 GetMyApiCollections(ctx, )
Get My Api Collections

<p>Get all the apiCollections for logged in user.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCollectionsJson400**](ApiCollectionsJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSharableApiCollectionById**
> ApiCollectionJson400 GetSharableApiCollectionById(ctx, )
Get Sharable Api Collection By Id

<p>Get Sharable Api Collection By Id.<br />Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCollectionJson400**](ApiCollectionJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMyApiCollection**
> ApiCollectionJson400 UpdateMyApiCollection(ctx, body)
Update My Api Collection By API_COLLECTION_ID

<p>Update Api Collection for logged in user.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostApiCollectionJson400**](PostApiCollectionJson400.md)| PostApiCollectionJson400 object that needs to be added. | 

### Return type

[**ApiCollectionJson400**](ApiCollectionJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

