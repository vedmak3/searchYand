package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var ur string

func main() {

	http.HandleFunc("/zapr", func(w http.ResponseWriter, r *http.Request) {
		ur = ""
		u1 := "https://yandex.ru/images/search?source=collections&rpt=imageview&url="
		u0 := r.FormValue("url")
		u2 := url.QueryEscape(u0)
		resp, _ := http.Get(u1 + u2)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		st := string(body)
		p0 := strings.Index(st, "link link_theme_normal cbir-other-sizes__link i-bem")
		if p0 > 0 {
			p0 = strings.Index(st[p0:], "href") + p0 + 6
			p1 := strings.Index(st[p0:], `"`) + p0
			ur = st[p0:p1]
		} else {
			ur = "error"
		}

	})
	http.HandleFunc("/img", func(w http.ResponseWriter, r *http.Request) {
		for {
			if ur != "" {
				break
			}
		}
		if ur != "error" {
			http.Redirect(w, r, ur, http.StatusSeeOther)
		} else {
			fmt.Fprintf(w, "Нету такой :((")
		}
	})
	http.ListenAndServe(":80", nil)

}
