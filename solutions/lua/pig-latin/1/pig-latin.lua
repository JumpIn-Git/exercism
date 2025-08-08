local function translate(word)
  local vowel = 'aeiou'
  local constant = 'bcdfghjklmnpqrstvwxyz'

  --Rule 1
  if 
    vowel:match(word:sub(1, 1))
    or word:sub(1, 2) == "xr"
    or word:sub(1, 2) == "yt"
  then
    return word .. "ay"
  end

  --Rule 3 (merges with rule 2)
  local _, index = word:find("qu")
  local index = index or 0
    
  --Rule 2
  local i = 1
  local constants = ""
  while i <= #word do
    local s = word:sub(i, i)
    
    --Rule 4
    if i > 1 and s == 'y' then
      break
    end
    
    --Rule 3
    if i == index - 1 then 
      constants = constants .. "qu"
      i = i + 2 --skip over qu
      break
    end
    
    if constant:match(s) then
      constants = constants .. s
    else
      break
    end
    i = i + 1
  end
  return word:sub(i) .. constants .. "ay"
end
return function(phrase)
  return phrase:gsub('[^%s]+', translate)
end
