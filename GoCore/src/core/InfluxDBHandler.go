package main


import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)

var INFLUX_DB_NAME,
INFLUX_DB_USER,
INFLUX_DB_PASS,
INFLUX_DB_PORT string

func initInfluxDB(){
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: INFLUX_DB_USER,
		Password: INFLUX_DB_PASS,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

}

func createPoint(){
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func batchPoint(app_name string, url string, method_name string){
	tags := map[string]string{"app_name": app_name, "method_name": method_name, "url": url}
}





