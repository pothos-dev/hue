package hue

// RuleID ...
type RuleID string

// Rule ...
type Rule struct {
	Name           string
	lastTriggered  string
	CreationTime   string
	TimesTriggered int
	Owner          string
	Status         string
	Conditions     []struct {
		Address  string
		Operator string
		Value    string
	}
	Actions []struct {
		Address string
		Method  string
		Body    struct {
			Scene string
		}
	}
}
