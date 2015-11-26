package main

import (
	"github.com/HuKeping/htmlemail"
)

func main() {
	server := "smpt.example.com"
	sender := "sender@example.com"
	auth := htmlemail.NewLoginAuth("username", "password")

	to := []string{"receiptor1@example.com", "receiptor2@example.com"}

	body := htmlemail.GetBody("http://", "logo_alert_string", "http://", "signature_alert_string", "http://", "somebody")

	header := make(map[string]string)
	header["Subject"] = "Golang发送邮件测试"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "quoted-printable"

	err := htmlemail.SendMail(server+":25", auth, sender, to, body, header)
	if err != nil {
		fmt.Println("with err:", err)
		return
	}
	fmt.Println("please check mailbox")
}
