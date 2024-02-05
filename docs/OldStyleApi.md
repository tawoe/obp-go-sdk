# \OldStyleApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ElasticSearchMetrics**](OldStyleApi.md#ElasticSearchMetrics) | **Get** /obp/v5.1.0/search/metrics | Search API Metrics via Elasticsearch


# **ElasticSearchMetrics**
> EmptyClassJson ElasticSearchMetrics(ctx, body)
Search API Metrics via Elasticsearch

<p>Search the API calls made to this API instance via Elastic Search.</p><p>Login is required.</p><p>CanSearchMetrics entitlement is required to search metrics data.</p><p>parameters:</p><p>esType  - elasticsearch type</p><p>simple query:</p><p>q       - plain_text_query</p><p>df      - default field to search</p><p>sort    - field to sort on</p><p>size    - number of hits returned, default 10</p><p>from    - show hits starting from</p><p>json query:</p><p>source  - JSON_query_(URL-escaped)</p><p>example usage:</p><p>/search/metrics/q=findThis</p><p>or:</p><p>/search/metrics/source={&quot;query&quot;:{&quot;query_string&quot;:{&quot;query&quot;:&quot;findThis&quot;}}}</p><p>Note!!</p><p>The whole JSON query string MUST be URL-encoded:</p><ul><li>For {  use %7B</li><li>For }  use %7D</li><li>For : use %3A</li><li>For &quot; use %22</li></ul><p>etc..</p><p>Only q, source and esType are passed to Elastic</p><p>Elastic simple query: <a href=\"https://www.elastic.co/guide/en/elasticsearch/reference/current/search-uri-request.html\">https://www.elastic.co/guide/en/elasticsearch/reference/current/search-uri-request.html</a></p><p>Elastic JSON query: <a href=\"https://www.elastic.co/guide/en/elasticsearch/reference/current/query-filter-context.html\">https://www.elastic.co/guide/en/elasticsearch/reference/current/query-filter-context.html</a></p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

