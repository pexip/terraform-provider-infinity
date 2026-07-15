resource "infinity_user_group_entity_mapping" "example" {
  entity_resource_uri = infinity_conference.example.id
  user_group          = infinity_user_group.example.id
}
