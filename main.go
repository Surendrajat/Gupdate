package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"runtime"

	"github.com/cavaliercoder/grab"
	"github.com/mholt/archiver/v3"
	"github.com/otiai10/copy"
)

const (
	goDwnURL   string = "https://dl.google.com/go/"
	golangURL1 string = "https://golang.org/"
	golangURL2 string = "https://golang.org/dl/"
)

var tmpDir = os.TempDir() + string(os.PathSeparator)

func main() {
	version, _ := getLatestVersion(golangURL1)
	version2, _ := getLatestVersion(golangURL2)
	if version == "" && version2 == "" {
		fmt.Println("Couldn't fetch latest version")
		return
	}
	if version < version2 || version == "" {
		version = version2
	}
	fmt.Println("Latest version:", version)
	fmt.Println("Current version:", runtime.Version())
	if version == runtime.Version() {
		fmt.Println("Already up to date")
		return
	}
	fmt.Println("downloading latest version...")
	fileURL, err := downloadGo(version)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("extracting downloaded archive...")
	err = archiver.Unarchive(fileURL, tmpDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("updating current go binaries...")
	goRootDir := runtime.GOROOT()
	err = os.RemoveAll(goRootDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.MkdirAll(goRootDir, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = copy.Copy(tmpDir+"go", goRootDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("updated go to the latest version!!!")
}

func getLatestVersion(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	versionRegex := regexp.MustCompile(`go1\.[0-9]+(.[0-9])?[0-9]*`)
	return versionRegex.FindString(string(body)), nil
}

func downloadGo(version string) (string, error) {
	dwnLink := fmt.Sprintf("%v%v.%v-%v.tar.gz", goDwnURL, version, runtime.GOOS, runtime.GOARCH)
	fileURL := fmt.Sprintf("%v%v.%v-%v.tar.gz", tmpDir, version, runtime.GOOS, runtime.GOARCH)
	resp, err := grab.Get(fileURL, dwnLink)
	if err != nil {
		return "", err
	}
	return resp.Filename, nil
}
