return function(sum)
	local results = {}
	for a = 1, sum // 2 do
		for b = 1, sum // 2 do
			for c = 1, sum // 2 do
				if a ^ 2 + b ^ 2 == c ^ 2 and a + b + c == sum then
					table.insert(results, { a, b, c })
				end
			end
		end
	end
	return results
end
