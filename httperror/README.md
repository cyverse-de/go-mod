# httperror

Handles some of the boilerplate logic needed for error handling with Labstack
Echo so that Echo errors can be treated like normal Go errors and will also be
marshalled into a standardized JSON format.

## Examples

Setting up the custom error handler in Echo:

```go
package main

import (
    "github.com/labstack/echo/v4"
    "github.com/cyverse-de/go-mod/httperror"
)

type App struct{
    router *echo.Echo
}

func New() *App {
    r := echo.New()

    // This is the important bit.
    r.HTTPErrorHandler = httperror.HTTPErrorHandler

    return &App{
        router: echo.New(),
    }
}

...
```

Raising an error from an Echo HTTP handler:

```go
func (a *App) ContrivedExampleHandler(c echo.Context) error {
    c := c.Param("id")
    if id == "" {
        // With the above error handler registered, echo Errors can be returned
        // like normal errors and they get marshalled into a consistent JSON
        // format.
        return echo.NewHTTPError(http.StatusBadRequest, errors.New("id was not set"))
    }
    ...do stuff...
}

```
