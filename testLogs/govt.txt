Most of these examples can be run with optional parameters like HospID, ICID as well. 
However, in some cases all possible combinations may be omitted. 

Total Claims - 

Total number of Claims
-------------------------
aql> AGGREGATE govt.count() ON uhc.claimsDB 
+---------+
| count   |
+---------+
| 1000000 |
+---------+
1 row in set (6.782 secs)

Average Processing Time
--------------------------

aql> AGGREGATE govt.avg_claim_time() ON uhc.claimsDB
+-----------------------------------------------+
| avg_claim_time                                  |
+-----------------------------------------------+
| MAP('{"sum":2610263577600, "count":984045}')  |
+-----------------------------------------------+
1 row in set (7.829 secs)

Therefore, average processing time = sum/count = 2652585.6 sec = 30 days, 16 hours, 49 minutes and 46 seconds.

aql> AGGREGATE govt.avg_claim_time('f9e7e0d2-ddcd-46a3-aa76-a8a8a16f4131') ON uhc.claimsDB
+------------------------------------+
| avg_claim_time                     |
+------------------------------------+
| MAP('{"sum":1728000, "count":1}')  |
+------------------------------------+
1 row in set (6.706 secs)

For this particular IC/TPA, processing time is about 20 days.

Outstanding Claims
--------------------

aql> AGGREGATE govt.outstanding_claim('f9e7e0d2-ddcd-46a3-aa76-a8a8a16f4131') ON uhc.claimsDB
+---------------------+
| outstanding...      |
+---------------------+
| MAP('{"count":1}')  |
+---------------------+
1 row in set (7.120 secs)

aql> AGGREGATE govt.outstanding_claim() ON uhc.claimsDB
+--------------------------+
| outstanding...           |
+--------------------------+
| MAP('{"count":875203}')  |
+--------------------------+
1 row in set (7.124 secs)

aql> 



