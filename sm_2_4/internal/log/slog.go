package log

import (
	"io"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func (l Logger) Err(msg string, err error) {
	l.Error(msg, slog.String("error", err.Error()))
}

func MustSetup(env string) Logger {
	var l *slog.Logger
	switch env {
	case "local":
		l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case "prod":
		f, err := os.OpenFile("logs.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("failed to create logs file")
		}

		l = slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, f), &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	default:
		panic("unknown env: " + env)
	}

	return Logger{l}
}
