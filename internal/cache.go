package internal

import (
	"areon546/nasa-wallpaper/pkg/net"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/areon546/go-files/files"
)

var (
	cacheDirectory string
)

func init() {
	// set cache directory
	setupCacheDirectory()
}

func setupCacheDirectory() {
	cache, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	cacheDirectory = filepath.Join(cache, "nasa-wallpapers") + "/"

	// ensure exists :
	err = files.MakeDirectory(cacheDirectory)
	if err != nil {
		panic(err)
	}
}

// makes sure that the filename being processed is relative to cache directory
func CacheFilename(fn string) string {
	fn = filepath.Clean(fn) // clean out any relative pathing
	if !strings.HasPrefix(fn, cacheDirectory) {
		return cacheDirectory + fn
	}

	return fn
}

// CacheURL caches the given url as the filename specified
func CacheURL(filename, url string) {
	net.DownloadFile(CacheFilename(filename), url)
}

// IsCached checks if a given filename exists within the
func IsCached(filename string) bool {
	filename = CacheFilename(filename)

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		// not exists
		return false
	}

	return true
}
