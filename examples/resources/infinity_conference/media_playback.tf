resource "infinity_media_library_entry" "example" {
  name        = "example"
  media_file  = "${path.module}/example.mp4"
}

resource "infinity_media_library_playlist" "example_playlist" {
  name        = "example_playlist"
}

resource "infinity_media_library_playlist_entry" "example_playlist_entry" {
  playlist = infinity_media_library_playlist.example_playlist.id
  media    = infinity_media_library_entry.example.id
}

resource "infinity_conference" "media_playback" {
  name        = "Media Playback Example"
  service_type = "media_playback"
  media_playlist = infinity_media_library_playlist.example_playlist.id
  on_completion = jsonencode({
    disconnect = true
  })
  aliases = [
    { alias = "exampleplayback" },
    { alias = "exampleplayback@example.com" },
  ]
}