return {
  encode = function(plaintext)
    plaintext = plaintext:lower()
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
      else goto continue end
      count = count + 1
      if count > 0 and count % 5 == 0 then
        text = text .. " "
      end
      ::continue::
    end
    if text:sub(-1) == " " then
      text = text:sub(1, -2)
    end
    return text
  end
}
 