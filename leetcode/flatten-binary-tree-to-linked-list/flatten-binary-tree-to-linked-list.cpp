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

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
    TreeNode() : val(-1), left(NULL), right(NULL) {}
};
//     1
// 2      5
//   3      6(rightTail)
//     4(leftTail)
class Solution {
public:
    void flatten(TreeNode* root) {
        flatten2(root);
    }
    TreeNode* flatten2(TreeNode* root) {
        if (root == NULL) {
            return NULL;
        }
        TreeNode* leftTail = flatten2(root->left);
        TreeNode* rightTail = flatten2(root->right);
        if (leftTail) {
            TreeNode* t = root->right;
            root->right = root->left;
            root->left = NULL;
            leftTail->right = t;            
        }
        if (rightTail)
            return rightTail;
        if (leftTail)
            return leftTail;
        return root;
    }
};

int val_array[100];
TreeNode root[100];
int init_val[] = {1,2,5,3,4,-1,6};
int len = (sizeof init_val) / (sizeof(int));
void build()
{
    memset(val_array, -1, sizeof val_array);
    memcpy(val_array, init_val, sizeof init_val);
    for (int i=0; i<len; i++) {        
        root[i].val = val_array[i];
        if (val_array[2*i+1] == -1) {
            // left null
            root[i].left = NULL;
        } else {
            root[i].left = &root[2*i+1];
        }
        if (val_array[2*i+2] == -1) {
            // left null
            root[i].right = NULL;
        } else {
            root[i].right = &root[2*i+2];
        }
    }
}
void p() {
    for (int i=0; i<len; i++) {
        cout << root[i].val << " ";
        if (root[i].left == NULL) {
            cout << "NULL" << " ";
        } else {
            cout << root[i].left->val << " ";
        }
        if (root[i].right == NULL) {
            cout << "NULL" << endl;
        } else {
            cout << root[i].right->val << endl;
        }
    }
}
int main()
{
    build();
    p();
    Solution s;
    s.flatten(root);
    p();
    return 0;
}

