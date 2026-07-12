package logger

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"
)

type Level int

const (
	LevelInfo  Level = 0
	LevelWarn  Level = 1
	LevelError Level = 2
)

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	}
	return "unknown"
}

type Logger struct {
	mu     sync.Mutex
	writer io.Writer
	level  Level
	fields map[string]any
}

func New(writer io.Writer, level Level) *Logger {
	if writer == nil {
		writer = os.Stdout
	}
	return &Logger{
		writer: writer,
		level:  level,
		fields: make(map[string]any),
	}
}

func (l *Logger) With(key string, value any) *Logger {
	fields := make(map[string]any, len(l.fields)+1)
	for k, v := range l.fields {
		fields[k] = v
	}
	fields[key] = value
	return &Logger{
		writer: l.writer,
		level:  l.level,
		fields: fields,
	}
}

func (l *Logger) log(level Level, msg string, fields map[string]any) {
	if level < l.level {
		return
	}
	entry := make(map[string]any, 3+len(l.fields)+len(fields))
	entry["level"] = level.String()
	entry["message"] = msg
	entry["timestamp"] = time.Now().UTC().Format(time.RFC3339)
	for k, v := range l.fields {
		entry[k] = v
	}
	for k, v := range fields {
		entry[k] = v
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	json.NewEncoder(l.writer).Encode(entry)
}

func (l *Logger) Info(msg string, fields map[string]any) {
	l.log(LevelInfo, msg, fields)
}

func (l *Logger) Warn(msg string, fields map[string]any) {
	l.log(LevelWarn, msg, fields)
}

func (l *Logger) Error(msg string, fields map[string]any) {
	l.log(LevelError, msg, fields)
}
