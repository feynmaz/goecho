package bindings

import "github.com/labstack/echo/v4"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (lr LoginRequest) Validate(ctx echo.Context) error {
	errs := new(RequestErrors)
	if lr.Username == "" {
		errs.Append(ErrUsernameEmpty)
	}
	if lr.Password == "" {
		errs.Append(ErrPasswordEmpty)
	}
	if errs.Len() == 0 {
		return nil
	}
	return errs
}