# \FXApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFx**](FXApi.md#CreateFx) | **Put** /obp/v5.1.0/banks/{BANK_ID}/fx | Create Fx
[**GetCurrenciesAtBank**](FXApi.md#GetCurrenciesAtBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/currencies | Get Currencies at a Bank
[**GetCurrentFxRate**](FXApi.md#GetCurrentFxRate) | **Get** /obp/v5.1.0/banks/{BANK_ID}/fx/{FROM_CURRENCY_CODE}/{TO_CURRENCY_CODE} | Get Current FxRate


# **CreateFx**
> FxRateJsonV220 CreateFx(ctx, body, bANKID)
Create Fx

<p>Create or Update Fx for the Bank.</p><p>Example:</p><p>“from_currency_code”:“EUR”,<br />“to_currency_code”:“USD”,<br />“conversion_value”: 1.136305,<br />“inverse_conversion_value”: 1 / 1.136305 = 0.8800454103431737,</p><p>Thus 1 Euro = 1.136305 US Dollar<br />and<br />1 US Dollar = 0.8800 Euro</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FxRateJsonV220**](FxRateJsonV220.md)| FXRateJsonV220 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**FxRateJsonV220**](FXRateJsonV220.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrenciesAtBank**
> CurrenciesJsonV510 GetCurrenciesAtBank(ctx, body, bANKID)
Get Currencies at a Bank

<p>Get Currencies specified by BANK_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CurrenciesJsonV510**](CurrenciesJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentFxRate**
> FxRateJsonV220 GetCurrentFxRate(ctx, body, tOCURRENCYCODE, fROMCURRENCYCODE, bANKID)
Get Current FxRate

<p>Get the latest FX rate specified by BANK_ID, FROM_CURRENCY_CODE and TO_CURRENCY_CODE</p><p>OBP may try different sources of FX rate information depending on the Connector in operation.</p><p>For example we want to convert EUR =&gt; USD:</p><p>OBP will:<br />1st try - Connector (database, core banking system or external FX service)<br />2nd try part 1 - fallbackexchangerates/eur.json<br />2nd try part 2 - fallbackexchangerates/usd.json (the inverse rate is used)<br />3rd try - Hardcoded map of FX rates.</p><p><img src=\"https://user-images.githubusercontent.com/485218/60005085-1eded600-966e-11e9-96fb-798b102d9ad0.png\" alt=\"FX Flow\" /></p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **tOCURRENCYCODE** | **string**| The to currency code | 
  **fROMCURRENCYCODE** | **string**| The from currency code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**FxRateJsonV220**](FXRateJsonV220.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

