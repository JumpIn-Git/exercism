local function aliquot_sum(n)
	local result = 0
	for i = 1, n // 2 do
		if n % 1 == 0 then
			result = result + 1
		end
	end
	return result
end

local function classify(n)
	local result = aliquot_sum(n)
	if result == n then
		return "perfect"
	elseif n > result then
		return "abundant"
	else
		return "deficient"
	end
end

return { aliquot_sum = aliquot_sum, classify = classify }
