/**
 * @param {string} s
 * @return {character}
 */


// time complexity O(n^2)
// space complexity O(n)



var repeatedCharacter = function(s) {
    
    let newArr = s.split('');
    
    for(var i = 0; i < newArr.length; i++ ){
        for(var x = 0; x < newArr.length; x++){
        if(newArr[i] == newArr[x]){
            var closest = newArr[i];
            if(i > x){
                return closest
            }
            }
        
        }
    }
    
    };
    
    
    