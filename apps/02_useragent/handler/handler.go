package handler

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, fmt.Sprintf("{\"status\":  200, \"message\": \"UserAgent Go application!\"}"))
	if err != nil {
		return
	}
}

func UserAgentHandler(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		visitorID := 0
		err := db.QueryRow("INSERT INTO visitors(user_agent, datetime) VALUES ($1, now()) RETURNING id",
			r.UserAgent(),
		).Scan(&visitorID)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal Server Error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, fmt.Sprintf("{\"status\":  200, \"message\": \"Hello visitor %d!\"}", visitorID))
	}
}

func ReportHandler(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var totalVisitors int
		err := db.QueryRow("SELECT COUNT(*) FROM visitors").Scan(&totalVisitors)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal Server Error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, fmt.Sprintf("{\"status\":  200, \"message\": \"There are %d visitors.\"}", totalVisitors))
	}
}
