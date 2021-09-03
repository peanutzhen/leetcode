/*
 *      Author: peanutzhen
 *      Created time: 2021/9/3 10:51:17
 */

#include <bits/stdc++.h>

using namespace std;

class LRUCache {
private:
    struct ListNode {
        int key;
        int val;
        ListNode *prev;
        ListNode *next;

        ListNode(int key, int value) : key(key), val(value), prev(nullptr), next(nullptr) {}
    };
    int size;
    int capacity;
    unordered_map<int, ListNode *> hashmap;
    ListNode *head;
    ListNode *tail;
public:
    LRUCache(int capacity) : size(0), capacity(capacity) {
        // 虚拟头指针
        head = new ListNode(0, 0);
        tail = new ListNode(0, 0);
        head->next = tail;
        tail->prev = head;
    }

    int get(int key) {
        if (hashmap.find(key) == hashmap.end())
            return -1;
        auto node = hashmap.find(key)->second;
        // 双向链表移除节点
        auto prev_node = node->prev;
        auto next_node = node->next;
        prev_node->next = next_node;
        next_node->prev = prev_node;
        // 移动至头部
        head->next->prev = node;
        node->next = head->next;
        node->prev = head;
        head->next = node;
        return node->val;
    }

    void put(int key, int value) {
        // 存在key则update
        if (hashmap.find(key) != hashmap.end()) {
            auto node = hashmap.find(key)->second;
            node->val = value;
            this->get(key); // 刷新至最近使用
            return;
        }
        // 是否超出容量限制 淘汰一个缓存
        if (this->size == this->capacity) {
            auto del_node = tail->prev;
            del_node->next = nullptr;
            del_node->prev->next = tail;
            tail->prev = del_node->prev;
            del_node->prev = nullptr;
            hashmap.erase(del_node->key);
            delete del_node;
            --this->size;
        }
        // put
        auto put_node = new ListNode(key, value);
        hashmap[key] = put_node;
        put_node->next = head->next;
        head->next->prev = put_node;
        head->next = put_node;
        put_node->prev = head;
        ++this->size;
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    auto *obj = new LRUCache(1);
    cout << obj->get(1) << "\n";    // -1
    obj->put(1, 1);
    cout << obj->get(1) << "\n";    // 1
    obj->put(1, 2);
    cout << obj->get(1) << "\n";    // 2
    obj->put(2, 3);
    cout << obj->get(1) << "\n";    // -1
    cout << obj->get(2) << "\n";    // 3
    return 0;
}

