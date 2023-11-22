def parse(tokens: list):
    AST = {"type": "Program", "body": []}
    return AST


def parser(tokens):
  current = 0
  def walk():
    token = tokens[current]
    if (token.type == 'dig'):
        current+=1
        return {
            'type': 'NumberLiteral',
            'value': token.value
        }
    if (token.type == 'paren' and token.value == '('):
        token = tokens[current+1]
        expression = {
            'type': 'CallExpression',
            'name': token.value,
            'params': []
        }
        token = tokens[current+1]
        while (token.value != ')'):
            expression.params.append(walk())
            token = tokens[current]
        current+=1
        return expression
    else:
        raise TypeError(f'Unknow Token {token}')
    # throw new TypeError(`Unknown token: '${token}'`);
  ast = {
    'type': 'Program',
    'body': [walk()]
  }
  return ast