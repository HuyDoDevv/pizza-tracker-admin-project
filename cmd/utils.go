package main

import (
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port   string
	DBPath string
}

func loadConfig() Config {
	return Config{
		Port:   getEnv("PORT", "8080"),
		DBPath: getEnv("DATABASE_URL", "./data/orders.db"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func loadTemplates(routers *gin.Engine) error {
	function := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}

	tmpl, err := template.New("").Funcs(function).ParseGlob("templates/*.tmpl")
	if err != nil {
		return err
	}

	routers.SetHTMLTemplate(tmpl)
	return nil
}
