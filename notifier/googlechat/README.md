# GoogleChat

> 文档 https://developers.google.com/chat/how-tos/webhooks#python

## 配置信息

| 关键配置    | 类型     | googlechat | yaml读取  | yaml可选 |
|---------|--------|------------|---------|--------|
| Webhook | string |            | webhook |        |

- Webhook: `https://chat.googleapis.com/v1/spaces/{{googlechat.spaceid}}/messages?key={{googlechat.key}}&token={{googlechat.token}}`

---

| 参数配置 | 类型     | googlechat | yaml读取 | yaml可选 |
|------|--------|------------|--------|--------|
| Text | string | text       | text   |        |

## 调用
### yaml

```yaml
googlechat:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://chat.googleapis.com/v1/spaces/AAAAxxxxxxx/messages?key=AIzaSyDdI0xxxxxxx&token=BBSS0tJT_7BVGSxxxxxxx"
    text: "Hello from notify. {{data}}"
```

### api

```go
opt := googlechat.Option{Webhook: "https://chat.googleapis.com/v1/spaces/AAAAxxxxxxx/messages?key=AIzaSyDdI0xxxxxxx&token=BBSS0tJT_7BVGSxxxxxxx"}
opt.Text = "Hello from notify."
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```