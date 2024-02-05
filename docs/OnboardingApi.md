# \OnboardingApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](OnboardingApi.md#CreateAccount) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID} | Create Account (PUT)
[**CreateUser**](OnboardingApi.md#CreateUser) | **Post** /obp/v5.1.0/users | Create User


# **CreateAccount**
> CreateAccountResponseJsonV310 CreateAccount(ctx, body, aCCOUNTID, bANKID)
Create Account (PUT)

<p>Create Account at bank specified by BANK_ID with Id specified by ACCOUNT_ID.</p><p>The User can create an Account for themself  - or -  the User that has the USER_ID specified in the POST body.</p><p>If the PUT body USER_ID <em>is</em> specified, the logged in user must have the Role canCreateAccount. Once created, the Account will be owned by the User specified by USER_ID.</p><p>If the PUT body USER_ID is <em>not</em> specified, the account will be owned by the logged in User.</p><p>The 'product_code' field SHOULD be a product_code from Product.<br />If the 'product_code' matches a product_code from Product, account attributes will be created that match the Product Attributes.</p><p>Note: The Amount MUST be zero.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateAccountRequestJsonV500**](CreateAccountRequestJsonV500.md)| CreateAccountRequestJsonV500 object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CreateAccountResponseJsonV310**](CreateAccountResponseJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> UserJsonV200 CreateUser(ctx, body)
Create User

<p>Creates OBP user.<br />No authorisation (currently) required.</p><p>Mimics current webform to Register.</p><p>Requires username(email) and password.</p><p>Returns 409 error if username not unique.</p><p>May require validation of email address.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserJson**](CreateUserJson.md)| CreateUserJson object that needs to be added. | 

### Return type

[**UserJsonV200**](UserJsonV200.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

