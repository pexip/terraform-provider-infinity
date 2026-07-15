/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_media_library_entry" "tf-test-media-entry" {
  name        = "tf-test-media-entry"
  description = "Test media entry for playlist"
  media_file  = "${path.module}/earth.mp4"
}

resource "infinity_media_library_playlist" "tf-test-playlist" {
  name        = "tf-test-playlist"
  description = "Test playlist"
  loop        = false
  shuffle     = false
}

resource "infinity_media_library_playlist_entry" "test" {
  entry_type = "MEDIA"
  media      = infinity_media_library_entry.tf-test-media-entry.id
  playlist   = infinity_media_library_playlist.tf-test-playlist.id
  position   = 5
  playcount  = 3
}
