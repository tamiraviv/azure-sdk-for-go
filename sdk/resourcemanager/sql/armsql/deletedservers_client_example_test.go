//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeletedServerListBySubscription.json
func ExampleDeletedServersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeletedServersClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DeletedServerListResult = armsql.DeletedServerListResult{
		// 	Value: []*armsql.DeletedServer{
		// 		{
		// 			Name: to.Ptr("sqlcrudtest-d-1414"),
		// 			Type: to.Ptr("Microsoft.Sql/deletedServers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/deletedServers/sqlcrudtest-d-1414"),
		// 			Properties: &armsql.DeletedServerProperties{
		// 				DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-15T20:20:00.345Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("sqlcrudtest-d-1414.database.windows.net"),
		// 				OriginalID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sqlcrudtest-d-1414"),
		// 				Version: to.Ptr("12.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sqlcrudtest-d-2424"),
		// 			Type: to.Ptr("Microsoft.Sql/deletedServers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/deletedServers/sqlcrudtest-d-2424"),
		// 			Properties: &armsql.DeletedServerProperties{
		// 				DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-13T10:10:00.678Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("sqlcrudtest-d-2424.database.windows.net"),
		// 				OriginalID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sqlcrudtest-d-2424"),
		// 				Version: to.Ptr("12.0"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeletedServerGet.json
func ExampleDeletedServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeletedServersClient().Get(ctx, "japaneast", "sqlcrudtest-d-1414", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeletedServer = armsql.DeletedServer{
	// 	Name: to.Ptr("sqlcrudtest-d-1414"),
	// 	Type: to.Ptr("Microsoft.Sql/deletedServers"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/deletedServers/sqlcrudtest-d-1414"),
	// 	Properties: &armsql.DeletedServerProperties{
	// 		DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-15T11:20:00.345Z"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("sqlcrudtest-d-1414.database.windows.net"),
	// 		OriginalID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sqlcrudtest-d-1414"),
	// 		Version: to.Ptr("12.0"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeletedServerList.json
func ExampleDeletedServersClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeletedServersClient().NewListByLocationPager("japaneast", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DeletedServerListResult = armsql.DeletedServerListResult{
		// 	Value: []*armsql.DeletedServer{
		// 		{
		// 			Name: to.Ptr("sqlcrudtest-d-1414"),
		// 			Type: to.Ptr("Microsoft.Sql/deletedServers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/deletedServers/sqlcrudtest-d-1414"),
		// 			Properties: &armsql.DeletedServerProperties{
		// 				DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-15T20:20:00.345Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("sqlcrudtest-d-1414.database.windows.net"),
		// 				OriginalID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sqlcrudtest-d-1414"),
		// 				Version: to.Ptr("12.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sqlcrudtest-d-2424"),
		// 			Type: to.Ptr("Microsoft.Sql/deletedServers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/deletedServers/sqlcrudtest-d-2424"),
		// 			Properties: &armsql.DeletedServerProperties{
		// 				DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-13T10:10:00.678Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("sqlcrudtest-d-2424.database.windows.net"),
		// 				OriginalID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sqlcrudtest-d-2424"),
		// 				Version: to.Ptr("12.0"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeletedServerRecover.json
func ExampleDeletedServersClient_BeginRecover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeletedServersClient().BeginRecover(ctx, "japaneast", "sqlcrudtest-d-1414", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeletedServer = armsql.DeletedServer{
	// 	Name: to.Ptr("sqlcrudtest-d-1414"),
	// 	Type: to.Ptr("Microsoft.Sql/deletedServers"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/deletedServers/sqlcrudtest-d-1414"),
	// 	Properties: &armsql.DeletedServerProperties{
	// 		DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-15T11:20:00.345Z"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("sqlcrudtest-d-1414.database.windows.net"),
	// 		OriginalID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sqlcrudtest-d-1414"),
	// 		Version: to.Ptr("12.0"),
	// 	},
	// }
}
