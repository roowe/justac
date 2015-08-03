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
    bool isPalindrome(ListNode* head) {
        if (head == NULL) {
            return true;
        }        
        int n = len(head);
        if (n == 1) {
            return true;
        }
        int m;
        if (n % 2 == 0) {
            m = n / 2;
        } else{
            m = n / 2  + 1;
        }
        ListNode* right = head;
        for (int i = 0; i < m; right=right->next, i++);
        // ListNode slow = head;  
        // ListNode fast = slow.next;  
        // while (fast != null && fast.next != null && fast.next.next != null){  
        //     slow = slow.next;  
        //     fast = fast.next.next;  
        // }
        // 可以用slow,fast
        right = reverse(right);
        // ListNode* cur = right;
        // while(cur != NULL) {
        //     cout << cur->val << endl;
        //     cur = cur->next;
        // }
        while (right != NULL && head->val == right->val) {
            head = head->next;
            right = right->next;
        }
        if (right == NULL) {
            return true;
        } else {
            return false;
        }
    }
    int len(ListNode* head) {
        int l = 0;
        for (;head != NULL; head=head->next, l++);
        return l;
    }
    // also slove https://leetcode.com/problems/reverse-linked-list/
    ListNode* reverse(ListNode* head) {
        if (head == NULL) {
            return NULL;
        }
        ListNode* cur = head->next;
        ListNode* pre = head;
        pre->next = NULL;
        while (cur != NULL ) {
            ListNode* post = cur->next;
            cur->next = pre;
            pre = cur;
            cur = post;
        }
        return pre;
    }
};

int main()
{
    ListNode node_list[100];
    for (int i=0; i<10; i++) {
        node_list[i].val=0;
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
    cout << s.isPalindrome(cur) << endl;
    return 0;
}

