package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/OSMIT-GmbH/terraform-provider-authentik/internal/provider"
)

// fix provider.Provider here to match whats in internal/provider
func Provider(version string, testing bool) *schema.Provider {
	return provider.Provider(version, testing)
}
