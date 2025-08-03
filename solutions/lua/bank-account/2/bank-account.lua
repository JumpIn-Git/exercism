local BankAccount = {}

function BankAccount:new()
	local M = { is_open = false, money = 0 }
	self.__index = self
	setmetatable(M, self)
	return M
end

function BankAccount:open()
	assert(not self.is_open)
	self.is_open = true
end
function BankAccount:close()
	assert(self.is_open)
	self.is_open = false
	self.money = 0
end
function BankAccount:deposit(x)
	assert(self.is_open)
	if x < 0 then
		error()
	end
	self.money = self.money + x
end
function BankAccount:withdraw(x)
	assert(self.is_open)
	assert(not (x < 0 or x > self.money))
	self.money = self.money - x
end
function BankAccount:balance()
	assert(self.is_open)
	return self.money
end

return BankAccount