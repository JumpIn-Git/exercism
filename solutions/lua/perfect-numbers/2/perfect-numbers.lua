local function aliquot_sum(n)
	local numbers = {}
	while n > 1 do
		n = n / 2
		table.insert(numbers, n)
	end
	local result = 0
	for _, v in ipairs(numbers) do
		result = result + v
	end
	return result
end

local function classify(n) end

return { aliquot_sum = aliquot_sum, classify = classify }
