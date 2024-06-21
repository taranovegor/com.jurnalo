package config

import "fmt"

type Http struct {
	Port int
}

func (http Http) Addr() string {
	return fmt.Sprintf(":%d", http.Port)
}
