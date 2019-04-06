package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	APPHASHKEY = "hash key string"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var (
	templateFunc = template.FuncMap{
		"timeToString": func(t time.Time) string {
			return t.Format("2006/01/02 15:04")
		},
	}
	t = &Template{
		templates: template.Must(template.New("t").Funcs(templateFunc).ParseGlob("templates/**.gohtml")),
	}
)

func main() {
	e := echo.New()
	e.Renderer = t
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/signs", signs)
	e.POST("/signup", signup)
	e.POST("/signin", signin)
	e.Logger.Fatal(e.Start(":8080"))
}

/*
	↓↓↓↓↓↓↓↓↓↓↓↓↓
	# controllerやadmin / usersなどとし下層に移動する
*/

type User struct {
	ID      int
	isLogin bool

	Name     string
	Email    string
	Password string

	CreatedAt time.Time
}

func signs(c echo.Context) error {
	return c.Render(http.StatusOK, "signs", nil)
}

func signup(c echo.Context) error {
	var u User
	if err := c.Bind(&u); err != nil {
		c.Logger().Error("bind error")
	}

	// check unique, set db
	// do_something()

	// send Email, wait retry login

	return c.Render(http.StatusOK, "signup", nil)
}

func signin(c echo.Context) error {
	var u User
	if err := c.Bind(&u); err != nil {
		c.Logger().Error("bind error")
	}

	// inquiry db, name & password
	ok := true
	if !ok {
		c.Logger().Error("login error")
		return c.Render(http.StatusUnauthorized, "signup", nil)
	}

	// set client cookie
	c.SetCookie(u.setCookie())

	// and session sets consider

	return c.Render(http.StatusOK, "service-in", nil)
}

/*
	# User utils
*/
func (u *User) setCookie() *http.Cookie {
	return &http.Cookie{
		Name:    u.Name,
		Value:   makeHmac(APPHASHKEY, u.Password),
		Expires: time.Now().Add(72 * time.Hour),
	}
}

func makeHmac(key, str string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}
