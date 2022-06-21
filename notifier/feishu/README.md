# FeiShu

> 文档 https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
> 
> 飞书加签有BUG，也是奇葩 

## 配置信息

| 关键配置    | 类型     | feishu | yaml读取  | yaml可选 |
|---------|--------|--------|---------|--------|
| Webhook | string |        | webhook |        |
| Secret  | string |        | secret  | ✔      |

- Webhook: `https://open.feishu.cn/open-apis/bot/v2/hook/{{feishu.token}}`

---

| 参数配置      | 类型             | feishu    | yaml读取    | yaml可选 |
|-----------|----------------|-----------|-----------|--------|
| MsgType   | string         | msg_type  | msgType   |        |
| Content   | map[string]any | content   | content   |        |
| Timestamp | *string        | timestamp | timestamp | ✔      |
| Sign      | *string        | sign      | sign      | ✔      |

## 调用
### yaml

```yaml
- notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://open.feishu.cn/open-apis/bot/v2/hook/f9fb75ca-4xxxxxxx"
    msgType: "text"
    content: {"text":"Hello from notify. {{data}}"}
```

### api

```go
opt := feishu.Option{Webhook: "https://open.feishu.cn/open-apis/bot/v2/hook/f9fb75ca-4xxxxxxx"}
opt.MsgType = "text"
opt.Content = map[string]any{"text": "Hello from notify."}
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```