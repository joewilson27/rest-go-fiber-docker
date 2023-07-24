package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joewilson27/rest-go-fiber-docker/database"
	"github.com/joewilson27/rest-go-fiber-docker/handlers"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,         // new config
		ViewsLayout: "layouts/main", // add this to config. Global layout (ex: header, navigation, menus)
	})

	setUpRoutes(app)

	// Serving static assets (untuk file static: image, etc.)
	app.Static("/", "./public") // you know it, public is pointing folder that file is accessible
	/*
		app.Static method to tell our app where to locate our static assets. In our case,
		we will be putting our images, CSS and JS in a folder called public.

		jadi nanti semua static files di folder public, tinggal panggil nama filenya saja dan
		jika berada di dalam folder, nama folder ditulis... misal javascript/app.js
		dimana app.js berada di dalam folder javascript yang ada di folder public
	*/

	// Set up 404 page
	app.Use(handlers.NotFound)
	/*
		Then we can go and update our app config in cmd/main.go to ALSO use
		the NotFound handler. Remember to place this after the other setup, but
		RIGHT BEFORE the call to app.Listen
	*/

	app.Listen(":3000")
}
