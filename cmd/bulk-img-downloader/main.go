// (c) 2020 Emir Erbasan (humanova)
// MIT License

package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tkanos/gonfig"
	"bulk-img-downloader/internal/downloader"
	"bulk-img-downloader/internal/pdfgen"
)

type Config struct {
	Url          string
	Extension    string
	StartIdx     int
	StopIdx      int
	MakePdf      bool
	GenerateUUID bool
}

func main() {
	conf := Config{}
	err := gonfig.GetConf("./config.json", &conf)
	if err != nil {
		panic(err)
	}

	dirName := "images"
	if conf.GenerateUUID {
		randId, _ := uuid.NewRandom()
		dirName = randId.String()
	}

	filenames, err := downloader.BulkDownload(conf.Url, conf.Extension, conf.StartIdx, conf.StopIdx, dirName)
	if err != nil {
		panic(err)
	}

	if conf.MakePdf {
		pdfName := fmt.Sprintf("%s_document.pdf", dirName)
		err = pdfgen.GeneratePdf(pdfName, filenames, conf.Extension)
		if err != nil {
			panic(err)
		}
	}

	return
}
