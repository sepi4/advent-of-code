local r = require("./lib/lib")
local ss = r.readInput("./04/example.txt")

print(ss)
for k, v in ipairs(ss) do
  print(k, v)
end
