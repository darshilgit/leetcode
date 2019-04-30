class Solution:
    def combinationSum2(self, candidates, target):
        self.res = []
        candidates.sort()
        if not candidates:
            return []

        self.helper(candidates, 0, [], target)
        return self.res

    def helper(self, candidates, start, items, target):
        # print('c: {} start: {} items:{} target:{}'.format(candidates, start, items, target))
        if target == 0 and items not in self.res:
            self.res.append(items)

        for i in range(start, len(candidates)):
            # print('candidate {}: {}'.format(i,candidates[i]))
            if candidates[i] > target:
                return
            self.helper(candidates, i+1, items + [candidates[i]], target - candidates[i])