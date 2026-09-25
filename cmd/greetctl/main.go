package main

import (
	"flag"
	"fmt"
	"greetctl-max/internal/logx"

	"github.com/fatih/color"

	gr "greetctl-max/pkg/greet"
)

func main() {
	name := flag.String("name", "mundo", "Nombre a saludar")
	flag.Parse()

	msg := gr.Hello(*name)
	fmt.Println("Proyecto listo desde CMD")
	fmt.Println(msg)

	color.Green("Hola, %s saludos", *name)
	color.Red("Hola, %s saludos", *name)

	logx.Info("FIN todo OK")
}
