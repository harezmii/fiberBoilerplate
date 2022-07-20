package main

import "fiberBoilerplate/pkg/app"

// @title Fiber Boilerplate API
// @version 1.0
// @description This is a fiber boilerplate server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1
// @BasePath :1453/
func main() {
	app.Run()
}
