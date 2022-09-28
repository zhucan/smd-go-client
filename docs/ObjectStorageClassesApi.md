# \ObjectStorageClassesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectStorageClass**](ObjectStorageClassesApi.md#CreateObjectStorageClass) | **Post** /object-storage-classes | create an object storage class
[**DeleteObjectStorageClass**](ObjectStorageClassesApi.md#DeleteObjectStorageClass) | **Delete** /object-storage-classes/{object-storage-class-id} | delete an object storage class
[**GetObjectStorageClass**](ObjectStorageClassesApi.md#GetObjectStorageClass) | **Get** /object-storage-classes/{object-storage-class-id} | get an object storage classs
[**ListObjectStorageClasses**](ObjectStorageClassesApi.md#ListObjectStorageClasses) | **Get** /object-storage-classes | get all object storage classes


# **CreateObjectStorageClass**
> ModelsObjectStorageClass CreateObjectStorageClass(ctx, body)
create an object storage class

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesObjectStorageClassRequestParams**](RoutesObjectStorageClassRequestParams.md)|  | 

### Return type

[**ModelsObjectStorageClass**](models.ObjectStorageClass.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteObjectStorageClass**
> ModelsObjectStorageClass DeleteObjectStorageClass(ctx, objectStorageClassId)
delete an object storage class

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectStorageClassId** | **int64**| identifier of the object storage class | 

### Return type

[**ModelsObjectStorageClass**](models.ObjectStorageClass.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectStorageClass**
> ModelsObjectStorageClass GetObjectStorageClass(ctx, objectStorageClassId)
get an object storage classs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectStorageClassId** | **int64**| identifier of the object storage class | 

### Return type

[**ModelsObjectStorageClass**](models.ObjectStorageClass.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListObjectStorageClasses**
> []ModelsObjectStorageClass ListObjectStorageClasses(ctx, optional)
get all object storage classes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ObjectStorageClassesApiListObjectStorageClassesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ObjectStorageClassesApiListObjectStorageClassesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsObjectStorageClass**](*models.ObjectStorageClass.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

