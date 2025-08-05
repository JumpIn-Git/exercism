return function(dna)
  local result = ""
  for i = 1, #dna do
    result = result .. ({
        G = 'C',
        C = 'G',
        T = 'A',
        A = 'U',
    })[dna:sub(i, i)]
  end
  return result
end
