resource "pexip_infinity_recurring_conference" "example" {
  conference    = pexip_infinity_conference.example.id
  current_index = 0
  ews_item_id   = "AAMkAGVmExampleRecurringEwsItemId"
  is_depleted   = false
  subject       = "Weekly Team Standup"
}
