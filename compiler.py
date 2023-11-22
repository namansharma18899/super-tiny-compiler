from tokenizer import tokenize
from token_parser import parse
from transformer import transformer


def compiler(input_str):
    # 1. Lexical Analysis -
    #     Breaks the input into basic syntax of the language (array of objects)
    tokens = tokenize(input_str)
    # 2. Syntax Analysis
    #   Converts the list of tokens to a AST(Abstract Syntax Tree) which represents our program
    list_ast = parse(tokens)
    # 3. Transformation
    jsAST = transformer(list_ast)
    # 4. Code Generation
    return tokens
