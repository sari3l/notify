# PushBullet

> 文档 https://docs.pushbullet.com/#push

## 配置信息

| 关键配置  | 类型     | pushbullet | yaml读取 | yaml可选 |
|-------|--------|------------|--------|--------|
| Token | string |            | token  |        |

---

| 参数配置                    | 类型        | pushbullet                | yaml读取                  | yaml可选 |
|-------------------------|-----------|---------------------------|-------------------------|--------|
| Type                    | string    | type                      | type                    |        |
| Iden                    | *string   | iden                      | iden                    | ✔      |
| Active                  | *bool     | active                    | active                  | ✔      |
| Created                 | *float32  | created                   | created                 | ✔      |
| Modified                | *float32  | modified                  | modified                | ✔      |
| Dismissed               | *bool     | dismissed                 | dismissed               | ✔      |
| Guid                    | *string   | guid                      | guid                    | ✔      |
| Direction               | *string   | direction                 | direction               | ✔      |
| SenderIden              | *string   | sender_iden               | senderIden              | ✔      |
| SenderEmail             | *string   | sender_email              | senderEmail             | ✔      |
| SenderEmailNormalized   | *string   | sender_email_normalized   | senderEmailNormalized   | ✔      |
| SenderName              | *string   | sender_name               | senderName              | ✔      |
| ReceiverIden            | *string   | receiver_iden             | receiverIden            | ✔      |
| ReceiverEmail           | *string   | receiver_email            | receiverEmail           | ✔      |
| ReceiverEmailNormalized | *string   | receiver_email_normalized | receiverEmailNormalized | ✔      |
| TargetDeviceIden        | *string   | target_device_iden        | targetDeviceIden        | ✔      |
| SourceDeviceIden        | *string   | source_device_iden        | sourceDeviceIden        | ✔      |
| ClientIden              | *string   | client_iden               | clientIden              | ✔      |
| ChannelIden             | *string   | channel_iden              | channelIden             | ✔      |
| AwakeAppGuids           | *[]string | awake_app_guids           | awakeAppGuids           | ✔      |
| Title                   | *string   | title                     | title                   | ✔      |
| Body                    | *string   | body                      | body                    | ✔      |
| Url                     | *string   | url                       | url                     | ✔      |
| FileName                | *string   | file_name                 | fileName                | ✔      |
| FileType                | *string   | file_type                 | fileType                | ✔      |
| FileUrl                 | *string   | file_url                  | fileUrl                 | ✔      |
| ImageUrl                | *string   | image_url                 | imageUrl                | ✔      |
| ImageWidth              | *int      | image_width               | imageWidth              | ✔      |
| ImageHeight             | *int      | image_height              | imageHeight             | ✔      |

- Type: `note | link | file`
- Direction: `"self | outgoing | incoming`

## 调用
### yaml

```yaml
pushbullet:
  - notifyLevel: 1
    notifyFormatter: ["{{data}}"]
    token: "o.crqLtgqLqdM0edqYs1xxxxxxx"
    type: "note"
    body: "Hello from notify. {{data}}"
```

### api

```go
body := "Hello from notify."
opt := pushbullet.Option{Token: "o.crqLtgqLqdM0edqYs1xxxxxxx"}
opt.Type = "note"
opt.Body = &body
err := opt.ToNotifier().Send(nil)
if err != nil {
    fmt.Println(err)
}
```