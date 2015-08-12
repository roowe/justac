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
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
class Solution {
private:
    ListNode* g_cur;
public:
    TreeNode* sortedListToBST(ListNode* head) {
        if (!head)
            return NULL;
        int size=0;
        ListNode* cur=head;
        for(; cur; cur=cur->next, size++);
        g_cur = head;
        return g(0, size-1);
    }
    TreeNode* g(int l, int r) {
        if (l>r) {
            return NULL;
        }
        int m = (l + r)/2;
        TreeNode* left = g(l, m-1);
        TreeNode* root = new TreeNode(g_cur->val);
        g_cur = g_cur->next;
        TreeNode* right = g(m+1,r);
        root->left = left;
        root->right = right;
        return root;
    }
};
void debug_bst(TreeNode* root)
{
    if (!root)
        return;
    debug_bst(root->left);
    cout << "root: " << root 
         << ", root->val: " << root->val;
    if (root->left)
        cout << ", left: " << root->left->val;
    else
        cout << ", left: NULL";
    if (root->right)
        cout << ", right: " << root->right->val << endl;
    else
        cout << ", right: NULL" << endl;
    debug_bst(root->right);    
}
int main()
{
    int array[100];
    int n;
    ListNode node_list[100];
    while (cin >> n) {
        for (int i=0; i<n; i++) {
            cin >> array[i];
        }
        cout << "n: " << n
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
        TreeNode* root = s.sortedListToBST(node_list);
        debug_bst(root);
        cout << "--------------------分割线--------------------" << endl;
    }
    return 0;
}

