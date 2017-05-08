function count_and_amount(s)
    function mapper(rec)
	  out['amt'] = (out['amt'] or 0) + rec['ClaimAmt']
	  out['count'] = (out['count'] or 0) + 1
      	  if rec['ClaimState'] == 6 then
              out['settled_count'] = (out['settled_count'] or 0) + 1
              out['settled_amt'] = (out['settled_amt'] or 0) + rec['ClaimAmt']
	  end
        return out
    end
    local function reducer(v1, v2)
        local out = map() 
        out['amt'] = a['amt'] + b['amt']
        out['count'] = a['count'] + b['count']
        out['settled_count'] = a['settled_count'] + b['settled_count']
        return out 
    end
    return s : map({count = 0, amt = 0, settled_count = 0, settle_amt =0},mapper) : reduce(reducer)
end


function avg_claim_time(s, icID)

    local function mapper(out, rec)
        if not icID or rec['InsurerID'] == icID then
		local diff = rec['AckTime'] - rec['ClaimFileTime']
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

function outstanding_claim(s, icID)

    local function mapper(out, rec)
        if rec['ClaimState'] ~= 6 then
          if icID then
            if rec['InsurerID'] == icID then
              out['count'] = (out['count'] or 0) + 1 
              out['amt'] = (out['amt'] or 0) + rec['ClaimAmt']
              if rec['ClaimType'] == 0 then
                out['no_action_count'] = (out['no_action_count'] or 0) + 1 
              end
              if rec['ClaimState'] == 7 then
                out['contested_count'] = (out['contested_count'] or 0) + 1 
              end
            end
          else
            out['count'] = (out['count'] or 0) + 1 
            out['amt'] = (out['amt'] or 0) + rec['ClaimAmt']
            if rec['ClaimType'] == 0 then
              out['no_action_count'] = (out['no_action_count'] or 0) + 1 
            end
            if rec['ClaimState'] == 7 then
              out['contested_count'] = (out['contested_count'] or 0) + 1 
            end
          end
        end
        return out 
    end 

    local function reducer(a, b)
        local out = map() 
        out['count'] = a['count'] + b['count']
        out['amt'] = a['amt'] + b['amt']
        out['no_action_count'] = a['no_action_count'] + b['no_action_count']
        out['contested_count'] = a['contested_count'] + b['contested_count']
        return out 
    end 

    return s : aggregate(map{count = 0, amt = 0, no_action_count = 0, contested_count =0 }, mapper) : reduce(reducer)
end

function acked_claim(s, icID)

    local function mapper(out, rec)
        if rec['ClaimState'] == 6 then
          if icID then
            if rec['InsurerID'] == icID then
              out['count'] = (out['count'] or 0) + 1 
            end
          else
            out['count'] = (out['count'] or 0) + 1 
          end
        end
        return out 
    end 

    local function reducer(a, b)
        local out = map() 
        out['count'] = a['count'] + b['count']
        return out 
    end 

    return s : aggregate(map{count = 0}, mapper) : reduce(reducer)
end

function rejected_claim(s, icID)

    local function mapper(out, rec)
        if rec['ClaimType'] == 2 then
          if icID then
            if rec['InsurerID'] == icID then
              out['count'] = (out['count'] or 0) + 1
              out['amt'] = (out['amt'] or 0) + rec['ClaimAmt'] 
              if rec['ClaimState'] == 6 then
                  out['settled'] = (out['settled'] or 0) + 1
                  out['settled_amt'] = (out['settled_amt'] or 0) + rec['ClaimAmt']
              end 
            end
          else
            out['count'] = (out['count'] or 0) + 1 
            out['amt'] = (out['amt'] or 0) + rec['ClaimAmt'] 
            if rec['ClaimState'] == 6 then
                out['settled'] = (out['settled'] or 0) + 1
                out['settled_amt'] = (out['settled_amt'] or 0) + rec['ClaimAmt']
            end 
          end
        end
        return out 
    end 

    local function reducer(a, b)
        local out = map() 
        out['count'] = a['count'] + b['count']
        out['amt'] = a['amt'] + b['amt']
        out['settled'] = a['settled'] + b['settled']
        out['settled_amt'] = a['settled_amt'] + b['settled_amt']
        return out 
    end 

    return s : aggregate(map{count = 0, amt = 0, settled = 0, settled_amt = 0}, mapper) : reduce(reducer)
end

function audited_claim(s, icID)

    local function mapper(out, rec)
        if rec['AuditStatus'] ~= 0 then
          if icID then
            if rec['InsurerID'] == icID then
              out['count'] = (out['count'] or 0) + 1
              out['amt'] = (out['amt'] or 0) + rec['ClaimAmt']
              if rec['AuditStatus'] == 2 then
                out['fraud_count'] = (out['fraud_count'] or 0) + 1
                out['fraud_amt'] = (out['fraud_amt'] or 0) + rec['ClaimAmt']
	      end 
            end
          else
            out['count'] = (out['count'] or 0) + 1 
            out['amt'] = (out['amt'] or 0) + rec['ClaimAmt'] 
            if rec['AuditStatus'] == 2 then
                out['fraud_count'] = (out['fraud_count'] or 0) + 1
                out['fraud_amt'] = (out['fraud_amt'] or 0) + rec['ClaimAmt']
	    end 
          end
        end
        return out 
    end 

    local function reducer(a, b)
        local out = map() 
        out['count'] = a['count'] + b['count']
        out['amt'] = a['amt'] + b['amt']
        out['fraud_count'] = a['fraud_count'] + b['fraud_count']
        out['fraud_amt'] = a['fraud_amt'] + b['fraud_amt']
        return out 
    end 

    return s : aggregate(map{count = 0, amt = 0, fraud_count =0, fraud_amt = 0}, mapper) : reduce(reducer)
end

function claim_detail(s)
    function mapper(rec)
	  out['amt'] = (out['amt'] or 0) + rec['ClaimAmt']
	  out['count'] = (out['count'] or 0) + 1
      	  if rec['ClaimState'] == 6 then
              out['settled_count'] = (out['settled_count'] or 0) + 1
              out['settled_amt'] = (out['settled_amt'] or 0) + rec['ClaimAmt']
	  end
        return out
    end
    local function reducer(v1, v2)
        local out = map() 
        out['amt'] = a['amt'] + b['amt']
        out['count'] = a['count'] + b['count']
        out['settled_count'] = a['settled_count'] + b['settled_count']
        return out 
    end
    return s : map({count = 0, amt = 0, settled_count = 0, settle_amt =0},mapper) : reduce(reducer)
end
