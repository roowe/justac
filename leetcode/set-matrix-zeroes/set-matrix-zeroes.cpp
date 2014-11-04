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
    void setZeroes(vector<vector<int> > &matrix) {
        int n = matrix.size();
        int m = matrix[0].size();
        vector<int> i_flag(n, 0);
        vector<int> j_flag(m, 0);
        for (int i=0; i<n; i++) {
            for (int j=0; j<m; j++) {
                if( matrix[i][j] == 0) {
                    i_flag[i] = 1;
                    j_flag[j] = 1;
                }
            }
        }
        for (int i=0; i<n; i++) {
            if (i_flag[i] == 1) {
                for (int j=0; j<m; j++) {
                    matrix[i][j] = 0;
                }                    
            }
        }
        for (int j=0; j<m; j++) {
            if (j_flag[j] == 1) {
                for (int i=0; i<n; i++) {
                    matrix[i][j] = 0;
                }                    
            }
        }
    }
};
int main()
{   
    Solution s;
    vector<vector<int> > matrix  {{1,2,3}, {4,0,56}};
    int n = matrix.size();
    int m = matrix[0].size();
    for (int i=0; i<n; i++) {
        for (int j=0; j<m; j++) {
            cout << matrix[i][j] << " ";
        }
        cout << endl;
    }
    cout << "--------------" << endl;
    s.setZeroes(matrix);
    for (int i=0; i<n; i++) {
        for (int j=0; j<m; j++) {
            cout << matrix[i][j] << " ";
        }
        cout << endl;
    }
    return 0;
}

