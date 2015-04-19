package main
import (
    "fmt"
)

func main() {
    performers, err := LoadPerformers()
    if err != nil {
        fmt.Println(err)
        return
    }

    for _, ps := range performers {
        fmt.Println(ps.Name)
    }
}