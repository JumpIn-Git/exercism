return function(n)
  local actions = {}
  for i, v in ipairs
    {'wink', 'double blink', 'close your eyes', 'jump'}
  do
    if n & (2 ^ (i - 1)) ~= 0 then
      table.insert(actions, v)
    end
  end
  if n & 16 ~= 0 then
    local t = {}
    for i = #actions, 1, -1 do
      table.insert(t, actions[i])
    end
    actions = t
  end
  return actions
end
