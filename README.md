# UHC - Universal Health Care
UniversalHealthCare for India

Introduction
--------------

This repository is WIP and the goal is to prototype a possible solution to India's Universal Health Care policy. 

The repository contains - 
1. Claims related data models, pipeline and analytics. 
2. Data for now will be artificially generated, however, in a real end-to-end system, data source will be the actual system which ties in the Health Care Providers, Insurance Companies/TPAs, and end beneficiaries.


New Claims Workflows -
-----------------------

Various possible flows for a claim -

Claim Accepted Type Flow
-------------------------
Claim Filed -> Claim Authorised -> Claim Documented ->Claim Approved -> Claim Paid -> Claim (Payment) Acknowledged 
Claim Filed -> Claim Authorised -> Claim Documented ->Claim Approved -> Claim Paid -> Claim (Payment) Contested
Claim Filed -> Claim Authorised -> Claim Documented ->Claim Approved -> Claim (Payment) Contested 


Claim Rejected Type Flow 
--------------------------
Claim Filed -> Claim Authorised -> Claim Documented ->Claim Rejected -> Claim (Rejected) Acknowledged 
Claim Filed -> Claim Authorised -> Claim Documented ->Claim Rejected -> Claim (Rejected) Contested

Claim No Action Flows -
------------------------
Claim Filed -> Claim Authorised -> Claim Documented ->Claim Hold -> Claim Documented -> ... -> Claim Approved/Claim Rejected flows

Claim Filed -> Claim Authorised -> Claim Documented -> Claim (No Action) Contested

Claim Contested -> Claim Accepted / Claim Rejected flows

Claim Unauthorised Flow -
-------------------------

Claim Filed -> Claim Failed Auth

1. Claim Filed - When care provider uses the system to file a new claim
2. Claim Authorised - This requires -
   a. Beneficiary Card Read / Authentication using Aadhar biometrics to prove she/he was physically present.
   b. Provider Card Read to prove the beneficiary indeed went to the care provider.
   
3. Claim Documented - When care provider updates an existing claim with additional information

At this point, IC/TPA is able to see the claim. The possible workflows here are -

4. Claim Accepted - IC/TPA reviews the claim and approves the claim for payment
5. Claim Paid - IC/TPA make the payment for the said claim. 

6. Claim Rejected - When IC/TPA reviews the claim and rejects it based on available information.

7. Claim On Hold - When IC/TPA reviews the claim and requires more information to process it.

The endflow of the claim is handled by the care provider -

8. Claim Acknowledged - The claim is acknowledged for IC/TPA response - Payment / Rejected claim. No response in closing a claim within the stipulated time will lead to auto-acknowledgement of the claim by the system.

9. Claim Contested - The claim is contested for IC/TPA response - Payment / Rejected / No action

If the claim fails authorisation then-

10. Claim Fail Auth - This end state indicates an attempt to file a claim with invalid authentication/authorization of the care provider and/or the beneficiary.

Audited Claim Workflows -
---------------------------------

Any Claim state beyond the Claims Filed state may set Claim Audited flag

Claim Audited Bit - Claim may be audited by IC/TPA for suspected fraud analysis, or by health care providers to contest at a later date. 
This state indicates the claim is under audit for further analysis underway, or was audited and has been cleared or found fraudulent.
A claim may be flagged in any state of the claim beyond the initial filing.


Outstanding Claims
-------------------

All claims not in the "Claim Acknowledged" state are considering outstanding.

Closed Claims
--------------

All claims in "Claims Acknowledged" state are considered closed, and cannot be re-opened. Only the Audit and Fraud flags may be set for a closed Claim, along with optional comments.


Claims Data Model
------------------

Claims data will contain the following parameters.

1. Claim ID - Unique 32-byte UUID
2. Hospital UHC ID - Unique 32-byte UUID assigned to Hospital during its UHC registration
3. IC/TPA UHC ID - Unique 32-byte UUID assigned to IC/TPA during its UHC registration
4. Amount of the Claim - Float
5. Beneficiary Discharge Timestamp - Date
6. Timestamp of Claim Filing - Date
7. Penalty (1% of amount claimed per 15 days of delay) - Float
8. Status of Claim - Enum (9 states as enumerated in the workflow)
9. Flow Type  - Accepted / Rejected / No Action / Unauthorized - Indicate what kind of claim closure was acknowledged / contested / audited. Default is No Action.
10. IsFlagged - Boolean (Default False) - Indicates if this claim is/was under audit. Once set, this flag cannot be reset. Any closure on audit results will still follow the usual closure workflows.
11. Log Trail - Array of Log Type {< From State > , < To State >, < Timestamp > , < UHC ID of the Modifier > , < Optional Comments > }. This will enumerate the trail of the claim lifecycle. 

Some of these data fields are interpreted according to states.

For all states indicating rejected claims -

12. Rejection Code - Indicates the reason for claim rejection, if flow type == Rejected. For all other flows, this code will be 0.

For all states indicating settled claims -

13. TDS Head - Name of the approving TPA authority, if flow type == Accepted. Null for all other flows.
14. Payment information - Amount, Timestamp, From IC/TPA UHC ID, To Provider UHC ID, Payment Transaction ID, if flow type == Accepted and state > = Claim Paid. Null for all other states.

For final state (Claim Acknowledged)

15. Timestamp of Claim Acknowledgement - Date. Null for all other states.

For claim that have IsFlagged == True
16. isFraud - Boolean - Indicating if the audit of a claim deems it as a fraudulent claim
17. Audit Log - Text - Optional field explaining results of audit in detail. It might be updated even in Claims Acknowledged state.


View of Data For Various Entities - Hospital / Health Care Provider
---------------------------------------------------------------------

1. Aggregate Claims View For a Time Period - Total number of claims filed, total acknowledged, total outstanding, total value of claims filed, total value of claims settled, total value of claims rejected, total value of claims still pending (No Action), Average time of claim settlement across all TPAs/ICs, Total number of cases audited, Average time taken to file a new claim after discharge of the beneficiary, Time taken to close on a claim, Number of claims contested, Number of claims deemed fraudulent,Number of cases unauthorized.

2. Per IC/TPA View For a Time Period - Number of claims filed, total acknowledged, total outstanding, total value of claims filed, total value of claims settled, total value of claims rejected, total value of claims still pending (No Action) with this IC/TPA, Average time of claim settlement for this IC/TPA, Number of clases audited by this TPA, Time taken to close on a claim, Number of claims contested, Number of claims deemed fraudulent,Number of cases unauthorized.

3. Detailed View of a claim filed.


View of Data For Various Entities - IC/TPA
-------------------------------------------

1. Aggregate Claims View For a Time Period - Total number of claims documented for this IC/TPA, total acknowledged, total outstanding, total value of claims filed, total value of claims settled, total value of claims rejected, total value of claims still pending (No Action), Average time of claim settlement across all care providers, Total number of cases audited, Time taken to close on a claim, Number of claims contested, Number of claims deemed fraudulent,Number of cases unauthorized.

2. Per Hospital/Care Provider View For a Time Period - Number of claims filed, total acknowledged, total outstanding, total value of claims filed, total value of claims settled, total value of claims rejected, total value of claims still pending (No Action) with this IC/TPA, Average time of claim settlement for this IC/TPA, Number of clases audited by this TPA, Time taken to close on a claim, Number of claims contested, Number of claims deemed fraudulent,Number of cases unauthorized.

3. Detailed view of a claim filed


View of Data For Various Entities - Government
-----------------------------------------------
