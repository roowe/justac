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

class LRUCache{
private:
    map<int, pair<int, int> > cache; // <key, <value, timestamp>>
    map<int, int> cache_timestamp; // <timestamp, key>
    int capacity;
    int timestamp = 0;
    
public:
    LRUCache(int capacity0) {
        capacity = capacity0;
    }
    friend ostream &operator<<(ostream &os, LRUCache &LRUCache) {
        os << "timestamp : " << LRUCache.timestamp;
        return os;
    }
    int get(int key) {
        // print();
        // cout << "key " << key << endl;
        cout << *this << endl;
        auto it = cache.find(key);
        if (it == cache.end()) {
            return -1;
        } else {
            // cout << "get it->first " << it->first << endl;
            timestamp += 1;
            // cout << "old timestamp " << (it->second).second << endl;
            cache_timestamp.erase(cache_timestamp.find((it->second).second));
            cache_timestamp[timestamp] = key;
            cache[key] = pair<int, int>((it->second).first, timestamp);
            return it->second.first;
        }
    }
    void print() {
        cout << "--------begin--------------" << endl;
        for (auto it = cache.begin(); it != cache.end(); ++it) {
            cout << it->first << endl;
        }
        cout << "--------end----------------" << endl;
    }
    void set(int key, int value) {
        auto it = cache.find(key);
        if (cache.size() == capacity && it == cache.end()) {
            auto cache_timestamp_begin = cache_timestamp.begin();
            cache_timestamp.erase(cache_timestamp_begin);
            // cout << cache_timestamp_begin->first << ", " << cache_timestamp_begin->second << endl;
            cache.erase(cache.find(cache_timestamp_begin->second));            
        }
        timestamp += 1;
        if (it != cache.end()) {
            cache_timestamp.erase(cache_timestamp.find((it->second).second));
        }
        cache_timestamp[timestamp] = key;
        cache[key] = pair<int, int>(value, timestamp);
    }
};


// Input:	1,[set(2,1),get(2),set(3,2),get(2),get(3)]
// Output:	[1,1,2]
// Expected:	[1,-1,2]
// Input:	2,[set(2,1),set(2,2),get(2),set(1,1),set(4,1),get(2)]
// Output:	[1,1]
// Expected:	[2,-1]
// Input:	2,[set(2,1),set(1,1),get(2),set(4,1),get(1),get(2)]
// Output:	[1,1,-1]
// Expected:	[1,-1,1]
int main()
{   
    LRUCache s(2);
    s.set(2,1);
    s.set(1,1);
    cout << "ans " << (s.get(2)) << endl;
    s.set(4,1);
    cout << "ans " << (s.get(1)) << endl;
    cout << "ans " << (s.get(2)) << endl;
    return 0;
}
// other solution http://fisherlei.blogspot.com/2013/11/leetcode-lru-cache-solution.html
