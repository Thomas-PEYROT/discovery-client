# Gateway

This a library of the [Go microservices project](https://github.com/Thomas-PEYROT/go-microservices-architecture).
It is used to create more easily microservices, by providing methods to register/unregister on the [discovery server](https://github.com/Thomas-PEYROT/discovery-server).

## Setup

First of all, import it in your project (in `go.mod`):

```
require github.com/Thomas-PEYROT/discovery-client v1.0.2
```

and then use `go mod tidy` to download. After that, you can use the `RegisterMicroservice` and `UnregisterMicroservice`
in your code. Here is a simple implementation with a Gin server (see the full code [here](https://github.com/Thomas-PEYROT/microservice-a)) :

```go
package main

import (
	"fmt"
	discovery "github.com/Thomas-PEYROT/discovery-client"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Registering the microservice
	discovery.RegisterMicroservice("my-service")

	// Creation of Gin router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world !")
	})

	// Start server in goroutinr
	port := discovery.ServiceInformations.Port
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error while launching server : %v", err)
		}
	}()

	// Listening system signals to execute code before closing
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Waiting for close signal
	<-quit
	log.Println("Stopping microservice...")

	// Unregistering microservice
	discovery.UnregisterMicroservice()
	log.Println("Microservice correctly unregistered.")
}
```