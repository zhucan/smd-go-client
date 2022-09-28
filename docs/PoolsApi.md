# \PoolsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePool**](PoolsApi.md#CreatePool) | **Post** /pools | create a pool
[**DeletePool**](PoolsApi.md#DeletePool) | **Delete** /pools/{pool-id} | delete a pool
[**GetPool**](PoolsApi.md#GetPool) | **Get** /pools/{pool-id} | get a pool
[**ListPools**](PoolsApi.md#ListPools) | **Get** /pools | get all pools


# **CreatePool**
> ModelsPool CreatePool(ctx, body)
create a pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesPoolCreateRequestParams**](RoutesPoolCreateRequestParams.md)|  | 

### Return type

[**ModelsPool**](models.Pool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePool**
> ModelsPool DeletePool(ctx, poolId)
delete a pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **int64**| identifier of the pool | 

### Return type

[**ModelsPool**](models.Pool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPool**
> ModelsPool GetPool(ctx, poolId)
get a pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **int64**| identifier of the pool | 

### Return type

[**ModelsPool**](models.Pool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPools**
> []ModelsPool ListPools(ctx, optional)
get all pools

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoolsApiListPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoolsApiListPoolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsPool**](*models.Pool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

