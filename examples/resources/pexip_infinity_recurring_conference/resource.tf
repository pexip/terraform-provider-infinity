resource "pexip_infinity_recurring_conference" "example" {
  conference      = "/configuration/v1/conference/1/"
  current_index   = 0
  ews_item_id     = "AAMkAGVm...ExampleRecurringEwsItemId"
  is_depleted     = false
  subject         = "Weekly Team Standup"
  scheduled_alias = "/configuration/v1/scheduled_alias/1/"
}
