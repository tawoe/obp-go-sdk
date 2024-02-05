# \TestDaniel707Api

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicEntityCreateMytestDaniel707**](TestDaniel707Api.md#DynamicEntityCreateMytestDaniel707) | **Post** /my/test_daniel707 | Create new My Test Daniel707
[**DynamicEntityCreatetestDaniel707**](TestDaniel707Api.md#DynamicEntityCreatetestDaniel707) | **Post** /test_daniel707 | Create new Test Daniel707
[**DynamicEntityDeleteMytestDaniel707**](TestDaniel707Api.md#DynamicEntityDeleteMytestDaniel707) | **Delete** /my/test_daniel707/TEST_DANIEL707_ID | Delete My Test Daniel707 by id
[**DynamicEntityDeletetestDaniel707**](TestDaniel707Api.md#DynamicEntityDeletetestDaniel707) | **Delete** /test_daniel707/TEST_DANIEL707_ID | Delete Test Daniel707 by id
[**DynamicEntityGetMytestDaniel707List**](TestDaniel707Api.md#DynamicEntityGetMytestDaniel707List) | **Get** /my/test_daniel707 | Get My Test Daniel707 List
[**DynamicEntityGetSingleMytestDaniel707**](TestDaniel707Api.md#DynamicEntityGetSingleMytestDaniel707) | **Get** /my/test_daniel707/TEST_DANIEL707_ID | Get My Test Daniel707 by id
[**DynamicEntityGetSingletestDaniel707**](TestDaniel707Api.md#DynamicEntityGetSingletestDaniel707) | **Get** /test_daniel707/TEST_DANIEL707_ID | Get Test Daniel707 by id
[**DynamicEntityGettestDaniel707List**](TestDaniel707Api.md#DynamicEntityGettestDaniel707List) | **Get** /test_daniel707 | Get Test Daniel707 List
[**DynamicEntityUpdateMytestDaniel707**](TestDaniel707Api.md#DynamicEntityUpdateMytestDaniel707) | **Put** /my/test_daniel707/TEST_DANIEL707_ID | Update My Test Daniel707
[**DynamicEntityUpdatetestDaniel707**](TestDaniel707Api.md#DynamicEntityUpdatetestDaniel707) | **Put** /test_daniel707/TEST_DANIEL707_ID | Update Test Daniel707


# **DynamicEntityCreateMytestDaniel707**
> InlineResponse20024TestDaniel707List DynamicEntityCreateMytestDaniel707(ctx, body)
Create new My Test Daniel707

<p>Create new My Test Daniel707.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;test_daniel707&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body39**](Body39.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20024TestDaniel707List**](inline_response_200_24_test_daniel707_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityCreatetestDaniel707**
> InlineResponse20024TestDaniel707List DynamicEntityCreatetestDaniel707(ctx, body)
Create new Test Daniel707

<p>Create new Test Daniel707.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;test_daniel707&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body48**](Body48.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20024TestDaniel707List**](inline_response_200_24_test_daniel707_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityDeleteMytestDaniel707**
> InlineResponse20024TestDaniel707List DynamicEntityDeleteMytestDaniel707(ctx, body)
Delete My Test Daniel707 by id

<p>Delete My Test Daniel707 by id</p><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;test_daniel707&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body41**](Body41.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20024TestDaniel707List**](inline_response_200_24_test_daniel707_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityDeletetestDaniel707**
> InlineResponse20024TestDaniel707List DynamicEntityDeletetestDaniel707(ctx, body)
Delete Test Daniel707 by id

<p>Delete Test Daniel707 by id</p><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;test_daniel707&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body50**](Body50.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20024TestDaniel707List**](inline_response_200_24_test_daniel707_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetMytestDaniel707List**
> InlineResponse20024 DynamicEntityGetMytestDaniel707List(ctx, )
Get My Test Daniel707 List

<p>Get My Test Daniel707 List.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;test_daniel707&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p><p>Can do filter on the fields<br />e.g: /test_daniel707?name=James%20Brown&amp;number=123.456&amp;number=11.11<br />Will do filter by this rule: name == &quot;James Brown&quot; &amp;&amp; (number==123.456 || number=11.11)</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetSingleMytestDaniel707**
> InlineResponse20024TestDaniel707List DynamicEntityGetSingleMytestDaniel707(ctx, )
Get My Test Daniel707 by id

<p>Get My Test Daniel707 by id.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;test_daniel707&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20024TestDaniel707List**](inline_response_200_24_test_daniel707_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetSingletestDaniel707**
> InlineResponse20024TestDaniel707List DynamicEntityGetSingletestDaniel707(ctx, )
Get Test Daniel707 by id

<p>Get Test Daniel707 by id.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;test_daniel707&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20024TestDaniel707List**](inline_response_200_24_test_daniel707_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGettestDaniel707List**
> InlineResponse20024 DynamicEntityGettestDaniel707List(ctx, )
Get Test Daniel707 List

<p>Get Test Daniel707 List.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;test_daniel707&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p><p>Can do filter on the fields<br />e.g: /test_daniel707?name=James%20Brown&amp;number=123.456&amp;number=11.11<br />Will do filter by this rule: name == &quot;James Brown&quot; &amp;&amp; (number==123.456 || number=11.11)</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityUpdateMytestDaniel707**
> InlineResponse20024TestDaniel707List DynamicEntityUpdateMytestDaniel707(ctx, body)
Update My Test Daniel707

<p>Update My Test Daniel707.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;test_daniel707&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body40**](Body40.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20024TestDaniel707List**](inline_response_200_24_test_daniel707_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityUpdatetestDaniel707**
> InlineResponse20024TestDaniel707List DynamicEntityUpdatetestDaniel707(ctx, body)
Update Test Daniel707

<p>Update Test Daniel707.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;test_daniel707&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body49**](Body49.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20024TestDaniel707List**](inline_response_200_24_test_daniel707_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

