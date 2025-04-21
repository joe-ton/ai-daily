package main

import "log/slog"

func main() {
	logger := slog.New(slog.NewJSONHandler(handler))
}
