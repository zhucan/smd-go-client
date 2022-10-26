# \ObjectGatewaysApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetObjectGateway**](ObjectGatewaysApi.md#GetObjectGateway) | **Get** /object-gateways/{object-gateway-id} | get an object gateway
[**ListObjectGateways**](ObjectGatewaysApi.md#ListObjectGateways) | **Get** /object-gateways | list all object gateways


# **GetObjectGateway**
> ModelsObjectGateway GetObjectGateway(ctx, objectGatewayId)
get an object gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectGatewayId** | **int64**| identifier of the object gateway | 

### Return type

[**ModelsObjectGateway**](models.ObjectGateway.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListObjectGateways**
> []ModelsObjectGateway ListObjectGateways(ctx, optional)
list all object gateways

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ObjectGatewaysApiListObjectGatewaysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ObjectGatewaysApiListObjectGatewaysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsObjectGateway**](*models.ObjectGateway.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

