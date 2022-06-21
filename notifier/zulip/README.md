# Zulip

> 文档 https://zulip.com/api/send-message

## 配置信息

| 关键配置     | 类型     | zulip | yaml读取   | yaml可选 |
|----------|--------|-------|----------|--------|
| Webhook  | string |       | webhook  |        |
| BotEmail | string |       | botEmail |        |
| BotKey   | string |       | botKey   |        |

- Webhook: `https://{{zulip.domain}}.zulipchat.com/api/v1/messages`

---

| 参数配置    | 类型      | zulip    | yaml读取  | yaml可选 |
|---------|---------|----------|---------|--------|
| Type    | string  | type     | type    |        |
| To      | string  | to       | to      |        |
| Content | string  | content  | content |        |
| Topic   | *string | topic    | topic   | ✔      |
| QueueId | *string | queue_id | queueId | ✔      |
| LocalId | *string | local_id | localId | ✔      |

- Type: `private | stream`
- To:
  - steam: "xxxxxx"
  - private: "[9,10]"

> For stream messages, either the name or integer ID of the stream. For private messages, either a list containing integer user IDs or a list containing string email addresses.

## 调用
### yaml

```yaml
zulip:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://xxxxxxx.zulipchat.com/api/v1/messages"
    botEmail: "test-bot@xxxxxxx.zulipchat.com"
    botKey: "iSaTzykqgY0xxxxxxx"
    type: "stream"
    to: "general"
    content: "Hello from notify. {{data}}"
```

### api

```go
opt := zulip.Option{
    Webhook:  "https://xxxxxxx.zulipchat.com/api/v1/messages",
    BotEmail: "test-bot@xxxxxxx.zulipchat.com",
    BotKey:   "iSaTzykqgY0xxxxxxx",
}
opt.Type = "stream"
opt.To = "general"
opt.Content = "Hello from notify."
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```