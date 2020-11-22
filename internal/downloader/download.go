// (c) 2020 Emir Erbasan (humanova)
// MIT License

package downloader

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"sync"
)

type ImageFile struct {
	path string
	id   int
}

func BulkDownload(fileUrl string, fileExt string, start int, stop int, dirName string) ([]string, error) {
	var filenames []ImageFile

	// create new directory if not exists for the images
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, os.ModeDir)
		if err != nil {
			return nil, err
		}
	}

	fmt.Printf("URL : %s\nfrom [%d] -> to [%d]\n", fileUrl, start, stop)

	wg := sync.WaitGroup{}
	for i := start; i <= stop; i++ {
		wg.Add(1)
		go func(i int, fileUrl string, fileExt string, filenames *[]ImageFile) {
			downloadUrl := fmt.Sprintf(fileUrl, strconv.Itoa(i))
			fileName := fmt.Sprintf("./%s/%d.%s", dirName, i, fileExt)

			err := DownloadFile(fileName, downloadUrl)
			if err != nil {
				log.Printf(err.Error() + "/ skipping\n")
			} else {
				fmt.Printf("downloaded : %s\n", fileName)
				*filenames = append(*filenames, ImageFile{path: fileName, id: i})
			}
			wg.Done()
		}(i, fileUrl, fileExt, &filenames)
	}
	wg.Wait()

	// sort and return downloaded image's filenames
	sort.Slice(filenames, func(i, j int) bool { return filenames[i].id < filenames[j].id })

	var strFilenames []string
	for _, img := range filenames {
		strFilenames = append(strFilenames, img.path)
	}
	return strFilenames, nil
}

func DownloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("status error")
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
