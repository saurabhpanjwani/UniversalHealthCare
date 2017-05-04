package shared

import (
	"flag"
	"log"
	"os"
	"runtime"

	as "github.com/aerospike/aerospike-client-go"
)

var WritePolicy = as.NewWritePolicy(0, 0)
var Policy = as.NewPolicy()

var Host = flag.String("h", "127.0.0.1", "Aerospike server seed hostnames or IP addresses")
var Port = flag.Int("p", 3000, "Aerospike server seed hostname or IP address port number.")
var Namespace = flag.String("n", "test", "Aerospike namespace.")
var Set = flag.String("s", "testset", "Aerospike set name.")
var showUsage = flag.Bool("u", false, "Show usage information.")
var Client *as.Client

func PanicOnError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// reads input flags and interprets the complex ones
func init() {
	// use all cpus in the system for concurrency
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.SetOutput(os.Stdout)

	flag.Parse()

	if *showUsage {
		flag.Usage()
		os.Exit(0)
	}

	var err error
	Client, err = as.NewClient(*Host, *Port)
	if err != nil {
		PanicOnError(err)
	}
}
