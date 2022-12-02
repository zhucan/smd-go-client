# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCephfsSubvolumeGroup**](CephfsSubvolumegroupsApi.md#CreateCephfsSubvolumeGroup) | **Post** /cephfs-subvolumegroups | create cephfs subvolumegroup
[**DeleteCephfsSubvolumeGroup**](CephfsSubvolumegroupsApi.md#DeleteCephfsSubvolumeGroup) | **Delete** /cephfs-subvolumegroups/{cephfs-subvolumegroup-id} | delete a cephfs subvolumegroup
[**GetCephfsSubvolumeGroup**](CephfsSubvolumegroupsApi.md#GetCephfsSubvolumeGroup) | **Get** /cephfs-subvolumegroups/{cephfs-subvolumegroup-id} | get a cephfs subvolumegroup
[**ListCephfsSubvolumeGroups**](CephfsSubvolumegroupsApi.md#ListCephfsSubvolumeGroups) | **Get** /cephfs-subvolumegroups | get all cephfs subvolumegroups

# **CreateCephfsSubvolumeGroup**
> ModelsCephfsSubvolumeGroup CreateCephfsSubvolumeGroup(ctx, body)
create cephfs subvolumegroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesCephfsSubvolumeGroupRequestParams**](RoutesCephfsSubvolumeGroupRequestParams.md)|  | 

### Return type

[**ModelsCephfsSubvolumeGroup**](models.CephfsSubvolumeGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCephfsSubvolumeGroup**
> ModelsCephfsSubvolumeGroup DeleteCephfsSubvolumeGroup(ctx, cephfsSubvolumegroupId)
delete a cephfs subvolumegroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cephfsSubvolumegroupId** | **int64**| identifier of the cephfs subvolumegroup | 

### Return type

[**ModelsCephfsSubvolumeGroup**](models.CephfsSubvolumeGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCephfsSubvolumeGroup**
> ModelsCephfsSubvolumeGroup GetCephfsSubvolumeGroup(ctx, cephfsSubvolumegroupId)
get a cephfs subvolumegroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cephfsSubvolumegroupId** | **int64**| identifier of the cephfs subvolumegroup | 

### Return type

[**ModelsCephfsSubvolumeGroup**](models.CephfsSubvolumeGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCephfsSubvolumeGroups**
> []ModelsCephfsSubvolumeGroup ListCephfsSubvolumeGroups(ctx, optional)
get all cephfs subvolumegroups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CephfsSubvolumegroupsApiListCephfsSubvolumeGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CephfsSubvolumegroupsApiListCephfsSubvolumeGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsCephfsSubvolumeGroup**](*models.CephfsSubvolumeGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

