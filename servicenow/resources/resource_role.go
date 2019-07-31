package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const roleSuffix = "suffix"
const roleDescription = "description"
const roleElevatedPrivilege = "elevated_privilege"
const roleAssignableBy = "assignable_by"
const roleName = "name"

// ResourceRole manages a Role in ServiceNow.
func ResourceRole() *schema.Resource {
	return &schema.Resource{
		Create: createResourceRole,
		Read:   readResourceRole,
		Update: updateResourceRole,
		Delete: deleteResourceRole,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			roleSuffix: {
				Type:     schema.TypeString,
				Required: true,
			},
			roleDescription: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			roleElevatedPrivilege: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			roleAssignableBy: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			roleName: {
				Type:     schema.TypeString,
				Computed: true,
			},
			commonProtectionPolicy: getProtectionPolicySchema(),
			commonScope:            getScopeSchema(),
		},
	}
}

func readResourceRole(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	role := &client.Role{}
	if err := snowClient.GetObject(client.EndpointRole, data.Id(), role); err != nil {
		data.SetId("")
		return err
	}

	resourceFromRole(data, role)

	return nil
}

func createResourceRole(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	role := resourceToRole(data)
	if err := snowClient.CreateObject(client.EndpointRole, role); err != nil {
		return err
	}

	resourceFromRole(data, role)

	return readResourceRole(data, serviceNowClient)
}

func updateResourceRole(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointRole, resourceToRole(data)); err != nil {
		return err
	}

	return readResourceRole(data, serviceNowClient)
}

func deleteResourceRole(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointRole, data.Id())
}

func resourceFromRole(data *schema.ResourceData, role *client.Role) {
	data.SetId(role.ID)
	data.Set(roleDescription, role.Description)
	data.Set(roleSuffix, role.Suffix)
	data.Set(roleElevatedPrivilege, role.ElevatedPrivilege)
	data.Set(roleAssignableBy, role.AssignableBy)
	data.Set(roleName, role.Name)
	data.Set(commonProtectionPolicy, role.ProtectionPolicy)
	data.Set(commonScope, role.Scope)
}

func resourceToRole(data *schema.ResourceData) *client.Role {
	role := client.Role{
		Suffix:            data.Get(roleSuffix).(string),
		Description:       data.Get(roleDescription).(string),
		ElevatedPrivilege: data.Get(roleElevatedPrivilege).(bool),
		AssignableBy:      data.Get(roleAssignableBy).(string),
	}
	role.ID = data.Id()
	role.ProtectionPolicy = data.Get(commonProtectionPolicy).(string)
	role.Scope = data.Get(commonScope).(string)
	return &role
}
