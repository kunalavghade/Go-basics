package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n\n%s\n\n", err, debug.Stack())

	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) getUserFromCtx(ctx context.Context) (*User, error) {
	u, ok := ctx.Value(contextAuthUser).(*User)
	if !ok || u == nil {
		return nil, errors.New("user not found in context")
	}
	return u, nil
}
