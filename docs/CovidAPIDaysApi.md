# \CovidAPIDaysApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicEntityCreateCovidAPIDays**](CovidAPIDaysApi.md#DynamicEntityCreateCovidAPIDays) | **Post** /CovidAPIDays | Create new Covid APIDays
[**DynamicEntityDeleteCovidAPIDays**](CovidAPIDaysApi.md#DynamicEntityDeleteCovidAPIDays) | **Delete** /CovidAPIDays/COVID_API_DAYS_ID | Delete Covid APIDays by id
[**DynamicEntityGetCovidAPIDaysList**](CovidAPIDaysApi.md#DynamicEntityGetCovidAPIDaysList) | **Get** /CovidAPIDays | Get Covid APIDays List
[**DynamicEntityGetSingleCovidAPIDays**](CovidAPIDaysApi.md#DynamicEntityGetSingleCovidAPIDays) | **Get** /CovidAPIDays/COVID_API_DAYS_ID | Get Covid APIDays by id
[**DynamicEntityUpdateCovidAPIDays**](CovidAPIDaysApi.md#DynamicEntityUpdateCovidAPIDays) | **Put** /CovidAPIDays/COVID_API_DAYS_ID | Update Covid APIDays


# **DynamicEntityCreateCovidAPIDays**
> InlineResponse20013CovidApiDaysList DynamicEntityCreateCovidAPIDays(ctx, body)
Create new Covid APIDays

<p>Create new Covid APIDays.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;CovidAPIDays&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body6**](Body6.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20013CovidApiDaysList**](inline_response_200_13_covid_api_days_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityDeleteCovidAPIDays**
> InlineResponse20013CovidApiDaysList DynamicEntityDeleteCovidAPIDays(ctx, body)
Delete Covid APIDays by id

<p>Delete Covid APIDays by id</p><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;CovidAPIDays&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body8**](Body8.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20013CovidApiDaysList**](inline_response_200_13_covid_api_days_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetCovidAPIDaysList**
> InlineResponse20013 DynamicEntityGetCovidAPIDaysList(ctx, )
Get Covid APIDays List

<p>Get Covid APIDays List.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;CovidAPIDays&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p><p>Can do filter on the fields<br />e.g: /CovidAPIDays?name=James%20Brown&amp;number=123.456&amp;number=11.11<br />Will do filter by this rule: name == &quot;James Brown&quot; &amp;&amp; (number==123.456 || number=11.11)</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetSingleCovidAPIDays**
> InlineResponse20013CovidApiDaysList DynamicEntityGetSingleCovidAPIDays(ctx, )
Get Covid APIDays by id

<p>Get Covid APIDays by id.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;CovidAPIDays&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20013CovidApiDaysList**](inline_response_200_13_covid_api_days_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityUpdateCovidAPIDays**
> InlineResponse20013CovidApiDaysList DynamicEntityUpdateCovidAPIDays(ctx, body)
Update Covid APIDays

<p>Update Covid APIDays.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;CovidAPIDays&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body7**](Body7.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20013CovidApiDaysList**](inline_response_200_13_covid_api_days_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

