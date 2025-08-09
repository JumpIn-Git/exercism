local triangle = {}

function triangle.kind(a, b, c)
  local err = 'Input Error'
  assert(a + b >= c, err)
  assert(b + c >= a, err)
  assert(a + c >= b, err)
  local seen = {}
  local count = 0
  for _, i in ipairs{a, b, c} do
    assert(i > 0, err)
    if not seen[i] then
      seen[i] = true
      count = count + 1
    end
  end
  if count == 1 then
    return 'equilateral'
  elseif count == 3 then
    return 'scalene'
  else
    return 'isosceles'
  end
end

return triangle
