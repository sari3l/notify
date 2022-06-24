# ShowDoc

> 文档 https://www.showdoc.com.cn/push

## 配置信息

| 关键配置    | 类型     | showdoc | yaml读取  | yaml可选 |
|---------|--------|---------|---------|--------|
| Webhook | string |         | webhook |        |

- Webhook: `https://push.showdoc.com.cn/server/api/push/{showdoc.token}`

---

| 参数配置    | 类型     | showdoc | yaml读取  | yaml可选 |
|---------|--------|---------|---------|--------|
| Title   | string | title   | title   |        |
| Content | string | content | content |        |

## 调用
### yaml

```yaml
showdoc:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://push.showdoc.com.cn/server/api/push/f7377c6f6f4a0xxxxxx"
    title: "demo"
    content: "Hello from notify. {{data}}"
```

### api

```go
opt := showdoc.Option{Webhook: "https://push.showdoc.com.cn/server/api/push/f7377c6f6f4a0xxxxxx"}
opt.Title = "demo"
opt.Content = "Hello from notify."
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```