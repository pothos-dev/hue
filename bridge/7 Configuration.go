package hue

import "fmt"

// Configuration ...
type Configuration struct {
	Name string

	Zigbeechannel int
	Bridgeid      string
	Mac           string
	Dhcp          bool
	Ipaddress     string
	Netmask       string
	Gateway       string
	Proxyaddress  string
	Proxyport     int

	Utc       string
	Localtime string
	Timezone  string

	Whitelist map[string]struct {
		LastUseDate string
		CreateDate  string
		Name        string
	}

	Modelid    string
	Swversion  string
	Apiversion string
	Swupdate   struct {
		Updatestate    int
		Checkforupdate bool
		Devicetypes    struct {
			Bridge  bool
			Lights  []interface{}
			Sensors []interface{}
		}
		URL    string
		Text   string
		Notify bool
	}

	Linkbutton bool

	Portalservices   bool
	Portalconnection string
	Portalstate      struct {
		Signedon      bool
		Incoming      bool
		Outgoing      bool
		Communication string
	}
}

// UserID ...
type UserID string

// CreateUser ...
func (b Bridge) CreateUser() (string, error) {
	var reply struct{ Username string }
	err := b.post("", map[string]string{"devicetype": "go_application"}, &reply)
	return reply.Username, err
}

// GetConfiguration ...
func (b Bridge) GetConfiguration() (Configuration, error) {
	var config Configuration
	err := b.get("config", &config)
	return config, err
}

// ModifyConfiguration ...
func (b Bridge) ModifyConfiguration(config map[string]interface{}) error {
	return b.put("config", config, nil)
}

// DeleteUser ...
func (b Bridge) DeleteUser(id UserID) error {
	return b.delete(fmt.Sprintf("config/whitelist/%s", id))
}
