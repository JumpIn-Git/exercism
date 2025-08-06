return function(s)
  s = s:lower()
  local abc = "abcdefghijklmnopqrstuvxyz"
  for i = 1, #abc do
    if not s:find(abc:sub(i, i)) then return false end
  end
  return true
end
