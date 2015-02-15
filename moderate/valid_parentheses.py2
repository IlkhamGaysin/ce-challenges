"""Given a string comprising just of the characters (,),{,},[,]
determine if it is well-formed or not."""
import fileinput

for line in fileinput.input():
    line = line.rstrip('\n')
    stack = []
    f = True
    for i in line:
        if i in ")]}":
            if (len(stack) < 1 or stack[-1] != i):
                f = False
                break
            else:
                stack.pop()
        else:
            if i == '(':
                stack.append(')')
            elif i == '[':
                stack.append(']')
            elif i == '{':
                stack.append('}')
            else:
                f = False
                break
    if len(stack) > 0:
        f = False
    print f
