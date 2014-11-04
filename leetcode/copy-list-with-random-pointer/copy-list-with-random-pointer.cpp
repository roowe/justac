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
// Definition for singly-linked list with a random pointer.

struct RandomListNode {
    int label;
    RandomListNode *next, *random;
    RandomListNode(int x) : label(x), next(NULL), random(NULL) {}
};

void print(RandomListNode *head) {
    cout << "--------------------begin--------------------" << endl;
    while (head != NULL) {
        cout << head << " label " << head->label << ", random " << head->random->label << endl;
        head = head->next;
    }
    cout << "-------------------- end --------------------" << endl;
}

class Solution {
    map<RandomListNode *, RandomListNode *> hash;
    
public:
    RandomListNode *copyRandomList(RandomListNode *head) {
        if (head == NULL) {
            return NULL;
        }
        // cout << head << endl;
        RandomListNode *copy = new RandomListNode(head->label);
        RandomListNode *copy_cur = copy;
        RandomListNode *cur = head;
        while (cur != NULL) {
            copy_cur->random = cur->random;
            hash[cur] = copy_cur;
            if (cur->next != NULL) {
                RandomListNode *next_copy = new RandomListNode(cur->next->label);
                copy_cur->next = next_copy;
                copy_cur = next_copy;
            }
            cur = cur->next;
        }
        // for (auto it=hash.begin(); it != hash.end(); ++it) {
        //     cout << "old " << it->first << ", new " << it->second << endl;
        // }
        // print(copy);
        copy_cur = copy;
        while (copy_cur != NULL) {
            copy_cur->random = hash[copy_cur->random];
            copy_cur = copy_cur->next;
        }
        return copy;
    }
};
RandomListNode *test_data() {
    RandomListNode *copy = new RandomListNode(1);
    RandomListNode *copy_cur = copy;
    vector<RandomListNode*> points;
    points.push_back(copy);
    for(int i=2; i<10; i++) {
        RandomListNode *next_copy = new RandomListNode(i);
        copy_cur->next = next_copy;
        copy_cur = next_copy;
        points.push_back(next_copy);
    }
    copy_cur = copy;
    int size = points.size();
    while (copy_cur != NULL) {
        int index = rand() % size;
        copy_cur->random = points[index];
        copy_cur = copy_cur->next;
    }
    return copy;
}
int main()
{   
    Solution s;
    RandomListNode *data = test_data();
    print(data);
    RandomListNode *data1 = s.copyRandomList(data);
    print(data1);    
    return 0;
}

