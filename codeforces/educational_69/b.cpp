#include <vector>
#include <algorithm>

#include "stdio.h"

#define N 200000

int main(int argc, char* argv[]) {
    int n, max, maxN;
    scanf("%d", &n);
    std::vector<int> columns(N, 0);
    max = 0; maxN = 0;
    for (int i = 0; i < n; i++) {
        scanf("%d", &columns[i]);
        if (columns[i] > max) {
            max = columns[i]; maxN = i;
        }
    }
    bool answer = true;
    for (int i = 0; i < n - 1; i++) {
        if (i < maxN && !(columns[i] < columns[i+1])) {
            answer = false;
            break;
        } else if (i > maxN && !(columns[i] > columns[i+1])) {
            answer = false;
            break;
        }
    }
    if (answer) {
        printf("YES\n");
    } else {
        printf("NO\n");
    }
}
