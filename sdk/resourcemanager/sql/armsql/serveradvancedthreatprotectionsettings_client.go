//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ServerAdvancedThreatProtectionSettingsClient contains the methods for the ServerAdvancedThreatProtectionSettings group.
// Don't use this type directly, use NewServerAdvancedThreatProtectionSettingsClient() instead.
type ServerAdvancedThreatProtectionSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerAdvancedThreatProtectionSettingsClient creates a new instance of ServerAdvancedThreatProtectionSettingsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerAdvancedThreatProtectionSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerAdvancedThreatProtectionSettingsClient, error) {
	cl, err := arm.NewClient(moduleName+".ServerAdvancedThreatProtectionSettingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerAdvancedThreatProtectionSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an Advanced Threat Protection state.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - advancedThreatProtectionName - The name of the Advanced Threat Protection state.
//   - parameters - The server Advanced Threat Protection state.
//   - options - ServerAdvancedThreatProtectionSettingsClientBeginCreateOrUpdateOptions contains the optional parameters for the
//     ServerAdvancedThreatProtectionSettingsClient.BeginCreateOrUpdate method.
func (client *ServerAdvancedThreatProtectionSettingsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, advancedThreatProtectionName AdvancedThreatProtectionName, parameters ServerAdvancedThreatProtection, options *ServerAdvancedThreatProtectionSettingsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, advancedThreatProtectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ServerAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ServerAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates an Advanced Threat Protection state.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
func (client *ServerAdvancedThreatProtectionSettingsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, advancedThreatProtectionName AdvancedThreatProtectionName, parameters ServerAdvancedThreatProtection, options *ServerAdvancedThreatProtectionSettingsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, advancedThreatProtectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerAdvancedThreatProtectionSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, advancedThreatProtectionName AdvancedThreatProtectionName, parameters ServerAdvancedThreatProtection, options *ServerAdvancedThreatProtectionSettingsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advancedThreatProtectionSettings/{advancedThreatProtectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if advancedThreatProtectionName == "" {
		return nil, errors.New("parameter advancedThreatProtectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advancedThreatProtectionName}", url.PathEscape(string(advancedThreatProtectionName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get a server's Advanced Threat Protection state.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - advancedThreatProtectionName - The name of the Advanced Threat Protection state.
//   - options - ServerAdvancedThreatProtectionSettingsClientGetOptions contains the optional parameters for the ServerAdvancedThreatProtectionSettingsClient.Get
//     method.
func (client *ServerAdvancedThreatProtectionSettingsClient) Get(ctx context.Context, resourceGroupName string, serverName string, advancedThreatProtectionName AdvancedThreatProtectionName, options *ServerAdvancedThreatProtectionSettingsClientGetOptions) (ServerAdvancedThreatProtectionSettingsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, advancedThreatProtectionName, options)
	if err != nil {
		return ServerAdvancedThreatProtectionSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerAdvancedThreatProtectionSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerAdvancedThreatProtectionSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerAdvancedThreatProtectionSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, advancedThreatProtectionName AdvancedThreatProtectionName, options *ServerAdvancedThreatProtectionSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advancedThreatProtectionSettings/{advancedThreatProtectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if advancedThreatProtectionName == "" {
		return nil, errors.New("parameter advancedThreatProtectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advancedThreatProtectionName}", url.PathEscape(string(advancedThreatProtectionName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerAdvancedThreatProtectionSettingsClient) getHandleResponse(resp *http.Response) (ServerAdvancedThreatProtectionSettingsClientGetResponse, error) {
	result := ServerAdvancedThreatProtectionSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAdvancedThreatProtection); err != nil {
		return ServerAdvancedThreatProtectionSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Get a list of the server's Advanced Threat Protection states.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - ServerAdvancedThreatProtectionSettingsClientListByServerOptions contains the optional parameters for the ServerAdvancedThreatProtectionSettingsClient.NewListByServerPager
//     method.
func (client *ServerAdvancedThreatProtectionSettingsClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerAdvancedThreatProtectionSettingsClientListByServerOptions) *runtime.Pager[ServerAdvancedThreatProtectionSettingsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerAdvancedThreatProtectionSettingsClientListByServerResponse]{
		More: func(page ServerAdvancedThreatProtectionSettingsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerAdvancedThreatProtectionSettingsClientListByServerResponse) (ServerAdvancedThreatProtectionSettingsClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServerAdvancedThreatProtectionSettingsClientListByServerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServerAdvancedThreatProtectionSettingsClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerAdvancedThreatProtectionSettingsClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerAdvancedThreatProtectionSettingsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdvancedThreatProtectionSettingsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advancedThreatProtectionSettings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerAdvancedThreatProtectionSettingsClient) listByServerHandleResponse(resp *http.Response) (ServerAdvancedThreatProtectionSettingsClientListByServerResponse, error) {
	result := ServerAdvancedThreatProtectionSettingsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogicalServerAdvancedThreatProtectionListResult); err != nil {
		return ServerAdvancedThreatProtectionSettingsClientListByServerResponse{}, err
	}
	return result, nil
}
