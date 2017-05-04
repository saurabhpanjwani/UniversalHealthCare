package main

import (
	"log"
	"strconv"
	"sync"

	as "github.com/aerospike/aerospike-client-go"
	shared "UniversalHealthCare/shared"
)

func main() {
	keyPrefix := "ClaimID"
	valuePrefix := "ClaimUUID"
	binName := "AllOtherColumns"
	size := 100
	loop := 10000

	// Write loop number of independent batched writes
	// Each batched write writes in size number of records
	// Wait for all go routines to finish before exiting
	var wg sync.WaitGroup

	for i := 1; i <= loop; i++ {
		//Increment wait group counter
		wg.Add(1)

		//Spawn a new batch
		go func() {
			defer wg.Done()
			writeRecords(shared.Client, keyPrefix, binName, valuePrefix, size)
		}()
	}

	// Wait for all go routines to finish
	wg.Wait()

	log.Println("Claims DB populated with %d records!", size*loop)
}

/**
 * Write batch records individually.
 */
func writeRecords(
	client *as.Client,
	keyPrefix string,
	binName string,
	valuePrefix string,
	size int,
) {
	for i := 1; i <= size; i++ {
		key, _ := as.NewKey(*shared.Namespace, *shared.Set, keyPrefix+strconv.Itoa(i))
		bin := as.NewBin(binName, valuePrefix+strconv.Itoa(i))

		log.Printf("Put: ns=%s set=%s key=%s bin=%s value=%s",
			key.Namespace(), key.SetName(), key.Value(), bin.Name, bin.Value)

		client.PutBins(shared.WritePolicy, key, bin)
	}
}
