# \ViewSystemApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSystemView**](ViewSystemApi.md#CreateSystemView) | **Post** /obp/v5.1.0/system-views | Create System View
[**DeleteSystemView**](ViewSystemApi.md#DeleteSystemView) | **Delete** /obp/v5.1.0/system-views/{VIEW_ID} | Delete System View
[**GetSystemView**](ViewSystemApi.md#GetSystemView) | **Get** /obp/v5.1.0/system-views/{VIEW_ID} | Get System View
[**GetSystemViewsIds**](ViewSystemApi.md#GetSystemViewsIds) | **Get** /obp/v5.1.0/system-views-ids | Get Ids of System Views
[**UpdateSystemView**](ViewSystemApi.md#UpdateSystemView) | **Put** /obp/v5.1.0/system-views/{VIEW_ID} | Update System View


# **CreateSystemView**
> ViewJsonV500 CreateSystemView(ctx, body)
Create System View

<p>Create a system view</p><p>Authentication is Mandatory and the user needs to have access to the CanCreateSystemView entitlement.<br />The 'alias' field in the JSON can take one of two values:</p><ul><li><em>public</em>: to use the public alias if there is one specified for the other account.</li><li><em>private</em>: to use the public alias if there is one specified for the other account.</li><li><p><em>''(empty string)</em>: to use no alias; the view shows the real name of the other account.</p></li></ul><p>The 'hide_metadata_if_alias_used' field in the JSON can take boolean values. If it is set to <code>true</code> and there is an alias on the other account then the other accounts' metadata (like more_info, url, image_url, open_corporates_url, etc.) will be hidden. Otherwise the metadata will be shown.</p><p>The 'allowed_actions' field is a list containing the name of the actions allowed on this view, all the actions contained will be set to <code>true</code> on the view creation, the rest will be set to <code>false</code>.</p><p>Please note that system views cannot be public. In case you try to set it you will get the error OBP-30258: System view cannot be public</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateViewJsonV500**](CreateViewJsonV500.md)| CreateViewJsonV500 object that needs to be added. | 

### Return type

[**ViewJsonV500**](ViewJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSystemView**
> DeleteSystemView(ctx, vIEWID)
Delete System View

<p>Deletes the system view specified by VIEW_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemView**
> ViewJsonV500 GetSystemView(ctx, vIEWID)
Get System View

<p>Get System View</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 

### Return type

[**ViewJsonV500**](ViewJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemViewsIds**
> ViewsIdsJsonV500 GetSystemViewsIds(ctx, )
Get Ids of System Views

<p>Get Ids of System Views</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ViewsIdsJsonV500**](ViewsIdsJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSystemView**
> ViewJsonV500 UpdateSystemView(ctx, body, vIEWID)
Update System View

<p>Update an existing view on a bank account</p><p>Authentication is Mandatory and the user needs to have access to the owner view.</p><p>The json sent is the same as during view creation (above), with one difference: the 'name' field<br />of a view is not editable (it is only set when a view is created)</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateViewJsonV500**](UpdateViewJsonV500.md)| UpdateViewJsonV500 object that needs to be added. | 
  **vIEWID** | **string**| The view id | 

### Return type

[**ViewJsonV500**](ViewJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

