# \DataWarehouseApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataWarehouseSearch**](DataWarehouseApi.md#DataWarehouseSearch) | **Post** /obp/v5.1.0/search/warehouse/{INDEX} | Data Warehouse Search
[**DataWarehouseStatistics**](DataWarehouseApi.md#DataWarehouseStatistics) | **Post** /obp/v5.1.0/search/warehouse/statistics/{INDEX}/{FIELD} | Data Warehouse Statistics


# **DataWarehouseSearch**
> EmptyClassJson DataWarehouseSearch(ctx, body, iNDEX)
Data Warehouse Search

<p>Search the data warehouse and get row level results.</p><p>Authentication is Mandatory</p><p>CanSearchWarehouse entitlement is required. You can request the Role below.</p><p>Elastic (search) is used in the background. See links below for syntax.</p><p>Examples of usage:</p><p>POST /search/warehouse/THE_INDEX_YOU_WANT_TO_USE</p><p>POST /search/warehouse/INDEX1,INDEX2</p><p>POST /search/warehouse/ALL</p><p>{ Any valid elasticsearch query DSL in the body }</p><p><a href=\"https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html\">Elasticsearch query DSL</a></p><p><a href=\"https://www.elastic.co/guide/en/elasticsearch/reference/6.2/search-request-body.html\">Elastic simple query</a></p><p><a href=\"https://www.elastic.co/guide/en/elasticsearch/reference/6.2/search-aggregations.html\">Elastic aggregations</a></p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ElasticSearchJsonV300**](ElasticSearchJsonV300.md)| ElasticSearchJsonV300 object that needs to be added. | 
  **iNDEX** | **string**| the elastic search index | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataWarehouseStatistics**
> EmptyClassJson DataWarehouseStatistics(ctx, body, fIELD, iNDEX)
Data Warehouse Statistics

<p>Search the data warehouse and get statistical aggregations over a warehouse field</p><p>Does a stats aggregation over some numeric field:</p><p><a href=\"https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-stats-aggregation.html\">https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-stats-aggregation.html</a></p><p>Authentication is Mandatory</p><p>CanSearchWarehouseStats Role is required. You can request this below.</p><p>Elastic (search) is used in the background. See links below for syntax.</p><p>Examples of usage:</p><p>POST /search/warehouse/statistics/INDEX/FIELD</p><p>POST /search/warehouse/statistics/ALL/FIELD</p><p>{ Any valid elasticsearch query DSL in the body }</p><p><a href=\"https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html\">Elasticsearch query DSL</a></p><p><a href=\"https://www.elastic.co/guide/en/elasticsearch/reference/6.2/search-request-body.html\">Elastic simple query</a></p><p><a href=\"https://www.elastic.co/guide/en/elasticsearch/reference/6.2/search-aggregations.html\">Elastic aggregations</a></p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ElasticSearchJsonV300**](ElasticSearchJsonV300.md)| ElasticSearchJsonV300 object that needs to be added. | 
  **fIELD** | **string**| the elastic search field | 
  **iNDEX** | **string**| the elastic search index | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

