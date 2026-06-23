func twoSum(nums []int, target int) []int {
   seen := make(map[int]int)

   // HashMap
   // Time Complexity - O(n)
   // Space Complexity - O(n)
   for i, num := range nums {
      value := target - num
      if _, exists := seen[value]; exists {
         return []int{seen[value], i}
      }
      seen[num] = i
   }
   return []int{}
}
