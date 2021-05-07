package lib

import (
	"strings"
	"time"
)

type Cookie struct {
	Name    string
	Value   string
	Expires time.Time
}

func GetCookie(cookRaw string, name string) *Cookie {
	if len(cookRaw) == 0 {
		return nil
	}
	res := &Cookie{Name: name}
	cookRaw = strings.ReplaceAll(cookRaw, " ", "")
	cookSplitted := strings.Split(cookRaw, ";")
	for _, c := range cookSplitted {
		i := strings.Index(c, "=")
		if n := c[:i]; n == name {
			res.Value = c[i+1:]
			return res
		}
	}
	return nil
}
