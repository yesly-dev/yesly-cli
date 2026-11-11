package cli

// examples are the Example section of each operation's help, by resource.method.
// TestEveryOperationHasAnExample keeps the list complete.
var examples = map[string]string{
	"agents.list":     `  yesly agents list`,
	"agents.retrieve": `  yesly agents retrieve agt_8Hd2kQ4mZ1LpXy`,
	"agents.pause":    `  yesly agents pause agt_8Hd2kQ4mZ1LpXy`,
	"agents.resume":   `  yesly agents resume agt_8Hd2kQ4mZ1LpXy`,
	"agents.update": `  yesly agents update agt_8Hd2kQ4mZ1LpXy --limits.messages-per-conversation 20 --recording.enabled
  yesly agents update agt_8Hd2kQ4mZ1LpXy --contact-hours @hours.json`,
	"senders.list": `  yesly senders list --agent agt_8Hd2kQ4mZ1LpXy --channel whatsapp`,
	"contacts.list": `  yesly contacts list --external-id customer-3391
  yesly contacts list --phone +34600000001
  yesly contacts list --all`,
	"contacts.create": `  yesly contacts create --external-id customer-3391 --name Marta \
    --phone +34600000001 --metadata segment=premium
  yesly contacts create --data @contact.json`,
	"contacts.retrieve":      `  yesly contacts retrieve ctc_3Kf9Lq2Wm7Rt4B`,
	"contacts.update":        `  yesly contacts update ctc_3Kf9Lq2Wm7Rt4B --phones +34600000001,+34600000002 --metadata segment=standard`,
	"contacts.delete":        `  yesly contacts delete ctc_3Kf9Lq2Wm7Rt4B`,
	"contacts.merge":         `  yesly contacts merge ctc_3Kf9Lq2Wm7Rt4B --contact ctc_…`,
	"contacts.list_memory":   `  yesly contacts list-memory ctc_3Kf9Lq2Wm7Rt4B`,
	"contacts.delete_memory": `  yesly contacts delete-memory ctc_3Kf9Lq2Wm7Rt4B mem-7f3a9c2e4b1d4e8a`,
	"conversations.list": `  yesly conversations list --agent agt_8Hd2kQ4mZ1LpXy --status no_reply --from 2026-12-01
  yesly conversations list --contact ctc_3Kf9Lq2Wm7Rt4B --all`,
	"conversations.create": `  yesly conversations create --agent agt_8Hd2kQ4mZ1LpXy \
    --contact '{"external_id": "customer-3391", "phone": "+34600000001", "name": "Marta"}' \
    --channels whatsapp,call --start-at 2026-12-01T10:00:00+01:00 \
    --context '{"order": "8841", "delivery_window": "2026-12-02T09:00/13:00"}' \
    --idempotency-key order-8841-delivery
  yesly conversations create --agent agt_8Hd2kQ4mZ1LpXy --contact ctc_3Kf9Lq2Wm7Rt4B \
    --channels call --context '{"order": "8902"}'`,
	"conversations.retrieve":      `  yesly conversations retrieve cnv_7Qp4Ns2Xk9Lm3D`,
	"conversations.delete":        `  yesly conversations delete cnv_7Qp4Ns2Xk9Lm3D`,
	"conversations.cancel":        `  yesly conversations cancel cnv_7Qp4Ns2Xk9Lm3D`,
	"conversations.list_messages": `  yesly conversations list-messages cnv_7Qp4Ns2Xk9Lm3D`,
	"conversations.recording":     `  yesly conversations recording cnv_7Qp4Ns2Xk9Lm3D --message msg_…`,
	"messages.list":               `  yesly messages list --conversation cnv_7Qp4Ns2Xk9Lm3D`,
	"messages.create": `  yesly messages create --agent agt_8Hd2kQ4mZ1LpXy --contact ctc_3Kf9Lq2Wm7Rt4B \
    --channel sms --text "Hola Marta, tu pedido 8841 sale hoy."
  yesly messages create --agent agt_8Hd2kQ4mZ1LpXy --contact ctc_3Kf9Lq2Wm7Rt4B --channel whatsapp \
    --template.name delivery_reminder --template.language es --template.variables 8841`,
	"messages.retrieve": `  yesly messages retrieve msg_…`,
	"templates.list":    `  yesly templates list --status approved`,
	"templates.create": `  yesly templates create --name delivery_reminder --language es --category utility \
    --body "Hola, tu pedido {{1}} llega mañana." --examples 8841`,
	"templates.delete":       `  yesly templates delete tpl_5Rk2Wd8Mq3Lx7B`,
	"opt_outs.list":          `  yesly opt-outs list --channel sms`,
	"opt_outs.create":        `  yesly opt-outs create --channel sms --address +34600000001`,
	"opt_outs.delete":        `  yesly opt-outs delete opt_…`,
	"webhook_endpoints.list": `  yesly webhook-endpoints list`,
	"webhook_endpoints.create": `  yesly webhook-endpoints create --url https://example.com/yesly/webhook \
    --enabled-events conversation.ended,conversation.no_reply
  yesly webhook-endpoints create --url https://example.com/yesly/webhook --enabled-events '*'`,
	"webhook_endpoints.update":      `  yesly webhook-endpoints update we_… --disabled`,
	"webhook_endpoints.delete":      `  yesly webhook-endpoints delete we_…`,
	"webhook_endpoints.test":        `  yesly webhook-endpoints test we_… --type conversation.ended`,
	"webhook_endpoints.roll_secret": `  yesly webhook-endpoints roll-secret we_…`,
	"events.list": `  yesly events list --type conversation.ended --from 2026-12-01T00:00:00+01:00
  yesly events list --conversation cnv_7Qp4Ns2Xk9Lm3D`,
	"events.retrieve":           `  yesly events retrieve evt_…`,
	"balance.retrieve":          `  yesly balance retrieve`,
	"balance.list_transactions": `  yesly balance list-transactions --type usage --agent agt_8Hd2kQ4mZ1LpXy --from 2026-12-01 --all`,
	"analytics.summary":         `  yesly analytics summary --range 30d --agent agt_8Hd2kQ4mZ1LpXy`,
}
