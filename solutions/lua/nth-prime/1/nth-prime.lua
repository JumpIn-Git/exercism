local function is_prime(n)
  for i = 2, math.floor(math.sqrt(n)) do
    if n % i == 0 then
      return false
    end
  end
  return true
end

return function(n)
  assert(n > 0)
  local count = 2
  local primes = {}
  while #primes <= n do
    if is_prime(count) then 
      table.insert(primes, count)
    end
    count = count + 1
  end
  return primes[n]
end