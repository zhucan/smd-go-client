# \OsdsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOsds**](OsdsApi.md#CreateOsds) | **Post** /osds | create osds
[**DeleteOsds**](OsdsApi.md#DeleteOsds) | **Delete** /osds | delete osds from specific ceph cluster
[**GetOsd**](OsdsApi.md#GetOsd) | **Get** /osds/{osd-id} | get a osd
[**ListOsds**](OsdsApi.md#ListOsds) | **Get** /osds | get all osds
[**UpdateOsd**](OsdsApi.md#UpdateOsd) | **Post** /osds/{osd-id} | Update osd


# **CreateOsds**
> []ModelsOsd CreateOsds(ctx, body)
create osds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesOsdsRequestParams**](RoutesOsdsRequestParams.md)|  | 

### Return type

[**[]ModelsOsd**](*models.Osd.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOsds**
> []ModelsOsd DeleteOsds(ctx, body)
delete osds from specific ceph cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesOsdsDeleteParams**](RoutesOsdsDeleteParams.md)|  | 

### Return type

[**[]ModelsOsd**](*models.Osd.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOsd**
> ModelsOsd GetOsd(ctx, osdId)
get a osd

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **osdId** | **int64**| identifier of the osd | 

### Return type

[**ModelsOsd**](models.Osd.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOsds**
> []ModelsOsd ListOsds(ctx, optional)
get all osds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OsdsApiListOsdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OsdsApiListOsdsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 
 **hostId** | **optional.Int64**| identifier of the host | 

### Return type

[**[]ModelsOsd**](*models.Osd.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOsd**
> ModelsOsd UpdateOsd(ctx, osdId, body)
Update osd

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **osdId** | **int64**| identifier of the osd | 
  **body** | [**RoutesOsdsUpdateParams**](RoutesOsdsUpdateParams.md)|  | 

### Return type

[**ModelsOsd**](models.Osd.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

