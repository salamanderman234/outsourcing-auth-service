package helper

import (
	"bytes"
	"encoding/json"
	"net/http"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
)

func SendMail(body string, subject string, targets ...string) error {
	bodyReq,_ := json.Marshal(map[string]any {
		"subject": subject,
		"body": body,
		"to": targets,
	})
	bodyBuff := bytes.NewBuffer(bodyReq)
	http.Post(domain.SEND_MAIL_SERVICE_URL, "application/json", bodyBuff)
	return nil
}