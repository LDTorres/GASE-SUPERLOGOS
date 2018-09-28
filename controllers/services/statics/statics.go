package statics

import (
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
)

var (
	RootDir, _      = filepath.Abs(beego.AppConfig.String("assets::jumps"))
	ImageFolderPath = beego.AppConfig.String("assets::imageFolderPath")
	ImageFolderDir  = RootDir + "/" + ImageFolderPath
)

func init() {
	checkOrCreateImagesFolder(ImageFolderDir)
}

func checkOrCreateImagesFolder(imageFolderDir string) (err error) {

	if _, err := os.Stat(imageFolderDir); os.IsNotExist(err) {
		os.MkdirAll(imageFolderDir, 644)
	}

	return

}
