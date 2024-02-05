# \AuthenticationTypeValidationApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthenticationTypeValidation**](AuthenticationTypeValidationApi.md#CreateAuthenticationTypeValidation) | **Post** /obp/v5.1.0/management/authentication-type-validations/OPERATION_ID | Create an Authentication Type Validation
[**DeleteAuthenticationTypeValidation**](AuthenticationTypeValidationApi.md#DeleteAuthenticationTypeValidation) | **Delete** /obp/v5.1.0/management/authentication-type-validations/OPERATION_ID | Delete an Authentication Type Validation
[**GetAllAuthenticationTypeValidations**](AuthenticationTypeValidationApi.md#GetAllAuthenticationTypeValidations) | **Get** /obp/v5.1.0/management/authentication-type-validations | Get all Authentication Type Validations
[**GetAllAuthenticationTypeValidationsPublic**](AuthenticationTypeValidationApi.md#GetAllAuthenticationTypeValidationsPublic) | **Get** /obp/v5.1.0/endpoints/authentication-type-validations | Get all Authentication Type Validations - public
[**GetAuthenticationTypeValidation**](AuthenticationTypeValidationApi.md#GetAuthenticationTypeValidation) | **Get** /obp/v5.1.0/management/authentication-type-validations/OPERATION_ID | Get an Authentication Type Validation
[**UpdateAuthenticationTypeValidation**](AuthenticationTypeValidationApi.md#UpdateAuthenticationTypeValidation) | **Put** /obp/v5.1.0/management/authentication-type-validations/OPERATION_ID | Update an Authentication Type Validation


# **CreateAuthenticationTypeValidation**
> JsonAuthTypeValidation CreateAuthenticationTypeValidation(ctx, body)
Create an Authentication Type Validation

<p>Create an Authentication Type Validation.</p><p>Please supply allowed authentication types.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotSupportedYet**](NotSupportedYet.md)| $colon$colon object that needs to be added. | 

### Return type

[**JsonAuthTypeValidation**](JsonAuthTypeValidation.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthenticationTypeValidation**
> bool DeleteAuthenticationTypeValidation(ctx, )
Delete an Authentication Type Validation

<p>Delete an Authentication Type Validation by operation_id.</p><p>Authentication is Mandatory</p>

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

# **GetAllAuthenticationTypeValidations**
> InlineResponse2001 GetAllAuthenticationTypeValidations(ctx, )
Get all Authentication Type Validations

<p>Get all Authentication Type Validations.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAuthenticationTypeValidationsPublic**
> InlineResponse2001 GetAllAuthenticationTypeValidationsPublic(ctx, )
Get all Authentication Type Validations - public

<p>Get all Authentication Type Validations - public.</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthenticationTypeValidation**
> JsonAuthTypeValidation GetAuthenticationTypeValidation(ctx, )
Get an Authentication Type Validation

<p>Get an Authentication Type Validation by operation_id.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JsonAuthTypeValidation**](JsonAuthTypeValidation.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthenticationTypeValidation**
> JsonAuthTypeValidation UpdateAuthenticationTypeValidation(ctx, body)
Update an Authentication Type Validation

<p>Update an Authentication Type Validation.</p><p>Please supply allowed authentication types.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotSupportedYet**](NotSupportedYet.md)| $colon$colon object that needs to be added. | 

### Return type

[**JsonAuthTypeValidation**](JsonAuthTypeValidation.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

