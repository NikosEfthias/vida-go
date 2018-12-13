package config

import (
	"encoding/json"
	"os"
)

var conf = map[string]string{
	"DB_ADDR":           "mongodb://localhost:27017",
	"DB":                "vida",
	"LISTEN_ADDR":       ":8080",
	"APP_EMAIL_ADDR":    "info@vida.events",
	"SMTP_ADDR":         "smtp.yandex.com:25",
	"APP_EMAIL_PASSWD":  "",
	"APP_BASE_URL":      "https://vida.events",
	"AWS_KEY":           "",
	"AWS_SECRET":        "",
	"PUBLIC_FILES_PATH": "public/",
	//app invitation template {{{
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
													<img src="data:image/png;base64, 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" style="width: 60px; margin:20px 0;">
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
													<p>Your friend {{.Name}} invited to join Vida. Vida is spanish for Life. Vida is a mobile and desktop app
														for seamless group & team event planning.</p>
													<p>Life is too short to spend planning, help {{.Name}} organize an event so you can go back to your Vida.</p>
													<a href="{{.Link}}">{{.Link}}</a>
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

</html>`,
	//}}}
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
	if d := os.Getenv(k); d != "" {
		return d
	}
	return conf[k]
}
