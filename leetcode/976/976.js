/**
 * @param {number[]} A
 * @return {number}
 */
var largestPerimeter = function(A) {
    A.sort((a,b)=>b-a);
    for (var i=2; i<A.length; ++i) 
        if (A[i-2] < A[i-1] + A[i])
            return A[i-2] + A[i-1] + A[i];
    return 0;
};