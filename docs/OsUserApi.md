# \OsUserApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOsUser**](OsUserApi.md#CreateOsUser) | **Post** /os-users | create os user
[**DeleteOsUser**](OsUserApi.md#DeleteOsUser) | **Delete** /os-users/{os-user-id} | delete os user
[**GetOsUser**](OsUserApi.md#GetOsUser) | **Get** /os-users/{os-user-id} | get an os user
[**ListOsUsers**](OsUserApi.md#ListOsUsers) | **Get** /os-users | list all kind of os users


# **CreateOsUser**
> ModelsOsUser CreateOsUser(ctx, body)
create os user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutesCreateOsUserRequestParams**](RoutesCreateOsUserRequestParams.md)|  | 

### Return type

[**ModelsOsUser**](models.OsUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOsUser**
> ModelsOsUser DeleteOsUser(ctx, osUserId)
delete os user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **osUserId** | **int64**| identifier of the os user | 

### Return type

[**ModelsOsUser**](models.OsUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOsUser**
> ModelsOsUser GetOsUser(ctx, osUserId)
get an os user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **osUserId** | **int64**| identifier of the os user | 

### Return type

[**ModelsOsUser**](models.OsUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOsUsers**
> []ModelsOsUser ListOsUsers(ctx, optional)
list all kind of os users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OsUserApiListOsUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OsUserApiListOsUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsOsUser**](*models.OsUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

