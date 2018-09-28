//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

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
