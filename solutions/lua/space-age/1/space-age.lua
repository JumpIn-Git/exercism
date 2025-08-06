local SpaceAge = {}

function SpaceAge:new(seconds)
  local earth = seconds / 31557600
  local M = {seconds = seconds}
  local planets = {
    mercury = 0.2408467,
    venus = 0.61519726,
    earth = 1.0,
    mars = 1.8808158,
    jupiter = 11.862615,
    saturn = 29.447498,
    uranus = 84.016846,
    neptune	= 164.79132
  }
  for k, v in pairs(planets) do
    M["on_" .. k] = function() return tonumber(string.format("%.2f", earth / v)) end
  end
  return M
end

return SpaceAge
