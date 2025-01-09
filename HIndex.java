public class HIndex {
    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.hIndex(new int[]{3, 0, 6, 1, 5}));
        System.out.println(s.hIndex(new int[]{1, 3, 1}));
        System.out.println(s.hIndex(new int[]{4, 4, 0, 0}));
    }
}
/*
// O(n log n)
// the treemap consumes O(log n) in the worst case to insert
// the loop is O(n)
// therefore, O(log n) * O(n) = O(n log n)
class Solution {
    public int hIndex(int[] citations) {
        TreeMap<Integer, Integer> tm = new TreeMap<>();
        int n = citations.length;
        for (int i = 0; i < n; i++) {
            int citation = citations[i];
            int key = Math.min(n, citation);
            Integer current = tm.putIfAbsent(key, 0);
            current = current == null ? 0: current;
            tm.put(key, current + 1);
        }
        Iterator<Entry<Integer,Integer>> it = tm.reversed().sequencedEntrySet().iterator();
        int count = 0;
        int hIndex = 0;
        while (it.hasNext() && hIndex == 0) {
            Entry<Integer, Integer> entry = it.next();
            Integer key = entry.getKey();
            if (count >= key) {
                hIndex = count;
                continue;
            }
            count += entry.getValue();
            if (count >= key) {
                hIndex = key;
            }
        }
        return hIndex;
    }
}
*/
// O(n)
class Solution {
    public int hIndex(int[] citations) {
        int n = citations.length;
        int[] tm = new int[n+1];
        for (int i = 0; i < n; i++) {
            int citation = citations[i];
            int key = Math.min(n, citation);
            tm[key]++;
        }
        int count = 0;
        int hIndex = 0;
        for (int i = n; i >= 0 && hIndex == 0; i--) {
            if(count >= i) {
                hIndex = count;
            } else {
                count += tm[i];
                if (count >= i) {
                    hIndex = i;
                }
            }
        }
        return hIndex;
    }
}
