package metrics

//var delayMap = make(map[string]int64)
//var maxCost int64 = 0
//
//var SendChan = make(chan xmpp.Chat, 8096)
//var ReceiveChan = make(chan xmpp.Chat, 8096)
//
//var sendCount int = 0
//var receiveCount int = 0
//var Lock = sync.Mutex{}
//
//var start, end int
//
//func recordSendMessage(message xmpp.Chat) {
//	go func() {
//		SendChan <- message
//	}()
//}
//
//func recordReceivedMessage(message xmpp.Chat) {
//	go func() {
//		ReceiveChan <- message
//	}()
//}
//
//func Start() {
//	go func() {
//		for {
//			select {
//			case msg := <-ReceiveChan:
//				Lock.Lock()
//				receiveCount = receiveCount + 1
//				startTimestamp := delayMap[msg.Text]
//				if startTimestamp > 0 {
//					endTimestamp := makeTimestampMilli()
//					if maxCost < endTimestamp-startTimestamp {
//						maxCost = endTimestamp - startTimestamp
//					}
//				}
//				Lock.Unlock()
//			}
//		}
//	}()
//	go func() {
//		for {
//			select {
//			case msg := <-SendChan:
//				Lock.Lock()
//				sendCount = sendCount + 1
//				delayMap[msg.Text] = makeTimestampMilli()
//				Lock.Unlock()
//			}
//		}
//	}()
//}
//
//func ShutDown() {
//	fmt.Printf("send message count: %d, receive message count: %d \n", sendCount, receiveCount)
//	//close(SendChan)
//	//close(ReceiveChan)
//
//	fmt.Printf("max message cost: %d ms\n", maxCost)
//
//}
//
//func makeTimestamp() int64 {
//	return time.Now().UnixNano() / 1e6
//}
//
//func unixMilli(t time.Time) int64 {
//	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
//}
//
//func makeTimestampMilli() int64 {
//	return unixMilli(time.Now())
//}
