return {
  encode = function(plaintext)
    plaintext = plaintext:lower():gsub("[^%w]", '')
    local abc = "abcdefghijklmnopqrstuvwxyz"
    local cipher = abc:reverse()

    local text = ""
    count = 0
    for i = 1, #plaintext do
      local s = plaintext:sub(i, i)
      local index = abc:find(s)
      if s:match("%d") then
        text = text .. s
      elseif s:match("%a") then
        text = text .. cipher:sub(index, index)
      end
      if i % 5 == 0 and i < #plaintext then
        text = text .. " "
      end
    end

    return text
  end
}
 