package shared

/* 
 * Define the claims data structure
 */
typedef Claims struct{
    ClaimID string
    HospitalID string
    InsurerTPAID string

    DischargeTime date
    ClaimFileTime date

    ClaimAmt float
    Penalty float

    ClaimStatus ClaimStatus
    ClaimType ClaimType

    AuditStatus AuditStatus
    AuditLog string // if AuditStatus != Undefined

    LogTrail []LogEntry

    RejectCode RejectCode

    TDSHead string
    PaymentInfo PaymentInfo

    AckTime date
}
