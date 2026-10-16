package main

import (
    "fmt"
    axonsec "github.com/albertoariza522/axonsec-sdk/go"
)

func main() {
    client := axonsec.NewClient("ax_beta_2026_public")
    result, _ := client.Classify("SSH brute force detected", map[string]interface{}{
        "freq": 0.9, "urgencia": 0.85,
    })
    fmt.Printf("clasificacion: %s, bias: %.4f\n", result.Clasificacion, result.BiasFinal)
}
