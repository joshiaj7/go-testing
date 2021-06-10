// main_test.go
package spotify

import (
	"testing"

	"github.com/zmb3/spotify"
)

// 1
type mockSpotifyClient struct{}

// 2
func (m *mockSpotifyClient) GetPlaylist(playlistID spotify.ID) (*spotify.FullPlaylist, error) {
	// 3
	return &spotify.FullPlaylist{
		SimplePlaylist: spotify.SimplePlaylist{
			Name: "whatever",
		},
	}, nil
}

func Test_NewGetPlaylistName(t *testing.T) {
	// 4
	client := &mockSpotifyClient{}

	// 5
	name := getPlaylistName(client, "whatever")

	if name != "whatever" {
		t.Errorf("expected %s, got %s", "whatever", name)
	}
}
