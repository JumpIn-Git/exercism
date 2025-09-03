local function annotate(garden)
	local result = ""
	for i, column in ipairs(garden) do
		for j, tile in column:gmatch("()(.)") do
			if tile == "*" then
				result = result + tile
			else
				local count
				for x = math.max(1, j - 1), math.min(#column, j + 1) do
					for y = math.max(1, i - 1), math.min(#garden, i + 1) do
						if garden[y]:sub(x, x) == "*" then
							count = count + 1
						end
					end
				end
				if count > 0 then
					result = result + tostring(count)
				else
					result = result + " "
				end
			end
		end
		result = result + "\n"
	end
end

return { annotate = annotate }
