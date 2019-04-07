package handlers

import (
	"encoding/json"
	"github.com/RyczkoDawid/school_app/cmd/web/helpers"
	"github.com/RyczkoDawid/school_app/cmd/web/structs"
	"github.com/RyczkoDawid/school_app/pkg/models/mysql"
	"net/http"
	"strconv"
)

func GetSummaryLesson(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get(":id"))

		if err != nil || id < 1 {
			helpers.NotFound(app, w, err)
			return
		}

		s, err := app.SummaryLesson.Get(id)

		if err == mysql.ErrNoRecord {
			helpers.NotFound(app, w, err)
			return

		} else if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		json.NewEncoder(w).Encode(s)
	}
}
