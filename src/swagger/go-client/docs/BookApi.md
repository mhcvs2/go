# \BookApi

All URIs are relative to *http://localhost/book*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBook**](BookApi.md#AddBook) | **Post** /insert | Add a new book to the store
[**DeleteBook**](BookApi.md#DeleteBook) | **Post** /delete | Add a new book to the store
[**FindBookByStatus**](BookApi.md#FindBookByStatus) | **Get** /list | Finds Book by status


# **AddBook**
> AddBook(ctx, book)
Add a new book to the store



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **book** | [**Book**](Book.md)| Pet object that needs to be added to the store | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBook**
> DeleteBook(ctx, book)
Add a new book to the store



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **book** | [**Book**](Book.md)| Pet object that needs to be added to the store | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindBookByStatus**
> FindBookByStatus(ctx, state)
Finds Book by status

Multiple status values can be provided with comma separated strings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **state** | [**[]string**](string.md)| Status values that need to be considered for filter | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

