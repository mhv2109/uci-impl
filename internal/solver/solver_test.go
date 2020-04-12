package solver

import "testing"

func TestSpinOptionToString(t *testing.T) {
	option := Option{}
	option.Name = "testoption"
	option.Type = OptionSpinType
	option.Default = "50"
	option.Min = "0"
	option.Max = "100"

	if a, e := option.String(),
		"option name testoption type spin default 50 min 0 max 100"; a != e {
		t.Errorf("Option.String() Expected: %s, Actual: %s", e, a)
	}
}

func TestStringOptionToString(t *testing.T) {
	option := Option{}
	option.Name = "testoption"
	option.Type = OptionStringType
	option.Default = "testdefault"

	if a, e := option.String(),
		"option name testoption type string default testdefault"; a != e {
		t.Errorf("Option.String() Expected: %s, Actual: %s", e, a)
	}
}
