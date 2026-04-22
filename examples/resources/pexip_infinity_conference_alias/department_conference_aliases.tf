variable "departments" {
  type = list(object({
    name        = string
    alias       = string
    conference  = string
    description = string
  }))
  default = [
    {
      name        = "Engineering"
      alias       = "engineering"
      conference  = "/api/admin/configuration/v1/conference/42/"
      description = "Engineering team conference"
    },
    {
      name        = "Marketing"
      alias       = "marketing"
      conference  = "/api/admin/configuration/v1/conference/43/"
      description = "Marketing team conference"
    },
    {
      name        = "Sales"
      alias       = "sales"
      conference  = "/api/admin/configuration/v1/conference/44/"
      description = "Sales team conference"
    }
  ]
}

resource "pexip_infinity_conference_alias" "department_aliases" {
  count       = length(var.departments)
  alias       = var.departments[count.index].alias
  description = var.departments[count.index].description
  conference  = var.departments[count.index].conference
}
