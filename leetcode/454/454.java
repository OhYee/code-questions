class Solution {
    public int fourSumCount(int[] A, int[] B, int[] C, int[] D) {
        Map<Integer, Integer> m = new HashMap();
        int res = 0;
        for (int a : A)
            for (int b : B)
                m.put(a + b, m.getOrDefault(a + b, 0) + 1);
        for (int c : C)
            for (int d : D)
                res += m.getOrDefault(-c - d, 0);
        return res;
    }
}