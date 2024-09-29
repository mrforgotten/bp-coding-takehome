package monitor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	service_type "txboardcast/service/type"
)

const MONITOR_URL = "https://mock-node-wgqbnxruha-as.a.run.app/check"

func MonitorTxChannel(tx string, monChannel chan service_type.MonitorResult) {

	go func() {

		for {
			var statusResult = service_type.MonitorResult{}

			result, err := GetMonitorTx(tx)
			statusResult.Result = result

			if err != nil {
				statusResult.Err = err
			}

			monChannel <- statusResult

			if result != "PENDING" {
				break
			}

			time.Sleep(1 * time.Second)

		}

		defer close(monChannel)
	}()
}

func GetMonitorTx(tx string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%v/%v", MONITOR_URL, tx))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}
	var monitorResult *service_type.MonitorStatusOutput

	err = json.Unmarshal(body, &monitorResult)
	if err != nil {
		return "", err
	}

	return monitorResult.TxStatus, nil
}

func HandlingMonitorResult(result string) {
	switch result {
	case "CONFIRMED":
		fmt.Println("Transaction has been processed and confirmed")
	case "FAILED":
		fmt.Println("Transaction failed to process")
		os.Exit(1)
	case "PENDING":
		fmt.Println("Transaction is awaiting processing")
	case "DNE":
		fmt.Println("Transaction does not exist")
		os.Exit(1)
	}

	// fmt.Println(result)
}
