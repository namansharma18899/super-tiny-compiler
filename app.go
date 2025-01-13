package main

import (
    "fmt"
    "strings"
)

type Token struct {
    Type  string
    Value string
}

type Node struct {
    Type     string
    Value    string
    Left     *Node
    Right    *Node
    Name     string
    Params   []*Node
    Body     []*Node
    Context  map[string]string
}

func tokenizer(input string) []Token {
    var tokens []Token
    current := 0
    
    for current < len(input) {
        char := string(input[current])
        
        if char == " " {
            current++
            continue
        }
        
        if strings.Contains("0123456789", char) {
            value := ""
            for current < len(input) && strings.Contains("0123456789", string(input[current])) {
                value += string(input[current])
                current++
            }
            tokens = append(tokens, Token{Type: "number", Value: value})
            continue
        }
        
        if char == "+" {
            tokens = append(tokens, Token{Type: "operator", Value: "+"})
            current++
            continue
        }
        
        if char == "(" {
            tokens = append(tokens, Token{Type: "paren", Value: "("})
            current++
            continue
        }
        
        if char == ")" {
            tokens = append(tokens, Token{Type: "paren", Value: ")"})
            current++
            continue
        }
        
        if strings.Contains("abcdefghijklmnopqrstuvwxyz", strings.ToLower(char)) {
            value := ""
            for current < len(input) && strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", string(input[current])) {
                value += string(input[current])
                current++
            }
            tokens = append(tokens, Token{Type: "name", Value: value})
            continue
        }
        
        panic(fmt.Sprintf("Unknown character: %s", char))
    }
    
    return tokens
}

func parser(tokens []Token) *Node {
    current := 0
    
    var walk func() *Node
    walk = func() *Node {
        token := tokens[current]
        
        if token.Type == "number" {
            current++
            return &Node{
                Type:  "NumberLiteral",
                Value: token.Value,
            }
        }
        
        if token.Type == "paren" && token.Value == "(" {
            current++
            token = tokens[current]
            
            node := &Node{
                Type:   "CallExpression",
                Name:   token.Value,
                Params: []*Node{},
            }
            
            current++
            
            for tokens[current].Type != "paren" || (tokens[current].Type == "paren" && tokens[current].Value != ")") {
                node.Params = append(node.Params, walk())
            }
            
            current++
            
            return node
        }
        
        panic(fmt.Sprintf("Unknown token type: %s", token.Type))
    }
    
    ast := &Node{
        Type: "Program",
        Body: []*Node{walk()},
    }
    
    return ast
}

func transformer(ast *Node) *Node {
    newAst := &Node{
        Type: "Program",
        Body: []* Node{},
    }
    
    context := ast
    var traverse func(node *Node, parent *Node)
    
    traverse = func(node *Node, parent *Node) {
        switch node.Type {
        case "NumberLiteral":
            parent.Body = append(parent.Body, &Node{
                Type:  "NumberLiteral",
                Value: node.Value,
            })
        case "CallExpression":
            expression := &Node{
                Type: "CallExpression",
                Callee: &Node{
                    Type: "Identifier",
                    Name: node.Name,
                },
                Arguments: []*Node{},
            }
            
            for _, param := range node.Params {
                traverse(param, expression)
            }
            
            parent.Body = append(parent.Body, expression)
        }
    }
    
    for _, node := range ast.Body {
        traverse(node, newAst)
    }
    
    return newAst
}

func codeGenerator(node *Node) string {
    switch node.Type {
    case "Program":
        var result []string
        for _, n := range node.Body {
            result = append(result, codeGenerator(n))
        }
        return strings.Join(result, "\n")
    case "ExpressionStatement":
        return codeGenerator(node.Expression) + ";"
    case "CallExpression":
        var params []string
        for _, param := range node.Arguments {
            params = append(params, codeGenerator(param))
        }
        return fmt.Sprintf("%s(%s)", node.Callee.Name, strings.Join(params, ", "))
    case "Identifier":
        return node.Name
    case "NumberLiteral":
        return node.Value
    default:
        panic(fmt.Sprintf("Unknown node type: %s", node.Type))
    }
}

func compiler(input string) string {
    tokens := tokenizer(input)
    ast := parser(tokens)
    newAst := transformer(ast)
    output := codeGenerator(newAst)
    return output
}

func main() {
    input := "(add 2 (subtract 4 2))"
    output := compiler(input)
    fmt.Println(output)
}
