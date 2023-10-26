from enum import Enum


class Tokens(Enum):
    paren = 'paren'
    alp = 'alp'
    dig = 'dig'

def tokenize(input_str):
    ind = 0
    tokens = []
    while ind < len(input_str):
        char = input_str[ind]
        if char == ' ':
            ind+=1
            continue
        if char == '(' or char== ')':
            tokens.append({
                'type':Tokens.paren.value,
                'value':char
            })
            ind+=1
            continue
        if char.isalpha():
            temp = ""
            while char.isalpha() and char!=' ' and ind < len(input_str):
                temp+=char
                ind+=1
                char = input_str[ind]
            tokens.append({
                'type':Tokens.alp.value,
                'value':temp
            })
            continue
        if char.isdigit():
            temp = ""
            while char.isdigit()  and ind < len(input_str):
                temp+=char
                ind+=1
                char = input_str[ind]
            tokens.append({
                'type':Tokens.dig.value,
                'value':temp
            })
            continue
        raise SyntaxError(f'Invalid keyword {char}')
    return tokens
 