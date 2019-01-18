package client

// EndpointWidget is the endpoint to manage widget records.
const EndpointWidget = "sp_widget.do"

// Widget is the json response for a Widget in ServiceNow.
type Widget struct {
	BaseResult
	CustomID     string `json:"id"`
	Name         string `json:"name"`
	Template     string `json:"template"`
	CSS          string `json:"css"`
	Public       bool   `json:"public,string"`
	Roles        string `json:"roles"`
	Link         string `json:"link"`
	Description  string `json:"description"`
	ClientScript string `json:"client_script"`
	ServerScript string `json:"script"`
	DemoData     string `json:"demo_data"`
	OptionSchema string `json:"option_schema"`
	HasPreview   bool   `json:"has_preview,string"`
	DataTable    string `json:"data_table"`
	ControllerAs string `json:"controller_as"`
}
