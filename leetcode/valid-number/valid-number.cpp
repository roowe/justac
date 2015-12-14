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
    if (array[m] == -1) {
        return NULL;
    }
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
    bool isNumber(string s) {
        int e_flag = 0;
        int dot_flag = 0;
        int num_flag = 0;
        int begin = 0;
        int end = s.length()-1;
        //myout << begin << ", " << end << endl;
        for (;s[begin] == ' '; ++begin);
        if (s[begin] == '-' || s[begin] == '+') {
            ++begin;
        }
        for (;s[end] == ' '; --end);
        myout << begin << ", " << end << endl;
        if (begin > end)
            return false;
        for (int i=begin; i<=end; i++) {     
            if (s[i] == 'e') {
                if (e_flag || num_flag == false) {
                    return false;
                } else {
                    e_flag = 1;
                    num_flag = 0;
                }
            } else if(s[i] == '.') {
                if (dot_flag || e_flag) {
                    return false;
                } else {
                    dot_flag = 1;
                }
            } else if(s[i] >= '0' && s[i] <= '9') {
                num_flag = 1;
            } else if (s[i] == '-' || s[i] == '+') {
                if (s[i-1] != 'e')
                    return false;
            } else {
                return false;
            }
        }
        return num_flag;
    }
};
int main()
{
    char s[1000];
    while(fgets(s, 1000, stdin) != NULL) {
        myout << "--------------------" << endl;
        Solution solution;
        s[strlen(s)-1] = 0;
        myout << s << endl;
        bool ret =  solution.isNumber(s);
        myout << ret << endl;
    }
    return 0;
}


