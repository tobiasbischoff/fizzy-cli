package cli

import "strings"

type multiString struct {
	values []string
}

func (m *multiString) String() string {
	return strings.Join(m.values, ",")
}

func (m *multiString) Set(value string) error {
	if value == "" {
		return nil
	}
	m.values = append(m.values, value)
	return nil
}

func (m *multiString) Values() []string {
	return append([]string(nil), m.values...)
}
