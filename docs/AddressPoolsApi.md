# \AddressPoolsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAddressPool**](AddressPoolsApi.md#CreateAddressPool) | **Post** /address-pools | create an address pool
[**DeleteAddressPool**](AddressPoolsApi.md#DeleteAddressPool) | **Delete** /address-pools/{address-pool-id} | delete an address pool
[**GetAddressPool**](AddressPoolsApi.md#GetAddressPool) | **Get** /address-pools/{address-pool-id} | get an address pool
[**ListAddressPools**](AddressPoolsApi.md#ListAddressPools) | **Get** /address-pools | list all address pools


# **CreateAddressPool**
> ModelsAddressPool CreateAddressPool(ctx, body)
create an address pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesCreateAddressPoolRequestParams**](RoutesCreateAddressPoolRequestParams.md)|  | 

### Return type

[**ModelsAddressPool**](models.AddressPool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAddressPool**
> ModelsAddressPool DeleteAddressPool(ctx, addressPoolId, optional)
delete an address pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addressPoolId** | **int64**| identifier of the address pool | 
 **optional** | ***AddressPoolsApiDeleteAddressPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressPoolsApiDeleteAddressPoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| force option will ignore errors | 

### Return type

[**ModelsAddressPool**](models.AddressPool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressPool**
> ModelsAddressPool GetAddressPool(ctx, addressPoolId)
get an address pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addressPoolId** | **int64**| identifier of the address pool | 

### Return type

[**ModelsAddressPool**](models.AddressPool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAddressPools**
> []ModelsAddressPool ListAddressPools(ctx, optional)
list all address pools

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddressPoolsApiListAddressPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressPoolsApiListAddressPoolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsAddressPool**](*models.AddressPool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

