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

//Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
    ListNode() : val(-1), next(NULL) {}    
};
ListNode* Add(int val, ListNode *node) {

}
class Solution {

public:
    // ListNode* partion(ListNode* begin, ListNode* end) {
    //     int v=begin->val;
    //     ListNode* i=begin;
    //     ListNode* j=begin->next;
    //     while (j != end) {
    //         if (j->val < v) {
    //             i = i->next;
    //             swap(i->val, j->val);
    //         }
    //         j = j->next;
    //     }
    //     swap(i->val, begin->val);
    //     return i;
    // }
    // void qc(ListNode* begin, ListNode* end) {
    //     if (begin != end) {
    //         ListNode* p = partion(begin, end);
    //         qc(begin, p);
    //         qc(p->next, end);
    //     }
    // }
    ListNode* sortList(ListNode* head) {
        int n = 0;
        ListNode* cur = head;
        while (cur != NULL) {
            cur = cur->next;
            n++;
        }
        // debug("n = %d \n", n);
        ListNode* tmp_head = new ListNode(-1);
        tmp_head->next = head;
        for (int sz=1; sz < n; sz = sz*2) {
            ListNode* pre = tmp_head;
            ListNode* slow = pre->next;
            ListNode* fast = pre->next;
            while (slow != NULL || fast != NULL) {
                for (int i=0; fast != NULL && i < sz; fast = fast->next, i++);
                // merge
                int si=0, fi=0;
                while (si < sz && fi < sz && slow != NULL && fast != NULL) {
                    if (slow->val > fast->val) {
                        pre->next = fast;
                        pre = fast;
                        fast = fast->next;
                        fi++;
                    } else {
                        pre->next = slow;
                        pre = slow;
                        slow = slow->next;
                        si++;
                    }                    
                }
                //debug("%d %d\n", si, fi);
                while (si < sz && slow != NULL) {
                    pre->next = slow;
                    pre = slow;
                    slow = slow->next;
                    si++;
                }
                while (fi < sz && fast != NULL) {
                    pre->next = fast;
                    pre = fast;
                    fast = fast->next;
                    fi++;
                }
                pre->next = fast;
                slow = fast;
            }            
        }
        ListNode* new_head = tmp_head->next;
        delete tmp_head;
        return new_head;
    }
};
int main()
{
    ListNode node_list[100];
    for (int i=0; i<10; i++) {
        node_list[i].val=10-i;
        node_list[i].next = &node_list[i+1];
    }
    ListNode* cur = node_list;
    while(cur != NULL) {
        cout << cur->val << endl;
        cur = cur->next;
    }
    Solution s;
    cur = s.sortList(node_list);
    while(cur != NULL) {
        cout << cur->val << endl;
        cur = cur->next;
    }
    return 0;
}

