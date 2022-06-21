# Discord

> 文档 https://discord.com/developers/docs/resources/webhook#execute-webhook

## 配置信息

| 关键配置    | 类型     | discord | yaml读取  | yaml可选 |
|---------|--------|---------|---------|--------|
| Webhook | string |         | webhook |        |

- Webhook: `https://discord.com/api/webhooks/{webhook.id}/{webhook.token}`

---

| 参数配置      | 类型      | discord   | yaml读取    | yaml可选 |
|-----------|---------|-----------|-----------|--------|
| Content   | string  | content   | content   |        |
| Username  | *string | username  | username  | ✔      |
| AvatarUrl | *string | avatarUrl | avatarUrl | ✔      |
| Tts       | *bool   | tts       | tts       | ✔      |

## 调用
### yaml

```yaml
discord:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://discord.com/api/webhooks/985040095xxxxxxx/uqTQQuwxxxxxxx-xxxxxxx"
    content: "Hello from notify. {{data}}"
```

### api

```go
opt := discord.Option{Webhook: "https://discord.com/api/webhooks/985040095xxxxxxx/uqTQQuwxxxxxxx-xxxxxxx"}
opt.Content = "Hello from notify."
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```