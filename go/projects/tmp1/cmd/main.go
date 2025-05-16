package main

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    ctx, done := signal.NotifyContext(
        context.Background(), os.Interrupt, syscall.SIGTERM)
    defer done()

    if err := run(ctx, logger); err != nil {
        logger.Error("Run application failed", "run", err)
    }
}

func run(ctx context.Context, *slog.Logger) error {
    <-ctx.Done()
    return nil
}
