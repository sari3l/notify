# Slack

> 文档 https://api.slack.com/apps/A03MKBNG5U0/incoming-webhooks

## 配置信息

| 关键配置    | 类型     | slack | yaml读取  | yaml可选 |
|---------|--------|-------|---------|--------|
| Webhook | string |       | webhook |        |

- Webhook: `https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX`

---

| 参数配置 | 类型     | slack | yaml读取 | yaml可选 |
|------|--------|-------|--------|--------|
| Text | string | text  | text   |        |

## 调用
### yaml

```yaml
slack:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://hooks.slack.com/services/Txxxxxx/Bxxxxxx/DWXKTfP0kEBjvvAxxxxxx"
    text: "Hello from notify. {{data}}"
```

### api

```go
opt := slack.Option{Webhook: "https://hooks.slack.com/services/Txxxxxx/Bxxxxxx/DWXKTfP0kEBjvvAxxxxxx"}
opt.Text = "Hello from notify."
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```