return {
  decode = function(c1, c2, c3)
    local colors = {
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
    local sum = tonumber("0" .. colors[c1] .. colors[c2]) * 10 ^ colors[c3]
    if sum > 10 ^ 9 then
      return sum / 10 ^ 9, "gigaohms"
    end
    if sum > 10 ^ 6 then
      return sum / 10 ^ 6, "megaohms"
    end
    if sum > 10 ^ 3 then
      return sum / 10 ^ 3, "kiloohms"
    end

    return sum, "ohms"
  end
}
