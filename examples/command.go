package main

import (
	"fmt"
	"github.com/namsral/envconfig"
	"os"
	"reflect"
)

type Specification struct {
	Debug bool    `env:",optional"`
	Port  int     `env:""`
	Rate  float32 `env:",optional"`
	User  string  `env:"USER"`
}

func main() {
	var spec Specification
	os.Clearenv()
	os.Setenv("DEBUG", "true")
	os.Setenv("PORT", "8080")
	os.Setenv("RATE", "0.5")
	os.Setenv("USER", "Kelsey")

	err := envconfig.Process(&spec)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	s := reflect.ValueOf(&spec).Elem()
	typeOfSpec := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.CanSet() {
			fieldName := typeOfSpec.Field(i).Name
			fmt.Printf("%s: %v (%s)\n", fieldName, f.Interface(), f.Kind())
		}
	}
}
