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
private:
    struct cmp{
        bool operator()(const ListNode* v1, const ListNode* v2){
            return v1->val > v2->val;
        }
    };
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        priority_queue<ListNode*, vector<ListNode*>, cmp> q;
        for (auto iter : lists) {
            if (iter) 
                q.push(iter);
        }
        ListNode* ret = new ListNode(-1);
        ListNode* cur = ret;
        while (!q.empty()) {
            auto node = q.top();
            q.pop();
            cur->next = node;
            cur = cur->next;
            if (node->next != NULL) {
                q.push(node->next);
            }
        }
        return ret->next;
    }
};
int main()
{
    int list_n;

    while (cin >> list_n) {
        int n;
        vector<ListNode*> data;
        for (int i=0; i<list_n; i++) {
            cin >> n;
            ListNode* node_list = new ListNode(-1);
            ListNode* cur = node_list;
            int val;
            for (int j=0; j<n; j++) {
                cin >> val;
                cur->next = new ListNode(val);
                cur = cur->next;
            }
            data.push_back(node_list->next);
        }
        for (int i=0; i<list_n; i++) {
            ListNode* cur = data[i];
            cout << cur->val << endl;
            while(cur->next != NULL) {
                cout << cur->val << "->";
                cur = cur->next;
            }
            cout << cur->val << endl;
        }
        cout << "--------------------" << endl;
        Solution s;
        vector<ListNode*> data1;
        ListNode* cur = s.mergeKLists(data1);
        cout << cur << endl;
        // while(cur->next != NULL) {
        //     cout << cur->val << "->";
        //     cur = cur->next;
        // }
        // cout << cur->val << endl;
    }
    return 0;
}

