package main

import (
	"log"
	"strconv"
	"sync"
    "os/exec"
    "time"

	as "github.com/aerospike/aerospike-client-go"
	shared "UniversalHealthCare/shared"
)

func main() {
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
			writeRecords(shared.Client, binName, size)
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
	binName string,
	size int,
) {
	for i := 1; i <= size; i++ {
        // A new claims record
        var claim shared.Claim

        /*
         * Populate the fields of this claim record
         */

        //ClaimID
        claimID, err := exec.Command("uuidgen").Output()
        if err != nil {
            shared.PanicOnError(err)
        }
        claim.ClaimID = claimID

        //HospitalID
        hospID, err := exec.Command("uuidgen").Output()
        if err != nil {
            shared.PanicOnError(err)
        }
        claim.HospitalID = hospID

        //InsurerID
        insID, err := exec.Command("uuidgen").Output()
        if err != nil {
            shared.PanicOnError(err)
        }
        claim.InsurerID = insID

        //ClaimFileTime - Subratract years,months and days from today
        rand.Seed(time.Now().Unix())
        claim.ClaimFileTime = time.now().AddDate(-1*rand.Intn(2), -1*rand.Intn(11), -1*rand.Intn(31))

        //DischargeTime - a random hour between [0,25] subtracted from claim file time
        //25 is chosen so that some records wil exceed the 24 hour discharge filing period
        claim.DischargeTime = claim.ClaimFileTime.Add(-1*rand.Intn(25))
        
        //ClaimAmt - Minimum is Rs.10, Maximum Rs. 10cr
        minAmt := 10
        maxAmt := 100000000
        claim.ClaimAmt = (minAmt + rand.Intn(maxAmt)) * rand.Float32()

        //Penalty, to a maximum of the claimAmt
        claim.Penalty = 0.01*claim.ClaimAmt*(rand.Intn(100))

        //ClaimState
        claim.ClaimState = ClaimFiled + rand.Intn(MaxClaimState)

        //ClaimType
        switch claim.ClaimState {
            ClaimFiled: 
                fallthrough
            ClaimDocumented:
                fallthrough
            ClaimOnHold:
                claim.ClaimType = ClaimNoActiontype

            ClaimApproved:
                fallthrough
            ClaimPaid:
                claim.ClaimType = ClaimAcceptedType

            ClaimRejected:
                claim.ClaimType = ClaimRejectedType
                

            default: // ClaimAcknowledged, or ClaimContested
            if rand.Float32() >= 0.5 {
                claim.ClaimType = ClaimAcknowledged
            } else {
                claim.ClaimType = ClaimContested
            }
        }

        //Audit Status
        if (claim.ClaimState != shared.ClaimFiled &&
            claim.ClaimState != shared.ClaimDocumented) {
            numAudited := 0.1   //10% get audited
            numFraud   := 0.01  //1% are fradulent
            auditRand := rand.Float32()
            if auditRand < 0.005 {
                claim.AuditType = AuditedAndFraud
            } else if auditRand < 0.01 {
                claim.AuditType = AuditedAndNotFraud
            } else if auditRand < 0.1 {
                claim.AuditType = AuditUnderway
            } else {
                claim.AuditType = NotAudited
            }
        }
            

        //AuditLog

        //logTrail

        //rejectCode

        //TDS Head

        //PaymentInfo

        //AckTime
		key, _ := as.NewKey(*shared.Namespace, *shared.Set, claim.ClaimID)
		bin := as.NewBin(binName, strconv.Itoa(i))

		log.Printf("Put: ns=%s set=%s key=%s bin=%s value=%s",
			key.Namespace(), key.SetName(), key.Value(), bin.Name, bin.Value)

		client.PutBins(shared.WritePolicy, key, bin)
	}
}
