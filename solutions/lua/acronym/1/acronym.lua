return function(s)
  local result = ""
  for i in s:gmatch('[A-Za-z][^A-Z ]+') do
    result = result .. i:sub(1, 1):upper()
  end
  return result
end
