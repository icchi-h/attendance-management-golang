/* Attendance Management Script with golang
 *
 * @author Haruyuki Ichino<mail@icchi.me>
 * @version 0.1
 * @date 2017/08/27
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/leekchan/timeutil"
)

// Constant
const LOG_DIR string = "./log/"
const OUTPUT_FILENAME_FORMAT string = "%Y-%m"
const OUTPUT_DATE_FORMAT string = "%Y/%m/%d(%a)"
const OUTPUT_TIME_FORMAT string = "%H:%M"
const START_COMMAND string = "start"
const FINISH_COMMAND string = "finish"
const SHOW_COMMAND string = "show"

func main() {

	// `os.Args` provides access to raw command-line
	var args []string = os.Args

	// コマンドライン引数の確認
	if len(args) != 2 {
		errorProcessing()
	}

	var now = time.Now()
	var outputFile string = filepath.Join(LOG_DIR, timeutil.Strftime(&now, OUTPUT_FILENAME_FORMAT)+".tsv")
	var today = timeutil.Strftime(&now, OUTPUT_DATE_FORMAT)
	var nowTime = timeutil.Strftime(&now, OUTPUT_TIME_FORMAT)

	if args[1] == START_COMMAND {
		// start command
		// fmt.Println("start processing")
		checkOutputFile(outputFile)

		if existTodaysData(outputFile, today) == 0 {
			// 通常の入社処理
			f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0755)
			check(err)
			defer f.Close()
			fmt.Fprint(f, today+"\t"+nowTime)
			fmt.Println("Recored: 本日の入社時刻 > " + nowTime)

		} else {
			fmt.Println("Error: 既に本日の入社時刻は入力されています")
		}
	} else if args[1] == FINISH_COMMAND {
		// finish command
		// fmt.Println("finish processing")
		checkOutputFile(outputFile)

		if existTodaysData(outputFile, today) == 2 {
			// 通常の退社処理
			f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0755)
			check(err)
			defer f.Close()
			fmt.Fprintln(f, "\t"+nowTime)
			fmt.Println("Recored: 本日の退社時刻 > " + nowTime)

		} else if existTodaysData(outputFile, today) == 3 {
			fmt.Println("Error: 既に本日の退社時刻は入力されています")
			os.Exit(0)
		} else {
			fmt.Println("Error: 本日の入社時刻が入力されていません")
			os.Exit(0)
		}
	} else if args[1] == SHOW_COMMAND {
		out, err := exec.Command("cat", outputFile).Output()
		check(err)
		fmt.Println(string(out))
	} else {
		fmt.Println("Command Error: Your command is '" + args[1] + "'\n")
		errorProcessing()
	}
}

func errorProcessing() {
	fmt.Println("Usage: \n\tams command\n\nThe commands are\n\t" + START_COMMAND + "\tRecord the start time\n\t" + FINISH_COMMAND + "\tRecord the finish time\n\t" + SHOW_COMMAND + "\tShow log file this month")
	os.Exit(0)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkOutputFile(file string) {
	// 出力ディレクトリの存在確認
	_, err := os.Stat(LOG_DIR)
	if err != nil {
		fmt.Println("Log directory does not exist.")
		os.Mkdir(LOG_DIR, 0755)
		fmt.Println("Created log directory '" + LOG_DIR + "'")
	}

	// 出力ファイルの存在確認
	_, err = os.Stat(file)
	if err != nil {
		// 存在しなければテンプレートファイルを作成
		fmt.Println("Log file does not exist.")
		f, err := os.Create(file)
		check(err)
		defer f.Close()
		fmt.Println("Created log file '" + file + "'")
		_, err = f.WriteString("Date\tStart Time\tFinish Time\n")
		f.Close()
	}
}

func existTodaysData(file string, date string) int {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error: Faild to load log file")
	}
	defer f.Close()

	// 最終行を取得
	sc := bufio.NewScanner(f)
	var lastLine = ""
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}
		lastLine = sc.Text()
	}

	// buf := make([]byte, 28)
	// stat, err := os.Stat(file)
	// start := stat.Size() - 28
	// _, err = f.ReadAt(buf, start)
	// var recentDay = strings.Split(string(buf), "\t")[0]

	// 最後の行から日付を取得
	var lastLineArr = strings.Split(lastLine, "\t")
	var recentDay = lastLineArr[0]
	// fmt.Println(recentDay)
	// fmt.Println(date)
	// fmt.Println(len(lastLineArr))

	if recentDay == date {
		return len(lastLineArr)
	} else {
		return 0
	}
}
