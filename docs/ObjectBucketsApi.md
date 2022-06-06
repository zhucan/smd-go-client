# \ObjectBucketsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectBucket**](ObjectBucketsApi.md#CreateObjectBucket) | **Post** /object-buckets | create an object bucket
[**DeleteObjectBucket**](ObjectBucketsApi.md#DeleteObjectBucket) | **Delete** /object-buckets/{object-bucket-id} | delete an object bucket
[**GetObjectBucket**](ObjectBucketsApi.md#GetObjectBucket) | **Get** /object-buckets/{object-bucket-id} | get an object bucket
[**ListObjectBuckets**](ObjectBucketsApi.md#ListObjectBuckets) | **Get** /object-buckets | get all object buckets


# **CreateObjectBucket**
> ModelsObjectBucket CreateObjectBucket(ctx, body)
create an object bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesObjectBucketRequestParams**](RoutesObjectBucketRequestParams.md)|  | 

### Return type

[**ModelsObjectBucket**](models.ObjectBucket.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteObjectBucket**
> ModelsObjectBucket DeleteObjectBucket(ctx, objectBucketId)
delete an object bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectBucketId** | **int32**| identifier of the object bucket | 

### Return type

[**ModelsObjectBucket**](models.ObjectBucket.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectBucket**
> ModelsObjectBucket GetObjectBucket(ctx, objectBucketId)
get an object bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectBucketId** | **int32**| identifier of the object bucket | 

### Return type

[**ModelsObjectBucket**](models.ObjectBucket.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListObjectBuckets**
> []ModelsObjectBucket ListObjectBuckets(ctx, optional)
get all object buckets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ObjectBucketsApiListObjectBucketsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ObjectBucketsApiListObjectBucketsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| limit | 

### Return type

[**[]ModelsObjectBucket**](*models.ObjectBucket.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

