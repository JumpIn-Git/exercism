return function(n)
  assert(n ~= 0)
  assert(n > 0)
  local count = 0
  while n ~= 1 do
    if n % 2 == 0 then
      n = n / 2
    else
      n = n * 3 + 1
    end
    count = count + 1
  end
  return count
end
