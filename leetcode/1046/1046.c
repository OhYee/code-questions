int cmp(void *a, void *b) {
    return *(int*)b - *(int*)a;
}

int lastStoneWeight(int* stones, int stonesSize){
    qsort(stones, stonesSize, sizeof(int), cmp);
    // for (int i=0; i<stonesSize; ++i) printf("%d ", stones[i]); printf("\n");
    while (stonesSize > 1) {
        if (stones[0] == stones[1]) {
            // 两个同时删除
            // printf("%d == %d\n", stones[0], stones[1]);
            for (int i=0; i<stonesSize-2; ++i)
                stones[i] = stones[i+2];
            stonesSize -= 2;
            // for (int i=0; i<stonesSize; ++i) printf("%d ", stones[i]); printf("\n");
        } else {
            // 删除两个，插入 stones[0] - stones[1]
            // printf("%d > %d\n", stones[0], stones[1]);
            int temp = 2;
            int newStone = stones[0] - stones[1];
            for (int i=0; i<stonesSize-1; ++i) {
                if (temp == 2 && (i+temp == stonesSize || newStone >= stones[i+temp])) {
                    stones[i] = newStone;
                    --temp;
                    ++i;
                }
                if (i+temp < stonesSize) stones[i] = stones[i+temp];
            }
            stonesSize -= 1;
            // for (int i=0; i<stonesSize; ++i) printf("%d ", stones[i]); printf("\n");
        }
    }
    return stonesSize == 1 ? stones[0] : 0;
}
