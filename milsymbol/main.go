package main

import (
	"log"
	"os"

	"github.com/dop251/goja"
)

func main() {

	jsFile, err := os.ReadFile("./milsymbol/assets/milsymbol.js")
	if err != nil {
		log.Fatalf("Error al leer el archivo JavaScript: %v", err)
	}

	vm := goja.New()

	//Cargamos la libreria
	_, err = vm.RunString(string(jsFile))
	if err != nil {
		log.Fatalf("Error al ejecutar milsymbol: %v", err)
	}

	script := `
        var symbol1 = new ms.Symbol("SFGPUCI---*****", { size: 100 });
        var svg1 = symbol1.asSVG();
    `

	// Ejecutar el script en goja
	result, err := vm.RunString(script)
	if err != nil {
		log.Fatalf("Error al ejecutar el script: %v", err)
	}

	svgData := result.String()
	err = os.WriteFile("assets/output_symbols.svg", []byte(svgData), 0644)
	if err != nil {
		log.Fatalf("Error al guardar el archivo SVG: %v", err)
	}
}
