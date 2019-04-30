class Solution:
    def anagramMappings(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: List[int]
        """
        dict_b = {}
        list_p = [0] * len(A)
        i = 0
        for ele in B:
            dict_b[ele] = i
            i += 1
        list_p = [dict_b[ele] for ele in A]
        return list_p

a = Solution()
a.anagramMappings([12,28,46,32,50],[50,12,32,46,28])
#print(a.anagramMappings([21,5,74,5,74,21],[21,5,74,74,5,21]))




#### Important to understand why not the below one!!! (Fails for the above commented test case)

# dict_a = {}
#         list_p = [0] * len(A)
#         i = 0
#         j = 0
#         for ele in A:
#             dict_a[ele] = i
#             i += 1
#         for ele_b in B:
#             if ele_b in dict_a.keys():
#                 index = dict_a[ele_b]
#                 list_p[index] = j
#                 j += 1
#         return list_p