//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsphere

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ProductsClient contains the methods for the Products group.
// Don't use this type directly, use NewProductsClient() instead.
type ProductsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProductsClient creates a new instance of ProductsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProductsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProductsClient, error) {
	cl, err := arm.NewClient(moduleName+".ProductsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProductsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CountDevices - Counts devices in product. '.default' and '.unassigned' are system defined values and cannot be used for
// product name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - options - ProductsClientCountDevicesOptions contains the optional parameters for the ProductsClient.CountDevices method.
func (client *ProductsClient) CountDevices(ctx context.Context, resourceGroupName string, catalogName string, productName string, options *ProductsClientCountDevicesOptions) (ProductsClientCountDevicesResponse, error) {
	req, err := client.countDevicesCreateRequest(ctx, resourceGroupName, catalogName, productName, options)
	if err != nil {
		return ProductsClientCountDevicesResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProductsClientCountDevicesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProductsClientCountDevicesResponse{}, runtime.NewResponseError(resp)
	}
	return client.countDevicesHandleResponse(resp)
}

// countDevicesCreateRequest creates the CountDevices request.
func (client *ProductsClient) countDevicesCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, options *ProductsClientCountDevicesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/countDevices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// countDevicesHandleResponse handles the CountDevices response.
func (client *ProductsClient) countDevicesHandleResponse(resp *http.Response) (ProductsClientCountDevicesResponse, error) {
	result := ProductsClientCountDevicesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CountDeviceResponse); err != nil {
		return ProductsClientCountDevicesResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create a Product. '.default' and '.unassigned' are system defined values and cannot be used for product
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - resource - Resource create parameters.
//   - options - ProductsClientBeginCreateOrUpdateOptions contains the optional parameters for the ProductsClient.BeginCreateOrUpdate
//     method.
func (client *ProductsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, catalogName string, productName string, resource Product, options *ProductsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ProductsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, catalogName, productName, resource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProductsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ProductsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a Product. '.default' and '.unassigned' are system defined values and cannot be used for product
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *ProductsClient) createOrUpdate(ctx context.Context, resourceGroupName string, catalogName string, productName string, resource Product, options *ProductsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, catalogName, productName, resource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProductsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, resource Product, options *ProductsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// BeginDelete - Delete a Product. '.default' and '.unassigned' are system defined values and cannot be used for product name'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - options - ProductsClientBeginDeleteOptions contains the optional parameters for the ProductsClient.BeginDelete method.
func (client *ProductsClient) BeginDelete(ctx context.Context, resourceGroupName string, catalogName string, productName string, options *ProductsClientBeginDeleteOptions) (*runtime.Poller[ProductsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, catalogName, productName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProductsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ProductsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a Product. '.default' and '.unassigned' are system defined values and cannot be used for product name'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *ProductsClient) deleteOperation(ctx context.Context, resourceGroupName string, catalogName string, productName string, options *ProductsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, catalogName, productName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProductsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, options *ProductsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewGenerateDefaultDeviceGroupsPager - Generates default device groups for the product. '.default' and '.unassigned' are
// system defined values and cannot be used for product name.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - options - ProductsClientGenerateDefaultDeviceGroupsOptions contains the optional parameters for the ProductsClient.NewGenerateDefaultDeviceGroupsPager
//     method.
func (client *ProductsClient) NewGenerateDefaultDeviceGroupsPager(resourceGroupName string, catalogName string, productName string, options *ProductsClientGenerateDefaultDeviceGroupsOptions) *runtime.Pager[ProductsClientGenerateDefaultDeviceGroupsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProductsClientGenerateDefaultDeviceGroupsResponse]{
		More: func(page ProductsClientGenerateDefaultDeviceGroupsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProductsClientGenerateDefaultDeviceGroupsResponse) (ProductsClientGenerateDefaultDeviceGroupsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.generateDefaultDeviceGroupsCreateRequest(ctx, resourceGroupName, catalogName, productName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProductsClientGenerateDefaultDeviceGroupsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ProductsClientGenerateDefaultDeviceGroupsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProductsClientGenerateDefaultDeviceGroupsResponse{}, runtime.NewResponseError(resp)
			}
			return client.generateDefaultDeviceGroupsHandleResponse(resp)
		},
	})
}

// generateDefaultDeviceGroupsCreateRequest creates the GenerateDefaultDeviceGroups request.
func (client *ProductsClient) generateDefaultDeviceGroupsCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, options *ProductsClientGenerateDefaultDeviceGroupsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/generateDefaultDeviceGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// generateDefaultDeviceGroupsHandleResponse handles the GenerateDefaultDeviceGroups response.
func (client *ProductsClient) generateDefaultDeviceGroupsHandleResponse(resp *http.Response) (ProductsClientGenerateDefaultDeviceGroupsResponse, error) {
	result := ProductsClientGenerateDefaultDeviceGroupsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceGroupListResult); err != nil {
		return ProductsClientGenerateDefaultDeviceGroupsResponse{}, err
	}
	return result, nil
}

// Get - Get a Product. '.default' and '.unassigned' are system defined values and cannot be used for product name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - options - ProductsClientGetOptions contains the optional parameters for the ProductsClient.Get method.
func (client *ProductsClient) Get(ctx context.Context, resourceGroupName string, catalogName string, productName string, options *ProductsClientGetOptions) (ProductsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, catalogName, productName, options)
	if err != nil {
		return ProductsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProductsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProductsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProductsClient) getCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, options *ProductsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProductsClient) getHandleResponse(resp *http.Response) (ProductsClientGetResponse, error) {
	result := ProductsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Product); err != nil {
		return ProductsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByCatalogPager - List Product resources by Catalog
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - options - ProductsClientListByCatalogOptions contains the optional parameters for the ProductsClient.NewListByCatalogPager
//     method.
func (client *ProductsClient) NewListByCatalogPager(resourceGroupName string, catalogName string, options *ProductsClientListByCatalogOptions) *runtime.Pager[ProductsClientListByCatalogResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProductsClientListByCatalogResponse]{
		More: func(page ProductsClientListByCatalogResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProductsClientListByCatalogResponse) (ProductsClientListByCatalogResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByCatalogCreateRequest(ctx, resourceGroupName, catalogName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProductsClientListByCatalogResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ProductsClientListByCatalogResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProductsClientListByCatalogResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByCatalogHandleResponse(resp)
		},
	})
}

// listByCatalogCreateRequest creates the ListByCatalog request.
func (client *ProductsClient) listByCatalogCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, options *ProductsClientListByCatalogOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByCatalogHandleResponse handles the ListByCatalog response.
func (client *ProductsClient) listByCatalogHandleResponse(resp *http.Response) (ProductsClientListByCatalogResponse, error) {
	result := ProductsClientListByCatalogResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductListResult); err != nil {
		return ProductsClientListByCatalogResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a Product. '.default' and '.unassigned' are system defined values and cannot be used for product name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - properties - The resource properties to be updated.
//   - options - ProductsClientBeginUpdateOptions contains the optional parameters for the ProductsClient.BeginUpdate method.
func (client *ProductsClient) BeginUpdate(ctx context.Context, resourceGroupName string, catalogName string, productName string, properties ProductUpdate, options *ProductsClientBeginUpdateOptions) (*runtime.Poller[ProductsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, catalogName, productName, properties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProductsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ProductsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update a Product. '.default' and '.unassigned' are system defined values and cannot be used for product name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *ProductsClient) update(ctx context.Context, resourceGroupName string, catalogName string, productName string, properties ProductUpdate, options *ProductsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, catalogName, productName, properties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ProductsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, properties ProductUpdate, options *ProductsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}
