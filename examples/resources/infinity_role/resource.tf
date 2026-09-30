resource "pexip_infinity_role" "example" {
  name = "Conference Manager"
  permissions = [
    "conference.create",
    "conference.read",
    "conference.update",
    "conference.delete"
  ]
}
