return function(n)
  local x = ""
  if n % 3 == 0 then
    x = x .. "Pling"
  end
  if n % 5 == 0 then
    x = x .. "Plang"
  end 
  if n % 7 == 0 then
    x = x .. "Plong"
  end
  if x == "" then
    x = tostring(n)
  end
  return x
end
