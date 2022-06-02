# \CephClustersApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCephCluster**](CephClustersApi.md#CreateCephCluster) | **Post** /ceph-clusters | create ceph cluster
[**DeleteCephCluster**](CephClustersApi.md#DeleteCephCluster) | **Delete** /ceph-clusters | delete ceph cluster
[**GetCephCluster**](CephClustersApi.md#GetCephCluster) | **Get** /ceph-clusters | get ceph cluster


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
> ModelsCephCluster DeleteCephCluster(ctx, )
delete ceph cluster

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ModelsCephCluster**](models.CephCluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCephCluster**
> ModelsCephCluster GetCephCluster(ctx, )
get ceph cluster

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ModelsCephCluster**](models.CephCluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

