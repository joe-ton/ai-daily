package main

import "fmt"

func main() {
    ctx := context.Background()

    if err := run(ctx); err != nil {
        fmt.Println!(os.Stderr)
    }
}
