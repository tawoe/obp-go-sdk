# \ATMApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAtm**](ATMApi.md#CreateAtm) | **Post** /obp/v5.1.0/banks/{BANK_ID}/atms | Create ATM
[**CreateAtmAttribute**](ATMApi.md#CreateAtmAttribute) | **Post** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID}/attributes | Create ATM Attribute
[**DeleteAtm**](ATMApi.md#DeleteAtm) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID} | Delete ATM
[**DeleteAtmAttribute**](ATMApi.md#DeleteAtmAttribute) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID}/attributes/{ATM_ATTRIBUTE_ID} | Delete ATM Attribute
[**GetAtm**](ATMApi.md#GetAtm) | **Get** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID} | Get Bank ATM
[**GetAtmAttribute**](ATMApi.md#GetAtmAttribute) | **Get** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID}/attributes/{ATM_ATTRIBUTE_ID} | Get ATM Attribute By ATM_ATTRIBUTE_ID
[**GetAtmAttributes**](ATMApi.md#GetAtmAttributes) | **Get** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID}/attributes | Get ATM Attributes
[**GetAtms**](ATMApi.md#GetAtms) | **Get** /obp/v5.1.0/banks/{BANK_ID}/atms | Get Bank ATMS
[**HeadAtms**](ATMApi.md#HeadAtms) | **Head** /obp/v5.1.0/banks/{BANK_ID}/atms | Head Bank ATMS
[**UpdateAtm**](ATMApi.md#UpdateAtm) | **Put** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID} | UPDATE ATM
[**UpdateAtmAccessibilityFeatures**](ATMApi.md#UpdateAtmAccessibilityFeatures) | **Put** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID}/accessibility-features | Update ATM Accessibility Features
[**UpdateAtmAttribute**](ATMApi.md#UpdateAtmAttribute) | **Put** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID}/attributes/{ATM_ATTRIBUTE_ID} | Update ATM Attribute
[**UpdateAtmLocationCategories**](ATMApi.md#UpdateAtmLocationCategories) | **Put** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID}/location-categories | Update ATM Location Categories
[**UpdateAtmNotes**](ATMApi.md#UpdateAtmNotes) | **Put** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID}/notes | Update ATM Notes
[**UpdateAtmServices**](ATMApi.md#UpdateAtmServices) | **Put** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID}/services | Update ATM Services
[**UpdateAtmSupportedCurrencies**](ATMApi.md#UpdateAtmSupportedCurrencies) | **Put** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID}/supported-currencies | Update ATM Supported Currencies
[**UpdateAtmSupportedLanguages**](ATMApi.md#UpdateAtmSupportedLanguages) | **Put** /obp/v5.1.0/banks/{BANK_ID}/atms/{ATM_ID}/supported-languages | Update ATM Supported Languages


# **CreateAtm**
> AtmJsonV510 CreateAtm(ctx, body, bANKID)
Create ATM

<p>Create ATM.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostAtmJsonV510**](PostAtmJsonV510.md)| PostAtmJsonV510 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmJsonV510**](AtmJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAtmAttribute**
> AtmAttributeResponseJsonV510 CreateAtmAttribute(ctx, body, aTMID, bANKID)
Create ATM Attribute

<p>Create ATM Attribute</p><p>The type field must be one of &quot;STRING&quot;, &quot;INTEGER&quot;, &quot;DOUBLE&quot; or DATE_WITH_DAY&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AtmAttributeJsonV510**](AtmAttributeJsonV510.md)| AtmAttributeJsonV510 object that needs to be added. | 
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmAttributeResponseJsonV510**](AtmAttributeResponseJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAtm**
> DeleteAtm(ctx, aTMID, bANKID)
Delete ATM

<p>Delete ATM.</p><p>This will also delete all its attributes.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAtmAttribute**
> DeleteAtmAttribute(ctx, aTMATTRIBUTEID, aTMID, bANKID)
Delete ATM Attribute

<p>Delete ATM Attribute</p><p>Delete a Atm Attribute by its id.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aTMATTRIBUTEID** | **string**| the atm attribute id | 
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAtm**
> AtmJsonV510 GetAtm(ctx, aTMID, bANKID)
Get Bank ATM

<p>Returns information about ATM for a single bank specified by BANK_ID and ATM_ID including:</p><ul><li>Address</li><li>Geo Location</li><li>License the data under this endpoint is released under</li><li>ATM Attributes</li></ul><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmJsonV510**](AtmJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAtmAttribute**
> AtmAttributeResponseJsonV510 GetAtmAttribute(ctx, aTMATTRIBUTEID, aTMID, bANKID)
Get ATM Attribute By ATM_ATTRIBUTE_ID

<p>Get ATM Attribute By ATM_ATTRIBUTE_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aTMATTRIBUTEID** | **string**| the atm attribute id | 
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmAttributeResponseJsonV510**](AtmAttributeResponseJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAtmAttributes**
> AtmAttributesResponseJsonV510 GetAtmAttributes(ctx, aTMID, bANKID)
Get ATM Attributes

<p>Get ATM Attributes</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmAttributesResponseJsonV510**](AtmAttributesResponseJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAtms**
> AtmsJsonV510 GetAtms(ctx, bANKID)
Get Bank ATMS

<p>Returns information about ATMs for a single bank specified by BANK_ID including:</p><ul><li>Address</li><li>Geo Location</li><li>License the data under this endpoint is released under</li></ul><p>Pagination:</p><p>By default, 100 records are returned.</p><p>You can use the url query parameters <em>limit</em> and <em>offset</em> for pagination</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmsJsonV510**](AtmsJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadAtms**
> AtmsJsonV400 HeadAtms(ctx, bANKID)
Head Bank ATMS

<p>Head Bank ATMS.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmsJsonV400**](AtmsJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAtm**
> AtmJsonV510 UpdateAtm(ctx, body, aTMID, bANKID)
UPDATE ATM

<p>Update ATM.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AtmJsonV510**](AtmJsonV510.md)| AtmJsonV510 object that needs to be added. | 
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmJsonV510**](AtmJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAtmAccessibilityFeatures**
> AtmAccessibilityFeaturesJson UpdateAtmAccessibilityFeatures(ctx, body, aTMID, bANKID)
Update ATM Accessibility Features

<p>Update ATM Accessibility Features.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccessibilityFeaturesJson**](AccessibilityFeaturesJson.md)| AccessibilityFeaturesJson object that needs to be added. | 
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmAccessibilityFeaturesJson**](AtmAccessibilityFeaturesJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAtmAttribute**
> AtmAttributeResponseJsonV510 UpdateAtmAttribute(ctx, body, aTMATTRIBUTEID, aTMID, bANKID)
Update ATM Attribute

<p>Update ATM Attribute.</p><p>Update an ATM Attribute by its id.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AtmAttributeJsonV510**](AtmAttributeJsonV510.md)| AtmAttributeJsonV510 object that needs to be added. | 
  **aTMATTRIBUTEID** | **string**| the atm attribute id | 
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmAttributeResponseJsonV510**](AtmAttributeResponseJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAtmLocationCategories**
> AtmLocationCategoriesResponseJsonV400 UpdateAtmLocationCategories(ctx, body, aTMID, bANKID)
Update ATM Location Categories

<p>Update ATM Location Categories.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AtmLocationCategoriesJsonV400**](AtmLocationCategoriesJsonV400.md)| AtmLocationCategoriesJsonV400 object that needs to be added. | 
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmLocationCategoriesResponseJsonV400**](AtmLocationCategoriesResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAtmNotes**
> AtmNotesResponseJsonV400 UpdateAtmNotes(ctx, body, aTMID, bANKID)
Update ATM Notes

<p>Update ATM Notes.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AtmNotesJsonV400**](AtmNotesJsonV400.md)| AtmNotesJsonV400 object that needs to be added. | 
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmNotesResponseJsonV400**](AtmNotesResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAtmServices**
> AtmServicesResponseJsonV400 UpdateAtmServices(ctx, body, aTMID, bANKID)
Update ATM Services

<p>Update ATM Services.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AtmServicesJsonV400**](AtmServicesJsonV400.md)| AtmServicesJsonV400 object that needs to be added. | 
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmServicesResponseJsonV400**](AtmServicesResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAtmSupportedCurrencies**
> AtmSupportedCurrenciesJson UpdateAtmSupportedCurrencies(ctx, body, aTMID, bANKID)
Update ATM Supported Currencies

<p>Update ATM Supported Currencies.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SupportedCurrenciesJson**](SupportedCurrenciesJson.md)| SupportedCurrenciesJson object that needs to be added. | 
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmSupportedCurrenciesJson**](AtmSupportedCurrenciesJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAtmSupportedLanguages**
> AtmSupportedLanguagesJson UpdateAtmSupportedLanguages(ctx, body, aTMID, bANKID)
Update ATM Supported Languages

<p>Update ATM Supported Languages.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SupportedLanguagesJson**](SupportedLanguagesJson.md)| SupportedLanguagesJson object that needs to be added. | 
  **aTMID** | **string**| the atm id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AtmSupportedLanguagesJson**](AtmSupportedLanguagesJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

