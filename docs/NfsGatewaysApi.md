# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNFSGateway**](NfsGatewaysApi.md#GetNFSGateway) | **Get** /nfs-gateways/{nfs-gateway-id} | get a nfs gateway
[**ListNFSGateways**](NfsGatewaysApi.md#ListNFSGateways) | **Get** /nfs-gateways | list all nfs gateways

# **GetNFSGateway**
> ModelsNfsGateway GetNFSGateway(ctx, nfsGatewayId)
get a nfs gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nfsGatewayId** | **int64**| identifier of the nfs gateway | 

### Return type

[**ModelsNfsGateway**](models.NFSGateway.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNFSGateways**
> []ModelsNfsGateway ListNFSGateways(ctx, optional)
list all nfs gateways

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NfsGatewaysApiListNFSGatewaysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NfsGatewaysApiListNFSGatewaysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsNfsGateway**](*models.NFSGateway.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

