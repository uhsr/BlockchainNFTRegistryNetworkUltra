// cmd/blockchainnftregistrynetworkultra/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftregistrynetworkultra/internal/blockchainnftregistrynetworkultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftregistrynetworkultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
