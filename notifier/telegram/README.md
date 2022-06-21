# Telegram

> 文档 https://core.telegram.org/bots/api#sendmessage

## 配置信息

| 关键配置    | 类型     | telegram | yaml读取  | yaml可选 |
|---------|--------|----------|---------|--------|
| Webhook | string |          | webhook |        |

- Webhook: `https://api.telegram.org/bot{{telegram.bottoken}}/sendMessage`

---

| 参数配置                     | 类型              | telegram                    | yaml读取                   | yaml可选 |
|--------------------------|-----------------|-----------------------------|--------------------------|--------|
| ChatId                   | string          | chat_id                     | chatId                   |        |
| Text                     | string          | text                        | text                     |        |
| ParseMode                | *string         | parse_mode                  | parseMode                | ✔      |
| Entities                 | *map[string]any | entities                    | entities                 | ✔      |
| DisableWebPagePreview    | *bool           | disable_web_page_preview    | disableWebPagePreview    | ✔      |
| DisableNotification      | *bool           | disable_notification        | disableNotification      | ✔      |
| ProtectContent           | *bool           | protect_content             | protectContent           | ✔      |
| ReplyToMessageId         | *bool           | reply_to_message_id         | replyToMessageId         | ✔      |
| AllowSendingWithoutReply | *bool           | allow_sending_without_reply | allowSendingWithoutReply | ✔      |
| ReplyMarkup              | *map[string]any | reply_markup                | replyMarkup              | ✔      |

备注：在telegram中提供了`Entity`结构体方便转化，但是同`ReplyMarkup`一样过于繁琐的键值设置不切合通知功能，所以均简化为`*map[string]any`，需要请自行转换

## 调用
### yaml

```yaml
telegram:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    webhook: "https://api.telegram.org/bot516xxxxxxx/sendMessage"
    chatId: "69xxxxxxx"
    text: "Hello from notify. {{data}}"
```

### api

```go
opt := telegram.Option{Webhook: "https://api.telegram.org/bot516xxxxxxx/sendMessage"}
opt.ChatId = "69xxxxxxx"
opt.Text = "Hello from notify."
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```