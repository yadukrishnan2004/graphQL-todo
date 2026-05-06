package middleware

import "context"

func GetUserID(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value(UserCtxKey).(uint)
	return userID, ok
}