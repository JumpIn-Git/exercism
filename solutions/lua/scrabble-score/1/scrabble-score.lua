local function score(word)
  if word == nil then return 0 end
  word = word:lower()
  local scores = {
    [1] = {'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't'},
    [2] = {'d', 'g'},
    [3] = {'b', 'c', 'm', 'p'},
    [4] = {'f', 'h', 'v', 'w', 'y'},
    [5] = {'k'},
    [8] = {'j', 'x'},
    [10] = {'q', 'z'}
  }
  
  local score = 0
  for s in word:gmatch('.') do
    
    for point, letters in pairs(scores) do
      local found = false
      for _, letter in ipairs(letters) do
        if letter == s then
          score = score + point
          found = true
          break
        end
      end
      
      if found then break end
    end
  end
  return score
end

return { score = score }
