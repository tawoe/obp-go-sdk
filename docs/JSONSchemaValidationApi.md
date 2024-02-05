# \JSONSchemaValidationApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJsonSchemaValidation**](JSONSchemaValidationApi.md#CreateJsonSchemaValidation) | **Post** /obp/v5.1.0/management/json-schema-validations/OPERATION_ID | Create a JSON Schema Validation
[**DeleteJsonSchemaValidation**](JSONSchemaValidationApi.md#DeleteJsonSchemaValidation) | **Delete** /obp/v5.1.0/management/json-schema-validations/OPERATION_ID | Delete a JSON Schema Validation
[**GetAllJsonSchemaValidations**](JSONSchemaValidationApi.md#GetAllJsonSchemaValidations) | **Get** /obp/v5.1.0/management/json-schema-validations | Get all JSON Schema Validations
[**GetAllJsonSchemaValidationsPublic**](JSONSchemaValidationApi.md#GetAllJsonSchemaValidationsPublic) | **Get** /obp/v5.1.0/endpoints/json-schema-validations | Get all JSON Schema Validations - public
[**GetJsonSchemaValidation**](JSONSchemaValidationApi.md#GetJsonSchemaValidation) | **Get** /obp/v5.1.0/management/json-schema-validations/OPERATION_ID | Get a JSON Schema Validation
[**UpdateJsonSchemaValidation**](JSONSchemaValidationApi.md#UpdateJsonSchemaValidation) | **Put** /obp/v5.1.0/management/json-schema-validations/OPERATION_ID | Update a JSON Schema Validation


# **CreateJsonSchemaValidation**
> JsonValidationV400 CreateJsonSchemaValidation(ctx, body)
Create a JSON Schema Validation

<p>Create a JSON Schema Validation.</p><p>Introduction:</p>  <p>JSON Schema is &quot;a vocabulary that allows you to annotate and validate JSON documents&quot;.</p><p>By applying JSON Schema Validation to your OBP endpoints you can constrain POST and PUT request bodies. For example, you can set minimum / maximum lengths of fields and constrain values to certain lists or regular expressions.</p><p>See <a href=\"https://json-schema.org/\">JSONSchema.org</a> for more information about the JSON Schema standard.</p><p>To create a JSON Schema from an any JSON Request body you can use <a href=\"https://jsonschema.net/app/schemas/0\">JSON Schema Net</a></p><p>(The video link below shows how to use that)</p><p>Note: OBP Dynamic Entities also use JSON Schema Validation so you don't need to additionally wrap the resulting endpoints with extra JSON Schema Validation but you could do.</p><p>You can apply JSON schema validations to any OBP endpoint's request body using the POST and PUT endpoints listed in the link below.</p><p>PLEASE SEE the following video explanation: <a href=\"https://vimeo.com/485287014\">JSON schema validation of request for Static and Dynamic Endpoints and Entities</a></p><p>To use this endpoint, please supply a valid json-schema in the request body.</p><p>Note: It might take a few minutes for the newly created JSON Schema to take effect!</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonSchemaV400**](JsonSchemaV400.md)| JsonSchemaV400 object that needs to be added. | 

### Return type

[**JsonValidationV400**](JsonValidationV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteJsonSchemaValidation**
> bool DeleteJsonSchemaValidation(ctx, )
Delete a JSON Schema Validation

<p>Delete a JSON Schema Validation by operation_id.</p><p>Authentication is Mandatory</p>

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

# **GetAllJsonSchemaValidations**
> InlineResponse2002 GetAllJsonSchemaValidations(ctx, )
Get all JSON Schema Validations

<p>Get all JSON Schema Validations.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllJsonSchemaValidationsPublic**
> InlineResponse2002 GetAllJsonSchemaValidationsPublic(ctx, )
Get all JSON Schema Validations - public

<p>Get all JSON Schema Validations - public.</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJsonSchemaValidation**
> JsonValidationV400 GetJsonSchemaValidation(ctx, )
Get a JSON Schema Validation

<p>Get a JSON Schema Validation by operation_id.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JsonValidationV400**](JsonValidationV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateJsonSchemaValidation**
> JsonValidationV400 UpdateJsonSchemaValidation(ctx, body)
Update a JSON Schema Validation

<p>Update a JSON Schema Validation.</p><p>Introduction:</p>  <p>JSON Schema is &quot;a vocabulary that allows you to annotate and validate JSON documents&quot;.</p><p>By applying JSON Schema Validation to your OBP endpoints you can constrain POST and PUT request bodies. For example, you can set minimum / maximum lengths of fields and constrain values to certain lists or regular expressions.</p><p>See <a href=\"https://json-schema.org/\">JSONSchema.org</a> for more information about the JSON Schema standard.</p><p>To create a JSON Schema from an any JSON Request body you can use <a href=\"https://jsonschema.net/app/schemas/0\">JSON Schema Net</a></p><p>(The video link below shows how to use that)</p><p>Note: OBP Dynamic Entities also use JSON Schema Validation so you don't need to additionally wrap the resulting endpoints with extra JSON Schema Validation but you could do.</p><p>You can apply JSON schema validations to any OBP endpoint's request body using the POST and PUT endpoints listed in the link below.</p><p>PLEASE SEE the following video explanation: <a href=\"https://vimeo.com/485287014\">JSON schema validation of request for Static and Dynamic Endpoints and Entities</a></p><p>To use this endpoint, please supply a valid json-schema in the request body.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonSchemaV400**](JsonSchemaV400.md)| JsonSchemaV400 object that needs to be added. | 

### Return type

[**JsonValidationV400**](JsonValidationV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

