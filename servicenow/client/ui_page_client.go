package client

import (
	"fmt"
	"encoding/json"
)

type UiPageResults struct {
	Results []UiPage `json:"records"`
}

type UiPage struct {
	Id   				string 		 `json:"sys_id,omitempty"`
	Name 				string		 `json:"name"`
	Description			string		 `json:"description"`
	Direct				bool		 `json:"direct,string"`
	Html				string		 `json:"html"`
	ProcessingScript	string 		 `json:"processing_script"`
	ClientScript 		string 		 `json:"client_script"`
	Category 			string		 `json:"category"`

	Status				string 		 `json:"__status,omitempty"`
	Error				*ErrorDetail `json:"__error,omitempty"`
}

type ErrorDetail struct {
	Reason	string `json:"reason"`
	Message	string `json:"message"`
}

func (client *ServiceNowClient) GetUiPage(id string) (*UiPage, error) {
	jsonResponse, err := client.requestJSON("GET", "sys_ui_page.do?JSONv2&sysparm_query=sys_id=" + id, nil)
	if err != nil {
		return nil, err
	}

	uiPageResults := UiPageResults{}
	if err := json.Unmarshal(jsonResponse, &uiPageResults); err != nil {
		return nil, err
	}

	if len(uiPageResults.Results) <= 0 {
		return nil, fmt.Errorf("No UI Page found using sys_id %s", id)
	}
	
	return &uiPageResults.Results[0], nil
}

func (client *ServiceNowClient) CreateUiPage(uiPage *UiPage) (page *UiPage, err error) {
	jsonResponse, err := client.requestJSON("POST", "sys_ui_page.do?JSONv2&sysparm_action=insert", uiPage)
	if err != nil {
		return
	}

	uiPageResults := UiPageResults{}
	if err = json.Unmarshal(jsonResponse, &uiPageResults); err != nil {
		return
	}

	if len(uiPageResults.Results) <= 0 {
		err = fmt.Errorf("No UI Page were inserted.")
	} else {
		page = &uiPageResults.Results[0]
		if page.Status != "success" {
			err = fmt.Errorf("Error during insert. %s: %s.", page.Error.Message, page.Error.Reason)
		}
	}

	return
}

func (client *ServiceNowClient) UpdateUiPage(uiPage *UiPage) error {
	_, err := client.requestJSON("POST", "sys_ui_page.do?JSONv2&sysparm_action=update&sysparm_query=sys_id=" + uiPage.Id, uiPage)
	return err
}


func (client *ServiceNowClient) DeleteUiPage(id string) error {
	_, err := client.requestJSON("POST", "sys_ui_page.do?JSONv2&sysparm_action=deleteRecord&sysparm_sys_id=" + id, nil)
	return err
}