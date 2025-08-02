return function(s)
  local seen = {}
  for i = 1, #s do
    local l = s:sub(i, i):lower()
    if seen[l] then
      return false
    end
    if l ~= "-" and l ~= " " then
      seen[l] = true
    end
  end
  return true
end
