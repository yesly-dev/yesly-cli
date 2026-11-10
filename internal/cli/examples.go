package cli

// examples are the Example section of each operation's help, by resource.method.
// TestEveryOperationHasAnExample keeps the list complete.
var examples = map[string]string{
	"appointments.list": `  yesly appointments list --status no_reply --date 2026-11-13
  yesly appointments list --metadata route_id=BIO-07 --all
  yesly appointments list --external-id ORD-48213`,
	"appointments.create": `  yesly appointments create --external-id ORD-48213 --type delivery \
    --customer.name Martha --customer.phone +34612345678 \
    --slot.start 2026-11-13T09:00:00+01:00 --slot.end 2026-11-13T11:00:00+01:00 \
    --metadata order_id=PV-18342 --metadata route_id=BIO-07
  yesly appointments create --data @appointment.json`,
	"appointments.retrieve": `  yesly appointments retrieve apt_…`,
	"appointments.update": `  yesly appointments update apt_… --slot.start 2026-11-14T09:00:00+01:00 --slot.end 2026-11-14T11:00:00+01:00
  yesly appointments update apt_… --metadata route_id=BIO-09`,
	"appointments.count": `  yesly appointments count --date 2026-11-13
  yesly appointments count --from 2026-11-01 --to 2026-11-30 --type delivery`,
	"appointments.recording": `  yesly appointments recording apt_… att_…`,
	"appointments.cancel":    `  yesly appointments cancel apt_… --reason "Shop closed"`,
	"appointments.import": `  yesly appointments import --appointments @appointments.json
  yesly appointments import --data '{"appointments": [ … ]}'`,
	"imports.retrieve":       `  yesly imports retrieve imp_…`,
	"slots.list":             `  yesly slots list --type delivery --from 2026-11-13 --to 2026-11-20 --available`,
	"slots.create":           `  yesly slots create --slots @slots.json`,
	"slots.delete":           `  yesly slots delete slot_…`,
	"appointment_types.list": `  yesly appointment-types list`,
	"appointment_types.create": `  yesly appointment-types create --id delivery --name "Home delivery" \
    --time-format window --place customer --first-contact '{"days_before": 1, "at": "18:00"}' \
    --reschedule.min-notice-hours 24 --reschedule.window-days 14
  yesly appointment-types create --id repair --name "Boiler repair" \
    --time-format exact --place customer --first-contact '{"hours_before": 3}'`,
	"appointment_types.update": `  yesly appointment-types update delivery --instructions "Have the order number at hand" --reschedulable=false`,
	"locations.list":           `  yesly locations list`,
	"locations.create": `  yesly locations create --id bilbao --name "Bilbao depot" \
    --address "Calle Mayor 1, 48001 Bilbao" --transfer-phone +34944000000`,
	"locations.update":       `  yesly locations update bilbao --transfer-phone +34944000001`,
	"webhook_endpoints.list": `  yesly webhook-endpoints list`,
	"webhook_endpoints.create": `  yesly webhook-endpoints create --url https://example.com/yesly/webhook \
    --enabled-events appointment.confirmed,appointment.no_reply
  yesly webhook-endpoints create --url https://example.com/yesly/webhook --enabled-events '*'`,
	"webhook_endpoints.update":      `  yesly webhook-endpoints update we_… --disabled`,
	"webhook_endpoints.delete":      `  yesly webhook-endpoints delete we_…`,
	"webhook_endpoints.test":        `  yesly webhook-endpoints test we_… --type appointment.confirmed`,
	"webhook_endpoints.roll_secret": `  yesly webhook-endpoints roll-secret we_…`,
	"events.list":                   `  yesly events list --type 'appointment.*' --from 2026-11-07T00:00:00Z`,
	"events.retrieve":               `  yesly events retrieve evt_…`,
	"balance.retrieve":              `  yesly balance retrieve`,
	"balance.list_transactions":     `  yesly balance list-transactions --type usage --from 2026-11-01 --all`,
	"analytics.retrieve_summary":    `  yesly analytics retrieve-summary --range 30d --type delivery`,
}
