package config

var ApplicationKey = "Application"

type ApplicationConfig struct {
	Environment string
	Port        int
}
