# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCephCluster**](CephClustersApi.md#CreateCephCluster) | **Post** /ceph-clusters | create ceph cluster
[**DeleteCephCluster**](CephClustersApi.md#DeleteCephCluster) | **Delete** /ceph-clusters/{ceph-cluster-id} | delete ceph cluster
[**GetCephCluster**](CephClustersApi.md#GetCephCluster) | **Get** /ceph-clusters/{ceph-cluster-id} | get ceph cluster
[**ListCephClusters**](CephClustersApi.md#ListCephClusters) | **Get** /ceph-clusters | get all ceph cluster

# **CreateCephCluster**
> ModelsCephCluster CreateCephCluster(ctx, body)
create ceph cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesCreateCephClusterRequestParams**](RoutesCreateCephClusterRequestParams.md)|  | 

### Return type

[**ModelsCephCluster**](models.CephCluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCephCluster**
> ModelsCephCluster DeleteCephCluster(ctx, cephClusterId)
delete ceph cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cephClusterId** | **int64**| identifier of the ceph cluster | 

### Return type

[**ModelsCephCluster**](models.CephCluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCephCluster**
> ModelsCephCluster GetCephCluster(ctx, cephClusterId)
get ceph cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cephClusterId** | **int64**| identifier of the ceph cluster | 

### Return type

[**ModelsCephCluster**](models.CephCluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCephClusters**
> []ModelsCephCluster ListCephClusters(ctx, optional)
get all ceph cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CephClustersApiListCephClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CephClustersApiListCephClustersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsCephCluster**](*models.CephCluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

