package docs

import "fmt"

// UpdateSwaggerHost updates the Swagger host based on configuration
func UpdateSwaggerHost(host, port string) {
	// Update the host directly in SwaggerInfo
	SwaggerInfo.Host = fmt.Sprintf("%s:%s", host, port)
}
