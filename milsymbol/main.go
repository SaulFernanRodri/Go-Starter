package main

import (
	"io/ioutil"
	"log"

	"github.com/dop251/goja"
)

func main() {

	jsFile, err := ioutil.ReadFile("assets/milsymbol.js")
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
        // Símbolo básico
        var symbol1 = new ms.Symbol("SFGPUCI---*****", { size: 100 });
        var svg1 = symbol1.asSVG();

        // Símbolo con etiquetas (label) y un borde diferente
        var symbol2 = new ms.Symbol("SFGPUCI---*****", {
            size: 100,
            fill: "#FF0000", // Relleno en rojo
            strokeWidth: 2,  // Grosor del borde
            strokeColor: "#000000", // Color del borde negro
            labels: {
                a: "Label A",
                b: "Label B"
            }
        });
        var svg2 = symbol2.asSVG();

        // Símbolo con diferentes colores
        var symbol3 = new ms.Symbol("SFGPUCI---*****", {
            size: 100,
            fill: "#00FF00",  // Relleno en verde
            frame: true,
            strokeColor: "#0000FF", // Color del borde azul
            strokeWidth: 1
        });
        var svg3 = symbol3.asSVG();

        // Combinar todos los símbolos en un solo string SVG
        var combinedSVG = svg1 + svg2 + svg3;
        combinedSVG;
    `

	// Ejecutar el script en goja
	result, err := vm.RunString(script)
	if err != nil {
		log.Fatalf("Error al ejecutar el script: %v", err)
	}

	svgData := result.String()
	err = ioutil.WriteFile("combined_symbols.svg", []byte(svgData), 0644)
	if err != nil {
		log.Fatalf("Error al guardar el archivo SVG: %v", err)
	}
}
