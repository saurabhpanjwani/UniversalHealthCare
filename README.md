# UHC - Universal Health Care
UniversalHealthCare for India

Introduction
--------------

This repository is WIP and the goal is to prototype a possible solution to India's Universal Health Care policy. 

The repository contains - 
1. Claims related data models, pipeline and analytics. 
2. Data for now will be artificially generated, however, in a real end-to-end system, data source will be the actual system which ties in the Health Care Providers, Insurance Companies/TPAs, and end beneficiaries.


Claims Lifecycle -
-------------------

New Claim Workflow -

Various possible flows for a claim -

Claim Filed -> Claim Documented ->Claim Approved -> Claim Paid -> Claim (Payment) Acknowledged 
Claim Filed -> Claim Documented ->Claim Approved -> Claim Paid -> Claim (Payment) Contested
Claim Filed -> Claim Documented ->Claim Approved -> Claim (Payment) Contested 

Claim Filed -> Claim Documented ->Claim Rejected -> Claim (Rejected) Acknowledged 
Claim Filed -> Claim Documented ->Claim Rejected -> Claim (Rejected) Contested

Claim Filed -> Claim Documented ->Claim Hold -> Claim Documented -> ... -> Claim Approved/Claim Rejected flows

Claim Filed -> Claim Documented -> Claim (No Action) Contested

Claim Contested -> Claim Accepted / Claim Rejected flows

1. Claim Filed - When care provider uses the system to file a new claim
2. Claim Documented - When care provider updates an existing claim with additional information

At this point, IC/TPA is able to see the claim. The possible workflows here are -

3. Claim Accepted - IC/TPA reviews the claim and approves the claim for payment
4. Claim Paid - IC/TPA make the payment for the said claim. 

5. Claim Rejected - When IC/TPA reviews the claim and rejects it based on available information.

6. Claim On Hold - When IC/TPA reviews the claim and requires more information to process it.

The endflow of the claim is handled by the care provider -

7. Claim Acknowledged - The claim is acknowledged for IC/TPA response - Payment / Rejected claim. No response in closing a claim within the stipulated time will lead to auto-acknowledgement of the claim by the system.

8. Claim Contested - The claim is contested for IC/TPA response - Payment / Rejected / No action


Re-opened Claim Workflows -
---------------------------
Claim Acknowledged -> Claim Flagged

9. Claim Flagged - Claim may be re-opened by IC/TPA for suspected fraud analysis, or by health care providers to contest on a later date. This state indicates the claim is flagged for further analysis underway.


Claims Data Model
------------------
Claims data will contain the following parameters.

1. Claim ID - Unique 32-byte UUID
2. Hospital UHC ID - Unique 32-byte UUID assigned to Hospital during its UHC registration
3. IC/TPA UHC ID - Unique 32-byte UUID assigned to IC/TPA during its UHC registration
4. Amount of the Claim - Float
5. Timestamp of Claim Filing - Date
Penalty (1% of amount claimed per 15 days of delay) - Float
6. Status of Claim - Enum (9 states as enumerated in the workflow)
