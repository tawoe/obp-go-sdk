# \InsurancePremiumGh29UkApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicEntityCreateInsurancePremiumGh29Uk**](InsurancePremiumGh29UkApi.md#DynamicEntityCreateInsurancePremiumGh29Uk) | **Post** /banks/gh.29.uk/InsurancePremium | Create new Insurance Premium
[**DynamicEntityDeleteInsurancePremiumGh29Uk**](InsurancePremiumGh29UkApi.md#DynamicEntityDeleteInsurancePremiumGh29Uk) | **Delete** /banks/gh.29.uk/InsurancePremium/INSURANCE_PREMIUM_ID | Delete Insurance Premium by id
[**DynamicEntityGetInsurancePremiumListGh29Uk**](InsurancePremiumGh29UkApi.md#DynamicEntityGetInsurancePremiumListGh29Uk) | **Get** /banks/gh.29.uk/InsurancePremium | Get Insurance Premium List
[**DynamicEntityGetSingleInsurancePremiumGh29Uk**](InsurancePremiumGh29UkApi.md#DynamicEntityGetSingleInsurancePremiumGh29Uk) | **Get** /banks/gh.29.uk/InsurancePremium/INSURANCE_PREMIUM_ID | Get Insurance Premium by id
[**DynamicEntityUpdateInsurancePremiumGh29Uk**](InsurancePremiumGh29UkApi.md#DynamicEntityUpdateInsurancePremiumGh29Uk) | **Put** /banks/gh.29.uk/InsurancePremium/INSURANCE_PREMIUM_ID | Update Insurance Premium


# **DynamicEntityCreateInsurancePremiumGh29Uk**
> InlineResponse20019InsurancePremiumList DynamicEntityCreateInsurancePremiumGh29Uk(ctx, body)
Create new Insurance Premium

<p>Create new Insurance Premium.</p><p>Retrive the premium for the customer.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;InsurancePremium&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body24**](Body24.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20019InsurancePremiumList**](inline_response_200_19_insurance_premium_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityDeleteInsurancePremiumGh29Uk**
> InlineResponse20019InsurancePremiumList DynamicEntityDeleteInsurancePremiumGh29Uk(ctx, body)
Delete Insurance Premium by id

<p>Delete Insurance Premium by id</p><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;InsurancePremium&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body26**](Body26.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20019InsurancePremiumList**](inline_response_200_19_insurance_premium_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetInsurancePremiumListGh29Uk**
> InlineResponse20019 DynamicEntityGetInsurancePremiumListGh29Uk(ctx, )
Get Insurance Premium List

<p>Get Insurance Premium List.</p><p>Retrive the premium for the customer.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;InsurancePremium&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p><p>Can do filter on the fields<br />e.g: /InsurancePremium?name=James%20Brown&amp;number=123.456&amp;number=11.11<br />Will do filter by this rule: name == &quot;James Brown&quot; &amp;&amp; (number==123.456 || number=11.11)</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetSingleInsurancePremiumGh29Uk**
> InlineResponse20019InsurancePremiumList DynamicEntityGetSingleInsurancePremiumGh29Uk(ctx, )
Get Insurance Premium by id

<p>Get Insurance Premium by id.</p><p>Retrive the premium for the customer.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;InsurancePremium&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20019InsurancePremiumList**](inline_response_200_19_insurance_premium_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityUpdateInsurancePremiumGh29Uk**
> InlineResponse20019InsurancePremiumList DynamicEntityUpdateInsurancePremiumGh29Uk(ctx, body)
Update Insurance Premium

<p>Update Insurance Premium.</p><p>Retrive the premium for the customer.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;InsurancePremium&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body25**](Body25.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20019InsurancePremiumList**](inline_response_200_19_insurance_premium_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

