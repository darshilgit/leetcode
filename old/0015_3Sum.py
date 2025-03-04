class Solution:
    def threeSum(self, nums):
        nums.sort()
        # print('sorted nums')
        # print(nums)
        # print('------------')
        ans = []
        for i in range(len(nums)):
            # print('current i value: {}'.format(i))
            # print('------------')
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            left = i + 1
            right = len(nums) - 1
            # print('left index: {}'.format(left))
            # print('right index: {}'.format(right))
            # print('------------')
            while (left < right):
                # print('inside while')
                # print('nums[i]: {}'.format(nums[i]))
                # print('nums[left]: {}'.format(nums[left]))
                # print('nums[right]: {}'.format(nums[right]))
                tsum = nums[i] + nums[left] + nums[right]
                # print('tsum: {}'.format(tsum))
                # print('------------')
                if tsum < 0:
                    # print('tsum < 0')
                    left += 1
                    # print('inside tsum<0, left index: {}'.format(left))
                    # print('------------')
                elif tsum > 0:
                    right -= 1
                    # print('inside tsum>0, right index: {}'.format(right))
                    # print('------------')
                elif tsum == 0:
                    temp = []
                    temp.append(nums[i])
                    temp.append(nums[left])
                    temp.append(nums[right])
                    ans.append(temp)
                    while(left<right and nums[left] == nums[left+1]):
                        # print('inside nums[left] == nums[left+1]')
                        # print('left index: {}'.format(left))
                        left +=1
                        # print('updated left index, left index: {}'.format(left))
                        # print('------------')
                    while (left<right and nums[right] == nums[right-1]):
                        # print('inside nums[right] == nums[right-1]')
                        # print('right index: {}'.format(right))
                        right-=1
                        # print('updated right index, right index: {}'.format(right))
                        # print('------------')
                    # print('left index: {}'.format(left))
                    left += 1
                    # print('updated left index, left index: {}'.format(left))
                    # print('------------')
                    # print('right index: {}'.format(right))
                    right -= 1
                    # print('updated right index, right index: {}'.format(right))
                    # print('------------')
        return ans

a = Solution()
print(a.threeSum([-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]))