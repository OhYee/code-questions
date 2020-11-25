/**
 * @param {number[]} A
 * @param {number[]} B
 * @param {number[]} C
 * @param {number[]} D
 * @return {number}
 */
var fourSumCount = function(A, B, C, D) {
    var m = {};
    var res = 0;
    A.map(a=>B.map(b=>m[a+b] = (m[a+b] || 0) + 1))
    C.map(c=>D.map(d=>res += (m[-c-d] || 0)))
    return res
};