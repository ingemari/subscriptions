package logs

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogger() *slog.Logger {
	// Файловый вывод
	fileWriter := &lumberjack.Logger{
		Filename: "logs/app.log",
		MaxSize:  100,
	}

	// Мультиплексор
	multiWriter := io.MultiWriter(os.Stdout, fileWriter)

	handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	return slog.New(handler)
}
