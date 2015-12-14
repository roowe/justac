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
#include <unordered_set>
#include <stdarg.h>
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
#define myout cout<< __LINE__ <<" "
using namespace std;

class Solution {
private:
    queue<string> q;
    queue<int> len;
public:
    int ladderLength(string start, string end, unordered_set<string> &dict) {
        q.push(start);
        len.push(1);
        while (!q.empty()) {
            string cur = q.front();
            int d = len.front();
            q.pop();
            len.pop();
            //myout << cur << ", " << d << endl;
            if(cur == end) {
                return d;
            }
            // bfs 找该单词一次转变的相邻的单词

            for (int i=0; i<cur.length(); i++) {
                int tmp = cur[i];
                for (int j='a'; j<='z'; j++) {
                    cur[i] = j;
                    auto got = dict.find(cur);
                    if (got != dict.end()) {
                        //myout <<  *got << endl;
                        q.push(*got);
                        len.push(d+1);
                        dict.erase(*got);
                    }
                    if (cur == end) {
                        return d+1;
                    }
                }
                cur[i] = tmp;
            }         
        }
        return 0;
    }    
};
// start = "hit"
// end = "cog"
// dict = ["hot","dot","dog","lot","log"]
int main()
{    
    Solution s;
    string start="hit";
    string end="cog";
    unordered_set<string> dict = {"hot","dot","dog","lot","log"};
    cout << s.ladderLength(start, end, dict) << endl;
    return 0;
}

