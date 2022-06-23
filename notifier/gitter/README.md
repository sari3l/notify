# Gitter

> 文档 https://developer.gitter.im/docs/messages-resource

## 配置信息

| 关键配置    | 类型     | gitter | yaml读取  | yaml可选 |
|---------|--------|--------|---------|--------|
| Webhook | string |        | webhook |        |
| Token   | string |        | token   |        |

- Webhook: `https://api.gitter.im/v1/rooms/:roomId/chatMessages`

---

| 参数配置 | 类型     | gitter | yaml读取 | yaml可选 |
|------|--------|--------|--------|--------|
| Text | string | text   | text   |        |

## 调用
### yaml

```yaml
gitter:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://api.gitter.im/v1/rooms/62b461356daxxxxxx/chatMessages"
    token: "3d0ebe3ce61c414xxxxxx"
    text: "Hello from notify. {{data}}"
```

### api

```go
opt := gitter.Option{
	Webhook: "https://api.gitter.im/v1/rooms/62b461356daxxxxxx/chatMessages",
	Token: "3d0ebe3ce61c414xxxxxx",
}
opt.Text = "Hello from notify."
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```