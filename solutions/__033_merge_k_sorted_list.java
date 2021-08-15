import java.util.PriorityQueue;

public class __033_merge_k_sorted_list {
    public static void main(String[] args) {

    }

    public ListNode mergeKLists(ListNode[] lists) {
        PriorityQueue<ListNode> pq = new PriorityQueue<>((ListNode a, ListNode b) -> {
            return a.val - b.val;
        });
        ListNode Head = new ListNode();
        for (int i = 0; i < lists.length; i++) {
            if (lists[i] != null)
                pq.offer(lists[i]);
        }
        ListNode tmp = Head;
        while (!pq.isEmpty()) {
            ListNode node = pq.poll();
            if (node.next != null) {
                pq.offer(node.next);
            }
            tmp.next = new ListNode(node.val);
            tmp = tmp.next;
        }
        return Head.next;
    }
}

class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}
