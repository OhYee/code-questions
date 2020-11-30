typedef struct _pair {
    char    c;
    int     count;
} pair;

int cmp(const void* a, const void* b) {
    return ((pair*)b)->count - ((pair*)a)->count;
}

void reSort(pair *pairs) {
    for(int i=0; i<25; ++i) {
        if(pairs[i].count <= pairs[i+1].count) {
            char temp_c = pairs[i].c;
            int temp_count = pairs[i].count;

            pairs[i].c = pairs[i+1].c;
            pairs[i].count = pairs[i+1].count;

            pairs[i+1].c = temp_c;
            pairs[i+1].count = temp_count;
        } else {
            break;
        }
    }
}

char * reorganizeString(char * S){
    pair pairs[26];
    for(int i=0; i<26; ++i) {
        pairs[i].c = 'a' + i;
        pairs[i].count = 0;
    }

    int length = 0;
    char *c = S;
    while(*c != '\0') {
        ++pairs[*c-'a'].count;
        ++length;
        ++c;
    }
    
    char lst = '\0';
    qsort(pairs, 26, sizeof(pair), cmp);

    for(int i=0; i<length; ++i) {
        int j;
        for(j=0; j<26; ++j) 
            if(pairs[j].c != lst && pairs[j].count != 0)
                break;   
        if (j < 26) {
            lst = S[i] = pairs[j].c;
            --pairs[j].count;

            reSort(pairs);
        } else {
            S[0] = '\0';
            break;
        }
    }

    return S;
}