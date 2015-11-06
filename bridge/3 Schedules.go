package hue

// ScheduleID ...
type ScheduleID string

// Schedule ...
type Schedule struct {
	Name        string
	Description string
	Command     struct {
		Address string
		Body    struct {
			Scene string
		}
		method string
	}
	Time       string
	Created    string
	Status     string
	AudoDelete bool
	StartTime  string
}

// GetAllSchedules ...
func GetAllSchedules() (map[ScheduleID]Schedule, error) {
	panic("TODO")
}

// CreateSchedule ...
func CreateSchedule(attributes map[string]interface{}) (ScheduleID, error) {
	panic("TODO")
}

// GetScheduleAttributes ...
func GetScheduleAttributes(id ScheduleID) (Schedule, error) {
	panic("TODO")
}

// SetScheduleAttributes ...
func SetScheduleAttributes(id ScheduleID, attributes map[string]interface{}) error {
	panic("TODO")
}

// DeleteSchedule ...
func DeleteSchedule(id ScheduleID) error {
	panic("TODO")
}
