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
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        // cout << l1->next <<" "<< l2->next << endl;
        ListNode* tmp_head = new ListNode(-1);
        ListNode* cur = tmp_head; 
        while (l1 != NULL && l2 != NULL) {
            // cout << l1->val << ", " << l2->val << endl;
            // cout << "cur->next " << cur->next << endl;
            if (l1->val < l2->val) {
                cur->next = l1;
                cur = l1;
                l1 = l1->next;                
            } else {
                cur->next = l2;
                cur = l2;                                
                l2 = l2->next;
            }
        }
        // cout << "cur->next " << cur->next << endl;
        while (l1 != NULL) {
            cur->next = l1;
            cur = l1;
            l1 = l1->next;            
        }
        while (l2 != NULL) {
            cur->next = l2;
            cur = l2;                                
            l2 = l2->next;
        }
        // cout << "cur->next " << cur->next << endl;
        return tmp_head->next;
    }
};
int main()
{
    ListNode node_list[100];
    for (int i=0; i<10; i++) {
        cout << i << " " << &node_list[i] << endl;
        node_list[i].val=2*i+1;
        node_list[i].next = &node_list[i+1];
    }
    node_list[9].next = NULL;
    for (int i=50; i<60; i++) {
        cout << i << " " << &node_list[i] << endl;
        node_list[i].val=2*i;
        node_list[i].next = &node_list[i+1];
    }
    node_list[59].next = NULL;
    Solution s;
    ListNode* cur = s.mergeTwoLists(node_list, node_list+50);
    cout << cur << endl;
    while(cur != NULL) {
        cout << cur->val << endl;
        cur = cur->next;
    }
    return 0;
}

