# \CrushRootsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCrushRoot**](CrushRootsApi.md#CreateCrushRoot) | **Post** /crush-roots | create crush root
[**DeleteCrushRoot**](CrushRootsApi.md#DeleteCrushRoot) | **Delete** /crush-roots/{crush-root-id} | delete a crush root
[**GetCrushRoot**](CrushRootsApi.md#GetCrushRoot) | **Get** /crush-roots/{crush-root-id} | get a crush root
[**ListCrushRoots**](CrushRootsApi.md#ListCrushRoots) | **Get** /crush-roots | get all crush roots


# **CreateCrushRoot**
> ModelsCrushRoot CreateCrushRoot(ctx, body)
create crush root

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesCrushRootRequestParams**](RoutesCrushRootRequestParams.md)|  | 

### Return type

[**ModelsCrushRoot**](models.CrushRoot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCrushRoot**
> ModelsCrushRoot DeleteCrushRoot(ctx, crushRootId, optional)
delete a crush root

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crushRootId** | **int32**| identifier of the crush root | 
 **optional** | ***CrushRootsApiDeleteCrushRootOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrushRootsApiDeleteCrushRootOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| force option will ignore errors | 

### Return type

[**ModelsCrushRoot**](models.CrushRoot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCrushRoot**
> ModelsCrushRoot GetCrushRoot(ctx, crushRootId)
get a crush root

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crushRootId** | **int32**| identifier of the crush root | 

### Return type

[**ModelsCrushRoot**](models.CrushRoot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCrushRoots**
> []ModelsCrushRoot ListCrushRoots(ctx, optional)
get all crush roots

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CrushRootsApiListCrushRootsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrushRootsApiListCrushRootsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| limit | 

### Return type

[**[]ModelsCrushRoot**](*models.CrushRoot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

