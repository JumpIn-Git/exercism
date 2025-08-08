local Clock = {}
Clock.__index = Clock

function Clock:plus(min)
  self.hours, self.minutes = self.translate(self.hours, self.minutes + min)
  return self
end
function Clock:minus(min)
  return self:plus(-min)
end
function Clock:equals(clock)
  return self.hours == clock.hours and self.minutes == clock.minutes
end

function Clock.translate(hours, minutes)
  minutes = minutes or 0
  hours = hours + minutes // 60
  hours = hours % 24
  minutes = minutes % 60 
  return hours, minutes
end
function Clock:__tostring()
  return string.format("%02d:%02d", self.hours, self.minutes)
end
function Clock.at(...)
  local obj = {}
  obj.hours, obj.minutes = Clock.translate(...)

  setmetatable(obj, Clock)
  return obj
end

return Clock