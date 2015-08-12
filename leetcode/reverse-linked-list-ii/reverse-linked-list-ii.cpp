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
struct Input{
    int n;
    int* arr;
    int l;
    int r;
};
class Solution {
public:
    ListNode* reverseBetween(ListNode* head, int m, int n) {
        ListNode* fake_head = new ListNode(-1);
        fake_head->next = head;
        ListNode* pre = fake_head;
        int i=1;
        for (;i<m; i++, pre=pre->next);
        if (!pre->next) {
            return fake_head->next;
        }

        ListNode* end = pre->next;
        ListNode* cur = end->next;
        // node_1...node_m
        //            pre,end,...cur...node_n
        //            pre,pre_next,...,end,...cur...node_n
        
        for (; i<n; i++) {
            ListNode* post = cur->next;            
            // 在pre插入cur
            cur->next = pre->next;
            pre->next = cur;
            //跳过cur
            end->next = post;
            cur = post;
        }
        return fake_head->next;
    }
};

int main()
{
    int array[100];
    int l, r;
    int n;
    ListNode node_list[100];
    while (cin >> n) {
        for (int i=0; i<n; i++) {
            cin >> array[i];
        }
        cin >> l >> r;
        cout << "n: " << n
             << ", l: " << l
             << ", r: " << r
             << endl;
        for (int i=0; i<n; i++) {
            node_list[i].val=array[i];
            node_list[i].next = &node_list[i+1];
        }
        node_list[n-1].next = NULL;
        ListNode* cur = node_list;
        while(cur != NULL) {
            cout << cur->val << endl;
            cur = cur->next;
        }
        cout << "-------" << endl;
        Solution s;
        cur = node_list;
        cur = s.reverseBetween(cur, l, r);
        while(cur != NULL) {
            cout << cur->val << endl;
            cur = cur->next;
        }
        cout << "--------------------分割线--------------------" << endl;
    }
    return 0;
}

