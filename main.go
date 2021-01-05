package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
)

func anyFileExists(dir string) (bool, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return false, err
	}

	if len(files) == 0 {
		return false, nil
	}

	return true, nil
}

func enumerateDirectories(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	ret := make([]string, 0)
	for _, file := range files {
		addondir := filepath.Join(dir, file.Name())
		ret = append(ret, addondir)
	}

	return ret, nil
}

func deployAddon(srcDir string, dstDir string) error {
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := copy.Copy(
			filepath.Join(srcDir, file.Name()),
			filepath.Join(dstDir, file.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	flag.Parse()
	numFlags := len(flag.Args())
	if numFlags != 2 {
		fmt.Fprintln(os.Stderr, "エラー: コマンドライン引数の数が不正です。")
		return
	}

	srcDir := flag.Arg(0)
	dstDir := flag.Arg(1)

	//すでにdstDirが示すディレクトリ内に何かファイルが存在する場合にはプログラムを終了する。
	fileExists, err := anyFileExists(dstDir)
	if err != nil {
		panic(err)
	}
	if fileExists == true {
		errMessage := fmt.Sprintf("エラー: %v にはすでにファイルが存在します。", dstDir)
		fmt.Fprintln(os.Stderr, errMessage)
		return
	}

	dirs, err := enumerateDirectories(srcDir)
	if err != nil {
		panic(err)
	}

	for _, dir := range dirs {
		err := deployAddon(filepath.Join(dir, "addon"), dstDir)
		if err != nil {
			fmt.Printf("処理をスキップします。 %v\n", dir)
			continue
		}
	}

	fmt.Println("アドオンの配備が完了しました。")
}
