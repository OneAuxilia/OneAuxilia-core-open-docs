package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mvrilo/go-redoc"
	fiberredoc "github.com/mvrilo/go-redoc/fiber"
)

func main() {
	doc := redoc.Redoc{
		Title:       "OneAuxilia API Document",
		Description: "OneAuxilia API Description",
		SpecFile:    "./docs/oneauxilia.json",
		SpecPath:    "/oneauxilia.json",
		DocsPath:    "/docs",
	}

	r := fiber.New()
	r.Use(fiberredoc.New(doc))
	godotenv.Load()
	serAdd := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	println(fmt.Sprintf("Documentation served at %s/docs", serAdd))
	panic(r.Listen(serAdd))
}
