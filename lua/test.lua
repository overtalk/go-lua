function GetNumber(n)
	-- Execute `reload` after code is updated
	-- return double(n)
	return triple(n)
end

function double(n)
	print("double attack")
	return n * 2
end

function triple(n)
	print("triple attack")
	return n * 3
end
