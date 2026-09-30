resource "pexip_infinity_user_group_entity_mapping" "example" {
  entity_resource_uri = pexip_infinity_conference.example.id
  user_group          = pexip_infinity_user_group.example.id
}
