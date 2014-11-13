// Copyright (c) 2013 Kelsey Hightower. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package envconfig

import (
	"os"
	"testing"
)

type Specification struct {
	Debug                        bool    `env:",optional"`
	Port                         int     `env:",optional"`
	Rate                         float32 `env:",optional"`
	User                         string  `env:",optional"`
	MultiWordVar                 string  `env:",optional"`
	MultiWordVarWithAlt          string  `env:"MULTI_WORD_VAR_WITH_ALT,optional"`
	MultiWordVarWithLowerCaseAlt string  `env:"multi_word_var_with_lower_case_alt,optional"`
}

func TestProcess(t *testing.T) {
	var s Specification
	os.Clearenv()
	os.Setenv("DEBUG", "true")
	os.Setenv("PORT", "8080")
	os.Setenv("RATE", "0.5")
	os.Setenv("USER", "Kelsey")
	err := Process(&s)
	if err != nil {
		t.Error(err.Error())
	}
	if !s.Debug {
		t.Errorf("expected %v, got %v", true, s.Debug)
	}
	if s.Port != 8080 {
		t.Errorf("expected %d, got %v", 8080, s.Port)
	}
	if s.Rate != 0.5 {
		t.Errorf("expected %f, got %v", 0.5, s.Rate)
	}
	if s.User != "Kelsey" {
		t.Errorf("expected %s, got %s", "Kelsey", s.User)
	}
}

func TestParseErrorBool(t *testing.T) {
	var s Specification
	os.Clearenv()
	os.Setenv("DEBUG", "string")
	err := Process(&s)
	v, ok := err.(*ParseError)
	if !ok {
		t.Errorf("expected ParseError, got %v", v)
	}
	if v.FieldName != "Debug" {
		t.Errorf("expected %s, got %v", "Debug", v.FieldName)
	}
	if s.Debug != false {
		t.Errorf("expected %v, got %v", false, s.Debug)
	}
}

func TestParseErrorFloat32(t *testing.T) {
	var s Specification
	os.Clearenv()
	os.Setenv("RATE", "string")
	err := Process(&s)
	v, ok := err.(*ParseError)
	if !ok {
		t.Errorf("expected ParseError, got %v", v)
	}
	if v.FieldName != "Rate" {
		t.Errorf("expected %s, got %v", "Rate", v.FieldName)
	}
	if s.Rate != 0 {
		t.Errorf("expected %v, got %v", 0, s.Rate)
	}
}

func TestParseErrorInt(t *testing.T) {
	var s Specification
	os.Clearenv()
	os.Setenv("PORT", "string")
	err := Process(&s)
	v, ok := err.(*ParseError)
	if !ok {
		t.Errorf("expected ParseError, got %v", v)
	}
	if v.FieldName != "Port" {
		t.Errorf("expected %s, got %v", "Port", v.FieldName)
	}
	if s.Port != 0 {
		t.Errorf("expected %v, got %v", 0, s.Port)
	}
}

func TestErrInvalidSpecification(t *testing.T) {
	m := make(map[string]string)
	err := Process(&m)
	if err != ErrInvalidSpecification {
		t.Errorf("expected %v, got %v", ErrInvalidSpecification, err)
	}
}

func TestAlternateVarNames(t *testing.T) {
	var s Specification
	os.Clearenv()
	os.Setenv("MULTI_WORD_VAR", "foo")
	os.Setenv("MULTI_WORD_VAR_WITH_ALT", "bar")
	os.Setenv("MULTI_WORD_VAR_WITH_LOWER_CASE_ALT", "baz")
	if err := Process(&s); err != nil {
		t.Error(err.Error())
	}

	// Setting the alt version of the var in the environment has no effect if
	// the struct tag is not supplied
	if s.MultiWordVar != "" {
		t.Errorf("expected %q, got %q", "", s.MultiWordVar)
	}

	// Setting the alt version of the var in the environment correctly sets
	// the value if the struct tag IS supplied
	if s.MultiWordVarWithAlt != "bar" {
		t.Errorf("expected %q, got %q", "bar", s.MultiWordVarWithAlt)
	}

	// Alt value is not case sensitive and is treated as all uppercase
	if s.MultiWordVarWithLowerCaseAlt != "baz" {
		t.Errorf("expected %q, got %q", "baz", s.MultiWordVarWithLowerCaseAlt)
	}
}

func TestMandatory(t *testing.T) {
	type Spec struct {
		Mandatory string `env:"MANDATORY"`
	}
	var s Spec
	os.Clearenv()
	os.Setenv("MANDATORY", "")

	err := Process(&s)
	if _, ok := err.(*EmptyMandatoryFieldError); !ok {
		t.Errorf("expected EmptyMandatoryFieldError, got %#v", err)
	}
}
