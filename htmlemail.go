package htmlemail

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

type LoginAuth struct {
	username string
	password string
}

func NewLoginAuth(username, password string) smtp.Auth {
	return &LoginAuth{username, password}
}

func (a *LoginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *LoginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown fromServer")
		}
	}
	return nil, nil
}

func SendMail(addr string, a smtp.Auth, from string, to []string, body string, header map[string]string) error {
	c, err := smtp.Dial(addr)
	host, _, _ := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	defer c.Close()

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: host, InsecureSkipVerify: true}
		if err = c.StartTLS(config); err != nil {
			return err
		}
	}

	if a != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(a); err != nil {
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	_, err = w.Write([]byte(message))

	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func GetBody(logoAddr, logoAlt, signatureAddr, signatureAlt, optionAddr, user string) string {

	logoAddr = strings.Replace(logoAddr, "=", "=3D", -1)
	signatureAddr = strings.Replace(signatureAddr, "=", "=3D", -1)
	optionAddr = strings.Replace(optionAddr, "=", "=3D", -1)

	body := `
	<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<html xmlns=3D"http://www.w3.org/1999/xhtml">
	<head>
		<meta http-equiv=3D"Content-Type" content=3D"text/html; charset=3Dutf-8">
		<meta name=3D"viewport" content=3D"width=3Ddevice-width, initial-scale=3D1.0">

		<style type=3D"text/css">
			#outlook a{padding:0;}
			.ReadMsgBody{width:100%;} .ExternalClass{width:100%;}
			.ExternalClass,
			.ExternalClass p,
			.ExternalClass span,
			.ExternalClass font,
			.ExternalClass td,
			.ExternalClass div {line-height: 100%;}
			body, table, td, p, a, li, blockquote{-webkit-text-size-adjust:100%; -ms-text-size-adjust:100%;}
			table, td{mso-table-lspace:0pt; mso-table-rspace:0pt;}
			img{-ms-interpolation-mode:bicubic;}

			body, #bodyTable, #bodyCell{height:100% !important; width:100% !important; margin:0; padding:0;}
			table{border-collapse:collapse !important;}
			=09

			@media only screen and (max-width: 480px){
			body{width:100% !important; min-width:100% !important;}

			td[id=3D"bodyCell"]{padding:30px 0 !important;}
			table[id=3D"emailContainer"]{max-width:600px !important; width:100% !im	portant;}
			td[class=3D"mobilePadding"], td[class=3D"bodyContent"]{padding-right:20	px !important; padding-left:20px !important;}

			h1{font-size:23px !important;}
			td[class=3D"bodyContent"]{font-size:18px !important;}
			table[class=3D"emailButton"]{max-width:480px !important; width:100% !im	portant;}
			td[class=3D"emailButtonContent"]{font-size:18px !important;}
			td[class=3D"emailButtonContent"] a{display:block;}
			td[class=3D"emailColumn"]{display:block !important; max-width:600px !im	portant; width:100% !important;}
			td[class=3D"footerContent"]{font-size:15px !important; padding-right:15	px; padding-bottom:30px; padding-left:15px;}
			table[id=3D"utilityLinkBlock"]{border-top:1px solid #E5E5E5; max-width:	600px !important; width:100% !important;}
			td[class=3D"utilityLink"]{background-color:#E1E1E1 !important; border-b	ottom:10px solid #F2F2F2;
			display:block !important; font-size:15px !important; padding:15px  !important; 
			text-align:center !important; width:100% !important;}
			td[class=3D"utilityLink"] a{color:#606060 !important; display:block !important; text-decoration:none !important;}
			}
		</style>
	</head>
	<body style=3D"background-color:#F2F2F2;">
		<center>
		<table border=3D"0" cellpadding=3D"0" cellspacing=3D"0" height=3D"100%" width=3D"100%" id=3D"bodyTable" style=3D"background-color:#F2F2F2;">
		<tr>
			<td align=3D"center" valign=3D"top" id=3D"bodyCell" style=3D"padding:40px 20px;">
				<table border=3D"0" cellpadding=3D"0" cellspacing=3D"0" id=3D"emailContainer" style=3D"width:600px;">
				<tr>
					<td align=3D"center" valign=3D"top" style=3D"padding-bottom:30px;">
						<table border=3D"0" cellpadding=3D"0" cellspacing=3D"0" width=3D"100%" id=3D"emailBody" style=3D"background-color:#FFFFFF; border-collapse:separate !important; border-radius:4px;">

						<!-----------------------------------First Part----------------------------------------->
						<tr>
							<td align=3D"center" valign=3D"top" class=3D"mobilePadding" style=3D" padding-top:40px; padding-right:40px; padding-bottom:0; padding-left:40px;">
								<table border=3D"0" cellpadding=3D"0" cellspacing=3D"0" width=3D"100%">
								<tr>
									<td align=3D"center" valign=3D"top" style=3D"padding-right:20px;">
										<a href=3D"" target=3D"_blank" title=3D"Europa" style=3D"text-decoration:none;">
										<img src=3D` + logoAddr + ` alt=3D` + logoAlt + ` height=3D"" width=3D"50" class=3D"logoImage" style=3D"border:0; color:#6DC6DD !important; font-family:Helvetica, Arial, sans-serif; font-size:60px; font-weight:bold; height:auto !important; letter-spacing:-4px; line-height:100%; outline:none; text-align:center; text-decoration:none;"/>
										</a>
									</td>
									<td valign=3D"middle" width=3D"100%" class=3D"headerContent" style=3D"color:#606060; font-family:Helvetica, Arial, sans-serif; font-size:15px; line-height:150%; text-align:left;">
									<h1 style=3D"color:#606060 !important; font-family:Helvetica, Arial, sans-serif; font-size:26px; font-weight:bold; letter-spacing:-1px; line-height:115%; margin:0; padding:0; text-align:left;">Action required: Please verify your email address.</h1>
									</td>
								</tr>
								</table>
							</td>
						</tr>

						<!-----------------------------------Second Part----------------------------------------->
						<tr>
							<td align=3D"left" valign=3D"top" class=3D"mobilePadding" style=3D"padding-top:40px; padding-right:40px; padding-bottom:40px; padding-left:40px;">
								<table border=3D"0" cellpadding=3D"0" cellspacing=3D"0" width=3D"100%">
								<tr>
									<td style=3D"border-top:1px dotted #CCCCCC; border-bottom:1px dotted #CCCCCC; padding-top:10px; padding-bottom:10px;">
										<h2 style=3D"color:#606060 !important; font-family:Helvetica, Arial, sans-serif; font-size:20px; letter-spacing:-.5px; line-height:115%; margin:0; padding:0; text-align:center;">` + user + `</h2>
									</td>
								</tr>
								</table>
							</td>
						</tr>
				
				
						<!-----------------------------------Third Part----------------------------------------->
						<tr>
							<td align=3D"center" valign=3D"middle" style=3D"padding-right:40px; padding-bottom:40px; padding-left:40px;">
								<table border=3D"0" cellpadding=3D"0" cellspacing=3D"0" class=3D"emailButton" style=3D"background-color:#6DC6DD; border-collapse:separate !important; border-radius:3px;">
								<tr>
									<td align=3D"center" valign=3D"middle" class=3D"emailButtonContent" style=3D"color:#FFFFFF; font-family:Helvetica, Arial, sans-serif; font-size:15px; font-weight:bold; line-height:100%; padding-top:18px; padding-right:15px; padding-bottom:15px;padding-left:15px;">
									<a href=3D"` + optionAddr + `" target=3D"_blank" style=3D"color:#FFFFFF; text-decoration:none;">Activate your account</a>
									</td>
								</tr>
								</table>
							</td>
						</tr>
				
				
						<!-----------------------------------Forth Part----------------------------------------->
						<tr>
							<td align=3D"center" valign=3D"top" class=3D"bodyContent" style=3D"border-top:2px solid #F2F2F2; color:#606060; font-family:Helvetica, Arial, sans-serif; font-size:15px; line-height:150%; padding-top:20px; padding-right:40px; padding-bottom:20px; padding-left:40px; text-align:center;">Help! 
							<a href=3D"#" target=4D"_blank" style=3D"color:#52BAD5;">I didn't request this.</a>
							</td>
						</tr>
						</table>
					</td>
				</tr>

				<tr>
					<td align=3D"center" valign=3D"top">
						<table border=3D"0" cellpadding=3D"0" cellspacing=3D"0" width=3D"100%" id=3D"emailFooter">
						<tr>
							<td align=3D"center" valign=3D"top" class=3D"footerContent" style=3D"color:#606060; font-family:Helvetica, Arial, sans-serif; font-size:13px; line-height:125%;"> 
								Please do not reply to this email. Emails sent to this address will not be answered.
							</td>
						</tr>
						<tr>
							<td align=3D"center" valign=3D"top" style=3D"padding-top:30px;">
							<a href=3D"" target=3D"_blank" title=3D` + signatureAlt + ` style=3D"text-decoration:none;">
							<img src=3D` + signatureAddr + ` height=3D"25" width=3D"100" style=3D"border:0; outline:none; text-decoration:none;"/>
								</a>
							</td>
						</tr>
						</table>
					</td>
				</tr>
				</table>
			</td>
		</tr>
		</table>
		</center>
	</body>
	</html>
	`
	return body
}
