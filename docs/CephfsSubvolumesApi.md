# \CephfsSubvolumesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCephfsSubvolume**](CephfsSubvolumesApi.md#CreateCephfsSubvolume) | **Post** /cephfs-subvolumes | create cephfs subvolume
[**DeleteCephfsSubvolume**](CephfsSubvolumesApi.md#DeleteCephfsSubvolume) | **Delete** /cephfs-subvolumes/{cephfs-subvolume-id} | delete a cephfs subvolume
[**GetCephfsSubvolume**](CephfsSubvolumesApi.md#GetCephfsSubvolume) | **Get** /cephfs-subvolumes/{cephfs-subvolume-id} | get a cephfs subvolume
[**ListCephfsSubvolumes**](CephfsSubvolumesApi.md#ListCephfsSubvolumes) | **Get** /cephfs-subvolumes | get all cephfs subvolumes


# **CreateCephfsSubvolume**
> ModelsCephfsSubvolume CreateCephfsSubvolume(ctx, body)
create cephfs subvolume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesCephfsSubvolumeRequestParams**](RoutesCephfsSubvolumeRequestParams.md)|  | 

### Return type

[**ModelsCephfsSubvolume**](models.CephfsSubvolume.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCephfsSubvolume**
> ModelsCephfsSubvolume DeleteCephfsSubvolume(ctx, cephfsSubvolumeId, optional)
delete a cephfs subvolume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cephfsSubvolumeId** | **int64**| identifier of the cephfs subvolume | 
 **optional** | ***CephfsSubvolumesApiDeleteCephfsSubvolumeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CephfsSubvolumesApiDeleteCephfsSubvolumeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| force option will ignore errors | 

### Return type

[**ModelsCephfsSubvolume**](models.CephfsSubvolume.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCephfsSubvolume**
> ModelsCephfsSubvolume GetCephfsSubvolume(ctx, cephfsSubvolumeId)
get a cephfs subvolume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cephfsSubvolumeId** | **int64**| identifier of the cephfs subvolume | 

### Return type

[**ModelsCephfsSubvolume**](models.CephfsSubvolume.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCephfsSubvolumes**
> []ModelsCephfsSubvolume ListCephfsSubvolumes(ctx, optional)
get all cephfs subvolumes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CephfsSubvolumesApiListCephfsSubvolumesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CephfsSubvolumesApiListCephfsSubvolumesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsCephfsSubvolume**](*models.CephfsSubvolume.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

