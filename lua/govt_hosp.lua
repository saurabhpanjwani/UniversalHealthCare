function avg_discharge_to_claim(s, hospID)

    local function mapper(out, rec)
        if not hospID or rec['HospitalID'] == hospID then
		local diff = rec['ClaimFileTime'] - rec['DischargeTime']
		if diff > 0 then
		  out['sum'] = (out['sum'] or 0) + diff
		  out['count'] = (out['count'] or 0) + 1 
		end
        end 
        return out
    end 

    local function reducer(a, b)
        local out = map() 

        out['sum'] = a['sum'] + b['sum']
        out['count'] = a['count'] + b['count']
        return out 
    end 

    return s : aggregate(map{sum = 0, count = 0}, mapper) : reduce(reducer)
end

local function map_claim(record)
  return map {ClaimID = record.ClaimID, 
	      HospitalID=record.HospitalID, 
	      InsurerID=record.InsurerID, 
	      DischargeTime=record.DischargeTime,
	      ClaimFileTime=record.ClaimFileTime,
	      ClaimAmount=record.ClaimAmt,
	      Penalty=record.Penalty,
	      ClaimState=record.ClaimState,
	      ClaimType=record.ClaimType,
	      AuditStatus=record.AuditStatus,
	      AuditLog=record.AuditLog,
	      LogTrail=record.LogTrail,
	      RejectCode=record.RejectCode,
	      TDSHead=record.TDSHead,
	      PaymentInfo=record.PaymentInfo,
	      AcknowledgedAt=record.AckTime,
  }
end

function claim_detail(stream,cId,hID, icID)
  local function filter_claim(record)
    return (record['ClaimID'] == cId) and
           (not hID or record['HospitalID'] == hID) and
           (not icID or record['InsurerID'] == icID)
  end
  return stream : filter(filter_claim) : map(map_claim)
end
