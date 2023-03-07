# \ObjectStoresApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectStore**](ObjectStoresApi.md#CreateObjectStore) | **Post** /object-stores | create an object store
[**DeleteObjectStore**](ObjectStoresApi.md#DeleteObjectStore) | **Delete** /object-stores/{object-store-id} | delete an object store
[**GetObjectStore**](ObjectStoresApi.md#GetObjectStore) | **Get** /object-stores/{object-store-id} | get an object store
[**ListObjectStores**](ObjectStoresApi.md#ListObjectStores) | **Get** /object-stores | get all object stores


# **CreateObjectStore**
> ModelsObjectStore CreateObjectStore(ctx, body)
create an object store

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesObjectStoreRequestParams**](RoutesObjectStoreRequestParams.md)|  | 

### Return type

[**ModelsObjectStore**](models.ObjectStore.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteObjectStore**
> ModelsObjectStore DeleteObjectStore(ctx, objectStoreId)
delete an object store

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectStoreId** | **int64**| identifier of the object store | 

### Return type

[**ModelsObjectStore**](models.ObjectStore.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectStore**
> ModelsObjectStore GetObjectStore(ctx, objectStoreId)
get an object store

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectStoreId** | **int64**| identifier of the object store | 

### Return type

[**ModelsObjectStore**](models.ObjectStore.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListObjectStores**
> []ModelsObjectStore ListObjectStores(ctx, optional)
get all object stores

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ObjectStoresApiListObjectStoresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ObjectStoresApiListObjectStoresOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsObjectStore**](*models.ObjectStore.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

