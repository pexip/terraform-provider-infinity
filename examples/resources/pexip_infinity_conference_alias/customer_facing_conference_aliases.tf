# Customer support conference
resource "pexip_infinity_conference_alias" "support_main" {
  alias       = "support"
  description = "Main customer support conference"
  conference  = data.pexip_infinity_conference.customer_support.id
}

# Sales demo conference
resource "pexip_infinity_conference_alias" "sales_demo" {
  alias       = "demo"
  description = "Sales demonstration conference"
  conference  = data.pexip_infinity_conference.sales_demo.id
}

# Training sessions
resource "pexip_infinity_conference_alias" "training" {
  alias       = "training"
  description = "Customer training sessions"
  conference  = data.pexip_infinity_conference.training_room.id
}

# Webinar alias
resource "pexip_infinity_conference_alias" "webinar" {
  alias       = "webinar"
  description = "Monthly product webinar"
  conference  = data.pexip_infinity_conference.monthly_webinar.id
}
