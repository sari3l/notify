# iGot

> 文档 http://hellyw.com/#/?id=推送消息

## 配置信息

| 关键配置    | 类型     | igot | yaml读取  | yaml可选 |
|---------|--------|------|---------|--------|
| Webhook | string |      | webhook |        |

- Webhook: `https://push.hellyw.com/{{igot.key}}`

---

| 参数配置    | 类型              | igot    | yaml读取  | yaml可选 |
|---------|-----------------|---------|---------|--------|
| Content | string          | content | content |        |
| Title   | *string         | title   | title   | ✔      |
| Url     | *string         | url     | url     | ✔      |
| Detail  | *map[string]any | detail  | detail  | ✔      |

## 调用
### yaml

```yaml
igot:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://push.hellyw.com/62aacbd09f3cfxxxxxxx"
    content: "Hello from notify. {{data}}"
```

### api

```go
opt := igot.Option{Webhook: "https://push.hellyw.com/62aacbd09f3cfxxxxxxx"}
opt.Content = "Hello from notify."
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```