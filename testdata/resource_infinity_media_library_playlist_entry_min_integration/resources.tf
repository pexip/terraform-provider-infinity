/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_media_library_entry" "tf-test-media-entry" {
  name       = "tf-test-media-entry"
  media_file = "${path.module}/earth.mp4"
}

resource "infinity_media_library_playlist" "tf-test-playlist" {
  name = "tf-test-playlist"
}

resource "infinity_media_library_playlist_entry" "test" {
  playlist = infinity_media_library_playlist.tf-test-playlist.id
}
