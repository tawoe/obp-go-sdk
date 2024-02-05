# \RateLimitsApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallsLimit**](RateLimitsApi.md#CallsLimit) | **Put** /obp/v5.1.0/management/consumers/{CONSUMER_ID}/consumer/call-limits | Set Rate Limits / Call Limits per Consumer
[**GetRateLimitingInfo**](RateLimitsApi.md#GetRateLimitingInfo) | **Get** /obp/v5.1.0/rate-limiting | Get Rate Limiting Info


# **CallsLimit**
> CallLimitPostJsonV400 CallsLimit(ctx, body, cONSUMERID)
Set Rate Limits / Call Limits per Consumer

<p>Set the API rate limits / call limits for a Consumer:</p><p>Rate limiting can be set:</p><p>Per Second<br />Per Minute<br />Per Hour<br />Per Week<br />Per Month</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CallLimitPostJsonV400**](CallLimitPostJsonV400.md)| CallLimitPostJsonV400 object that needs to be added. | 
  **cONSUMERID** | **string**| new consumer id | 

### Return type

[**CallLimitPostJsonV400**](CallLimitPostJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRateLimitingInfo**
> RateLimitingInfoV310 GetRateLimitingInfo(ctx, )
Get Rate Limiting Info

<p>Get information about the Rate Limiting setup on this OBP Instance such as:</p><p>Is rate limiting enabled and active?<br />What backend is used to keep track of the API calls (e.g. REDIS).</p><p>Note: Rate limiting can be set at the Consumer level and also for anonymous calls.</p><p>See the consumer rate limits / call limits endpoints.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RateLimitingInfoV310**](RateLimitingInfoV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

