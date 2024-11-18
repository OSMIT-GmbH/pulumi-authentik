// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// Optional HTTP headers sent with every request
func GetHeaders(ctx *pulumi.Context) string {
	return config.Get(ctx, "authentik:headers")
}

// Whether to skip TLS verification, can optionally be passed as `AUTHENTIK_INSECURE` environmental variable
func GetInsecure(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "authentik:insecure")
}

// The authentik API token, can optionally be passed as `AUTHENTIK_TOKEN` environmental variable
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "authentik:token")
}

// The authentik API endpoint, can optionally be passed as `AUTHENTIK_URL` environmental variable
func GetUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "authentik:url")
}
