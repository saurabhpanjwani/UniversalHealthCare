package main

import (
	"log"
	"sync"
    "os/exec"
    "time"
    "math/rand"
    "reflect"
    "encoding/json"

	as "github.com/aerospike/aerospike-client-go"
	shared "UniversalHealthCare/shared"
)

func main() {
	size := 2000
	loop := 20

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
			writeRecords(shared.Client, size)
		}()
	}

	// Wait for all go routines to finish
	wg.Wait()

	log.Println("Claims DB populated with ", size*loop, " records!")
}

/**
 * Write batch records individually.
 */
func writeRecords(
	client *as.Client,
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
        claimTime := time.Now().AddDate(-1*rand.Intn(2), -1*rand.Intn(11), -1*rand.Intn(31))
        claim.ClaimFileTime = claimTime.Unix() 

        //DischargeTime - a random hour between [0,25] subtracted from claim file time
        //25 is chosen so that some records wil exceed the 24 hour discharge filing period
        claim.DischargeTime = claimTime.Add(time.Duration(-1*rand.Intn(25))*time.Hour).Unix()
        
        //ClaimAmt - Minimum is Rs.10, Maximum Rs. 10cr
        minAmt := 10
        maxAmt := 100000000
        claim.ClaimAmt = float32(minAmt + rand.Intn(maxAmt)) * rand.Float32()

        //Penalty, to a maximum of the claimAmt
        claim.Penalty = 0.01*claim.ClaimAmt*float32(rand.Intn(100))

        //ClaimState
        claim.ClaimState = shared.ClaimState(int(shared.ClaimFiled) + rand.Intn(int(shared.MaxClaimState)))

        //ClaimType
        switch claim.ClaimState {
            case shared.ClaimFiled: 
                fallthrough
            case shared.ClaimDocumented:
                fallthrough
            case shared.ClaimOnHold:
                claim.ClaimType = shared.ClaimNoActionType

            case shared.ClaimApproved:
                fallthrough
            case shared.ClaimPaid:
                claim.ClaimType = shared.ClaimAcceptedType

            case shared.ClaimRejected:
                claim.ClaimType = shared.ClaimRejectedType
                

            default: // ClaimAcknowledged, or ClaimContested
                if rand.Float32() >= 0.5 {
                    claim.ClaimType = shared.ClaimAcceptedType
                } else {
                    claim.ClaimType = shared.ClaimRejectedType
                }
        }

        //Audit Status
        if (claim.ClaimState != shared.ClaimFiled &&
            claim.ClaimState != shared.ClaimDocumented) {
            numAudited := float32(0.1)   //10% get audited
            numFraud   := float32(0.01)  //1% are fradulent
            auditRand := rand.Float32()
            if auditRand < numFraud/2 {
                claim.AuditStatus = shared.AuditedAndFraud
            } else if auditRand < numFraud {
                claim.AuditStatus = shared.AuditedAndNotFraud
            } else if auditRand < numAudited {
                claim.AuditStatus = shared.AuditUnderway
            } else {
                claim.AuditStatus = shared.NotAudited
            }
        }

        //AuditLog
        if (claim.AuditStatus != shared.NotAudited) {
            claim.AuditLog = "The case was audited"
        }
    
        //logTrail - This is updated as the claim's state changes.
        // This is used in debugging, but will be left empty in this dummy DB

        //rejectCode
        if claim.ClaimType == shared.ClaimRejectedType { 
           claim.RejectCode = shared.RejectCode(rand.Intn(int(shared.MaxRejectCodes)) + 1)
        }

        //TDS Head
        if claim.ClaimType == shared.ClaimAcceptedType {
            claim.TDSHead = "Dr. Rajeev Kapoor"
        }

        //AckTime
        ackTime := time.Now().Add(time.Duration(-1 - rand.Intn(10))*time.Hour)
        claim.AckTime = ackTime.Unix()

        //PaymentInfo
        paymentInfo := shared.PaymentInfo{123456.70, ackTime.Add(time.Duration(-3)*time.Hour).Unix(), 
            claim.InsurerID, claim.HospitalID, "YHO2648721KSA", "Paid and Approved by Admin of Insurer" }

        //Marshal PaymentInfo
        claim.PaymentInfo, err = json.Marshal(paymentInfo)
        if err != nil {
            log.Println("Marshalling error:", err)
        }

		key, _ := as.NewKey(*shared.Namespace, *shared.Set, claim.ClaimID)

        // Write all field names and values into the corresponding index of a binMap
        binsMap := make(map[string]interface{})
		val := reflect.Indirect(reflect.ValueOf(claim))
        for i := 0; i < val.NumField(); i++ {
            binName := val.Type().Field(i).Name
            binValue := val.Field(i).Interface()
            binsMap[binName] = binValue
		    //log.Printf("Put: ns=%s set=%s key=%s bin=%s value=%s",
			//    key.Namespace(), key.SetName(), key.Value(), 
            //    binName, binsMap[binName])
        }
		err = client.Put(shared.WritePolicy, key, binsMap)
        if err != nil {
            shared.PanicOnError(err)
        }
	}
}
