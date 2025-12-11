local r = require('read_input')
local input = r.readInput('input_01.txt')
-- local input = r.readInput('example_01.txt')

local moves = {}
local s = ''
-- print(string.sub(input, 1, 1))
for i = 1, #input do
  local c = input:sub(i, i)
  -- print(i, c)
  if c == '\n' or i == #input then
    table.insert(moves, {
      dir = s:sub(1, 1),
      amount = tonumber(s:sub(2)),
    })
    s = ''
  else
    s = s .. c
  end
end

-- game
local position = 50
local times = 0

for _, x in pairs(moves) do
  local d = x.dir == 'L' and -1 or 1
  local a = x.amount

  while a > 0 do
    position = position + d

    if position < 0 then
      position = position + 100
    elseif position > 99 then
      position = position - 100
    end

    if position == 0 then
      times = times + 1
    end

    a = a - 1
  end
end

print('times', times)
