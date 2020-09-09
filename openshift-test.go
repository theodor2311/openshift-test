package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", OpenShiftTest)
	http.ListenAndServe(":8080", nil)
}

// OpenShiftTest Handler
func OpenShiftTest(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>OpenShift Test Version 1</h1>\n\nPod Name: %s \nNamespace: %s\nIP Address: %s\nNode Name: %s\nService Account: %s", os.Getenv("POD_NAME"), os.Getenv("POD_NAMESPACE"), os.Getenv("POD_IP"), os.Getenv("NODE_NAME"), os.Getenv("POD_SERVICE_ACCOUNT"))
}
