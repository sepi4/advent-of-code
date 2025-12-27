local M = {}
M.readInput = function(path)
  local file = io.open(path, 'r')
  if file then
    local input = file:read('*a')
    io.close(file)
    return input
  end
  return ''
end

return M
