# PushOver

> 文档 https://pushover.net/

## 配置信息

| 参数配置     | 类型      | pushover | yaml读取   | yaml可选 |
|----------|---------|----------|----------|--------|
| Token    | string  | token    | token    |        |
| User     | string  | user     | user     |        |
| Message  | string  | message  | message  |        |
| Device   | *string | device   | device   | ✔      |
| Html     | *string | html     | html     | ✔      |
| Priority | *string | priority | priority | ✔      | 
| Sound    | *string | sound    | sound    | ✔      |
| Title    | *string | title    | title    | ✔      |
| Url      | *string | url      | url      | ✔      |
| UrlTitle | *string | urlTitle | urlTitle | ✔      |

## 调用
### yaml

```yaml
pushover:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    token: "aefgdwtxs1xoxjxxxxxxx"
    user: "uacn4m8sq3cm7jxxxxxxx"
    message: "Hello from notify. {{data}}"
```

### api

```go
opt := pushover.Option{}
opt.Token = "aefgdwtxs1xoxjxxxxxxx"
opt.User = "uacn4m8sq3cm7jxxxxxxx"
opt.Message = "Hello from notify."
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```