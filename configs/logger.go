package configs

import (
	"context"
	"fmt"
	"log"
	"path"
	"runtime"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
)

type ctxKey int

const loggerKey ctxKey = iota

const (
	LogInfo  string = "INFO"
	LogError string = "ERROR"
	LogDebug string = "DEBUG"
)

func WithLogger(ctx context.Context) context.Context {
	requestID := middleware.GetReqID(ctx)
	if requestID == "" {
		requestID = "no-request-id"
	}

	prefix := fmt.Sprintf("[%s] ", requestID)
	logger := log.New(log.Writer(), prefix, log.LstdFlags)

	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) *log.Logger {
	if logger, ok := ctx.Value(loggerKey).(*log.Logger); ok {
		return logger
	}
	return log.Default()
}

func Logger(ctx context.Context, level string, msg string, args ...any) {
	level = strings.ToUpper(level)

	pc, _, _, ok := runtime.Caller(1)
	funcName := "unknown"

	if ok {
		fullFuncName := runtime.FuncForPC(pc).Name()
		fullFuncName = strings.ReplaceAll(fullFuncName, "(*", "")
		fullFuncName = strings.ReplaceAll(fullFuncName, ")", "")

		funcName = path.Base(fullFuncName)
		parts := strings.Split(funcName, ".")
		n := len(parts)
		if n >= 3 {
			funcName = fmt.Sprintf("Package: %s - Layer: %s - Function: %s", parts[n-3], parts[n-2], parts[n-1])
		}
	}

	requestID := middleware.GetReqID(ctx)
	if requestID == "" {
		requestID = "no-request-id"
	}

	logMsg := fmt.Sprintf("[%s - Request: %s - %s] %s", level, requestID, funcName, fmt.Sprintf(msg, args...))
	log.Println(logMsg)
}
