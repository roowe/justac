#include<cstdio>
#include<iostream>
#include<set>
#include<map>
#include<cstring>
#include<string>
#include<algorithm>
#include<cmath>
#include<vector>
#include<stack>
#include<queue>
#include<cctype>
#define LL long long int
#define INF 0x3f3f3f3f
#define EPS 1.0e-8
#define RPE(i,s,e) for(int i = s; i < (int)(e); i++)
#define FOR(i,c) for(__typeof((c).begin()) i = (c).begin(); i != (c).end(); ++i)
#define MEM(c,i) memset(c, i, sizeof(c)) // i爲0或者-1
#define QC(q) while(!q.empty())q.pop()
#define IN(x,low,high) (x>low&&x<high)//low<x<high
#ifdef ONLINE_JUDGE  //判断是不是OJ系统上编译
#define debug(format, ...)  0
#else
char _red_buff[65535];
// {debug,     "\e[0;38m" },
// {info,      "\e[1;37m" },
// {notice,    "\e[1;36m" },
// {warning,   "\e[1;33m" },
// {error,     "\e[1;31m" },
// {critical,  "\e[1;35m" },
// {alert,     "\e[1;44m" },
// {emergency, "\e[1;41m" }
#define debug(format, ...) strcpy(_red_buff, "\e[1;36m");strcat(_red_buff, format);strcat(_red_buff, "\e[0m\r\n");printf(_red_buff, ##__VA_ARGS__)
#endif
using namespace std;

int p[3001];
int v[3001];
void set_v(int i)
{
    v[i] = 1;
    // i->p[i]...j->i
    for (int j=p[i]; j!=i; j=p[j]) 
        v[j]=1;
}

int main()
{
    int n, m;
    cin >> n;
    for (int i=1; i<=n; i++) {
        cin >> p[i];
    }
    cin >> m;
    int cycle_want = n - m;                  // 环的个数
    int cycle_origin = 0;
    for (int i=1; i<=n; i++) {
        if (!v[i]) {
            set_v(i);
            cycle_origin++;
        }
    }
    int cycle_diff = abs(cycle_origin - cycle_want);
    cout << cycle_diff << endl;

    if (cycle_origin < cycle_want) {
        debug("add_cycle");
        // 加环，将单个拆成两个
        while (cycle_diff--) {
            int i, j, min_i;
            for (i=1; p[i]==i; i++){}
            // 查找除去i之外的最小值
            for (j=p[i], min_i=p[i]; p[j]!=i; j=p[j]) {
                min_i = min(p[j], min_i);
            }
            cout << " " << i << " " << min_i;
            swap(p[min_i], p[i]);
        }
    } else {
        // 减环，其他环连到1去
        debug("sub_cycle");
        while (cycle_diff--){
            MEM(v, 0);
            set_v(1);
            int i;
            for (i=1; v[i]==1; i++) {}
            cout << " 1 " << i;
            swap(p[1],p[i]);
        }
    }
    cout << endl;
    return 0;
}

