/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_media_library_playlist" "media_library_playlist-test" {
  name        = "tf-test-media-library-playlist"
  description = "tf-test media library playlist description"
  loop        = true
  shuffle     = true
}
