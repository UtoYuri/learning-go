package main

import (
	"fmt"
	"go-module/internal/models"
	"go-module/pkg/config"
	"go-module/web/utils"
	"os"
	"path/filepath"
)

type Files map[string]os.FileInfo

func traversalFolder(folder string) (Files, error) {
	var files = make(Files)
	err := filepath.Walk(folder, func(pathname string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files[pathname] = info
		return nil
	})
	return files, err
}

func main() {
	if err := config.Init(); err != nil {
		fmt.Println("Init failed,", err)
		return
	}

	localResourceDir := os.Getenv("LOCAL_RESOURCE_DIR")
	if localResourceDir == "" {
		fmt.Println("LOCAL_RESOURCE_DIR not provided")
		return
	}

	files, err := traversalFolder(localResourceDir)
	if err != nil {
		fmt.Println("Index failed,", err)
		return
	}

	for pathname := range files {
		kind, suffix := utils.ParseFileFormat(files[pathname].Name())
		media := &models.Media{
			Title: files[pathname].Name(),
			Kind: kind,
			Suffix: suffix,
			Size: uint64(files[pathname].Size()),
			Url: fmt.Sprintf("/%s", pathname),
		}
		err := models.CreateMedia(media)
		if err != nil {
			fmt.Printf("Create media<%s> failed, %v\n", pathname, err)
		}

		fmt.Printf("Index media success, <%s>\n", pathname)
	}
}
