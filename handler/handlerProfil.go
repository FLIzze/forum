package forum

import (
	"database/sql"
	t2 "forum/logOutSessionHtml"
	t "forum/users"
	"html/template"
	"net/http"
)

func Handler_profil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/profil.html"))
	databaseForum, _ := sql.Open("sqlite3", "../forum.db")

	if r.FormValue("logOutButton") == "logout" {
		t.Logout(r)
		t.LogOutCookie(r, w)
		t2.LogOutSession()
	} else if r.FormValue("delete") != "" {
		t.DeleteAccount(r, databaseForum, w)
	} else if r.FormValue("username") != "" || r.FormValue("email") != "" {
		t.UserEdit(r, databaseForum)
	}
	t.MpSendOrReceivedByUser(databaseForum)
	t.MessagesSendByUser(databaseForum)
	t.TopicCreatedByUser(databaseForum)

	t.GetCookieHandler(w, r)

	tmpl.Execute(w, t.USER)
}
