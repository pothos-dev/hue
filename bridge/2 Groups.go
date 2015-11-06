package hue

// GroupID ...
type GroupID string

// Group ...
type Group struct {
	Name   string
	Lights []LightID
	Type   string
	Action LightState
}

// GetAllGroups ...
func (b Bridge) GetAllGroups() (map[GroupID]Group, error) {
	panic("TODO")
}

// CreateGroup ...
func (b Bridge) CreateGroup(name string, lights []LightID) (GroupID, error) {
	panic("TODO")
}

// GetGroupAttributes ...
func (b Bridge) GetGroupAttributes(id GroupID) (Group, error) {
	panic("TODO")
}

// SetGroupAttributes ...
func (b Bridge) SetGroupAttributes(id GroupID, attributes map[string]interface{}) error {
	panic("TODO")
}

// SetGroupState ...
func (b Bridge) SetGroupState(id GroupID, lightstate map[string]interface{}) error {
	panic("TODO")
}

// DeleteGroup ...
func (b Bridge) DeleteGroup(id GroupID) error {
	panic("TODO")
}
