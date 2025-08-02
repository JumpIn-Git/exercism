return function(numbers)
  return {
   to = function(level)
      local list = {}
      for _, number in ipairs(numbers) do
        local count = number
        while count < level do
          if not list[count] then
            list[count] = true
          end
          count = count + number
        end
      end

      local sum = 0
      for k, _ in pairs(list) do
        sum = sum + k
      end
      return sum
    end
  }
end
