// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This file enables 'go generate' to regenerate this specific SDK
//go:generate pwsh ../../../../eng/scripts/build.ps1 -goExtension "@autorest/go@4.0.0-preview.54" -skipBuild -cleanGenerated -format -tidy -generate resourcemanager/containerregistry/armcontainerregistry

package armcontainerregistry
