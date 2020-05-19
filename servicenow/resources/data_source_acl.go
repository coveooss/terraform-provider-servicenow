package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const aclType = "type"
const aclOperation = "operation"
const aclAdminOverrides = "admin_overrides"
const aclName = "name"
const aclDescription = "description"
const aclActive = "active"
const aclAdvanced = "advanced"
const aclCondition = "condition"
const aclScript = "script"

// DataSourceACL reads the informations about a single ACL in ServiceNow.
func DataSourceACL() *schema.Resource {
	resourceSchema := map[string]*schema.Schema{
		aclName: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Enter the name of the object being secured, either the record name or the table and field names.",
		},
		aclType: {
			Type:        schema.TypeString,
			Description: "Select what kind of object this ACL rule secures.",
			Computed:    true,
		},
		aclOperation: {
			Type:        schema.TypeString,
			Description: "Select the operation this ACL rule secures.",
			Computed:    true,
		},
		aclAdminOverrides: {
			Type:        schema.TypeBool,
			Description: "Users with admin override this rule",
			Computed:    true,
		},
		aclDescription: {
			Type:        schema.TypeString,
			Description: "Enter a description of the object or permissions this ACL rule secures.",
			Computed:    true,
		},
		aclActive: {
			Type:        schema.TypeBool,
			Description: "Activates the ACL rule.",
			Computed:    true,
		},
		aclAdvanced: {
			Type:        schema.TypeBool,
			Description: "Displays the Script field when active.",
			Computed:    true,
		},
		aclCondition: {
			Type:        schema.TypeString,
			Description: "Selects the fields and values that must be true for users to access the object.",
			Computed:    true,
		},
		aclScript: {
			Type:        schema.TypeString,
			Description: "Custom script describing the permissions required to access the object.",
			Computed:    true,
		},
		commonProtectionPolicy: getProtectionPolicySchema(),
		commonScope:            getScopeSchema(),
	}

	setOnlyRequiredSchema(resourceSchema, aclName)

	return &schema.Resource{
		Read:   readDataSourceACL,
		Schema: resourceSchema,
	}
}

func readDataSourceACL(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	acl := &client.ACL{}
	if err := snowClient.GetObjectByName(client.EndpointACL, data.Get(aclName).(string), acl); err != nil {
		data.SetId("")
		return err
	}

	resourceFromACL(data, acl)

	return nil
}

func resourceFromACL(data *schema.ResourceData, acl *client.ACL) {
	data.SetId(acl.ID)
	data.Set(aclType, acl.Type)
	data.Set(aclOperation, acl.Operation)
	data.Set(aclAdminOverrides, acl.AdminOverrides)
	data.Set(aclName, acl.Name)
	data.Set(aclDescription, acl.Description)
	data.Set(aclActive, acl.Active)
	data.Set(aclAdvanced, acl.Advanced)
	data.Set(aclCondition, acl.Condition)
	data.Set(aclScript, acl.Script)
	data.Set(commonProtectionPolicy, acl.ProtectionPolicy)
	data.Set(commonScope, acl.Scope)
}
