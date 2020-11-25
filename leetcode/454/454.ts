 // @ts-ignore
function fourSumCount(A: number[], B: number[], C: number[], D: number[]): number {
    var m:{[key:number]:number} = {};
    var res = 0;
    A.map(a=>B.map(b=>m[a+b] = (m[a+b] || 0) + 1))
    C.map(c=>D.map(d=>res += (m[-c-d] || 0)))
    return res
};