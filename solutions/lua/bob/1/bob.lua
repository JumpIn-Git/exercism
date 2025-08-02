local bob = {}

function bob.hey(say)
  say = say:gsub("%s+", "") --remove whitespace
  local upper = say:match("%a") and say:upper() == say --atleast one char
  local q = say:sub(-1, -1) == "?"
  local shout = say:sub(-1, -1) == "!"
  local empty = say == ""

  if empty then 
    return "Fine. Be that way!"
  elseif upper and q then
    return "Calm down, I know what I'm doing!"
  elseif q then
    return "Sure."
  elseif upper then
    return "Whoa, chill out!"
  else return "Whatever." end
end

return bob
