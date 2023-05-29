package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getMinimumHealth' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY initial_players
 *  2. INTEGER_ARRAY new_players
 *  3. INTEGER rank
 */

func addPlayerToList(initial_players []int32, num int32) []int32 {
	if num < initial_players[0] {
		initial_players = append([]int32{num}, initial_players...)
		return initial_players
	}
	for i, player := range initial_players {
		if num >= player && i < len(initial_players)-1 && num <= initial_players[i+1] {
			initial_players = append(initial_players[:i+1], initial_players[i:]...)
			initial_players[i+1] = num
			break
		} else if num >= player && i == len(initial_players)-1 {
			initial_players = append(initial_players, num)
			break
		}
	}
	return initial_players
}

func getMinimumHealth(initial_players []int32, new_players []int32, rank int32) int64 {
	size := len(initial_players)
	total := initial_players[size-int(rank)]
	for _, nPlayer := range new_players {
		initial_players = addPlayerToList(initial_players, nPlayer)
		size := len(initial_players)
		total += initial_players[size-int(rank)]
	}
	return int64(total)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	initial_playersCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var initial_players []int32

	for i := 0; i < int(initial_playersCount); i++ {
		initial_playersItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		initial_playersItem := int32(initial_playersItemTemp)
		initial_players = append(initial_players, initial_playersItem)
	}

	new_playersCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var new_players []int32

	for i := 0; i < int(new_playersCount); i++ {
		new_playersItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		new_playersItem := int32(new_playersItemTemp)
		new_players = append(new_players, new_playersItem)
	}

	rankTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	rank := int32(rankTemp)

	result := getMinimumHealth(initial_players, new_players, rank)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
