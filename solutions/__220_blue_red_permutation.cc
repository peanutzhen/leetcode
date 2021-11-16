/*
 *      Copyright 2021 peanutzhen
 *      Created time: 2021/11/15 14:51:48
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
  ios::sync_with_stdio(false);
  cin.tie(nullptr);

  int t;
  cin >> t;
  while (t--) {
    int n;
    cin >> n;

    vector<int> a(n);
    for (int i = 0; i < n; i++)
      cin >> a[i];

    string color;
    cin >> color;

    vector<int> blue;
    vector<int> red;
    for (int i = 0; i < n; i++)
      if (color[i] == 'B')
        blue.push_back(a[i]);
      else
        red.push_back(a[i]);

    sort(blue.begin(), blue.end());
    sort(red.begin(), red.end());

    bool ok = true;
    int num = 1;
    for (int i = 0; i < blue.size(); i++) {
      if (blue[i] < num)
        ok = false;
      ++num;
    }
    for (int i = 0; i < red.size(); i++) {
      if (red[i] > num)
        ok = false;
      ++num;
    }
    if (ok)
      cout << "YES\n";
    else
      cout << "NO\n";
  }
  return 0;
}
