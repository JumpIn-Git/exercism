local Character = {}

local function ability()
    local result = 0
    local smallest = 7
    for _ = 1, 4 do
      local rand = math.random(1, 6)
      result = result + rand
      smallest = math.min(smallest, rand)
    end
    return result - smallest
end

local function modifier(input)
  return (input - 10) // 2
end

function Character:new(name)
  local obj = {name = name}
  for _, v in ipairs
    {'strength', 'dexterity', 'constitution', 'intelligence', 'wisdom', 'charisma'} 
  do
    obj[v] = ability()
  end
  obj.hitpoints = 10 + modifier(obj.constitution)
  return obj
end

return { Character = Character, ability = ability, modifier = modifier }
