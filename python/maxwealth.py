## Traversing An Array


class Solution:
    def maximumWealth(self, accounts: List[List[int]]) -> int:
        max_wealth_so_far = 0
        for account in accounts:
            curr_customer_wealth = sum(account)
            max_wealth_so_far = max(max_wealth_so_far, curr_customer_wealth)
        return max_wealth_so_far