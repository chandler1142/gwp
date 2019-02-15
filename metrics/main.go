package main

import (
	"github.com/go-metrics"
	"log"
	"os"
	"time"
)

func testGauge() {
	g := metrics.NewGauge()
	metrics.Register("bar", g)
	g.Update(1)

	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 1)
		g.Update(j)
		j++
	}
}

func testCounter()  {
	g := metrics.NewGauge()
	metrics.Register("bar", g)
	g.Update(1)


	c := metrics.NewCounter()
	metrics.Register("foo", c)
	c.Inc(45)

	c.Dec(3)

	go metrics.Log(metrics.DefaultRegistry,
		1 * time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	time.Sleep(3 * time.Second)
}

func testHistogram() {
	s := metrics.NewExpDecaySample(1024, 0.015) // or metrics.NewUniformSample(1028)

	h := metrics.NewHistogram(s)

	metrics.Register("baz", h)
	h.Update(1)


	go metrics.Log(metrics.DefaultRegistry,
		1 * time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))


	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 1)
		j++
		h.Update(j)
	}
}

func testMeters() {
	m := metrics.NewMeter()
	metrics.Register("quux", m)
	m.Mark(1)


	go metrics.Log(metrics.DefaultRegistry,
		1 * time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))


	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 1)
		j++
		m.Mark(1)
	}
}

func main() {
	testMeters()
}
