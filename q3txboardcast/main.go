package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
	"txboardcast/service/boardcast"
	"txboardcast/service/monitor"
	service_type "txboardcast/service/type"
	"txboardcast/service/util"
)

func main() {
	flags := &service_type.AvailableFlag{}
	flags.Symbol = flag.String("symbol", "", "[string] Any symbol in string") // api allowed empty string "" somehow
	flags.Price = flag.String("price", "", fmt.Sprintf("(required) [string] Price in between 0 - %v (or 2^128 - 1).", util.MAX_ALLOW_STRING_NUMBER))
	flags.Timestamp = flag.String("timestamp", fmt.Sprintf("%v", time.Now().Unix()), "(optional) Time in unix time based, current time will be used if not provided.")
	flags.Mon = flag.Bool("mon", true, "(optional) [boolean] default true. Monitoring the transaction, if not provide will return only transaction hash. It is advised to use -mon=false instead of -mon false")
	flags.Tx = flag.String("tx", "", "(optional) [string] Will override symbol,price, and timestamp flag if provided with transaction. Monitor only once, if -mon=false.")
	flag.Parse()

	var b *service_type.BoardcastInput
	args := os.Args

	argsLen := len(args)

	if *flags.Tx == "" {
		// no transaction mode

		if argsLen == 1 {
			// case no flag
			buffReader := bufio.NewReader(os.Stdin)
			b = boardcast.ReadBoardcastInput(buffReader)
		} else {
			// case there is flag
			b = boardcast.ReadBoardcastInputFromFlag(flags)
		}
		tx, err := boardcast.BoardcastTransaction(*b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v", tx)
		fmt.Println()
		flags.Tx = &tx
	} else {
		// case tx is provided in flag, but not monitor. So it should get only one time result
		if !*flags.Mon {
			resultChan := make(chan service_type.MonitorResult)

			monitor.MonitorTxChannel(*flags.Tx, resultChan)
			result := <-resultChan
			monitor.HandlingMonitorResult(result.Result)
			if result.Err != nil {
				log.Fatalf("Error during monitor \"%v\"\nerror: %v", *flags.Tx, result.Err)
			}
		}
	}

	if *flags.Mon {
		tx := *flags.Tx

		resultChan := make(chan service_type.MonitorResult)
		go monitor.MonitorTxChannel(tx, resultChan)
		for resultChan := range resultChan {
			monitor.HandlingMonitorResult(resultChan.Result)
			if resultChan.Err != nil {
				log.Fatalf("Error during monitor \"%v\"\nerror: %v", tx, resultChan.Err)
				break
			}
		}
	}

}
