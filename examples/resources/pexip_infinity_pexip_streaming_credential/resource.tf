resource "pexip_infinity_pexip_streaming_credential" "example" {
  kid        = "streaming-key-001"
  public_key = file("${path.module}/keys/streaming.pub")
}
