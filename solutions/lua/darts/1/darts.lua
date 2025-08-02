local Darts = {}

function Darts.score(x, y)
  local dist = math.sqrt(x ^ 2 + y ^ 2)
  return (dist <= 1 and 10 or
    dist <= 5 and 5 or
    dist <= 10 and 1
    or 0)
end

return Darts
