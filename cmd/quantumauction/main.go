// cmd/quantumauction/main.go
package main

import (
"flag"
"log"
"os"

"quantumauction/internal/quantumauction"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := quantumauction.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
