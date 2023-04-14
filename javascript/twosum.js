/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    var newArr = new Array();
    for(var i = 0; i < nums.length ; i++){
        for( var x = i+1; x < nums.length; x++){
        if(nums[i] + nums[x] == target){
            newArr.push(i , x)
        }
        }
    }
    return newArr
};