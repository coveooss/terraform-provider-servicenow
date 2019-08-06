package resources_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/coveooss/terraform-provider-servicenow/servicenow/resources"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mock.Mock
}

func (m *ClientMock) GetObject(endpoint string, id string, responseObjectOut client.Record) error {
	args := m.Called(endpoint, id, responseObjectOut)
	return args.Error(0)
}

func (m *ClientMock) GetObjectByName(endpoint string, name string, responseObjectOut client.Record) error {
	args := m.Called(endpoint, name, responseObjectOut)
	return args.Error(0)
}

func (m *ClientMock) CreateObject(endpoint string, record client.Record) error {
	args := m.Called(endpoint, record)
	return args.Error(0)
}

func (m *ClientMock) UpdateObject(endpoint string, record client.Record) error {
	args := m.Called(endpoint, record)
	return args.Error(0)
}

func (m *ClientMock) DeleteObject(endpoint string, id string) error {
	args := m.Called(endpoint, id)
	return args.Error(0)
}

type RecordMock struct {
	mock.Mock
}

func (m *RecordMock) GetID() string {
	args := m.Called()
	return args.String(0)
}

func (m *RecordMock) GetScope() string {
	args := m.Called()
	return args.String(0)
}

func (m *RecordMock) GetStatus() string {
	args := m.Called()
	return args.String(0)
}

func (m *RecordMock) GetError() string {
	args := m.Called()
	return args.String(0)
}

var resourcesToTest = []*schema.Resource{
	resources.ResourceApplication(),
	resources.ResourceApplicationMenu(),
	resources.ResourceApplicationModule(),
	resources.ResourceContentCSS(),
	resources.ResourceCSSInclude(),
	resources.ResourceCSSIncludeRelation(),
	resources.ResourceDBTable(),
	resources.ResourceExtensionPoint(),
	resources.ResourceJsInclude(),
	resources.ResourceJsIncludeRelation(),
	resources.ResourceOAuthEntity(),
	resources.ResourceRole(),
	resources.ResourceRestMessage(),
	resources.ResourceRestMessageHeader(),
	resources.ResourceRestMethod(),
	resources.ResourceRestMethodHeader(),
	resources.ResourceScriptInclude(),
	resources.ResourceSystemProperty(),
	resources.ResourceSystemPropertyCategory(),
	resources.ResourceSystemPropertyRelation(),
	resources.ResourceUIMacro(),
	resources.ResourceUIPage(),
	resources.ResourceUIScript(),
	resources.ResourceWidget(),
	resources.ResourceWidgetDependency(),
	resources.ResourceWidgetDependencyRelation(),
}

var dataSourcesToTest = []*schema.Resource{
	resources.DataSourceApplication(),
	resources.DataSourceApplicationCategory(),
	resources.DataSourceDBTable(),
	resources.DataSourceRole(),
	resources.DataSourceSystemProperty(),
	resources.DataSourceSystemPropertyCategory(),
}

func TestResourcesCanRead(t *testing.T) {
	for _, res := range resourcesToTest {
		data := schema.ResourceData{}
		data.SetId("hello")
		clientMock := new(ClientMock)
		clientMock.
			On("GetObject", mock.AnythingOfType("string"), "hello", mock.Anything).
			Return(nil)

		res.Read(&data, clientMock)
		clientMock.AssertExpectations(t)
	}
}

func TestResourceRestMessageHandleReadError(t *testing.T) {
	for _, res := range resourcesToTest {
		data := schema.ResourceData{}
		data.SetId("hello")
		clientMock := new(ClientMock)
		clientMock.
			On("GetObject", mock.AnythingOfType("string"), "hello", mock.Anything).
			Return(fmt.Errorf("nothing to see here"))

		res.Read(&data, clientMock)
		clientMock.AssertExpectations(t)
		assert.Equal(t, "", data.Id())
	}
}

func TestDataSourcesCanRead(t *testing.T) {
	for _, res := range dataSourcesToTest {
		data := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{
			"name": "oi",
		})

		clientMock := new(ClientMock)
		clientMock.
			On("GetObjectByName", mock.AnythingOfType("string"), "oi", mock.Anything).
			Return(nil)

		res.Read(data, clientMock)
		clientMock.AssertExpectations(t)
	}
}

func TestResourcesCanUpdate(t *testing.T) {
	for _, res := range resourcesToTest {
		fakeData := map[string]interface{}{}
		// Fill the data with stuff following the schema.
		for key, prop := range res.Schema {
			if !prop.Computed {
				switch prop.Type {
				case schema.TypeString:
					fakeData[key] = "hello"
				case schema.TypeBool:
					fakeData[key] = true
				case schema.TypeInt:
					fakeData[key] = 42
				}
			}
		}

		data := schema.TestResourceDataRaw(t, res.Schema, fakeData)
		data.SetId("fenouille")

		clientMock := new(ClientMock)
		clientMock.
			On("UpdateObject", mock.AnythingOfType("string"), mock.Anything).
			Return(nil)
		clientMock.
			On("GetObject", mock.AnythingOfType("string"), "fenouille", mock.Anything).
			Return(nil)

		res.Update(data, clientMock)
		clientMock.AssertExpectations(t)
	}
}

func TestResourcesCanDelete(t *testing.T) {
	for _, res := range resourcesToTest {
		data := schema.ResourceData{}
		data.SetId("fenouille")
		clientMock := new(ClientMock)
		clientMock.
			On("DeleteObject", mock.AnythingOfType("string"), "fenouille").
			Return(nil)

		res.Delete(&data, clientMock)
		clientMock.AssertExpectations(t)
	}
}
