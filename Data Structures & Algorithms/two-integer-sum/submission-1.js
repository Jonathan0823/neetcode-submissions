class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {
        const map = new Map()
        let index = 0
        for(let num of nums) {
            const completion = target - num
            if (map.has(completion)) return [map.get(completion), index]
            map.set(num, index)
            index++;
        } 
    }
}
