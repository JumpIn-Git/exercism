local grains = {}

function grains.square(n)
  return (
    n == 1 and 1 or 2 ^ (n - 1)
  )
end

function grains.total()
  local sum = 0
  local grains = 1.0
  for i = 1, 64 do
    sum = sum + grains
    grains = grains * 2
  end
  return sum
end

return grains
