package main

import (
	_ "soedirman-digital-school/pertemuan_1"
	_ "soedirman-digital-school/pertemuan_2"
	_ "soedirman-digital-school/pertemuan_3"

	"github.com/gofiber/fiber/v2"
)

type RequestData struct {
	JariJariLingkaran float64 `json:"jari_jari_lingkaran"`
	SisiPersegi       int     `json:"sisi_persegi"`
	AlasSegitiga      int     `json:"alas_segitiga"`
	TinggiSegitiga    int     `json:"tinggi_segitiga"`
}

type ResponseData struct {
	LuasLingkaran     float64 `json:"luas_lingkaran"`
	LuasPersegi       int     `json:"luas_persegi"`
	LuasSegitiga      float64 `json:"luas_segitiga"`
	KelilingLingkaran float64 `json:"keliling_lingkaran"`
	KelilingPersegi   int     `json:"keliling_persegi"`
	KelilingSegitiga  int     `json:"keliling_segitiga"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/calculate", func(c *fiber.Ctx) error {
		request := new(RequestData)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		response := new(ResponseData)
		response.LuasLingkaran = 3.14 * request.JariJariLingkaran * request.JariJariLingkaran
		response.LuasPersegi = request.SisiPersegi * request.SisiPersegi
		response.LuasSegitiga = 0.5 * float64(request.AlasSegitiga) * float64(request.TinggiSegitiga)
		response.KelilingLingkaran = 2 * 3.14 * request.JariJariLingkaran
		response.KelilingPersegi = 4 * request.SisiPersegi
		response.KelilingSegitiga = 3 * request.AlasSegitiga

		return c.JSON(response)
	})

	app.Listen(":3000")
}
