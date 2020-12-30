void push(int w, int *heap, int *heapSize) {
    heap[*heapSize] = w;
    int t = (*heapSize)++, temp;

    while (t) {
        int p = (t - 1) / 2;
        if (heap[p] < heap[t]) {
            temp = heap[p]; heap[p] = heap[t]; heap[t] = temp;
            t = p;
        } else {
            break;
        }
    }
}

int pop(int *heap, int *heapSize) {
    int res = heap[0];
    heap[0] = heap[--(*heapSize)];

    int p = 0, temp;
    while (p < *heapSize) {
        int l = p * 2 + 1, r = l + 1, t = p;
        if (l < *heapSize && heap[l] > heap[t]) t = l;
        if (r < *heapSize && heap[r] > heap[t]) t = r;
        if (p != t) {
            temp = heap[p]; heap[p] = heap[t]; heap[t] = temp;
            p = t;
        } else {
            break;
        }
    }
    return res;
}

int lastStoneWeight(int* stones, int stonesSize){
    if (stonesSize == 0) return 0;
    if (stonesSize == 1) return stones[0];
    int heap[stonesSize * 2], heapSize = 0;

    for (int i=0; i<stonesSize; ++i)
        push(stones[i], heap, &heapSize);

    while (heapSize > 1) {
        int x = pop(heap, &heapSize);
        int y = pop(heap, &heapSize);
        if (x != y) push(x-y, heap, &heapSize);
    }
    return heapSize == 0 ? 0 : heap[0];
}
