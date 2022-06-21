# RocketChat

> 文档 https://docs.rocket.chat/guides/administration/admin-panel/integrations

## 配置信息

| 关键配置    | 类型     | rockatchat | yaml读取  | yaml可选 |
|---------|--------|------------|---------|--------|
| Webhook | string |            | webhook |        |

- Webhook: `http://{{rockatchat.server}}/hooks/{{rocketchat.token}}`

---

| 参数配置      | 类型      | rockatchat | yaml读取   | yaml可选 |
|-----------|---------|------------|----------|--------|
| Text      | string  | text       | text     | 	      |
| title     | *string | title      | title    | 	✔     |
| TitleLink | *string | title_link | title    | 	✔     |
| ImageUrl  | *string | image_url  | imageUrl | 	✔     |
| Color     | *string | color      | color    | 	✔     |

## 调用
### yaml

```yaml
rocketchat:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "http://127.0.0.1:3000/hooks/62a6f649axxxxxxx/SdWX8fXKEs8xMBFxxxxxxx"
    text: "Hello from notify. {{data}}"
```

### api

```go
opt := rocketchat.Option{Webhook: "http://127.0.0.1:3000/hooks/62a6f649axxxxxxx/SdWX8fXKEs8xMBFxxxxxxx"}
opt.Text = "Hello from notify."
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```