package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/chfanghr/librespot/Spotify"
)

func GetCoverArtUrls(track *Spotify.Track) (res map[Spotify.Image_Size]string) {
	res = make(map[Spotify.Image_Size]string)
	for _, v := range track.GetAlbum().GetCoverGroup().GetImage() {
		url := fmt.Sprintf("https://i.scdn.co/image/%032s", hex.EncodeToString(v.GetFileId()))
		res[v.GetSize()] = url
	}
	return
}
