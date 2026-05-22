resource "pexip_infinity_policy_server" "local_media_policy_burst" {
  name                                    = "Local Media Policy Server Overflow with Burst Locations"
  internal_media_location_policy_template = <<-EOF
    {
      {% if call_info.location == "DC_EAST" %}
        "result" : {
        "location" : "DC_EAST",
        "overflow_locations" : [ "DC_EAST_BURST", "DC_WEST", "DC_WEST_BURST" ]
        }
      {% elif call_info.location == "DC_WEST" %}
        "result" : {
        "location" : "DC_WEST",
        "overflow_locations" : [ "DC_WEST_BURST", "DC_EAST", "DC_EAST_BURST" ]
        }
      {% else %}
        "result" : {{suggested_media_overflow_locations|pex_to_json}}
      {% endif %}
    }
  EOF
}
