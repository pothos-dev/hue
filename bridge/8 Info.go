package hue

// GetAllTimezones ...
func (b Bridge) GetAllTimezones() ([]string, error) {
	var out []string
	err := b.get("info/timezones", &out)
	return out, err
}
