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
class Solution {
public:
    string addBinary(string a, string b) {
        int len = max(a.length(), b.length());
        string ret(len, '0');
        int i = a.length() - 1;
        int j = b.length() - 1;
        int k = 0;
        int in = 0;
        while (i >= 0 || j >= 0) {
            int av = 0, bv = 0;
            if (i >= 0) {
                av = a[i--] - '0';
            }
            if (j >= 0) {
                bv = b[j--] - '0';
            }
            int v = av+bv+in;
            if (v >= 2) {
                v = v % 2;
                in = 1;
            } else {
                in = 0;
            }
            ret[k++] = v + '0';
        }
        if (in)
            ret.resize(len+1,'1');
        reverse(ret.begin(), ret.end());
        return ret;
    }
};
int main()
{
    Solution s;
    cout << s.addBinary("11", "0") << endl;
    cout << s.addBinary("11", "1") << endl;
    cout << s.addBinary("1111", "1111") << endl;
    cout << s.addBinary("1010", "1011") << endl;
    return 0;
}

