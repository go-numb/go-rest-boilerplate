# go-REST-boilerplate

It is boilerplate for Golang REST API .

## Description

go-rest-boilerplate is the base to make REST API quickly for me.

## TODO
1. requests DB


## Installation

```
$ go get -u github.com/go-numb/go-rest-boilerplate
```


## Uses struct
```
type User struct {
    ID      int
	isLogin bool

	Name     string
	Email    string
	Password string

	CreatedAt time.Time
}
```

## Decisions
```
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
```

## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-rest-boilerplate/blob/master/LICENSE)