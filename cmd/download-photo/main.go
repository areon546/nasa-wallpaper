package downloadphoto

import (
	"areon546/nasa-wallpaper/internal"
)

var c internal.Config

func main() {
	c = internal.ReadConfig()
	internal.DownloadPhoto(c)
}
