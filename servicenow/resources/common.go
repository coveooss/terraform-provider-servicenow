package resources

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

const commonProtectionPolicy = "protection_policy"

func getProtectionPolicySchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "read",
		Description: "Determines how application files are protected when downloaded or installed. Can be empty for no protection, 'read' for read-only protection or 'protected'.",
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
