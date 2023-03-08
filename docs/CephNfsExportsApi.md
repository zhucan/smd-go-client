# \CephNfsExportsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCephNFSExport**](CephNfsExportsApi.md#CreateCephNFSExport) | **Post** /ceph-nfs-exports | create ceph nfs export
[**DeleteCephNFSExport**](CephNfsExportsApi.md#DeleteCephNFSExport) | **Delete** /ceph-nfs-exports/{ceph-nfs-export-id} | delete a ceph nfs export
[**GetCephNFSExport**](CephNfsExportsApi.md#GetCephNFSExport) | **Get** /ceph-nfs-exports/{ceph-nfs-export-id} | get a ceph nfs export
[**ListCephNFSExports**](CephNfsExportsApi.md#ListCephNFSExports) | **Get** /ceph-nfs-exports | get all ceph nfs exports


# **CreateCephNFSExport**
> ModelsCephNfsExport CreateCephNFSExport(ctx, body)
create ceph nfs export

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesCephNfsExportRequestParams**](RoutesCephNfsExportRequestParams.md)|  | 

### Return type

[**ModelsCephNfsExport**](models.CephNFSExport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCephNFSExport**
> ModelsCephNfsExport DeleteCephNFSExport(ctx, cephNfsExportId, optional)
delete a ceph nfs export

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cephNfsExportId** | **int64**| identifier of the ceph nfs export | 
 **optional** | ***CephNfsExportsApiDeleteCephNFSExportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CephNfsExportsApiDeleteCephNFSExportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| force option will ignore errors | 

### Return type

[**ModelsCephNfsExport**](models.CephNFSExport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCephNFSExport**
> ModelsCephNfsExport GetCephNFSExport(ctx, cephNfsExportId)
get a ceph nfs export

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cephNfsExportId** | **int64**| identifier of the ceph nfs export | 

### Return type

[**ModelsCephNfsExport**](models.CephNFSExport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCephNFSExports**
> []ModelsCephNfsExport ListCephNFSExports(ctx, optional)
get all ceph nfs exports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CephNfsExportsApiListCephNFSExportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CephNfsExportsApiListCephNFSExportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsCephNfsExport**](*models.CephNFSExport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

