package resolver

import (
	"context"
	"errors"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("not found")
	ErrForbidden    = errors.New("forbidden")
	ErrInternal     = errors.New("internal server error")
)

func GetUserIDByContext(ctx context.Context) (int, error) {
	for _, key := range []any{"user_id", "userId"} {
		if val := ctx.Value(key); val != nil {
			switch v := val.(type) {
			case int:
				return v, nil
			case uint:
				return int(v), nil
			case int64:
				return int(v), nil
			case float64:
				return int(v), nil
			}
		}
	}
	return 0, ErrUnauthorized
}

func GetRoleByContext(ctx context.Context) (string, error) {
	for _, key := range []any{"user_role", "role", "userRole"} {
		if val := ctx.Value(key); val != nil {
			if role, ok := val.(string); ok {
				return role, nil
			}
		}
	}
	return "", ErrUnauthorized
}

func IsAdminFromContext(ctx context.Context) bool {
	role, err := GetRoleByContext(ctx)
	if err != nil {
		return false
	}
	return role == "admin"
}
