/*
* jmu may 2021, oracle streaming credit car movs random producer
*/
package main

import (
	"fmt"
	"time"
	"os"
	"context"
	"github.com/google/uuid"
	"math/rand"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/streaming"
	"github.com/ShiraazMoollatjie/goluhn"
	"github.com/icrowley/fake"
)

func main() {

    intro()

	stream := os.Getenv("stream")
    fmt.Printf("oci stream ocid: %s\n", stream)

	tenancy := os.Getenv("tenancy")
    fmt.Printf("oci tenancy ocid: %s\n", tenancy)

    user := os.Getenv("user")
    fmt.Printf("oci user ocid: %s\n", user)

    region := os.Getenv("region")
    fmt.Printf("oci region: %s\n", region)

    fingerprint := os.Getenv("fingerprint")
    fmt.Printf("fingerprint: %s\n", fingerprint)

    ppkcontent := os.Getenv("ppkfile")

	ppkpasswd := os.Getenv("ppkpassword")
    fmt.Printf("ppk passwd: %s\n", "****")

	nrcp := common.NewRawConfigurationProvider(tenancy, user, region, fingerprint, string(ppkcontent), &ppkpasswd)

	var endpoint = "https://cell-1.streaming." + region + ".oci.oraclecloud.com"
	fmt.Printf("Streaming endpoint: %s\n", endpoint)
	sclient, err := streaming.NewStreamClientWithConfigurationProvider(nrcp, endpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var message = ""

	for 1 >=0  {
		cnumber := goluhn.Generate(16)
		name := fake.FirstName()
		surname := fake.LastName()
		product := fake.Product()
		message = fmt.Sprintf("{card:\"%s\",holder:\"%s\",amount:%d,prod:\"%s\"}", cnumber, name + " " + surname,rand.Intn(99999),product)
		putMessage(sclient, stream, region, uuid.New().String(), message)
		time.Sleep(1000000000)
	}

	fmt.Println("Sample Subscriber Disconnected")
}

/*
* put a meessage in a topic
*/
func putMessage(client streaming.StreamClient, stream string, region string, key string, value string) int{
	var req streaming.PutMessagesRequest
	req.StreamId = &stream
	var entry streaming.PutMessagesDetailsEntry
	entry.Key = []byte(key)
	entry.Value = []byte(value)
	var entryarray [] streaming.PutMessagesDetailsEntry
	entryarray = append(entryarray, entry)
	var det streaming.PutMessagesDetails
	det.Messages = entryarray
	req.PutMessagesDetails = det
	//client.SetRegion(region)
	_, err := client.PutMessages(context.Background(), req)
	if err != nil {
		fmt.Println("Error: ", err)
		return  -1
	}
	fmt.Println("Sent to OCI streaming -> ", value)
	return  0
}
/*
*
*/
func intro(){
    fmt.Fprintf(os.Stderr, "\n (c) jmu 2021 | ccg, produces random credit card movements and publishes to oci streaming\n")
    fmt.Printf("--------------------------------------------------------------------------------------\n")
}
/*
*
*/
func check(e error) {
    if e != nil {
        panic(e)
    }
}