return {
  transform = function(dataset)
    local result = {}
    for k, v in pairs(dataset) do
      for _, letter in ipairs(v) do
        result[letter:lower()] = k
      end
    end
    return result
  end
}
