class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        max_profit, min_price = 0, float('inf')
        for price in prices:
            # min_price stores the minimum cost so far
            min_price = min(min_price, price)
            profit = price - min_price
            #max_profit so far
            max_profit = max(max_profit, profit)
        return max_profit