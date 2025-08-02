local Hamming = {}

function Hamming.compute(a, b)
  if #a ~= #b then
    error("strands must be of equal length")
  end
  local sum = 0
  for i = 1, #a do
    if a:sub(i, i) ~= b:sub(i, i) then
      sum = sum + 1
    end
  end
  return sum
end

return Hamming
