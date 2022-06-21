# PushDeer

> 文档 https://github.com/easychen/pushdeer#推送消息

## 配置信息

| 参数配置        | 类型      | pushdeer | yaml读取  | yaml可选 |
|-------------|---------|----------|---------|--------|
| PushKey     | string  | pushkey  | pushKey |        |
| Type        | *string | type     | type    | ✔      |
| Text        | *string | text     | text    | ✔      |
| Description | *string | desc     | desc    | ✔      |

## 调用
### yaml

```yaml
pushdeer:
 - notifyLevel: 1
   notifyFormatter: ["{{data}}"]
   pushKey: "PDU12208TFNJw3weI89xGxxxxxxx"
   text: "Hello from notify. {{data}}"
```

### api

```go
text := "Hello from notify."
opt := pushdeer.Option{}
opt.PushKey = "PDU12208TFNJw3weI89xGxxxxxxx"
opt.Text = &text
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```