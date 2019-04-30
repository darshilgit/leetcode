class Solution:
    def numUniqueEmails(self, emails):
        """
        :type emails: List[str]
        :rtype: int
        """
        valid_email_set = set()

        for email in emails:
            local_name,domain_name = email.split('@')
            if '+' in local_name:
                local_name = local_name.split('+',1)[0]
                local_name = local_name.replace('.','')
                valid_email_set.add(local_name + '@' + domain_name)
        return valid_email_set


a = Solution()
print(a.numUniqueEmails(["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]))
