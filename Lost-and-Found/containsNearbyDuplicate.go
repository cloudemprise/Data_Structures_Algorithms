// Solution 2: Time Complexity Space Complexity
func containsNearbyDuplicate(nums []int, k int) bool {

	var ln int = len(nums)

	//if ln == 2 && nums[0] == nums[1] { return true }

	for n := 0; n < ln-1; n++ {
		for m := 1; m <= k; m++ {
			if idx:= n+m; ln-1 < idx {
				break
			} else if nums[n] == nums[idx] {
				return true
			}
		}
	}
	return false
}