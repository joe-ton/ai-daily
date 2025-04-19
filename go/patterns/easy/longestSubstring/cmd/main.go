package main

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    ctx, stop := signal.NotifyContext(
        context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    if 

}
