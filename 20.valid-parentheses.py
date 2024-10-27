class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        rule = {")":"(", "}":"{", "]":"["}
        for letter in s:
            if letter in rule:
                if len(stack) > 0:
                    top = stack.pop()
                    if top != rule[letter]:
                        return False
                else:
                    return False
            else:
                stack.append(letter)

        return not stack
