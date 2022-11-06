/**
 * @param {number[]} nums
 */
var Solution = function(nums) {
    this.nums = nums;
    this.random = [...nums];
};

/**
 * @return {number[]}
 */
Solution.prototype.reset = function() {
    return this.nums;
};

/**
 * @return {number[]}
 */
Solution.prototype.shuffle = function() {
    for(let i = 0; i < this.random.length; i++) {
        const rand = Math.floor(Math.random() * this.nums.length);
        this.swap(i, rand);
    }
    return this.random;
};

Solution.prototype.swap = function(i,j) {
    [this.random[i], this.random[j]] = [this.random[j], this.random[i]]
}