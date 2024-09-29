package boardcast

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	service_type "txboardcast/service/type"
	"txboardcast/service/util"
)

const BOARDCAST_URL = "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"

func BoardcastTransaction(boardcastInput service_type.BoardcastInput) (string, error) {
	jsonBytePayload := []byte(fmt.Sprintf(`{"symbol": "%v", "price": %v, "timestamp": %v}`, boardcastInput.Symbol, boardcastInput.Price, boardcastInput.Timestamp))

	resp, err := http.Post(BOARDCAST_URL, "application/json", bytes.NewReader(jsonBytePayload))
	if err != nil {
		log.Fatalf("Error making Post boardcast request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// log.Println("Error transform body: ", err)
		return "", err
	}

	var result *service_type.BoardcastOutput
	err = json.Unmarshal(body, &result)
	if err != nil {
		// log.Println("Error transform body: ", err)
		return "", err
	}

	return result.TxHash, nil
}

func ReadBoardcastInputFromFlag(flags *service_type.AvailableFlag) *service_type.BoardcastInput {
	input := &service_type.BoardcastInput{}
	if *flags.Price == "" {
		log.Fatal("error price flag is required\n")
		os.Exit(1)
	}

	isInvalidateInput := false
	_, err := util.ValidateLongIntString(*flags.Price)
	if err != nil {
		log.Fatal("error price input: ", err.Error())
		isInvalidateInput = true
	}
	_, err = util.ValidateLongIntString(*flags.Timestamp)
	if err != nil {
		log.Fatal("error timestamp input: ", err.Error())
		isInvalidateInput = true
	}
	if isInvalidateInput {
		os.Exit(1)
	}

	input = &service_type.BoardcastInput{
		Symbol:    *flags.Symbol,
		Price:     *flags.Price,
		Timestamp: *flags.Timestamp,
	}

	return input
}

func ReadBoardcastInput(reader *bufio.Reader) *service_type.BoardcastInput {
	input := &service_type.BoardcastInput{}
	var symbol, priceString, timestampString string
	fmt.Print("symbol: ")
	symbol = util.ReadLine(reader)
	for {
		fmt.Print("price: ")
		priceString = util.ReadLine(reader)
		isValid, err := util.ValidateLongIntString(priceString)
		if err != nil {
			fmt.Println("Error price input: ", err.Error())
		}
		if isValid {
			break
		}
	}

	for {
		fmt.Print("timestamp (unix-time): ")
		timestampString = util.ReadLine(reader)
		isValid, err := util.ValidateLongIntString(timestampString)
		if err != nil {
			fmt.Println("Error timestamp input: ", err.Error())
		}
		if isValid {
			break
		}
	}

	input.Symbol = symbol
	input.Price = priceString
	input.Timestamp = timestampString

	return input
}
