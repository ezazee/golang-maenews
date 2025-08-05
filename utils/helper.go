package utils

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

// Fungsi helper untuk mengirim respons JSON
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Fungsi helper untuk membuat slug dari string
var nonAlphanumericRegex = regexp.MustCompile(`[^a-z0-9]+`)

func Slugify(s string) string {
	return strings.Trim(nonAlphanumericRegex.ReplaceAllString(strings.ToLower(s), "-"), "-")
}
