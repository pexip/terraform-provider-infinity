resource "pexip_infinity_media_library_entry" "example" {
  name        = "example"
  media_file  = "${path.module}/example.mp4"
}

resource "pexip_infinity_media_library_playlist" "example_playlist" {
  name        = "example_playlist"
}

resource "pexip_infinity_media_library_playlist_entry" "example_playlist_entry" {
  playlist = pexip_infinity_media_library_playlist.example_playlist.id
  media    = pexip_infinity_media_library_entry.example.id
}

resource "pexip_infinity_conference" "media_playback" {
  name        = "Media Playback Example"
  service_type = "media_playback"
  media_playlist = pexip_infinity_media_library_playlist.example_playlist.id
  on_completion = jsonencode({
    disconnect = true
  })
  aliases = [
    { alias = "exampleplayback" },
    { alias = "exampleplayback@example.com" },
  ]
}