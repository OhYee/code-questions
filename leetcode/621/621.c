typedef struct {
    int name;
    int count;
} pair;

int compare(void *a, void *b) {
    return ((pair *)b)->count - ((pair *)a)->count;
}

int leastInterval(char *tasks, int tasksSize, int n) {
    // 先统计每个任务的个数
    pair pairs[26];
    for (int i = 0; i < 26; ++i) {
        pairs[i].name = i;
        pairs[i].count = 0;
    }
    for (int i = 0; i < tasksSize; ++i) {
        int taskName = tasks[i] - 'A';
        ++pairs[taskName].count;
    }

    // 排序
    qsort(pairs, 26, sizeof(pair), compare);

    int loop = 0;   // 当前任务轮
    int finish[26]; // 任务完成轮数
    memset(finish, -1, sizeof(int) * 26);
    // 执行任务
    while (1) {
        int index = -1; // 当前轮要执行的任务下标

        for (int i = 0; i < 26; ++i) {
            if (pairs[i].count <= 0)
                break;
            if (finish[pairs[i].name] < loop) {
                index = i;
                break;
            }
        }
        if (index == -1) {
            // 当前轮没有任务需要执行
            if (pairs[0].count == 0) {
                // 没有任务了
                return loop;
            }
        } else {
            // 执行下标为 index 的任务
            finish[pairs[index].name] = loop + n; // 记录下次可执行该任务的时间
            --pairs[index].count;
            // 更新队列，重新排序
            for (int i = index + 1; i < 26; ++i) {
                if (pairs[i].count >= pairs[i - 1].count) {
                    int tname = pairs[i - 1].name;
                    int tcount = pairs[i - 1].count;

                    pairs[i - 1].name = pairs[i].name;
                    pairs[i - 1].count = pairs[i].count;

                    pairs[i].name = tname;
                    pairs[i].count = tcount;
                } else
                    break;
            }
        }
        ++loop;
    }
}