return {
  valid = function(isbn)
    isbn = isbn:gsub('-', '')
    if #isbn ~= 10 then return false end
    local result = 0
    for i = 1, #isbn do
      local num = isbn:sub(i, i)
      if num == "X" and i == 10 then
        num = 10
      elseif num:match("%d") then
        num = tonumber(num)
      else return false end
      result = result + (num * i)
    end
    return result % 11 == 0
  end
}
