package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const dbTableLabel = "label"
const dbTableUserRole = "user_role"
const dbTableAccess = "access"
const dbTableReadAccess = "read_access"
const dbTableCreateAccess = "create_access"
const dbTableAlterAccess = "alter_access"
const dbTableDeleteAccess = "delete_access"
const dbTableWebServiceAccess = "web_service_access"
const dbTableConfigurationAccess = "configuration_access"
const dbTableExtendable = "extendable"
const dbTableLiveFeed = "live_feed"
const dbTableName = "name"

// ResourceDBTable manages a DBTable in ServiceNow.
func ResourceDBTable() *schema.Resource {
	return &schema.Resource{
		Create: createResourceDBTable,
		Read:   readResourceDBTable,
		Update: updateResourceDBTable,
		Delete: deleteResourceDBTable,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			dbTableLabel: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The display name for this table that can be localized.",
			},
			dbTableUserRole: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Role ID required for end users to access this table.",
			},
			dbTableAccess: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "public",
				Description: "Whether this Script can be accessed from only this application scope or all application scopes. Values can be 'package_private' or 'public'.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"package_private", "public"})
					return
				},
			},
			dbTableReadAccess: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Used when access is set to 'public'. Allow other application scoped to run scripts that read data from this table.",
			},
			dbTableCreateAccess: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Used when access is set to 'public'. Allow other application scopes to run scripts that create data in this table.",
			},
			dbTableAlterAccess: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Used when access is set to 'public'. Allow other application scoped to run scripts that update data in this table.",
			},
			dbTableDeleteAccess: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Used when access is set to 'public'. Allow other application scoped to run scripts that delete data in this table.",
			},
			dbTableWebServiceAccess: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Used when access is set to 'public'. Flag to determine if web service calls can be made to this table.",
			},
			dbTableConfigurationAccess: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Used when access is set to 'public'. Allow design time configuration of this table from other application scopes.",
			},
			dbTableExtendable: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Allow other tables to extend this table.",
			},
			dbTableLiveFeed: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Flag to determine if live feed should be enabled for this table.",
			},
			dbTableName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The internal name of the table.",
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceDBTable(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	dbTable := &client.DBTable{}
	if err := snowClient.GetObject(client.EndpointDBTable, data.Id(), dbTable); err != nil {
		data.SetId("")
		return err
	}

	resourceFromDBTable(data, dbTable)

	return nil
}

func createResourceDBTable(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	dbTable := resourceToDBTable(data)
	if err := snowClient.CreateObject(client.EndpointDBTable, dbTable); err != nil {
		return err
	}

	resourceFromDBTable(data, dbTable)

	return readResourceDBTable(data, serviceNowClient)
}

func updateResourceDBTable(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointDBTable, resourceToDBTable(data)); err != nil {
		return err
	}

	return readResourceDBTable(data, serviceNowClient)
}

func deleteResourceDBTable(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointDBTable, data.Id())
}

func resourceFromDBTable(data *schema.ResourceData, dbTable *client.DBTable) {
	data.SetId(dbTable.ID)
	data.Set(dbTableLabel, dbTable.Label)
	data.Set(dbTableUserRole, dbTable.UserRole)
	data.Set(dbTableAccess, dbTable.Access)
	data.Set(dbTableReadAccess, dbTable.ReadAccess)
	data.Set(dbTableCreateAccess, dbTable.CreateAccess)
	data.Set(dbTableAlterAccess, dbTable.AlterAccess)
	data.Set(dbTableDeleteAccess, dbTable.DeleteAccess)
	data.Set(dbTableWebServiceAccess, dbTable.WebServiceAccess)
	data.Set(dbTableConfigurationAccess, dbTable.ConfigurationAccess)
	data.Set(dbTableExtendable, dbTable.Extendable)
	data.Set(dbTableLiveFeed, dbTable.LiveFeed)
	data.Set(dbTableName, dbTable.Name)
	data.Set(commonScope, dbTable.Scope)
}

func resourceToDBTable(data *schema.ResourceData) *client.DBTable {
	dbTable := client.DBTable{
		Label:                data.Get(dbTableLabel).(string),
		UserRole:             data.Get(dbTableUserRole).(string),
		Access:               data.Get(dbTableAccess).(string),
		ReadAccess:           data.Get(dbTableReadAccess).(bool),
		CreateAccess:         data.Get(dbTableCreateAccess).(bool),
		AlterAccess:          data.Get(dbTableAlterAccess).(bool),
		DeleteAccess:         data.Get(dbTableDeleteAccess).(bool),
		WebServiceAccess:     data.Get(dbTableWebServiceAccess).(bool),
		ConfigurationAccess:  data.Get(dbTableConfigurationAccess).(bool),
		Extendable:           data.Get(dbTableExtendable).(bool),
		LiveFeed:             data.Get(dbTableLiveFeed).(bool),
		CreateAccessControls: true,
		CreateModule:         false,
		CreateMobileModule:   false,
	}
	dbTable.ID = data.Id()
	dbTable.Scope = data.Get(commonScope).(string)
	return &dbTable
}
