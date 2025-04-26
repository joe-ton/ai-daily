package main

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    ctx, done := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer done()

    tp, err := initTelemetry()

    if err := run()
}

func initTelemetry() (*trace.TracerProvider, error) {
    exporter, err := 
}
