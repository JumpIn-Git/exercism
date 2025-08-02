local BankAccount = {}

function BankAccount.new()
  local M = {is_open = false, money = 0}
  function M:check() if not self.is_open then error() end end

  function M:open()
    if self.is_open then error() end
    self.is_open = true
  end
  function M:close()
    self:check()
    self.is_open = false
    self.money = 0
  end
  function M:deposit(x)
    self:check()
    if x < 0 then error() end
    self.money = self.money + x
  end
  function M:withdraw(x)
    self:check()
    if x < 0 or x > self.money then error() end
    self.money = self.money - x
  end
  function M:balance() 
    self:check()
    return self.money 
  end

  return M
end

return BankAccount
