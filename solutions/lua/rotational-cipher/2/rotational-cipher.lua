return {
  rotate = function(input, key)
    local abc = "abcdefghijklmnopqrstuvwxyz"
    local cipher = abc:sub(key + 1) .. abc:sub(1, key)
    
    local text = ""
    for i = 1, #input do
      local s = input:sub(i, i)
      if not s:match("%a") then
          text = text .. s
      else
        local index = abc:find(s:lower())
        local ciphered = cipher:sub(index, index)
        text = text .. (s:match("%u") and ciphered:upper() or ciphered)
      end
    end
    
    return text
  end
}
