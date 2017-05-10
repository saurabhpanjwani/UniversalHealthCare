Claims Management
-------------------

This document enumerates the relation between various modules in the Claims process, in order to -
1. Reduce or eliminate fraud by relying on historical trends (Fraud Management)
2. Hasten the process of claims approval while ensuring #1 (Adjudication Recommender)
3. Associate a trust score to aid in fraud detection and auditing (Trust Score Management) - During claims processing

Let us look at each of them in detail.

Claims Trust Score Management Engine
--------------------------------------
A trust score is the probability that a given claim is not a fraud. 
A mathematical predictor model will be used in order to arrive at a trust score for any claim that is under processing. 
Trust score for a claim is computed on both the entities as well as past anonymised trends -

Entities Related Historic Trends -
1. Percentage of claims filed by the entities (health care provider, Doctor, or the beneficiary) Accepted with high confidence (Adjudication engine score is >80%)
2. Percentage of claims filed by the entities (health care provider, Doctor, or the beneficiary) Not Accepted with high confidence (<30% score by Adjucation engine)
3. Percentage of claims filed by the entities (health care provider, Doctor, or the beneficiary) audited and found fradulent (in retrospect or during claims filing process)
4. Average billed amount / Billed amount for this procedure
5. Number of patients treated (if health care provider) on this day/week/month
6. Reporting Lags
7. Treatment Characteristics and Procedures
8. Years of Experience of Doctor
9. Number of Prior Incidents from Doctor 
10. Number of Unauthorized claims from the entities

General Historic Trends (coming from Fraud Management Engine) -
1. Average number of claims filed and accepted with high confidence
2. Average number of claims filed and not accepted with high confidence
3. Average number of claims audited and found fraudulent
4. Average Billed amount for this procedure
5. Average number of patients treated per day/week/month
6. Average Reporting Lags
7. Average Treatment Characteristics and Procedures
8. Average Years of Experience of Doctor
9. Average Number of Prior Incidents from Doctor 
10. Average Number of Unauthorized claims from the entities


Adjudication Recommender Engine
--------------------------------
The Adjudication engine acts as a recommender system for Insurers. It will assist the ICs and TPAs to make faster and more data-driven decisions when a new claim is presented.

The adjudication engine will do the following -
1. Parse the Policy DB and apply it on the claim presented. If not found in order, the claim is strongly recommended to be rejected.

2. Based on the Claims Trust Score Engine, the Adjudication Engine will either -
   a. Strongly recommend Accepting the claim (>80% confidence) citing relevant parameters from Claim Trust Score Engine
   b. Strongly recommend Rejecting the claim (<30% confidence) citing relevant parameters from Claim Trust Score Engine
   c. Recommend TPA/IC intervention (30-80% confidence levels)

A TPA sign off is still required in order to set the status of the claims to accepted or rejected. The claims trust score engine is expected to learn from TPA action on the claim for future.


Fraud Management Engine
-------------------------
Fraud Management System is a generic system which will collect anonymised claims data in order to identify current trends and be able to suggest to the Trust Score Engine with higher accuracy.

Some examples of data analysed by Fraud Management System are (courtesy Trendwise Analytics) -

1. Doctors, who treated whopping, say 50+ patients in a day. 
2. Providers administering far higher rates of tests than others. 
3. Providers costing far more, per patient basis, than others. 
4. Providers with high ratio of distance patients. 
5. Providers prescribing certain drugs at higher rate than others. 
6. Distance between claimant’s home address and medical provider 
7. Multiple medical opinions/providers 
8. Changing providers for the same treatment (possibly correlated with other claim activity) 
9. High number of treatments for type of injury 
10. Abnormally long treatment time off for the type of injury 
11. Accident severity does not correlate with severity of injury 
12. Providers billing for services not provided. 
13. Providers administering (more) tests and treatments or providing equipments that are not medically necessary. 
14. Providers administering more expensive tests and equipments (up-coding). 
15. Providers multiple-billing for services rendered. 
16. Providers unbundling or billing separately for laboratory tests performed together to get higher reimbursements. 
17. Providers charging more than peers for the same services. 
18. Providers conducting medically unrelated procedures and services. 
19. Policy holders traveling long distance for treatment which may be available nearby. (Possibly scams by bogus providers.) 
20. Policy holders letting others use their healthcare cards. 

This analysis may be applied across Claims, Entities, Procedures, and at geographical units like District, City, State, Country, etc.

This module will also enable data analytics to be run on top of the UHC system in order for various government agencies, insurers and hospitals to monitor the general health of the public.

Incentive to Report Fraud -
-----------------------------
In order to assist fraud detection, it is suggested that beneficiaries be given incentives to report fraudulent claims filing. This necessrily means the identity of the reporter should be confidential. This will boost the Fraud management system, and also bring in resistance toengaging in unlawful practises.

References -
-------------
1. Healthcare Fraud Management using BigData - Whitepaper by Trendwise Analytics
2. 10 popular health care provider fraud schemes - Fraud Magazine
