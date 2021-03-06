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
    ListNode() : val(0), next(NULL) {}    
};
class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        if (headA == NULL || headB == NULL) {
            return NULL;
        }
        ListNode* pA = headA;
        ListNode* pB = headB;
        ListNode* tailA = NULL;
        ListNode* tailB = NULL;
        while(true) {
            if (pA == NULL) {
                pA = headB;
            }
            if (pB == NULL) {
                pB = headA;
            }
            if (pA->next == NULL) {
                tailA = pA;
            }
            if (pB->next == NULL) {
                tailB = pB;
            }
            if (tailA != NULL && tailB != NULL && tailA != tailB) {
                return NULL;
            }
            if (pA == pB) {
                return pA;
            }
            pA = pA->next;
            pB = pB->next;
        }
    }
};
int main()
{
    ListNode node_list[100];
    for (int i=0; i<10; i++) {
        node_list[i].val=0;
        node_list[i].next = &node_list[i+1];
    }
    ListNode node_list_b[100];
    for (int i=0; i<3; i++) {
        node_list_b[i].val=0;
        node_list_b[i].next = &node_list_b[i+1];
    }
    node_list_b[2].next = &node_list[5];
    ListNode* cur = node_list;
    while(cur != NULL) {
        cout << cur << endl;
        cur = cur->next;
    }
    cout << "--------------------" << endl;
    cur = node_list_b;
    while(cur != NULL) {
        cout << cur << endl;
        cur = cur->next;
    }
    cout << "--------------------" << endl;
    Solution s;
    cur = s.getIntersectionNode(node_list, node_list_b);
    cout << "ret " << cur << endl;
    while(cur != NULL) {
        cout << cur->val << endl;
        cur = cur->next;
    }
    return 0;
}

