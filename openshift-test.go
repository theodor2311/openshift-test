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
	fmt.Fprintf(res, "OpenShift Test, Hostname: %s", os.Getenv("HOSTNAME"))
}
