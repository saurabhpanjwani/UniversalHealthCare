package shared

/* 
 * Define state of a claim
 */
type ClaimState int

const (
    ClaimFiled ClaimState = iota + 1 //1
    ClaimDocumented                  //2
    ClaimApproved                    //3
    ClaimPaid                        //4
    ClaimRejected                    //5
    ClaimAcknowledged                //6
    ClaimContested                   //7
    ClaimOnHold                      //8
    MaxClaimState = ClaimOnHold      //8 ---> Last
)

/* 
 * Define type of claim flow
 */
type ClaimType int

const (
    ClaimNoActionType ClaimType = iota
    ClaimAcceptedType 
    ClaimRejectedType
)

/* 
 * Define structure for capturing state transitions
 */
type LogEntry struct {
    FromState ClaimState
    ToState ClaimState
    Timestamp int64
    ModifierID string
    Comments string
}

/* 
 * Define rejection codes - Annexure 6.1
 */
type RejectCode int

const (
    RejectUndefined RejectCode = iota
    R001 //Data not uploaded within 7 days of transaction (transaction done within 24 hours of discharging patient).
    R002 //Data not uploaded within 7 days of transaction and also Transaction not done within 24 hours of discharging patient.
    R003 //Transaction not done within 24 hours after discharge
    R004 //Patient stay in hospital is >24 hours (It is not a Day care package). 
    R005 //Normal pre and postsurgical expenses are included in package
    R006 //No investigation brief, indications of admissions, disease description, and line of treatment mentioned with package blocked.
    R007 //The surgery package blocked is already included in previous major surgery blocked
    R008 //Registering, Blocking, Discharging transactions in 0 LOS.
    R009 //Registering, blocking, discharging transactions in ___ duration for Day Care procedures
    R0010 //During the course of investigation it reveals that the hospital involves in the practice of charging money to patient. 
    R0011 //Patient and hospital admission documents (medical documents, registration of patient in admission register etc.) not found in hospital of URN and package blocked.
    R0012 //Procedure done does not match the diagnosis.
    R0013 //Rejected as deficiency of medical documents (specify the kind medical documents required, basis also).
    R0014 //Duplicate claim – twice uploaded for same Hospitalization
    R0015 //Hospital was found to be involved in a major fraud.
    R0016 //Claims rejected in case of policy exclusions.
    R0017 //Member not enrolled.
    R0018 //Patient not required hospitalization.
    R0019 //OVERWRITING / MALPRACTICE
    R0020 //FP/BCP code Mismatch
    R0021 //Age/Gender Mismatch
    R0022 //Directly Discharge from ICU
    R0023 //Claim Rejected under Hotlist Card
    R0024 //Rejected as deficiency of medical documents (specify the kind medical documents required, basis also).
    R0025 //Rejected as the diagnostic package is blocked independently which is against RSBY guidelines
    R0026 //Rejected due to as per package list length of stay (LOS) is ------------- days but patient admitted & discharged in ------------- days
    R0027 //Rejected due to pre-Authorization was not taken prior to performing hysterectomy on patient below 40 yrs of age as per MOLE guideline
    MaxRejectCodes = R0027
)

type AuditStatus int

const (
    AuditUndefined AuditStatus = iota
    AuditUnderway
    AuditedAndFraud
    AuditedAndNotFraud
    NotAudited
)

type PaymentInfo struct {
    Amount float32
    TimeOfPayment int64 
    PayerID string
    PayeeID string
    TxnID string
    Comments string
}
