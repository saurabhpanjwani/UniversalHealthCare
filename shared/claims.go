package shared

import(
    "time"
)

/* 
 * Define the claims data structure
 */
type Claims struct{
    ClaimID string
    HospitalID string
    InsurerTPAID string

    DischargeTime time.Time 
    ClaimFileTime time.Time

    ClaimAmt float32
    Penalty float32

    ClaimState  ClaimState
    ClaimType ClaimType

    AuditStatus AuditStatus
    AuditLog string // if AuditStatus != Undefined

    LogTrail []LogEntry

    RejectCode RejectCode

    TDSHead string
    PaymentInfo PaymentInfo

    AckTime time.Time
}
