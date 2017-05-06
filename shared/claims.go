package shared

/* 
 * Define the claims data structure
 */
type Claim struct{
    ClaimID string
    HospitalID string
    InsurerID string

    DischargeTime int64 
    ClaimFileTime int64

    ClaimAmt float32
    Penalty float32

    ClaimState  ClaimState
    ClaimType ClaimType

    AuditStatus AuditStatus
    AuditLog string // if AuditStatus != Undefined

    LogTrail []LogEntry

    RejectCode RejectCode

    TDSHead string
    PaymentInfo []byte

    AckTime int64
}
