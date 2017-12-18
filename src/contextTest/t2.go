package main

import (
	"context"
	"fmt"
)

func t21() {
	ctx := context.WithValue(context.Background(), "key", "value2222")
	fmt.Println(ctx.Value("key").(string))
}

func main() {
	t21()
}
