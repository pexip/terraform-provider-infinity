resource "pexip_infinity_media_library_playlist_entry" "example" {
  playlist = pexip_infinity_media_library_playlist.example.id
  media    = pexip_infinity_media_library_entry.example.id
  position = 1
}
