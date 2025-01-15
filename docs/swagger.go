package docs

import (
	"fmt"
	"os"
)

// UpdateSwaggerHost updates the Swagger host based on configuration
func UpdateSwaggerHost(host, port string) {
	env := os.Getenv("ENV")
	prodHost := os.Getenv("DOMAIN")

	if env == "production" {
		SwaggerInfo.Host = prodHost
		SwaggerInfo.Schemes = []string{"https"}
	}else{
		SwaggerInfo.Host = fmt.Sprintf("%s:%s", host, port)
	}
}