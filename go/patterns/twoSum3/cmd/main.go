package main

func main() {
    logger := slog.New(slog.newJSONHandler(os.Stdout, nil))

    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

}

func initTelemetry() (trace *tracer.Telemetry) (err error) {

}
