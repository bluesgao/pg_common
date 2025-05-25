package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	SMTPHost string
	SMTPPort int
	From     string
	Password string // 授权码
}

// 发送注册验证码邮件
func SendRegisterCodeEmail(cfg EmailConfig, to string, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "您的注册验证码")

	// 构建 HTML 邮件内容
	body := fmt.Sprintf(`
		<h2>欢迎您</h2>
		<p>您的注册验证码是：<strong style="color:#2d8cf0;font-size:24px;">%s</strong></p>
		<p>请在 5 分钟内使用，勿泄露给他人。</p>
	`, code)

	m.SetBody("text/html", body)

	d := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.From, cfg.Password)
	return d.DialAndSend(m)
}
