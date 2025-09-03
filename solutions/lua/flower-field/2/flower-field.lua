local function annotate(garden)
	for i, row in ipairs(garden) do
		for j, tile in row:gmatch("()(.)") do
			if tile == " " then
				local count = 0
				for x = math.max(1, j - 1), math.min(#garden, j + 1) do
					for y = math.max(1, i - 1), math.min(#row, i + 1) do
						if garden[x]:sub(y, y + 1) == "*" then
							count = count + 1
						end
					end
				end
				if count > 0 then
					row = row:sub(i, j - 1) .. tostring(count) .. row:sub(j + 1)
					garden[i] = row
				end
			end
		end
	end
	return garden
end

return { annotate = annotate }
