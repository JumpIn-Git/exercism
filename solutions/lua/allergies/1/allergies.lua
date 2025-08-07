local function list(score)
  score = score % 256
  local scores = {
    { "cats", 128 },
    { "pollen", 64 },
    { "chocolate", 32 },
    { "tomatoes", 16 },
    { "strawberries", 8 },
    { "shellfish", 4 },
    { "peanuts", 2 },
    { "eggs", 1 }
  }
  local allergies = {}
  for _, v in ipairs(scores) do
    if score >= v[2] then
      table.insert(allergies, 1, v[1])
      score = score - v[2]
    end
  end
  return allergies
end

local function allergic_to(score, which)
  local x = false
  for _, v in ipairs(list(score)) do
    if v == which then x = true end
  end
  return score < 0 or x
end

return { list = list, allergic_to = allergic_to }
