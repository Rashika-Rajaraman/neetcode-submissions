/*
func twoSum(nums []int, target int) []int {
   
   // Brute Force
   // Time Complexity - O(n2)
   // Space Complexity - O(1)
    for i:=0; i<len(nums);i++{
      for j:=i+1; j<len(nums);j++{
         if nums[i]+nums[j] == target {
           return []int{i, j}
         }
      }
    }
    return []int{}
}
*/

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
