# Saludos en Go
Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Instalación
Ejecuta el siguiente comando para instalar el paquete:
```bash
go get -u github.com/carlap20/greetings
```

## Uso
Aquí tienes un ejemplo de cómo utilizar el paquete en tu código:

``` go
package main

import (
    "fmt"
    "github.com/carlap20/greetings"
)

func main() {
    message, err := greetings.Hello("Juan")

    if err != nil {
        fmt.Println("Ocurrió un error:", err)
        return
    }

    fmt.Println(message)
}
