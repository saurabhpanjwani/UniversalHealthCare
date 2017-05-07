function count(s)
    function mapper(rec)
            return 1
    end
    local function reducer(v1, v2)
        return v1 + v2
    end
    return s : map(mapper) : reduce(reducer)
end

function average_claim(s)

    local function mapper(out, rec)
        local diff = rec['AckTime'] - rec['ClaimFileTime']
        if diff > 0 then
          out['sum'] = (out['sum'] or 0) + diff
          out['count'] = (out['count'] or 0) + 1
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

    return s : aggregate(map{sum = 0, count = 0}, mapper) : reduce(reducer)
end
