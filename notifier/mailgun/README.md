# MailGun

> 文档 https://app.mailgun.com/app/account/setup

## 配置信息

| 关键配置    | 类型     | mailgun | yaml读取  | yaml可选 |
|---------|--------|---------|---------|--------|
| Webhook | string |         | webhook |        |
| ApiKey  | string |         | apiKey  |        |

- Webhook: `https://api.mailgun.net/v3/{{mailgun.domain}}/messages`

---

| 参数配置    | 类型     | mailgun | yaml读取  | yaml可选 |
|---------|--------|---------|---------|--------|
| From    | string | from    | from    |        |
| To      | string | to      | to      |        |
| Subject | string | subject | subject |        |
| Text    | string | text    | text    |        |

## 调用
### yaml

```yaml
mailgun:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://api.mailgun.net/v3/sandboxxxxxxx.mailgun.org/messages"
    apiKey: "3393aa7e232d9xxxxxxx"
    from: "Mailgun Sandbox <postmaster@sandboxxxxxxx.mailgun.org>"
    to: "Ss <xxxxxxx@krunsea.com>"
    subject: "demo"
    text: "Hello from notify. {{data2}}"
```

### api

```go
opt := mailgun.Option{
    Webhook: "https://api.mailgun.net/v3/sandboxxxxxxx.mailgun.org/messages",
    ApiKey:  "3393aa7e232d9xxxxxxx",
}
opt.From = "Mailgun Sandbox <postmaster@sandboxxxxxxx.mailgun.org>"
opt.To = "Ss <xxxxxxx@krunsea.com>"
opt.Subject = "demo"
opt.Text = "Hello from notify."
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```