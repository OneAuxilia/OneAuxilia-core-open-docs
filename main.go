package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mvrilo/go-redoc"
	fiberredoc "github.com/mvrilo/go-redoc/fiber"
)

func main() {
	doc := redoc.Redoc{
		Title:       "Logarin API Document",
		Description: "Logarin API Description",
		SpecFile:    "./docs/swagger.json",
		SpecPath:    "/swagger.json",
		DocsPath:    "/docs",
	}

	r := fiber.New()
	r.Use(fiberredoc.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Listen(":8000"))
}
