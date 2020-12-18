char findTheDifference(char * s, char * t){
    int cnt[26];
    memset(cnt, 0, sizeof(cnt));
    char *c = s;
    while (*c != '\0') {
        ++cnt[*c-'a'];
        c++;
    }
    c = t;
    while (*c != '\0') {
        --cnt[*c-'a'];
        if (cnt[*c-'a']<0) return *c;
        c++;
    }
    return ' ';
}
