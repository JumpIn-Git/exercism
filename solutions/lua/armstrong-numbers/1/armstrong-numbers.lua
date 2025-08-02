local ArmstrongNumbers = {}

function ArmstrongNumbers.is_armstrong_number(number)
  local s = tostring(number)
  local mult = #s
  local sum = 0
  for i = 1, #s do
    sum = sum + tonumber(s:sub(i, i)) ^ mult
  end
  return sum == number
end

return ArmstrongNumbers
