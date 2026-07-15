resource "infinity_media_library_playlist_entry" "example" {
  playlist = infinity_media_library_playlist.example.id
  media    = infinity_media_library_entry.example.id
  position = 1
}
