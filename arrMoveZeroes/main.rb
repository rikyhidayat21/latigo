# nums = [0, 1, 0, 3, 12]
nums = [0, 0, 1]

j = 0 # tracking index non 0 value
nums.each_with_index do |v, i|
  if v != 0
    nums[i], nums[j] = nums[j], nums[i]
    j += 1
  end
end

p nums
