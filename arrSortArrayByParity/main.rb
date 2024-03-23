nums = [3, 1, 2, 4]

# not in place
ans = []

nums.each do |val|
  if val % 2 != 0
    ans << val
  else
    ans.unshift(val)
  end
end

p ans
