# \WebUiPropsApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebUiProps**](WebUiPropsApi.md#CreateWebUiProps) | **Post** /obp/v5.1.0/management/webui_props | Create WebUiProps
[**DeleteWebUiProps**](WebUiPropsApi.md#DeleteWebUiProps) | **Delete** /obp/v5.1.0/management/webui_props/{WEB_UI_PROPS_ID} | Delete WebUiProps
[**GetWebUiProps**](WebUiPropsApi.md#GetWebUiProps) | **Get** /obp/v5.1.0/management/webui_props | Get WebUiProps


# **CreateWebUiProps**
> WebUiPropsCommons CreateWebUiProps(ctx, body)
Create WebUiProps

<p>Create a WebUiProps.</p><p>Authentication is Mandatory</p><p>Explaination of Fields:</p><ul><li>name is required String value</li><li>value is required String value</li></ul><p>The line break and double quotations should do escape, example:</p><pre><code>{&quot;name&quot;: &quot;webui_some&quot;, &quot;value&quot;: &quot;this valuehave &quot;line break&quot; and double quotations.&quot;}</code></pre><p>should do escape like this:</p><pre><code>{&quot;name&quot;: &quot;webui_some&quot;, &quot;value&quot;: &quot;this value\\nhave \\&quot;line break\\&quot; and double quotations.&quot;}</code></pre><p>Insert image examples:</p><pre><code>// set width=100 and height=50{&quot;name&quot;: &quot;webui_some_pic&quot;, &quot;value&quot;: &quot;here is a picture &lt;img alt=&quot;hello&quot; src=&quot;http://somedomain.com/images/pic.png&quot; width=&quot;100&quot; height=&quot;50&quot; /&gt;&quot;}// only set height=50{&quot;name&quot;: &quot;webui_some_pic&quot;, &quot;value&quot;: &quot;here is a picture &lt;img alt=&quot;hello&quot; src=&quot;http://somedomain.com/images/pic.png&quot; width=&quot;&quot; height=&quot;50&quot; /&gt;&quot;}// only width=20%{&quot;name&quot;: &quot;webui_some_pic&quot;, &quot;value&quot;: &quot;here is a picture &lt;img alt=&quot;hello&quot; src=&quot;http://somedomain.com/images/pic.png&quot; width=&quot;20%&quot; height=&quot;&quot; /&gt;&quot;}</code></pre>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebUiPropsCommons**](WebUiPropsCommons.md)| WebUiPropsCommons object that needs to be added. | 

### Return type

[**WebUiPropsCommons**](WebUiPropsCommons.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebUiProps**
> DeleteWebUiProps(ctx, wEBUIPROPSID)
Delete WebUiProps

<p>Delete a WebUiProps specified by WEB_UI_PROPS_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **wEBUIPROPSID** | **string**| the web ui props id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebUiProps**
> InlineResponse20011 GetWebUiProps(ctx, )
Get WebUiProps

<p>Get the all WebUiProps key values, those props key with &quot;webui_&quot; can be stored in DB, this endpoint get all from DB.</p><p>url query parameter:<br />active: It must be a boolean string. and If active = true, it will show<br />combination of explicit (inserted) + implicit (default)  method_routings.</p><p>eg:<br /><a href=\"https://apisandbox.openbankproject.com/obp/v3.1.0/management/webui_props\">https://apisandbox.openbankproject.com/obp/v3.1.0/management/webui_props</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v3.1.0/management/webui_props?active=true\">https://apisandbox.openbankproject.com/obp/v3.1.0/management/webui_props?active=true</a></p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

