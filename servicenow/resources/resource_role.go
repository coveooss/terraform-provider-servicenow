package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
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
		},
	}
}

func readResourceRole(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	role, err := client.GetRole(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromRole(data, role)

	return nil
}

func createResourceRole(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdRole, err := client.CreateRole(resourceToRole(data))
	if err != nil {
		return err
	}

	resourceFromRole(data, createdRole)

	return readResourceRole(data, serviceNowClient)
}

func updateResourceRole(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateRole(resourceToRole(data))
	if err != nil {
		return err
	}

	return readResourceRole(data, serviceNowClient)
}

func deleteResourceRole(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteRole(data.Id())
}

func resourceFromRole(data *schema.ResourceData, role *client.Role) {
	data.SetId(role.Id)
	data.Set(roleDescription, role.Description)
	data.Set(roleSuffix, role.Suffix)
	data.Set(roleElevatedPrivilege, role.ElevatedPrivilege)
	data.Set(roleAssignableBy, role.AssignableBy)
	data.Set(roleName, role.Name)
}

func resourceToRole(data *schema.ResourceData) *client.Role {
	role := client.Role{
		Suffix:            data.Get(roleSuffix).(string),
		Description:       data.Get(roleDescription).(string),
		ElevatedPrivilege: data.Get(roleElevatedPrivilege).(bool),
		AssignableBy:      data.Get(roleAssignableBy).(string),
	}
	role.Id = data.Id()
	return &role
}
