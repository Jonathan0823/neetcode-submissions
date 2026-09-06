class Solution {
    /**
     * @param {number[]} digits
     * @return {number[]}
     */
    plusOne(digits) {
        let i = digits.length - 1;
        while (i >= 0) {
            if (digits[i] < 9) {
                digits[i] += 1;
                return digits; 
            } else {
                digits[i] = 0;
            }
            i--
        }

        digits.unshift(1);
        return digits
    }
}
