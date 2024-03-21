def replace_elements(arr)
    mx = -1
    ans = [-1]
    
    if arr.length == ans.length
        return ans
    end

    arr.reverse_each do |val|
        if val > mx
            ans.unshift(val)
            mx = val
            next
        end
        ans.unshift(mx)
    end
    ans.shift
    ans
end

arr = [17, 18, 5, 4, 6, 1]
print replace_elements(arr)
