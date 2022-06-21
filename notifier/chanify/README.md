# Chanify

> 文档 https://github.com/chanify/chanify#send-text

## 配置信息

| 关键配置    | 类型     | chanify | yaml读取  | yaml可选 |
|---------|--------|---------|---------|--------|
| Webhook | string |         | webhook |        |

- Webhook: `http://<address>:<port>/v1/sender/<token>`

---

| 参数配置              | 类型                   | chanify           | yaml读取            | yaml可选 |
|-------------------|----------------------|-------------------|-------------------|--------|
| Text              | *string              | text              | text              | ✔      |
| Title             | *string              | title             | title             | ✔      |
| Copy              | *string              | copy              | copy              | ✔      |
| AutoCopy          | *int                 | autocopy          | autoCopy          | ✔      |
| Sound             | *int                 | sound             | sound             | ✔      |
| Priority          | *int                 | priority          | priority          | ✔      |
| InterruptionLevel | *string              | interruptionlevel | interruptionLevel | ✔      |
| Actions           | *map[string][]string | actions           | actions           | ✔      |
| Timeline          | *map[string]any      | timeline          | timeline          | ✔      |
| Link              | *string              | link              | link              | ✔      |

## 调用
### yaml

```yaml
chanify:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://api.chanify.net/v1/sender/CID5jpkGEiJBQTZaVxxxxxx"
    text: "Hello from notify. {{data}}"
```

### api

```go
text := "Hello from notify."
opt := chanify.Option{Webhook: "https://api.chanify.net/v1/sender/CID5jpkGEiJBQTZaVxxxxxx"}
opt.Text = &text
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```