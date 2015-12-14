#include<cstdio>
#include<iostream>
#include<set>
#include<map>
#include <unordered_map>
#include <unordered_set>
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
    vector<string> path;
    vector<vector<string>> ret;
    unordered_map<string, vector<string>> prevMap;
public:
    bool bfs(string beginWord, string endWord, unordered_set<string> &wordList) {
        for (auto word:wordList) {
            prevMap[word] = vector<string>();            
        }
        vector<unordered_set<string>> rows(2);
        int cur = 0, prev = 1;
        rows[cur].insert(beginWord);
        while(1) {
            cur = 1 - cur;
            prev = 1 - prev;
            for (auto word:rows[prev]) {
                // cout << prev << ":" << word << endl;
                wordList.erase(word);
            }
            rows[cur].clear();
            for (auto word:rows[prev]) {
                string w = word;
                for (int i=0; i<w.length(); i++) {
                    for (int j='a'; j<='z'; j++) {
                        if (word[i] == j) {
                            continue;
                        }
                        w[i] = j;
                        if(wordList.count(w) > 0 || w == endWord) {
                            prevMap[w].push_back(word);
                            rows[cur].insert(w);                            
                        }
                    }
                    w[i] = word[i];
                }                
            }
            if (rows[cur].size() == 0) {
                return false;
            }
            if (rows[cur].count(endWord) > 0) {
                return true;
            }
        }
    }
    void dfs(string word) {
        // cout << "dfs: " << word << endl;
        if (prevMap[word].size() == 0) {
            path.push_back(word);
            vector<string> path1 = path;
            reverse(path1.begin(), path1.end());
            ret.push_back(path1);
            path.pop_back();
            return;
        }
        path.push_back(word);
        for (auto s:prevMap[word]) {
            dfs(s);
        }
        path.pop_back();
    }
    vector<vector<string>> findLadders(string beginWord, string endWord, unordered_set<string> &wordList) {
        if (bfs(beginWord, endWord, wordList) == true) {
            // 从后面开始搜索是因为，从头开始搜的话，是没办法走到end，但是从end开始的话，由于记录的原因，肯定能走到start
            dfs(endWord);
        }            
        return ret;
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
    vector<vector<string>> ret =  s.findLadders(start, end, dict);
    for (auto vs: ret) {
        for (auto s: vs) {
            cout << s << ", ";
        }
        cout << endl;
    }
    return 0;
}

