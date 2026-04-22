locals {
  departments = {
    engineering = { alias = "engineering", description = "Engineering team conference" }
    marketing   = { alias = "marketing", description = "Marketing team conference" }
    sales       = { alias = "sales", description = "Sales team conference" }
  }
}

resource "pexip_infinity_conference" "departments" {
  for_each     = local.departments
  name         = each.key
  service_type = "conference"
}

resource "pexip_infinity_conference_alias" "department_aliases" {
  for_each    = local.departments
  alias       = each.value.alias
  description = each.value.description
  conference  = pexip_infinity_conference.departments[each.key].id
}
