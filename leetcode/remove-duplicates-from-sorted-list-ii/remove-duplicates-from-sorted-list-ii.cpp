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
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
    ListNode() : val(-1), next(NULL) {}    
};
class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        if (!head) {
            return NULL;
        }
        ListNode* fake_head = new ListNode(-1);
        ListNode* cur = fake_head;
        ListNode* i = head;
        while (i){
            if (i->next && i->next->val == i->val) {
                while (i->next && i->next->val == i->val)
                    i = i->next;
                if (i)
                    i = i->next;                
            } else{
                cur->next = i;
                cur = cur->next;
                i=i->next;
            }
        }
        cur->next = NULL;
        return fake_head->next;
    }
};
int main()
{
    ListNode node_list[100];
    for (int i=0; i<10; i++) {
        if (i < 4) {
            node_list[i].val=1;
        } else if(i >=4 && i < 8) {
            node_list[i].val = i;
        } else {
            node_list[i].val = 9;
        }
        node_list[i].next = &node_list[i+1];
    }
    node_list[9].next = NULL;
    ListNode* cur = node_list;
    while(cur != NULL) {
        cout << cur->val << endl;
        cur = cur->next;
    }
    cout << "--------------------" << endl;
    Solution s;
    cur = s.deleteDuplicates(node_list);
    while(cur != NULL) {
        cout << cur->val << endl;
        cur = cur->next;
    }
    return 0;
}

