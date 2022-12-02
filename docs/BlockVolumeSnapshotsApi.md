# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockVolumeSnapshot**](BlockVolumeSnapshotsApi.md#GetBlockVolumeSnapshot) | **Get** /block-volume-snapshots/{snapshot-id} | get a snapshot of block volume
[**ListBlockVolumeSnapshots**](BlockVolumeSnapshotsApi.md#ListBlockVolumeSnapshots) | **Get** /block-volume-snapshots | list all snapshots of block volume

# **GetBlockVolumeSnapshot**
> ModelsBlockVolumeSnapshot GetBlockVolumeSnapshot(ctx, snapshotId)
get a snapshot of block volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **int64**| identifier of the snapshot | 

### Return type

[**ModelsBlockVolumeSnapshot**](models.BlockVolumeSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBlockVolumeSnapshots**
> []ModelsBlockVolumeSnapshot ListBlockVolumeSnapshots(ctx, optional)
list all snapshots of block volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlockVolumeSnapshotsApiListBlockVolumeSnapshotsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlockVolumeSnapshotsApiListBlockVolumeSnapshotsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsBlockVolumeSnapshot**](*models.BlockVolumeSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

