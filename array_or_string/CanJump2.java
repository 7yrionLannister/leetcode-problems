import java.util.LinkedList;
import java.util.Queue;

public class CanJump2 {
    public static void main(String[] args) {
        System.out.println(jump(new int[]{2, 3, 1, 1, 4}));
        System.out.println(jump(new int[]{2, 3, 0, 1, 4}));
        System.out.println(jump(new int[]{3, 5, 0, 0}));
    }

	public static Queue<Integer> queue;
	public static int n = 0;

    public static int jump(int[] nums) {
        n = nums.length;
        int[] jumps = new int[n];
        canJumpRecursive(nums, jumps);
		return jumps[n - 1];
    }

    private static void canJumpRecursive(int[] nums, int[] jumps) {
		queue = new LinkedList<>();
		queue.add(0);
		while (!queue.isEmpty()) {
			int index = queue.poll();
			int power = nums[index];
			int targetIndex = Math.min(index + power, n - 1);
			while (targetIndex > index) {
				// visit all neighbors
				if (jumps[targetIndex] == 0) {
					int newJumps = jumps[index] + 1;
					jumps[targetIndex] = newJumps;
					queue.add(targetIndex);
				}
				targetIndex--;
			}	
		}
    }
}