package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

// ref: https://github.com/spf13/pflag/issues/236#issuecomment-931600452
type enum struct {
	Allowed []string
	Value   string
}

func NewEnum(allowed []string, defaultValue string) *enum {
	return &enum{
		Allowed: allowed,
		Value:   defaultValue,
	}
}

// Set implements pflag.Value.
func (e *enum) Set(s string) error {
	for _, v := range e.Allowed {
		if v == s {
			e.Value = s
			return nil
		}
	}
	return fmt.Errorf("%s is not in [%s]", s, strings.Join(e.Allowed, ","))
}

// String implements pflag.Value.
func (e *enum) String() string {
	return e.Value
}

// Type implements pflag.Value.
func (e *enum) Type() string {
	return "string"
}

var _ pflag.Value = (*enum)(nil)
