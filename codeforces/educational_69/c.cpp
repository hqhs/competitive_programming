#include <vector>
#include <algorithm>

#include "stdio.h"

#define N 300000

typedef long long ll;

int main(int argc, char* argv[]) {
    int n, k;
    scanf("%d %d", &n, &k);
    std::vector<ll> arr(n, 0);
    std::vector<ll> dif(n-1, 0);
    for (int i = 0; i < n; i++) {
        scanf("%lld", &arr[i]);
        if (i > 0) {
            dif[i-1] = arr[i-1] - arr[i];
        }
    }
    // split into N groups and join them until where's K group based on principle:
    // find minimal diff between every two items, join them, repeat until done
    ll answer = arr[n-1] - arr[0];
    // printf("dif\n");
    // debug
    // for (ll e : dif) {
    //     printf("%lld ", e);
    // }
    // printf("\n");
    std::sort(dif.begin(), dif.end());
    // debug
    // printf("sorted\n");
    // for (ll e : dif) {
    //     printf("%lld ", e);
    // }
    // printf("\n");
    for (int i = 0; i < k-1; i++) {
        answer += dif[i];
    }
    printf("%lld\n", answer);
}
