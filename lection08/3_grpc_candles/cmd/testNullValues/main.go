package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"tfs-grpc/3_grpc_candles/internal/candlespb"
	"time"
)

func main() {
	messRaw, _ := proto.Marshal(&candlespb.CandleResponse{
		Instrument: "APL",
		Open:       65,
	})
	var mes candlespb.CandleResponse
	err := proto.Unmarshal(messRaw, &mes)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Proto: instr: %s, open: %f, time: %+v\n", mes.GetInstrument(), mes.GetOpen(), mes.GetTs().AsTime())
	if mes.Ts == nil {
		fmt.Println("Timestamp is nill\n")
	}
	prTs := mes.GetTs().AsTime()
	if prTs.IsZero() {
		fmt.Println("Timestamp haz a zero value\n")
	}

	var ts time.Time
	if ts.IsZero() {
		fmt.Println("Ts haz a zero value\n")
	}

	fmt.Println(prTs.IsZero())
	fmt.Printf("time: %+v\n", ts)
}
