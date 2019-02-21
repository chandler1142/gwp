package metrics

import (
	"github.com/go-metrics"
	"log"
	"os"
	"time"
)

var sendMessageCount metrics.Counter
var receivedMessageCount metrics.Counter
var SendRate metrics.Meter
var ReceivedRate metrics.Meter

func init() {
	sendMessageCount = metrics.NewCounter()
	receivedMessageCount = metrics.NewCounter()
	SendRate = metrics.NewMeter()
	ReceivedRate = metrics.NewMeter()

	metrics.Register("sendMessageCount", sendMessageCount)
	metrics.Register("receivedMessageCount", receivedMessageCount)
	metrics.Register("SendRate", SendRate)
	metrics.Register("receivedRate", ReceivedRate)


}

func RecordSendMessageCount() {
	sendMessageCount.Inc(1)
	SendRate.Mark(1)
}

func RecordReceivedMessageCount() {
	receivedMessageCount.Inc(1)
	ReceivedRate.Mark(1)
}

func PrintLog() {
	go metrics.Log(metrics.DefaultRegistry,
		1 * time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))
}
