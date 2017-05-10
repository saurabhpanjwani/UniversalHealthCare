Claims Management
-------------------

This document enumerates the relation between various modules in the Claims process, in order to -
1. Reduce or eliminate fraud by relying on historical trends (Fraud Management)
2. Hasten the process of claims approval while ensuring #1 (Adjudication Recommender)
3. Associate a trust score to aid in fraud detection and auditing (Trust Score Management)

Let us look at each of them in detail.

Claims Trust Score Management Engine
--------------------------------------
A trust score is the probability that a given claim is not a fraud. 
A mathematical predictor model will be used in order to arrive at a trust score. 
Trust score for a claim is computed as -

1. Percentage of claims filed by the entities (health care provider, Doctor, or the beneficiary) Accepted with high confidence (Adjudication engine score is >80%)
2. Percentage of claims filed by the entities (health care provider, Doctor, or the beneficiary) Not Accepted with high confidence (<30% score by Adjucation engine)
3. Percentage of claims filed by the entities (health care provider, Doctor, or the beneficiary) audited and found fradulent (in retrospect or during claims filing process)
4. Average billed amount / Billed amount for this procedure
5. Number of patients treated (if health care provider)
6. Reporting Lags
7. Treatment Characteristics and Procedures
8. Years of Experience of Doctor
9. Number of Prior Incidents from Doctor 
10. Number of Unauthorized claims from the entities



Adjudication Recommender Engine
--------------------------------
Based on the Claims Trust Score Engine, the Adjudication Engine will either -
1. Strongly recommend Accepting the claim (>80% confidence) citing relevant parameters from Claim Trust Score Engine
2. Strongly recommend Rejecting the claim (<30% confidence) citing relevant parameters from Claim Trust Score Engine
3. Recommend TPA/IC intervention (30-80% confidence levels)



Fraud Management Engine
-------------------------



References -
-------------
1. Healthcare Fraud Management using BigData - Whitepaper by Trendwise Analytics
2. 10 popular health care provider fraud schemes - Fraud Magazine
