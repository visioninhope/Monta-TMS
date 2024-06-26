package util

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LogFromContext returns a request-specific zerolog instance using the provided context.
// The returned logger will have the request ID as well as some other value predefined.
// If no logger is associated with the context provided, the global zerolog instance
// will be returned instead - this function will _always_ return a valid (enabled) logger.
// Should you ever need to force a disabled logger for a context, use `util.DisableLogger(ctx, true)`
// and pass the context returned to other code/`LogFromContext`.
func LogFromContext(ctx context.Context) *zerolog.Logger {
	l := log.Ctx(ctx)
	if l.GetLevel() == zerolog.Disabled {
		if ShouldDisableLogger(ctx) {
			return l
		}
		l = &log.Logger
	}
	return l
}

// LogFromFiberContext returns a request-specific zerolog instance using the fiber.Ctx of the request.
// The returned logger will have the request ID as well as some other value predefined.
func LogFromFiberContext(c *fiber.Ctx) *zerolog.Logger {
	return LogFromContext(c.UserContext())
}

func LogLevelFromString(s string) zerolog.Level {
	l, err := zerolog.ParseLevel(s)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to parse log level, defaulting to %s", zerolog.DebugLevel)
		return zerolog.DebugLevel
	}

	return l
}
