class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isPalindrome(s) {
        const filteredS = s.toLocaleLowerCase().replace(/\s+/g, "").replace(/[^a-z0-9]/g, '')
        let i = 0;
        let j = filteredS.length - 1;

        while (i < j) {
            if (filteredS[i] !== filteredS[j]) return false;
            i++;
            j--;
        }

        return true
    }
}
