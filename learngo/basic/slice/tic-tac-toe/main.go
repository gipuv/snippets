package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

const (
	Reset = "\033[0m"  // 终端颜色重置
	Red   = "\033[31m" // 红色，用于显示玩家X
	Blue  = "\033[34m" // 蓝色，用于显示玩家O
	Green = "\033[32m" // 绿色，用于显示标题和胜利信息
)

func main() {
	// 捕获 Ctrl+C 等信号，实现优雅退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\n\n👋 棋盘退出，再见！")
		os.Exit(0)
	}()

	// 用于从标准输入（键盘输入）逐行读取文本
	scanner := bufio.NewScanner(os.Stdin)

	for {
		clearScreen()
		fmt.Println(Green + "=== 欢迎来到井字棋游戏 ===" + Reset)
		fmt.Println("1. 开始游戏")
		fmt.Println("2. 退出")
		fmt.Print("请选择: ")

		if !scanner.Scan() {
			fmt.Println("\n👋 棋盘退出，再见！")
			return
		}

		// TrimSpace 去掉这行字符串首尾的所有空白字符（包括空格、换行、制表符等）
		// 获取当前 Scanner 读取到的那一行字符串内容
		// （注意：调用 scanner.Text() 之前，必须先调用 scanner.Scan() 成功读取一行，否则 scanner.Text() 为空）
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "2":
			fmt.Println("\n👋 棋盘退出，再见！")
			return
		case "1":
			playGame(scanner)
		default:
			fmt.Println("无效输入，请输入 1 或 2。按回车继续...")
			scanner.Scan()
		}
	}
}

// 游戏主循环
func playGame(scanner *bufio.Scanner) {
	// 初始化3x3棋盘，使用"_"表示空格
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	currentPlayer := "X" // 先手玩家X
	turn := 0            // 计数回合数

	for {
		clearScreen()
		fmt.Printf("🎮 玩家 %s 的回合\n", currentPlayer)
		printBoardWithCoords(board)

		row, col, err := getPlayerMove(scanner, board)
		if err != nil {
			// 输入被中断或错误，退出程序
			fmt.Println("\n👋 棋盘退出，再见！")
			os.Exit(0)
		}

		// 落子
		board[row][col] = currentPlayer
		turn++

		// 判断是否有赢家
		if winner := checkWinner(board); winner != "" {
			clearScreen()
			printBoardWithCoords(board)
			fmt.Printf(Green+"🎉 玩家 %s 获胜！\n"+Reset, winner)
			break
		} else if turn == 9 {
			// 棋盘满，平局
			clearScreen()
			printBoardWithCoords(board)
			fmt.Println("😐 平局！")
			break
		}

		// 切换玩家
		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}

	// 游戏结束，询问是否重开
	fmt.Print("\n是否再来一局？(y/n): ")
	if !scanner.Scan() {
		fmt.Println("\n👋 棋盘退出，再见！")
		return
	}
	if strings.ToLower(scanner.Text()) == "y" {
		playGame(scanner)
	}
}

// 获取玩家输入的行列坐标，并做合法性校验
func getPlayerMove(scanner *bufio.Scanner, board [][]string) (int, int, error) {
	for {
		fmt.Print("请输入 行 列 (0~2，以空格分隔): ")
		if !scanner.Scan() {
			// 输入被中断（例如 Ctrl+C）
			return 0, 0, errors.New("输入中断")
		}

		input := scanner.Text()
		parts := strings.Fields(input)

		// 输入格式必须是两个数字
		if len(parts) != 2 {
			fmt.Println("❌ 输入格式错误，请输入两个数字，例如: 1 2")
			continue
		}

		row, err1 := strconv.Atoi(parts[0])
		col, err2 := strconv.Atoi(parts[1])

		// 数字范围必须在0到2之间
		if err1 != nil || err2 != nil || row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("❌ 无效坐标，请输入 0 到 2 之间的数字。")
			continue
		}

		// 该位置必须为空
		if board[row][col] != "_" {
			fmt.Println("❌ 该位置已被占用，请重新选择。")
			continue
		}

		return row, col, nil
	}
}

// 带坐标打印棋盘，并为X、O着色
func printBoardWithCoords(board [][]string) {
	fmt.Println("\n   0 1 2")
	for i, row := range board {
		fmt.Printf("%d  ", i)
		for _, cell := range row {
			switch cell {
			case "X":
				fmt.Print(Red + "X" + Reset + " ")
			case "O":
				fmt.Print(Blue + "O" + Reset + " ")
			default:
				fmt.Print("_ ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// 判断是否有玩家获胜，返回获胜玩家 "X" 或 "O"，否则返回空字符串
func checkWinner(board [][]string) string {
	// 所有可能的获胜三连线坐标
	lines := [][][2]int{
		{{0, 0}, {0, 1}, {0, 2}}, // 行
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}}, // 列
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}}, // 对角线
		{{0, 2}, {1, 1}, {2, 0}},
	}

	for _, line := range lines {
		a, b, c := line[0], line[1], line[2]
		if board[a[0]][a[1]] != "_" &&
			board[a[0]][a[1]] == board[b[0]][b[1]] &&
			board[b[0]][b[1]] == board[c[0]][c[1]] {
			return board[a[0]][a[1]]
		}
	}
	return ""
}

// 清屏，支持Windows和类Unix系统
func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
