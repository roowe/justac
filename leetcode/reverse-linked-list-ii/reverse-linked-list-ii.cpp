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
    ListNode* reverseBetween(ListNode* head, int m, int n) {
        ListNode* fake_head = new ListNode(-1);
        ListNode* cur = fake_head;
        int i=1;
        while(head) {
            // cout << "head -> " << head << endl;
            if (i == m) {
                ListNode* reverse_last = head;
                ListNode* pre = head;
                head = head->next;
                i++;
                while (i<=n) {
                    ListNode* post = head->next;
                    head->next = pre;
                    pre = head;
                    head = post;
                    i++;
                }
                // cout << "cur --> " << cur << endl;
                // cout << "pre --> " << pre << endl;
                // cout << "head --> " << head << endl;
                // cout << "head->next --> " << head->next << endl; 
                cur->next = pre;
                reverse_last->next=head;
                cur = head;
                if (head)
                    head = head->next;
            } else {
                cur->next = head;
                cur = cur->next;
                head = head->next;
                i++;
            }
            // if (i>20) break;
        }
        return fake_head->next;
    }
};

int main()
{
    ListNode node_list[100];
    for (int i=0; i<10; i++) {
        node_list[i].val=i;
        node_list[i].next = &node_list[i+1];
    }
    ListNode* cur = node_list;
    while(cur != NULL) {
        cout << cur << endl;
        cur = cur->next;
    }
    cout << "-------" << endl;
    Solution s;
    cur = node_list;
    cur = s.reverseBetween(cur,1,11);
    while(cur != NULL) {
        cout << cur << endl;
        cur = cur->next;
    }
    return 0;
}

