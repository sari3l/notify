# Bark

## 配置信息

| 关键配置    | 类型     | bark | yaml读取  | yaml可选 |
|---------|--------|------|---------|--------|
| Webhook | string |      | webhook |        |

- Webhook:
  - `https://api.day.app/{{bark.token}}/{{bark.data}}`
  - `https://api.day.app/{{bark.token}}/{{bark.title}}/{{bark.data}}`

---

| 参数配置      | 类型      | bark      | yaml读取    | yaml可选 |
|-----------|---------|-----------|-----------|--------|
| Sound     | *string | sound     | sound     | ✔      |
| IsArchive | *int    | isArchive | isArchive | ✔      |
| Icon      | *string | icon      | icon      | ✔      |
| Group     | *string | group     | group     | ✔      |
| Level     | *string | level     | level     | ✔      |
| Url       | *string | url       | url       | ✔      |
| Copy      | *string | copy      | copy      | ✔      |
| Badge     | *int    | badge     | badge     | ✔      |
| AutoCopy  | *int    | autoCopy  | autoCopy  | ✔      |

## 调用
### yaml

```yaml
bark:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://api.day.app/qr3XSFqvNxxxxxxx/{{data}}"
    sound: "minuet"
```

### api

```go
opt := bark.Option{Webhook: "https://api.day.app/qr3XSFqvNxxxxxxx/Hello+from+notify."}
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```