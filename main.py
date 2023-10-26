import json
from compiler import compiler


if __name__=='__main__':
    inpstr = '(add a (sub b c))'
    output = compiler(inpstr)
    print(json.dumps(output, indent=3))