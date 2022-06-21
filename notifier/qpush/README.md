# QPush

> 文档 http://qpush.me/zh_tw/

## 配置信息

| 参数配置 | 类型      | qpush     | yaml读取 | yaml可选 |
|------|---------|-----------|--------|--------|
| Name | string  | name      | name   |        |
| Code | string  | code      | code   |        |
| Msg  | *string | msg[text] | msg    | ✔      |

## 调用
### yaml

```yaml
qpush:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    name: "xxxxxxx"
    code: "xxxxxxx"
    msg: "Hello from notify. {{data}}"
```

### api

```go
msg := "Hello from notify."
opt := qpush.Option{}
opt.Name = "xxxxxxxxxxxxxx"
opt.Code = "xxxxxxxxxxxxxx"
opt.Msg = &msg
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```