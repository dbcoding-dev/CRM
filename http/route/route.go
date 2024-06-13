package route

import (
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/cors"
	"goyave.dev/goyave/v5/middleware/parse"
)

func Register(server *goyave.Server, router *goyave.Router) {
	router.CORS(cors.Default())
	router.GlobalMiddleware(&parse.Middleware{})
	router.Get("/", func(response *goyave.Response, request *goyave.Request) {
		response.String(200, "Welcome to Goyave!")
	})
}
