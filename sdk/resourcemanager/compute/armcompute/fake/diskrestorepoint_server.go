//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"net/http"
	"net/url"
	"regexp"
)

// DiskRestorePointServer is a fake server for instances of the armcompute.DiskRestorePointClient type.
type DiskRestorePointServer struct {
	// Get is the fake for method DiskRestorePointClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *armcompute.DiskRestorePointClientGetOptions) (resp azfake.Responder[armcompute.DiskRestorePointClientGetResponse], errResp azfake.ErrorResponder)

	// BeginGrantAccess is the fake for method DiskRestorePointClient.BeginGrantAccess
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGrantAccess func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, grantAccessData armcompute.GrantAccessData, options *armcompute.DiskRestorePointClientBeginGrantAccessOptions) (resp azfake.PollerResponder[armcompute.DiskRestorePointClientGrantAccessResponse], errResp azfake.ErrorResponder)

	// NewListByRestorePointPager is the fake for method DiskRestorePointClient.NewListByRestorePointPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByRestorePointPager func(resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, options *armcompute.DiskRestorePointClientListByRestorePointOptions) (resp azfake.PagerResponder[armcompute.DiskRestorePointClientListByRestorePointResponse])

	// BeginRevokeAccess is the fake for method DiskRestorePointClient.BeginRevokeAccess
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRevokeAccess func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *armcompute.DiskRestorePointClientBeginRevokeAccessOptions) (resp azfake.PollerResponder[armcompute.DiskRestorePointClientRevokeAccessResponse], errResp azfake.ErrorResponder)
}

// NewDiskRestorePointServerTransport creates a new instance of DiskRestorePointServerTransport with the provided implementation.
// The returned DiskRestorePointServerTransport instance is connected to an instance of armcompute.DiskRestorePointClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDiskRestorePointServerTransport(srv *DiskRestorePointServer) *DiskRestorePointServerTransport {
	return &DiskRestorePointServerTransport{
		srv:                        srv,
		beginGrantAccess:           newTracker[azfake.PollerResponder[armcompute.DiskRestorePointClientGrantAccessResponse]](),
		newListByRestorePointPager: newTracker[azfake.PagerResponder[armcompute.DiskRestorePointClientListByRestorePointResponse]](),
		beginRevokeAccess:          newTracker[azfake.PollerResponder[armcompute.DiskRestorePointClientRevokeAccessResponse]](),
	}
}

// DiskRestorePointServerTransport connects instances of armcompute.DiskRestorePointClient to instances of DiskRestorePointServer.
// Don't use this type directly, use NewDiskRestorePointServerTransport instead.
type DiskRestorePointServerTransport struct {
	srv                        *DiskRestorePointServer
	beginGrantAccess           *tracker[azfake.PollerResponder[armcompute.DiskRestorePointClientGrantAccessResponse]]
	newListByRestorePointPager *tracker[azfake.PagerResponder[armcompute.DiskRestorePointClientListByRestorePointResponse]]
	beginRevokeAccess          *tracker[azfake.PollerResponder[armcompute.DiskRestorePointClientRevokeAccessResponse]]
}

// Do implements the policy.Transporter interface for DiskRestorePointServerTransport.
func (d *DiskRestorePointServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DiskRestorePointClient.Get":
		resp, err = d.dispatchGet(req)
	case "DiskRestorePointClient.BeginGrantAccess":
		resp, err = d.dispatchBeginGrantAccess(req)
	case "DiskRestorePointClient.NewListByRestorePointPager":
		resp, err = d.dispatchNewListByRestorePointPager(req)
	case "DiskRestorePointClient.BeginRevokeAccess":
		resp, err = d.dispatchBeginRevokeAccess(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DiskRestorePointServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/restorePointCollections/(?P<restorePointCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorePoints/(?P<vmRestorePointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/diskRestorePoints/(?P<diskRestorePointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	restorePointCollectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointCollectionName")])
	if err != nil {
		return nil, err
	}
	vmRestorePointNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmRestorePointName")])
	if err != nil {
		return nil, err
	}
	diskRestorePointNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskRestorePointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameUnescaped, restorePointCollectionNameUnescaped, vmRestorePointNameUnescaped, diskRestorePointNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DiskRestorePoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DiskRestorePointServerTransport) dispatchBeginGrantAccess(req *http.Request) (*http.Response, error) {
	if d.srv.BeginGrantAccess == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGrantAccess not implemented")}
	}
	beginGrantAccess := d.beginGrantAccess.get(req)
	if beginGrantAccess == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/restorePointCollections/(?P<restorePointCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorePoints/(?P<vmRestorePointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/diskRestorePoints/(?P<diskRestorePointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/beginGetAccess`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.GrantAccessData](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		restorePointCollectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointCollectionName")])
		if err != nil {
			return nil, err
		}
		vmRestorePointNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmRestorePointName")])
		if err != nil {
			return nil, err
		}
		diskRestorePointNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskRestorePointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginGrantAccess(req.Context(), resourceGroupNameUnescaped, restorePointCollectionNameUnescaped, vmRestorePointNameUnescaped, diskRestorePointNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGrantAccess = &respr
		d.beginGrantAccess.add(req, beginGrantAccess)
	}

	resp, err := server.PollerResponderNext(beginGrantAccess, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginGrantAccess.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGrantAccess) {
		d.beginGrantAccess.remove(req)
	}

	return resp, nil
}

func (d *DiskRestorePointServerTransport) dispatchNewListByRestorePointPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByRestorePointPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByRestorePointPager not implemented")}
	}
	newListByRestorePointPager := d.newListByRestorePointPager.get(req)
	if newListByRestorePointPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/restorePointCollections/(?P<restorePointCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorePoints/(?P<vmRestorePointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/diskRestorePoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		restorePointCollectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointCollectionName")])
		if err != nil {
			return nil, err
		}
		vmRestorePointNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmRestorePointName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByRestorePointPager(resourceGroupNameUnescaped, restorePointCollectionNameUnescaped, vmRestorePointNameUnescaped, nil)
		newListByRestorePointPager = &resp
		d.newListByRestorePointPager.add(req, newListByRestorePointPager)
		server.PagerResponderInjectNextLinks(newListByRestorePointPager, req, func(page *armcompute.DiskRestorePointClientListByRestorePointResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByRestorePointPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByRestorePointPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByRestorePointPager) {
		d.newListByRestorePointPager.remove(req)
	}
	return resp, nil
}

func (d *DiskRestorePointServerTransport) dispatchBeginRevokeAccess(req *http.Request) (*http.Response, error) {
	if d.srv.BeginRevokeAccess == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRevokeAccess not implemented")}
	}
	beginRevokeAccess := d.beginRevokeAccess.get(req)
	if beginRevokeAccess == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/restorePointCollections/(?P<restorePointCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorePoints/(?P<vmRestorePointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/diskRestorePoints/(?P<diskRestorePointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endGetAccess`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		restorePointCollectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointCollectionName")])
		if err != nil {
			return nil, err
		}
		vmRestorePointNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmRestorePointName")])
		if err != nil {
			return nil, err
		}
		diskRestorePointNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskRestorePointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginRevokeAccess(req.Context(), resourceGroupNameUnescaped, restorePointCollectionNameUnescaped, vmRestorePointNameUnescaped, diskRestorePointNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRevokeAccess = &respr
		d.beginRevokeAccess.add(req, beginRevokeAccess)
	}

	resp, err := server.PollerResponderNext(beginRevokeAccess, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginRevokeAccess.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRevokeAccess) {
		d.beginRevokeAccess.remove(req)
	}

	return resp, nil
}
