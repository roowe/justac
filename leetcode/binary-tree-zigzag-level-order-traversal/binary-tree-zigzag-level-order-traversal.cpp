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
#define myout cout << ">> "
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;

    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};


TreeNode* build(vector<int>& array, int l, int r) {
    if (l > r) {
        return NULL;
    }
    int m = (l+r)/2;
    TreeNode* root = new TreeNode(array[m]);
    root->left = build(array, l, m-1);
    root->right = build(array, m+1, r);
    return root;
}

TreeNode* init(vector<int>& array) {
    return build(array, 0, array.size()-1);
}

void debug_bst(TreeNode* root)
{
    if (!root)
        return;
    if (!root->left && !root->right)
        return;

    myout << "root: " << root 
         << ", root->val: " << root->val;
    if (root->left) 
        cout << ", left: " << root->left->val;
    else
        cout << ", left: NULL";
    if (root->right)
        cout << ", right: " << root->right->val << endl;
    else
        cout << ", right: NULL" << endl;

    debug_bst(root->left);
    
    debug_bst(root->right);    
}

class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<TreeNode*>> vv(2);
        vector<vector<int>> ret;

        if (root == NULL) {
            return ret;
        }

        vector<int> row;

        vector<TreeNode*> v;
        v.push_back(root);
        vv[0] = v;

        int cur = 0;
        while(vv[cur].size() > 0) {
            int next = (cur+1)%2;
            v.clear();
            row.clear();
            // 考虑使用Stack去维护
            for (int i=0; i<vv[cur].size(); i++) {
                row.push_back(vv[cur][i]->val);
                if (vv[cur][i]->left)
                    v.push_back(vv[cur][i]->left);
                if (vv[cur][i]->right)
                    v.push_back(vv[cur][i]->right);
            }
            
            if (cur == 1) {
                vector<int> row1;
                for (auto iter=row.rbegin(); iter!=row.rend(); iter++) {
                    row1.push_back(*iter);
                }
                ret.push_back(row1);
            } else {
                ret.push_back(row);
            }
            
            vv[next] = v;
            cur = next;
        }        
        return ret;
    }
};
int main()
{
    int n;
    while(cin >> n) {
        vector<int> array(n);
        for (int i=0; i<n; i++)
            cin >> array[i];
        TreeNode* root = init(array);
        debug_bst(root);
        Solution s;
        vector<vector<int>> ret = s.zigzagLevelOrder(root);
        for (auto v:ret) {
            myout;
            for (auto i:v) {
                cout << i << " ";                
            }
            cout << endl;
        }        
    }
    return 0;
}




