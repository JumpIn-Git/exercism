local function aliquot_sum(n)
	local result = 0
	for i = 1, math.floor(n / 2) do
		if n % i == 0 then
			result = result + i
		end
	end
	return result
end

local function classify(n)
	local result = aliquot_sum(n)
	if result == n then
		return "perfect"
	elseif n > result then
		return "deficient"
	end
end

return { aliquot_sum = aliquot_sum, classify = classify }
