# \CustomerCarsApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicEntityCreatecustomerCars**](CustomerCarsApi.md#DynamicEntityCreatecustomerCars) | **Post** /customer_cars | Create new Customer Cars
[**DynamicEntityDeletecustomerCars**](CustomerCarsApi.md#DynamicEntityDeletecustomerCars) | **Delete** /customer_cars/CUSTOMER_CARS_ID | Delete Customer Cars by id
[**DynamicEntityGetSinglecustomerCars**](CustomerCarsApi.md#DynamicEntityGetSinglecustomerCars) | **Get** /customer_cars/CUSTOMER_CARS_ID | Get Customer Cars by id
[**DynamicEntityGetcustomerCarsList**](CustomerCarsApi.md#DynamicEntityGetcustomerCarsList) | **Get** /customer_cars | Get Customer Cars List
[**DynamicEntityUpdatecustomerCars**](CustomerCarsApi.md#DynamicEntityUpdatecustomerCars) | **Put** /customer_cars/CUSTOMER_CARS_ID | Update Customer Cars


# **DynamicEntityCreatecustomerCars**
> InlineResponse20023CustomerCarsList DynamicEntityCreatecustomerCars(ctx, body)
Create new Customer Cars

<p>Create new Customer Cars.</p><p>The car the customer arrived in</p><p><strong>Property List:</strong></p><ul><li>customer_identifier: description of <strong>customer_identifier</strong> field, can be markdown text.</li><li>manufacturer: description of <strong>manufacturer</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;customer_cars&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body36**](Body36.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20023CustomerCarsList**](inline_response_200_23_customer_cars_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityDeletecustomerCars**
> InlineResponse20023CustomerCarsList DynamicEntityDeletecustomerCars(ctx, body)
Delete Customer Cars by id

<p>Delete Customer Cars by id</p><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;customer_cars&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body38**](Body38.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20023CustomerCarsList**](inline_response_200_23_customer_cars_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetSinglecustomerCars**
> InlineResponse20023CustomerCarsList DynamicEntityGetSinglecustomerCars(ctx, )
Get Customer Cars by id

<p>Get Customer Cars by id.</p><p>The car the customer arrived in</p><p><strong>Property List:</strong></p><ul><li>customer_identifier: description of <strong>customer_identifier</strong> field, can be markdown text.</li><li>manufacturer: description of <strong>manufacturer</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;customer_cars&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20023CustomerCarsList**](inline_response_200_23_customer_cars_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetcustomerCarsList**
> InlineResponse20023 DynamicEntityGetcustomerCarsList(ctx, )
Get Customer Cars List

<p>Get Customer Cars List.</p><p>The car the customer arrived in</p><p><strong>Property List:</strong></p><ul><li>customer_identifier: description of <strong>customer_identifier</strong> field, can be markdown text.</li><li>manufacturer: description of <strong>manufacturer</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;customer_cars&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p><p>Can do filter on the fields<br />e.g: /customer_cars?name=James%20Brown&amp;number=123.456&amp;number=11.11<br />Will do filter by this rule: name == &quot;James Brown&quot; &amp;&amp; (number==123.456 || number=11.11)</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityUpdatecustomerCars**
> InlineResponse20023CustomerCarsList DynamicEntityUpdatecustomerCars(ctx, body)
Update Customer Cars

<p>Update Customer Cars.</p><p>The car the customer arrived in</p><p><strong>Property List:</strong></p><ul><li>customer_identifier: description of <strong>customer_identifier</strong> field, can be markdown text.</li><li>manufacturer: description of <strong>manufacturer</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;customer_cars&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body37**](Body37.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20023CustomerCarsList**](inline_response_200_23_customer_cars_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

