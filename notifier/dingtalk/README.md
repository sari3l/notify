# DingTalk

> 文档 https://open.dingtalk.com/document/robots/custom-robot-access

## 配置信息

| 关键配置        | 类型     | dingtalk | yaml读取      | yaml可选 |
|-------------|--------|----------|-------------|--------|
| Token       | string |          | token       |        |
| MessageType | string |          | messageType | ✔      |
| Secret      | string |          | secret      | ✔      |

- Webhook: `https://open.feishu.cn/open-apis/bot/v2/hook/{{feishu.token}}`

---

| 参数配置           | 类型                | dingtalk       | yaml读取         | yaml可选 |
|----------------|-------------------|----------------|----------------|--------|
| Text           | *string           | text           | text           | ✔      |
| Title          | *string           | title          | title          | ✔      |
| SingleTitle    | *string           | singleTitle    | singleTitle    | ✔      |
| SingleUrl      | *string           | singleUrl      | singleUrl      | ✔      |
| BtnOrientation | *string           | btnOrientation | btnOrientation | ✔      |
| Content        | *string           | content        | content        | ✔      |
| PicUrl         | *string           | picUrl         | picUrl         | ✔      |
| MessageUrl     | *string           | messageUrl     | messageUrl     | ✔      |
| Btns           | *[]map[string]any | btns           | btns           | ✔      |
| Links          | *[]map[string]any | links          | links          | ✔      |

备注：在telegram中提供了`Btns | Links`对应结构体`方便转化，但是过于繁琐的键值设置不切合通知功能，所以均简化为`*map[string]any`，需要请自行转换

## 调用
### yaml

```yaml
dingtalk:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    token: "5ff195509f5653f6bfffxxxxxxx"
    content: "Hello from notify. {{data}}"
```

### api

```go
content := "Hello from notify."
opt := dingtalk.Option{Token: "5ff195509f5653f6bfffxxxxxxx"}
opt.Content = &content
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```