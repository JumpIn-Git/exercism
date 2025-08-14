return {
  valid = function(s)
    if select(2, s:gsub('[%s%d]', '')) ~= #s then return false end
    local sum = 0
    local double = false
    for i in s:reverse():gmatch('%d') do
      n = tonumber(i)
       if double then
        n = n * 2
        if n > 9 then
          n = n - 9
        end
      end
      sum = sum + n
      double = not double
    end
    return sum % 10 == 0 and (
      sum > 0 or select(2, s:gsub('[%d]', '')) > 1
    )
  end
}
