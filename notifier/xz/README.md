# XZ

> 文档 https://xz.qqoq.net/#/index

## 配置信息

| 关键配置    | 类型     | xz  | yaml读取  | yaml可选 |
|---------|--------|-----|---------|--------|
| Webhook | string |     | webhook |        |

- Webhook: `https://xizhi.qqoq.net/{{xz.token}}.send`
---

| 参数配置    | 类型      | xz      | yaml读取  | yaml可选 |
|---------|---------|---------|---------|--------|
| Title   | string  | title   | title   |        | 
| Content | *string | content | content | ✔      |
| Type    | *string | type    | type    | ✔      |
| Time    | *string | time    | time    | ✔      | 
| Date    | *string | date    | date    | ✔      |


- Type: 
  - 立即: `null`
  - 每月: `m`
  - 每周: `w`
  - 每天: `d`
  - 每小时: `h`
  - 每分钟: `mm`
- Time: 
  - 每月 / 每周 / 每天: `24:00`
  - 每小时: `00`

## 调用
### yaml

```yaml
xz:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://xizhi.qqoq.net/XZf282c9538b8axxxxxxx.send"
    title: "demo"
    content: "Hello from notify. {{data}}"
```

### api

```go
content := "Hello from notify."
opt := xz.Option{Webhook: "https://xizhi.qqoq.net/XZf282c9538b8axxxxxxx.send"}
opt.Title = "demo"
opt.Content = &content
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```