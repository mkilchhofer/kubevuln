package main

import (
	colim "ca-vuln-scan/goroutinelimits"
	"ca-vuln-scan/process_request"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	sysreport "github.com/armosec/logger-go/system-reports/datastructures"
	pkgcautils "github.com/armosec/utils-k8s-go/armometadata"
	"github.com/golang/glog"

	wssc "github.com/armosec/armoapi-go/apis"
)

var goroutineLimit *colim.CoroutineGuardian

func startScanImage(scanCmd *wssc.WebsocketScanCommand) {
	log.Printf("Scan request to image %s is put on processing queue", scanCmd.ImageTag)

	goroutineLimit.Wait()
	go func() {
		log.Printf("ProcessScanRequest for jobid %v/%v %s image: %s starting", scanCmd.ParentJobID, scanCmd.JobID, scanCmd.Wlid, scanCmd.ImageTag)

		_, err := process_request.ProcessScanRequest(scanCmd)
		if err != nil {
			log.Printf("ProcessScanRequest for jobid %v/%v %s image: %s failed due to: %s", scanCmd.ParentJobID, scanCmd.JobID, scanCmd.Wlid, scanCmd.ImageTag, err.Error())
		}
		goroutineLimit.Release()
	}()

}

func scanImage(w http.ResponseWriter, req *http.Request) {
	var WebsocketScan wssc.WebsocketScanCommand

	if req.Method == http.MethodPost {
		err := json.NewDecoder(req.Body).Decode(&WebsocketScan)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("fail decode json from web socket, error %v", err)
			return
		}
		if WebsocketScan.ImageTag == "" && WebsocketScan.ImageHash == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("image tag and image hash are missing")
			return
		}
		if WebsocketScan.IsScanned {
			w.WriteHeader(http.StatusAccepted)
			log.Printf("the image %s already scanned", WebsocketScan.ImageTag)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "scan request accepted\n")
		// Backend aggregation depends on this report!!!
		// don't change any parameter before checking with BE side first!!!!
		report := &sysreport.BaseReport{
			CustomerGUID: os.Getenv("CA_CUSTOMER_GUID"),
			Reporter:     "ca-vuln-scan",
			Status:       sysreport.JobStarted,
			Target:       fmt.Sprintf("vuln scan:: scanning wlid: %v ,containerName: %v imageTag: %v imageHash: %s", WebsocketScan.Wlid, WebsocketScan.ContainerName, WebsocketScan.ImageTag, WebsocketScan.ImageHash),
			ActionID:     "1",
			ActionIDN:    1,
			ActionName:   "vuln scan",
			JobID:        WebsocketScan.JobID,
			ParentAction: WebsocketScan.ParentJobID,
			Details:      "Inqueueing",
		}
		report.SendAsRoutine([]string{}, true)
		// End of Backend must not change report
		startScanImage(&WebsocketScan)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "unsupported method\n")
	}

}

func main() {
	process_request.CreateAnchoreResourcesDirectoryAndFiles()
	flag.Parse()

	displayBuildTag()

	pkgcautils.LoadConfig("", true)

	scanRoutinslimitStr := os.Getenv("CA_MAX_VULN_SCAN_ROUTINS")
	scanRoutinslimit := colim.MAX_VULN_SCAN_ROUTINS
	if len(scanRoutinslimitStr) != 0 {
		if i, err := strconv.Atoi(scanRoutinslimitStr); err == nil {
			scanRoutinslimit = i
		}
	}
	goroutineLimit, _ = colim.CreateCoroutineGuardian(scanRoutinslimit)

	uri := "/" + wssc.WebsocketScanCommandVersion + "/" + wssc.WebsocketScanCommandPath
	log.Printf("uri %v", uri)
	http.HandleFunc(uri, scanImage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func displayBuildTag() {
	imageVersion := "local build"
	dat, err := ioutil.ReadFile("./build_number.txt")
	if err == nil {
		imageVersion = string(dat)
	} else {
		dat, err = ioutil.ReadFile("./build_date.txt")
		if err == nil {
			imageVersion = fmt.Sprintf("%s, date: %s", imageVersion, string(dat))
		}
	}
	glog.Infof("Image version: %s", imageVersion)
}
