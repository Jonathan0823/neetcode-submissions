class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {
        if (strs.length === 1) return [strs];
        const res = new Map();
        for (const s of strs) {
            const sortedS = s.split("").sort().join("")
            if (!res.has(sortedS)) {
                res.set(sortedS, [])
            }
            res.get(sortedS).push(s);
        }

        return Array.from(res.values())

    }
}
