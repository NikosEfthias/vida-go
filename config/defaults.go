package config

import (
	"encoding/json"
	"os"
)

var conf = map[string]string{
	"DB_ADDR":          "mongodb://localhost:27017",
	"DB":               "vida",
	"LISTEN_ADDR":      ":8080",
	"APP_EMAIL_ADDR":   "info@vidaevents.org",
	"SMTP_ADDR":        "smtp.gmail.com:587",
	"APP_EMAIL_PASSWD": "<secret>",
	"APP_BASE_URL":     "https://devo.vidavidavida.com",
	"APP_INVITATION_TEMPLATE": `<html>
<head>
	<title>Email Title</title>
	<meta content="text/html; charset=utf-8" http-equiv="Content-Type">
	<meta content="width=device-width" name="viewport">
	<style media="screen and (max-width: 680px)">
		@media screen and (max-width: 680px) {
			.page-center {
				padding-left: 0 !important;
				padding-right: 0 !important;
			}
			.footer-center {
				padding-left: 20px !important;
				padding-right: 20px !important;
			}
		}
	</style>
</head>

<body style="background-color: #f4f4f5;">
	<table cellpadding="0" cellspacing="0" style="width: 100%; height: 100%; background-color: #f4f4f5; text-align: center;">
		<tbody>
			<tr>
				<td style="text-align: center;">
					<table align="center" cellpadding="0" cellspacing="0" id="body" style="background-color: #fff; width: 100%; max-width: 680px; height: 100%;">
						<tbody>
							<tr>
								<td>
									<table align="center" cellpadding="0" cellspacing="0" class="page-center" style="text-align: left; padding-bottom: 88px; width: 100%; padding-left: 120px; padding-right: 120px;">
										<tbody>
											<tr>
												<td style="padding-top: 24px;">
													<img src="https://devo.vidavidavida.com/apple-touch-icon-60x60.png" style="width: 60px; margin:20px 0;">
												</td>
											</tr>
											<tr>
												<td colspan="2" style="padding-top: 72px; -ms-text-size-adjust: 100%; -webkit-font-smoothing: antialiased; -webkit-text-size-adjust: 100%; color: #000000; font-family: 'Postmates Std', 'Helvetica', -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif; font-size: 48px; font-smoothing: always; font-style: normal; font-weight: 600; letter-spacing: -2.6px; line-height: 52px; mso-line-height-rule: exactly; text-decoration: none;">Hello!</td>
											</tr>
											<tr>
												<td style="padding-top: 48px; padding-bottom: 48px;">
													<table cellpadding="0" cellspacing="0" style="width: 100%">
														<tbody>
															<tr>
																<td style="width: 100%; height: 1px; max-height: 1px; background-color: #d9dbe0; opacity: 0.81"></td>
															</tr>
														</tbody>
													</table>
												</td>
											</tr>
											<tr>
												<td style="-ms-text-size-adjust: 100%; -ms-text-size-adjust: 100%; -webkit-font-smoothing: antialiased; -webkit-text-size-adjust: 100%; color: #9095a2; font-family: 'Postmates Std', 'Helvetica', -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif; font-size: 16px; font-smoothing: always; font-style: normal; font-weight: 400; letter-spacing: -0.18px; line-height: 24px; mso-line-height-rule: exactly; text-decoration: none; vertical-align: top; width: 100%;">
													<p>Your friend {{.Name}} invited to join Vida. Vida is spanish for Life. Vida is a mobile and desktop app for seamless group & team event planning.</p>
													<p>Life is too short to spend planning, help {{.Name}} organize an event so you can go back to your Vida.</p>
													<a href="#">{{.Link}}</a>
													<p>Looking forward to seeing you on Vida.</p>
													<p>Cheers!</p>
													<p>The Vida Team</p>
												</td>
											</tr>
										</tbody>
									</table>
								</td>
							</tr>
						</tbody>
					</table>
				</td>
			</tr>
		</tbody>
	</table>
</body>
</html>
`,
}

func init() {
	f, err := os.Open("conf.json")
	if nil != err {
		switch {
		case os.IsNotExist(err):
			d, _ := json.MarshalIndent(conf, "", "	")
			f, err = os.OpenFile("conf.json", os.O_CREATE|os.O_WRONLY, 0644)
			if nil != err {
				panic(err)
			}
			_, _ = f.Write(d)
			_ = f.Close()
		default:
			panic(err)
		}
		return
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&conf)
	if nil != err {
		panic(err)
	}
}

//Get config
func Get(k string) string {
	return conf[k]
}
