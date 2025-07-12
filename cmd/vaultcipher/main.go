// cmd/vaultcipher/main.go
package main

import (
"flag"
"log"
"os"

"vaultcipher/internal/vaultcipher"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := vaultcipher.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
