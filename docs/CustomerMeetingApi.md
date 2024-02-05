# \CustomerMeetingApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMeeting**](CustomerMeetingApi.md#CreateMeeting) | **Post** /obp/v5.1.0/banks/{BANK_ID}/meetings | Create Meeting (video conference/call)
[**GetMeeting**](CustomerMeetingApi.md#GetMeeting) | **Get** /obp/v5.1.0/banks/{BANK_ID}/meetings/{MEETING_ID} | Get Meeting
[**GetMeetings**](CustomerMeetingApi.md#GetMeetings) | **Get** /obp/v5.1.0/banks/{BANK_ID}/meetings | Get Meetings


# **CreateMeeting**
> MeetingJsonV310 CreateMeeting(ctx, body, bANKID)
Create Meeting (video conference/call)

<p>Create Meeting: Initiate a video conference/call with the bank.</p><p>The Meetings resource contains meta data about video/other conference sessions</p><p>provider_id determines the provider of the meeting / video chat service. MUST be url friendly (no spaces).</p><p>purpose_id explains the purpose of the chat. onboarding | mortgage | complaint etc. MUST be url friendly (no spaces).</p><p>Login is required.</p><p>This call is <strong>experimental</strong>. Currently staff_user_id is not set. Further calls will be needed to correctly set this.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMeetingJsonV310**](CreateMeetingJsonV310.md)| CreateMeetingJsonV310 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**MeetingJsonV310**](MeetingJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMeeting**
> MeetingJsonV310 GetMeeting(ctx, mEETINGID, bANKID)
Get Meeting

<p>Get Meeting specified by BANK_ID / MEETING_ID<br />Meetings contain meta data about, and are used to facilitate, video conferences / chats etc.</p><p>The actual conference/chats are handled by external services.</p><p>Login is required.</p><p>This call is <strong>experimental</strong> and will require further authorisation in the future.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mEETINGID** | **string**| the meeting id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**MeetingJsonV310**](MeetingJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMeetings**
> MeetingsJsonV310 GetMeetings(ctx, bANKID)
Get Meetings

<p>Meetings contain meta data about, and are used to facilitate, video conferences / chats etc.</p><p>The actual conference/chats are handled by external services.</p><p>Login is required.</p><p>This call is <strong>experimental</strong> and will require further authorisation in the future.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**MeetingsJsonV310**](MeetingsJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

