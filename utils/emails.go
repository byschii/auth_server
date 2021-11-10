package utils

import (
	"github.com/smtp2go-oss/smtp2go-go"
)

func SendEmailVerficationMail(to string, code string, cfg Configuration) {
	var verificatonUrl = ""
	var email = smtp2go.Email{
		From:     cfg.Smpt2goEmailSender + " <" + cfg.Smpt2goEmailSender + "@" + cfg.ServiceDomain + ">",
		To:       []string{"<" + to + ">"},
		Subject:  "Email Verification",
		HtmlBody: "",
	}
}
