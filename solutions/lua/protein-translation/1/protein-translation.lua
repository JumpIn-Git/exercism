local function proteins(strand)
  local result = {}
  local codon_map = {
    Methionine = {"AUG"},
    Phenylalanine = {"UUU", "UUC"},
    Leucine = {"UUA", "UUG"},
    Serine = {"UCU", "UCC", "UCA", "UCG"},
    Tyrosine = {"UAU", "UAC"},
    Cysteine = {"UGU", "UGC"},
    Tryptophan = {"UGG"},
    STOP = {"UAA", "UAG", "UGA"}
  }
  local i = 1
  while i <= #strand do
    local s = strand:sub(i, i + 2)
    local found
    for acid, codons in pairs(codon_map) do
      for _, codon in ipairs(codons) do
        if s == codon then
          found = acid
          break
        end
      end
    end
    if found == nil then error()
    elseif found == 'STOP' then break
    else
      table.insert(result, found)
    end
    i = i + 3
  end
  return result
end

return { proteins = proteins }
