# \CephNfsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCephNFS**](CephNfsApi.md#CreateCephNFS) | **Post** /ceph-nfs | create an ceph nfs
[**DeleteCephNFS**](CephNfsApi.md#DeleteCephNFS) | **Delete** /ceph-nfs/{ceph-nfs-id} | delete an ceph nfs
[**GetCephNFS**](CephNfsApi.md#GetCephNFS) | **Get** /ceph-nfs/{ceph-nfs-id} | get an ceph nfs
[**ListCephNFSes**](CephNfsApi.md#ListCephNFSes) | **Get** /ceph-nfs | get all ceph nfs


# **CreateCephNFS**
> ModelsCephNfs CreateCephNFS(ctx, body)
create an ceph nfs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesCephNfsRequestParams**](RoutesCephNfsRequestParams.md)|  | 

### Return type

[**ModelsCephNfs**](models.CephNFS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCephNFS**
> ModelsCephNfs DeleteCephNFS(ctx, cephNfsId)
delete an ceph nfs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cephNfsId** | **int64**| identifier of the ceph nfs | 

### Return type

[**ModelsCephNfs**](models.CephNFS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCephNFS**
> ModelsCephNfs GetCephNFS(ctx, cephNfsId)
get an ceph nfs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cephNfsId** | **int64**| identifier of the ceph nfs | 

### Return type

[**ModelsCephNfs**](models.CephNFS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCephNFSes**
> []ModelsCephNfs ListCephNFSes(ctx, optional)
get all ceph nfs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CephNfsApiListCephNFSesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CephNfsApiListCephNFSesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsCephNfs**](*models.CephNFS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

