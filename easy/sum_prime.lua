primes = {2, 3, 5, 7, 11, 13}
function isPrime(a)
  local b = 1
  while a % primes[b] > 0 do
    if primes[b] * primes[b] > a then
      primes[#primes + 1] = a
      return true
    end
    b = b + 1
  end
  return false
end

function primeSeq()
  local p0, i = 1, 0
  return function()
           if p0 <= #primes then
             i, p0 = primes[p0], p0+1
             return i
           end
           i = i + 2
           while not isPrime(i) do
             i = i + 2
           end
           p0 = p0 + 1
           return i
         end
end

local primesum, nextPrime = 0, primeSeq()
for _ = 1, 1000 do
  primesum = primesum + nextPrime()
end
print(primesum)
