resource "pexip_infinity_user_group" "example" {
  name        = "Administrators"
  description = "System administrators group"
  users = [
    "/configuration/v1/end_user/1/",
    "/configuration/v1/end_user/2/"
  ]
  user_group_entity_mappings = [
    "/configuration/v1/user_group_entity_mapping/1/"
  ]
}
