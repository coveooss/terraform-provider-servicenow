package resources

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

const commonProtectionPolicy = "protection_policy"
const commonScope = "scope"

func getProtectionPolicySchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "read",
		Description: "Determines how application files are protected when downloaded or installed. Can be empty for no protection, 'read' for read-only protection or 'protected'.",
	}
}

func getScopeSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "global",
		ForceNew:    true,
		Description: "Associates a resource to a specific application ID in ServiceNow.",
	}
}

// setOnlyRequiredSchema Changes required parameters. For data sources, only one attribute is normally required and everything else is computed.
func setOnlyRequiredSchema(schema map[string]*schema.Schema, requiredName string) {
	for key, val := range schema {
		val.Computed = true
		val.Required = false
		val.Optional = false
		val.ForceNew = false
		val.Default = nil

		if key == requiredName {
			val.Computed = false
			val.Required = true
		}
	}
}

func validateStringValue(actual string, key string, expectedValues []string) (warns []string, errs []error) {
	for _, expected := range expectedValues {
		if actual == expected {
			return
		}
	}
	var message = ""
	for i, expected := range expectedValues {
		if i != 0 {
			message += " or "
		}
		message += "'" + expected + "'"
	}
	errs = append(errs, fmt.Errorf("%q must be %s, got: %s", key, message, actual))
	return
}
