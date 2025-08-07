local DNA = {}
DNA.__index = DNA

function DNA:count(s)
  local x = self.nucleotideCounts[s]
  if x then return x else error('Invalid Nucleotide') end
end

function DNA:new(s)
  local M = {nucleotideCounts = {}}
  setmetatable(M, self)
  local x = {'A', 'C', 'G', 'T'}
  local y = 0
  for i = 1, #x do
    local _, count = s:gsub(x[i], "")
    M.nucleotideCounts[x[i]] = count
    y = y + count
  end
  if y <  #s then error('Invalid Sequence') end
  return M
end

return DNA
