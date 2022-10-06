package security

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	c "github.com/rafaelbfs/GoConvenience/Convenience"
	t "github.com/rafaelbfs/eSkills/Frontend/templates"
	ht "html/template"
	"net/http"
)

var (
	googleSuccessPage, googleSignInPage *ht.Template
)

func init() {
	googleSuccessPage = c.Try(ht.ParseFiles(t.GetPath("google_success.html"))).ResultOrPanic()
	googleSignInPage = c.Try(ht.ParseFiles(t.GetPath("google_authenticate.html"))).ResultOrPanic()
}

func ConfigureGoogleAuth(p *http.ServeMux) {

	key := CfgGet(SESSION_SECRET) // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30          // 30 days
	isProd := false               // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New(CfgGet(GOOGLE_CLIENT_ID), CfgGet(GOOGLE_CLIENT_SECRET),
			"http://localhost:3000/auth/google/callback",
			"email", "profile"),
	)

	p.HandleFunc("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		googleSuccessPage.Execute(res, user)
	})

	p.HandleFunc("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	p.HandleFunc("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	})

	p.HandleFunc("/googleauth", func(res http.ResponseWriter, req *http.Request) {
		googleSignInPage.Execute(res, false)
	})
}
