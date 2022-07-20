# FTQQ

> 文档 https://sct.ftqq.com/sendkey

## 配置信息

| 关键配置    | 类型     | ftqq  | yaml读取  | yaml可选 |
|---------|--------|-------|---------|--------|
| Webhook | string |       | webhook |        |

- Webhook: `https://sctapi.ftqq.com/{{ftqq.SendKey}}.send`

---

| 参数配置        | 类型      | ftqq    | yaml读取  | yaml可选 |
|-------------|---------|---------|---------|--------|
| Title       | string  | title   | title   |        |
| Description | *string | desp    | desp    | ✔      |
| Channel     | *string | channel | channel | ✔      |
| OpenId      | *string | openId  | openId  | ✔      |

## 调用
### yaml

```yaml
ftqq:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://sctapi.ftqq.com/SCT140140Txxxxxxx.send"
    title: "demo"
    desp: "Hello from notify. {{data}}"
```

### api

```go
desp := "Hello from notify""
opt := ftqq.Option{Webhook: "https://sctapi.ftqq.com/SCT140140Txxxxxxx.send"}
opt.Title = "demo"
opt.Description = &desp
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```