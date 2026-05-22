# Corporate domain aliases
resource "pexip_infinity_conference_alias" "corp_main" {
  alias       = "corp.main"
  description = "Corporate main conference"
  conference  = data.pexip_infinity_conference.corporate_main.id
}

# Subsidiary domain aliases
resource "pexip_infinity_conference_alias" "subsidiary_a" {
  alias       = "suba.meeting"
  description = "Subsidiary A conference"
  conference  = data.pexip_infinity_conference.subsidiary_a.id
}

# Partner domain aliases
resource "pexip_infinity_conference_alias" "partner_collab" {
  alias       = "partner.collaboration"
  description = "Partner collaboration conference"
  conference  = data.pexip_infinity_conference.partner_room.id
}
