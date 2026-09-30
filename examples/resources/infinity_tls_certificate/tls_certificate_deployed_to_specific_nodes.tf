# Get node resource URIs
data "pexip_infinity_node" "conferencing_nodes" {
  count = length(var.conferencing_node_ids)
  id    = var.conferencing_node_ids[count.index]
}

resource "pexip_infinity_tls_certificate" "node_specific" {
  certificate = file("${path.module}/certificates/conferencing.crt")
  private_key = file("${path.module}/certificates/conferencing.key")
  parameters  = "TLS certificate for conferencing nodes"
  nodes = [
    for node in data.pexip_infinity_node.conferencing_nodes : node.id
  ]
}
