package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	console "github.com/asynkron/goconsole"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.StaticFS("/", http.Dir("./"))
	go router.Run(":9737")
	go openbrowser("http://localhost:9737")
	_, _ = console.ReadLine()

}

func openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
