return {
  value = function(colors)
    local ids = {
      black = 0,
      brown = 1,
      red = 2,
      orange = 3,
      yellow = 4,
      green = 5,
      blue = 6,
      violet = 7,
      grey = 8,
      white = 9
    }
    return tonumber("" .. ids[colors[1]] .. ids[colors[2]])
  end
}
