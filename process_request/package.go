package process_request

import (
	"fmt"
	"strings"
)

type PackageHandler interface {
	initPackageHandler(image *OciImage) error
	readFileListForPackage(packageName string) (*[]string, error)
}

type dpkgPackageHandler struct {
	dpkgDirectoryList *OciImageFsList
	image             *OciImage
}

func readFileListForPackageDpkg(packageName string, image *OciImage) (*[]string, error) {
	var packageFileList = []string{"/var/lib/dpkg/info/%s.list", "/var/lib/dpkg/info/%s:amd64.list"}
	for _, name := range packageFileList {
		listFilePath := fmt.Sprintf(name, packageName)
		fileContent, err := image.GetFile(listFilePath)
		if err == nil {
			fileList := strings.Split(string(*fileContent), "\n")
			return &fileList, nil
		}
	}
	return nil, fmt.Errorf("Could not find package %s", packageName)
}

func readFileListForPackageApk(packageName string, image *OciImage) (*[]string, error) {
	return nil, fmt.Errorf("unsupported Apk")
}

func readFileListForPackage(packageName string, packageManagerType string, image *OciImage) (*[]string, error) {
	var err error
	var fileList *[]string
	switch packageManagerType {
	case "dpkg":
		fileList, err = readFileListForPackageDpkg(packageName, image)
	case "apk":
		fileList, err = readFileListForPackageApk(packageName, image)
	default:
		fileList = nil
		err = fmt.Errorf("Unsupported packager type %s", packageManagerType)
	}
	return fileList, err
}

func CreatePackageHandler(packageManagerType string, image *OciImage) (PackageHandler, error) {
	var err error
	var packageHandler PackageHandler
	switch packageManagerType {
	case "dpkg":
		packageHandler = &dpkgPackageHandler{}
		err = packageHandler.initPackageHandler(image)
	//case "apk":
	//	err = nil
	default:
		err = fmt.Errorf("Unsupported packager type %s", packageManagerType)
	}
	return packageHandler, err
}

func (ph *dpkgPackageHandler) initPackageHandler(image *OciImage) error {
	ph.image = image
	// Pre-read the list of files in /var/lib/dpkg/info
	fileContent, err := image.ListDirectoryFile("/var/lib/dpkg/info", false, false)
	if err == nil {
		ph.dpkgDirectoryList = fileContent
		return nil
	}
	return fmt.Errorf("Error initialization in dpkg handler %s", err)
}

func (ph *dpkgPackageHandler) readFileListForPackage(packageName string) (*[]string, error) {
	// Check package name
	for _, fsEntry := range *ph.dpkgDirectoryList {
		var packageFileList = []string{fmt.Sprintf("var/lib/dpkg/info/%s.list", packageName), fmt.Sprintf("var/lib/dpkg/info/%s:amd64.list", packageName)}
		for _, packageFileName := range packageFileList {
			//log.Printf("%s ?= %s", packageFileName, fsEntry.path)
			if packageFileName == fsEntry.Path {
				// gotcha!
				fileContent, err := ph.image.GetFile("/" + fsEntry.Path)
				if err != nil {
					return nil, err
				} else {
					fileList := strings.Split(string(*fileContent), "\n")
					return &fileList, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("Not found package %s", packageName)
}
