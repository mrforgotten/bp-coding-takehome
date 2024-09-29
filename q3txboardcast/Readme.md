# Transaction Broadcasting and Monitoring Client

## Prerequisite
* go cli checkout [https://go.dev/doc/install](https://go.dev/doc/install)

## Usage
### Command and running
There are flag command available, you can use it with `--help` flag.
```
# eg.
$ go run main.go --help

Usage 
  -mon
        (optional) default true. Monitoring the transaction, if not provide will return only transaction hash. It is advised to use -mon=false instead of -mon false (default true)
  -price string
        (required) Price in between 0 - 340282366920938463463374607431768211455 (or 2^128 - 1).
  -symbol string
        Any symbol in string
  -timestamp string
        (optional) Time in unix time based, current time will be used if not provided. (default "1727627963")
  -tx string
        (optional) Will override symbol,price, and timestamp flag if provided with transaction. Monitor only once, if -mon false.
```
**NOTE:** `timestamp` and `price` can be as big as 2^128 - 1, but since golang only support upto 64 bit int, so input will be string instead of int.

You can use prompt input by simple run the executable. It will prompt 3 input including `symbol`, `price`, and `timestamp`. After that it will output transaction hash and then will output result from monitoring. The monitoring will run every seconds and will stop when the output not `PENDING`.
```
# eg. run from source
$ go run main.go
symbol: ETH
price: 4500
timestamp (unix-time): 1727628861
e1b135ae9fdda9ec382c022a2a5f602492e5de8c4aa77cd2a46f5e74035b7f48
PENDING
PENDING
..
CONFIRMED
```

If you want to use command directly, you can use these example.
```
# eg. run from source
# This will output transaction with monitoring until the output not `PENDING`
$ go run main.go -symbol ETH -price 4500 -timestamp 1727628861

# in case you only need the transaction, provide `-mon=false`. This will only output transaction
$ go run main.go -symbol ETH -price 4500 -timestamp 1727628861 -mon=false

# in case you want to monitoring only, provide `-tx <hash>`
$ go run main.go -tx e1b135ae9fdda9ec382c022a2a5f602492e5de8c4aa77cd2a46f5e74035b7f48

# If you only want to check transsaction once, provide -mon=false will result in monitor only once.
$ go run main.go -tx e1b135ae9fdda9ec382c022a2a5f602492e5de8c4aa77cd2a46f5e74035b7f48 -mon=false
```

### build
```
# This will output as `txboardcast` and will sometime have extension like .exe for windows. 

go build 
```
* after build you can use the executable like the running from source command.
```
# eg.
txboardcast -symbol ETH -price 6000 -timestamp 17276123456
```

## Integration
Since this package is not available in public github, you can instead put `service` folder into your golang application module and refractor it according to your application design.<br> 
* `service/type/boardcastType.go` have 2 struct types
    1. `BoardcastInput` for using as `BoardcastTransaction` parameter
        * `Symbol` string - symbol of transaction
        * `Price` string - price number string
        * `Timestamp` string - Timestamp in unix date string
    2. `BoardcastOutput` - the output response from boardcasting
        * TxHash string 

* `service/boardcast/boardcast.go` have 3 functions
    1. `BoardcastTransaction` - using `BoardcastInput` as parameter input
    2. `ReadBoardcastInputFromFlag` - return `BoardcastInput` from flag and using `AvailableFlag` as input
    3. `ReadBoardcastInput` - use `*bufio.Reader` as input to read input and return `BoardcastInput`. in case of `price` and `timestamp`, if the input is incorrect, it will reprompt until cancel or input correctly.

* `service/type/monitorType.go` have 2 struct types
    1. `MonitorResult` - monitoring result that used as channel for yield result from monitor. It will be use to store channel in `MonitorTxChannel`
        * `Result` string
        * `Err` error
    2. `MonitorStatusOutput` - result from getting transaction monitor using `GetMonitorTx`
        * `TxStatus` string

* `service/monitor/monitor.go` have 3 functions
    1. `MonitorTxChannel` - have 2 input, transaction string, and `MonitorResult` as channel. With this, you can yield result from using this function. eg.
    * ```go
        // make channel
        resultChan := make(chan service_type.MonitorResult)
		
        // run monitor
        go monitor.MonitorTxChannel(tx, resultChan)
		
        // handle result until there is break from `MonitorTxChannel` result in no more yield output.
        for resultChan := range resultChan {
            // handle...
			monitor.PrintMonitorResult(resultChan.Result)
			if resultChan.Err != nil {
                // break if error
				log.Fatalf("Error during monitor \"%v\"\nerror: %v", tx, resultChan.Err)
				break
			}
		}
       ```
    2. `GetMonitorTx` - use string transaction as input. This function use http get and return result in string or error if there is any. 
    3. `HandlingMonitorResult` - use result string as input and print result. It can be modified to handle differently as of the code using `os.Exit(1)` which is not ideal for application use.
