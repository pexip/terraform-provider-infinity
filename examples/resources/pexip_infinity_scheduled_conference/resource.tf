resource "pexip_infinity_scheduled_conference" "example" {
  conference   = "/configuration/v1/conference/1/"
  start_time   = "2024-12-01T10:00:00Z"
  end_time     = "2024-12-01T11:00:00Z"
  subject      = "Weekly Team Meeting"
  ews_item_id  = "AAMkAGVm...ExampleEwsItemId"
  ews_item_uid = "040000008200E00074C5B7101A82E00800000000"
}
