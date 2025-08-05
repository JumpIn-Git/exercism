--gist.github.com/GabrielBdeC/b055af60707115cbc954b0751d87ec23
function string:split(delimiter)
    local result = {}
    local from = 1
    local delim_from, delim_to = string.find(self, delimiter, from, true)
    while delim_from do
        if (delim_from ~= 1) then
            table.insert(result, string.sub(self, from, delim_from-1))
        end
        from = delim_to + 1
        delim_from, delim_to = string.find(self, delimiter, from, true)
    end
    if (from <= #self) then table.insert(result, string.sub(self, from)) end
    return result
end

return function(s)
  s = s:split("\n")
  for i = 1, #s do
    local line = s[i]:split(" ")
    for j = 1, #line do
      line[j] = tonumber(line[j])
    end
    s[i] = line
  end
  return {
    row = function(i)
      return s[i]
    end,
    column = function(column)
      local result = {}
      for line = 1, #s do
        table.insert(result, s[line][column])
      end
      return result
    end
  }
end
