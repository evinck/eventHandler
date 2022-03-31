package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	log.Print("Beginning of myHander function")
	//
	event := cloudevents.NewEvent()
	json.NewDecoder(in).Decode(&event)
	log.Print("event ID = " + event.ID())
	// log.Print("resourceId =" + event.)

	//var datas []byte
	// event.Data is a JSON byte
	var datas = event.Data.([]byte)

	var result map[string]interface{}
	json.Unmarshal([]byte(datas), &result)

	// resourceID is like "/n/example_namespace/b/my_bucket/o/my_object"
	log.Print("resourceId = " + result["resourceId"].(string))

	msg := struct {
		Msg string `json:"message"`
	}{
		Msg: fmt.Sprintf("output = %s", result["resourceId"]),
		// Msg: fmt.Sprintf("output = %s", datas),
	}
	json.NewEncoder(out).Encode(&msg)
	log.Print("End of myHander function")
}
