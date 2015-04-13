function permgen(a, n)
  n = n or #a
  if n <= 1 then
    coroutine.yield(a)
  else
    for i = 1, n do
      a[n], a[i] = a[i], a[n]
      permgen(a, n-1)
      a[n], a[i] = a[i], a[n]
    end
  end
end

function permutations(a)
  local co = coroutine.create(function() permgen(a) end)
  return function()
    local code, res = coroutine.resume(co)
    return res
  end
end

for line in io.lines(arg[1]) do
  a, f = {}, false
  for i in line:gmatch("%d+") do
    a[#a+1] = tonumber(i)
  end
  for i in permutations(a) do
    for o1 = 1, 3 do
      if o1 == 1 then
        r1 = i[1] + i[2]
      elseif o1 == 2 then
        r1 = i[1] - i[2]
      else
        r1 = i[1] * i[2]
      end
      for o2 = 1, 3 do
        if o2 == 1 then
          r2 = r1 + i[3]
        elseif o2 == 2 then
          r2 = r1 - i[3]
        else
          r2 = r1 * i[3]
        end
        for o3 = 1, 3 do
          if o3 == 1 then
            r3 = r2 + i[4]
          elseif o3 == 2 then
            r3 = r2 - i[4]
          else
            r3 = r2 * i[4]
          end
          for o4 = 1, 3 do
            if o4 == 1 then
              r4 = r3 + i[5]
            elseif o4 == 2 then
              r4 = r3 - i[5]
            else
              r4 = r3 * i[5]
            end
            if r4 == 42 then
              f = true
              break
            end
          end
          if f then break end
        end
        if f then break end
      end
      if f then break end
    end
  end
  if f then
    print("YES")
  else
    print("NO")
  end
end
