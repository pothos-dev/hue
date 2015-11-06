package hue

// SceneID ...
type SceneID string

// Scene ...
type Scene struct {
	Name   string
	Lights []LightID
	Active bool
}

// GetAllScenes ...
func (b Bridge) GetAllScenes() (map[SceneID]Scene, error) {
	panic("TODO")
}

// CreateScene ...
func (b Bridge) CreateScene(id SceneID) (Scene, error) {
	panic("TODO")
}

// ModifyScene ...
func (b Bridge) ModifyScene(id SceneID, attributes map[string]interface{}) error {
	panic("TODO")
}

// RecallScene ...
func (b Bridge) RecallScene(id SceneID) error {
	panic("TODO")
}

// DeleteScene ...
func (b Bridge) DeleteScene(id SceneID) error {
	panic("TODO")
}
