package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	filename    string
	downloadUrl string
)

func main() {
	checkOs()
	downloadUpdate()

	fmt.Println("Unzipping file...")

	Unzip(filename, strings.ReplaceAll(filename, ".zip", ""))

	os.Remove(filename)

	fmt.Println("Done")
}

func checkOs() {
	if runtime.GOOS == "windows" {
		if runtime.GOARCH == "amd64" {
			filename = "Random-Good-Hanime-windows-amd64"
		}
		if runtime.GOARCH == "arm" {
			filename = "Random-Good-Hanime-windows-arm"
		}
	}
	if runtime.GOOS == "linux" {
		if runtime.GOARCH == "amd64" {
			filename = "Random-Good-Hanime-linux-amd64"
		}
		if runtime.GOARCH == "arm" {
			filename = "Random-Good-Hanime-linux-arm"
		}
	}
	if runtime.GOOS == "darwin" {
		if runtime.GOARCH == "amd64" {
			filename = "Random-Good-Hanime-darwin-amd64"
		}
		if runtime.GOARCH == "arm" {
			fmt.Println("Error, Darwin arm not supported")
			os.Exit(1)
		}
	}

	fmt.Println("Currently running on: " + runtime.GOOS + " " + runtime.GOARCH)

	filename = filename + ".zip"
}

func downloadUpdate() {
	api := "https://api.github.com/repos/superredstone/random-good-hanime/releases"

	response, err := http.Get(api)
	if err != nil {
		fmt.Println(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var data []Githubapi
	json.Unmarshal([]byte(responseData), &data)

	for i := 0; i < len(data[0].Assets); i++ {
		if data[0].Assets[i].Name == filename {
			fmt.Println("Downloading: " + data[0].Assets[i].Url)
			downloadUrl = data[0].Assets[i].Url

			break
		}
	}

	err = DownloadFile(filename, downloadUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded new version: ", filename)

	fmt.Println("Download complete")
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

type Githubapi struct {
	Assets []Assets `json:"assets"`
}

type Assets struct {
	Url  string `json:"browser_download_url"`
	Name string `json:"name"`
}
