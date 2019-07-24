#include <vector>
#include <algorithm>

#include "stdio.h"

int main(int argc, char* argv[]) {
    int t, n, m1, m2, tmp;
    scanf("%d", &t); // requests
    std::vector<int> len(100000, 0);
    for (int i = 0; i < t; i++) {
        m1 = 0; m2 = 0;
        scanf("%d", &n);
        for (int f = 0; f < n; f++) {
            scanf("%d", &tmp);
            // printf("%d\n", tmp); // debug
            if (tmp > m1) {
                m2 = m1;
                m1 = tmp;
            } else if (tmp > m2) {
                m2 = tmp;
            }
        }
        // find max0, max1 from W
        // min(min(max0, max1) - 1, N - 2) is max length
        len[i] = std::min(std::min(m1, m2) - 1, n - 2);
        //debug
        // printf("m1: %d, m2: %d\n", m1, m2);
    }
    for (int i = 0; i < t; i++) {
        printf("%d\n", len[i]);
    }
}
