package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func GetAnswer(content string) (answer int) {
	defer func() {
		// recover
		if err := recover(); err != nil {
			fmt.Println("calculator", err)
			panic(err)
			answer = 0
		}
	}()
	answer = 0
	index := strings.Index(content, "=")
	if index < 0 {
		return
	}
	answerData := content[:index]
	calData := content[index+1:]
	fmt.Println("getanswer:", answerData, "\t", calData, "\tcal:", Count(answerData))
	calanswer, err := strconv.ParseFloat(calData, 64)
	if err != nil {
		return
	}
	if Count(answerData) == calanswer {
		answer = 1
	}
	return
}

func Count(content string) float64 {
	var arr []string = genPostfixExpression(content)
	return countPostfixExpression(arr)
}

type StackNode struct {
	Data interface{}
	next *StackNode
}

type LinkStack struct {
	top   *StackNode
	Count int
}

func (linkStack *LinkStack) Init() {
	linkStack.top = nil
	linkStack.Count = 0
}

func (linkStack *LinkStack) Push(data interface{}) {
	var node *StackNode = new(StackNode)
	node.Data = data
	node.next = linkStack.top
	linkStack.top = node
	linkStack.Count++
}

func (linkStack *LinkStack) Pop() interface{} {
	if linkStack.top == nil {
		return nil
	}
	returnData := linkStack.top.Data
	linkStack.top = linkStack.top.next
	linkStack.Count--
	return returnData
}

//Look up the top element in the stack, but not pop.
func (linkStack *LinkStack) TopData() interface{} {
	if linkStack.top == nil {
		return nil
	}
	return linkStack.top.Data
}

func (linkStack *LinkStack) FMT() {
	fmt.Print("linkstack")
	i := linkStack.top
	for i != nil {
		fmt.Print("\t", i.Data.(string))
		i = i.next
	}
	fmt.Println("end")
}

/*
generate postfix expression
*/
func genPostfixExpression(content string) (result []string) {
	// 符号栈
	var stack LinkStack
	stack.Init()
	// 中缀表达式
	var splitstr []string = genInfixExpresion(content)
	// fmt.Println(splitstr)
	// 数据存储
	// var datas []string

	for _, str := range splitstr {
		// stack.FMT()
		// fmt.Println(result)
		if isNumberString(str) { //数字直接处理
			result = append(result, str)
		} else { //字符入栈处理
			/*
			   四种情况入栈
			   1 左括号直接入栈
			   2 栈内为空直接入栈
			   3 栈顶为左括号，直接入栈
			   4 当前元素不为右括号时，在比较栈顶元素与当前元素，如果当前元素大，直接入栈。
			*/
			if str == "(" || stack.TopData() == nil || stack.TopData().(string) == "(" || (str != ")" && compareOperator(str, stack.TopData().(string)) == 1) {
				stack.Push(str)
			} else if str == ")" {
				/*
				   当前元素为右括号时，提取操作符，直到碰见左括号
				*/
				flag := true
				for {
					pop := stack.Pop()
					if pop == nil {
						break
					}
					if pop.(string) == "(" {
						flag = false
						break
					} else {
						result = append(result, pop.(string))
					}
				}
				if flag {
					panic("bad match for brackets!")
				}
			} else {
				/*
				 */
				for {
					pop := stack.TopData()

					if pop == nil || pop.(string) == "(" {
						stack.Push(str)
						break
					} else {
						if compareOperator(str, pop.(string)) != 1 {
							result = append(result, stack.Pop().(string))
						}
					}
					// if pop != nil && compareOperator(str, pop.(string)) != 1 {
					//  result = append(result, stack.Pop().(string))
					// } else {
					//  stack.Push(str)
					//  break
					// }
				}
			}
		}
	}

	for {
		pop := stack.Pop()
		if pop != nil {
			result = append(result, pop.(string))
		} else {
			break
		}
	}
	fmt.Println("result", result)
	return
}

/*
Count postfix expression
*/
func countPostfixExpression(arr []string) float64 {
	var stack LinkStack
	stack.Init()
	for _, v := range arr {
		if isNumberString(v) {
			if f, err := strconv.ParseFloat(v, 64); err != nil {
				panic("operatin process go wrong.")
			} else {
				stack.Push(f)
			}
		} else {
			p1 := stack.Pop()
			p2 := stack.Pop()
			if p1 == nil || p2 == nil {
				panic("bad expression")
			}

			p3 := normalCalculate(p2.(float64), p1.(float64), v)
			stack.Push(p3)
		}
	}
	res := stack.Pop().(float64)
	return res
}

/*
infix expression
*/
func genInfixExpresion(content string) (result []string) {
	bys := []byte(content)
	var tmp string
	for i := 0; i < len(bys); i++ {
		if !isNumber(bys[i]) {
			if tmp != "" {
				result = append(result, tmp)
				tmp = ""
			}
			result = append(result, string(bys[i]))
		} else {
			tmp = tmp + string(bys[i])
		}
	}
	if tmp != "" {
		result = append(result, tmp)
	}
	return
}

/*
juge is number
*/
func isNumber(o1 byte) bool {
	if o1 == '+' || o1 == '-' || o1 == '*' || o1 == '/' || o1 == '(' || o1 == ')' {
		return false
	} else {
		return true
	}
}

func isNumberString(o1 string) bool {
	if o1 == "+" || o1 == "-" || o1 == "*" || o1 == "/" || o1 == "(" || o1 == ")" {
		return false
	} else {
		return true
	}
}

/*
if return 1, o1 > o2.
if return 0, o1 = 02
if return -1, o1 < o2
*/
func compareOperator(o1, o2 string) int {
	// + - * /
	var o1Priority int
	if o1 == "+" || o1 == "-" {
		o1Priority = 1
	} else {
		o1Priority = 2
	}
	var o2Priority int
	if o2 == "+" || o2 == "-" {
		o2Priority = 1
	} else {
		o2Priority = 2
	}
	if o1Priority > o2Priority {
		return 1
	} else if o1Priority == o2Priority {
		return 0
	} else {
		return -1
	}
}

func normalCalculate(a, b float64, operation string) float64 {
	switch operation {
	case "*":
		return a * b
	case "-":
		return a - b
	case "+":
		return a + b
	case "/":
		return a / b
	default:
		panic("invalid operator")
	}
}
