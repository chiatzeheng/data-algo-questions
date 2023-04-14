/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    if(x < 0|| (x % 10 == 0 && x != 0)){
        return false;
    }
  
    var number1 = 0 
        while(x > number1){
            number1 = number1 * 10 + x % 10
             x = ~~(x/10);
        }
    return x === number1 || x === ~~(number1/10);
};