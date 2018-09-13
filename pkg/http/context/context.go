package reqcontext

import (
	"context"

	"github.com/jinsenglin/prototype-go/pkg/http/session"
)

type contextKey string

const (
	userIndexKey contextKey = "UserIndex"
	sessionKey   contextKey = "Session"
)

// SetUserIndex ...
func SetUserIndex(ctx context.Context, userIndex int) context.Context {
	return context.WithValue(ctx, userIndexKey, userIndex)
}

// GetUserIndex ...
func GetUserIndex(ctx context.Context) (int, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// the string type assertion returns ok=false for nil.
	userIndex, ok := ctx.Value(userIndexKey).(int)
	return userIndex, ok
}

// SetSession ...
func SetSession(ctx context.Context, s session.Session) context.Context {
	return context.WithValue(ctx, sessionKey, s)
}

// GetSession ...
func GetSession(ctx context.Context) (session.Session, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// the string type assertion returns ok=false for nil.
	s, ok := ctx.Value(sessionKey).(session.Session)
	return s, ok
}
