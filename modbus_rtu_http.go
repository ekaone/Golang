// HTTP to Modbus RTU

package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"net/http"

	"github.com/goburrow/modbus"
	"github.com/gorilla/mux"
)

func main() {
	/*Configure RTU Client*/
	rtu := modbus.NewRTUClientHandler("/dev/ttyUSB0")
	rtu.SlaveId = 1
	rtu.BaudRate = 38400
	rtu.DataBits = 8
	rtu.Parity = "O"
	rtu.StopBits = 1
	rtu.Timeout = 3 * time.Second

	err := rtu.Connect()
	if err != nil {
		log.Fatalf("RTU Connect: %v", err)
	}
	defer rtu.Close()

	cli := modbus.NewClient(rtu)

	/*Define HTTP Handlers*/
	router := mux.NewRouter()

	/*Input Handler*/
	router.HandleFunc("/inputs/{input:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			/*Parse Variables*/
			vars := mux.Vars(r)
			input, _ := strconv.Atoi(vars["input"])

			/*Read Single Input*/
			res, err := cli.ReadDiscreteInputs(uint16(input), 1)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("Error: %v", err)
				return
			}

			/*Write response (LSB of return byte)*/
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%v", res[0]&0x01)
		}).Methods("GET")

	/*Output (coil) Handler*/
	router.HandleFunc("/outputs/{output:[0-9]+}/{value:[01]}",
		func(w http.ResponseWriter, r *http.Request) {
			/*Parse Variables*/
			vars := mux.Vars(r)
			output, _ := strconv.Atoi(vars["output"])
			value, _ := strconv.Atoi(vars["value"])
			value *= 0xFF00 //0xFF00 = ON, 0x0000 = OFF

			/*Write Single Coil*/
			_, err := cli.WriteSingleCoil(uint16(output), uint16(value))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("Error: %v", err)
				return
			}
			w.WriteHeader(http.StatusOK)
		}).Methods("POST")

	err = http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatalf("HTTP Server : %v", err)
	}
}
