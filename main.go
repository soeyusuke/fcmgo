package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()
	title := "abc123"
	body := "cba321"

	if err := SendPushMessage(ctx, title, body); err != nil {
		fmt.Fprintf(os.Stderr, "send message err: %v", err)
		os.Exit(1)
	}
}
