package alerts

import (
	"alertmanager-receivers/pkg/config"
	"encoding/json"
	"fmt"
	"strings"
)
import "gopkg.in/gomail.v2"

func SendMail(subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Conf.Smtp.Username)
	m.SetHeader("To", strings.Split(config.Conf.Smtp.MailTo, ",")...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	d := gomail.NewDialer(config.Conf.Smtp.Host, config.Conf.Smtp.Port, config.Conf.Smtp.Username, config.Conf.Smtp.Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendAlert(record AlertRecord) error {
	for _, alert := range record.Alerts {
		body, err := json.MarshalIndent(alert, "", "    ")
		if err != nil {
			return err
		}
		if err = SendMail(alert.Status+":"+alert.Labels.Instance+" " + alert.Labels.Job +" "+alert.Labels.Alertname, fmt.Sprintf("%s\n", string(body))); err != nil {
			return err
		}
	}
	return nil
}