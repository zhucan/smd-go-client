# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCephFilesystem**](CephFilesystemsApi.md#CreateCephFilesystem) | **Post** /ceph-filesystems | create a ceph filesystem
[**DeleteCephFilesystem**](CephFilesystemsApi.md#DeleteCephFilesystem) | **Delete** /ceph-filesystems/{cephfs-id} | delete a ceph filesystem
[**GetCephFilesystem**](CephFilesystemsApi.md#GetCephFilesystem) | **Get** /ceph-filesystems/{cephfs-id} | get a ceph filesystem
[**ListCephFilesystems**](CephFilesystemsApi.md#ListCephFilesystems) | **Get** /ceph-filesystems | get all ceph filesystems

# **CreateCephFilesystem**
> ModelsCephfs CreateCephFilesystem(ctx, body)
create a ceph filesystem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesCephfsRequestParams**](RoutesCephfsRequestParams.md)|  | 

### Return type

[**ModelsCephfs**](models.Cephfs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCephFilesystem**
> ModelsCephfs DeleteCephFilesystem(ctx, cephfsId)
delete a ceph filesystem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cephfsId** | **int64**| identifier of the ceph filesystem | 

### Return type

[**ModelsCephfs**](models.Cephfs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCephFilesystem**
> ModelsCephfs GetCephFilesystem(ctx, cephfsId)
get a ceph filesystem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cephfsId** | **int64**| identifier of the ceph filesystem | 

### Return type

[**ModelsCephfs**](models.Cephfs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCephFilesystems**
> []ModelsCephfs ListCephFilesystems(ctx, optional)
get all ceph filesystems

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CephFilesystemsApiListCephFilesystemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CephFilesystemsApiListCephFilesystemsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsCephfs**](*models.Cephfs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

