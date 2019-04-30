# You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
#
# Example 1:
#
# Input: coins = [1, 2, 5], amount = 11
#
# Output: 3
#
# Explanation: 11 = 5 + 5 + 1
# Example 2:
#
# Input: coins = [2], amount = 3
#
# Output: -1
class Solution:
    def coinChange(self, coins, amount):
        """
        :type coins: List[int]
        :type amount: int
        :rtype: int
        """
        amounts = []
        for i in range(amount+1):
            amounts.append(i)
        cache = [0] + [amount + 1] * amount
        print("Amounts:       " + str(amounts))
        print("Initial cache: " + str(cache))
        for i in range(1,amount+1):
            for coin in coins:
                print("i: " + str(i))
                print("coin: " + str(coin))
                if (i-coin) >= 0:
                    print("i-coin: " + str(i-coin))
                    print("cache[i]: " + str(cache[i]))
                    print("cache[i-coin] + 1: " + str(cache[i-coin] + 1))
                    cache[i] = min(cache[i],cache[i-coin] + 1)
                    print("after updating cache[i], cache: " + str(cache))
                    print("----------------------")
                else:
                    print("i-coin !>= 0")
        if cache[amount] > amount:
            return -1
        else:
            return cache[amount]


a = Solution()
print(a.coinChange([1,2,5],11))