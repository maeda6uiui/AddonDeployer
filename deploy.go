package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
	"github.com/urfave/cli/v2"
)

func anyFileExists(targetDir string) (bool, error) {
	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		return false, err
	}

	if len(files) == 0 {
		return false, nil
	}

	return true, nil
}

func enumerateDirectories(targetRootDir string) ([]string, error) {
	files, err := ioutil.ReadDir(targetRootDir)
	if err != nil {
		return nil, err
	}

	ret := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			addonDir := filepath.Join(targetRootDir, file.Name())
			ret = append(ret, addonDir)
		}
	}

	return ret, nil
}

func deployAddon(inputRootDir string, outputRootDir string) error {
	subdirs, err := ioutil.ReadDir(inputRootDir)
	if err != nil {
		return err
	}

	for _, subdir := range subdirs {
		err := copy.Copy(
			filepath.Join(inputRootDir, subdir.Name()),
			filepath.Join(outputRootDir, subdir.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func appAction(c *cli.Context) error {
	inputRootDir := c.String("inputRootDir")
	outputRootDir := c.String("outputRootDir")

	if inputRootDir == "" {
		fmt.Println("エラー: 入力ディレクトリを指定してください")
		return nil
	}
	if outputRootDir == "" {
		fmt.Println("エラー: 出力ディレクトリを指定してください")
		return nil
	}

	//すでにoutputRootDirが示すディレクトリ内に何かファイルが存在する場合にはプログラムを終了する
	fileExists, err := anyFileExists(outputRootDir)
	if err != nil {
		return err
	}
	if fileExists {
		errMessage := fmt.Sprintf("エラー: %v にはすでにファイルが存在します", outputRootDir)
		fmt.Fprintln(os.Stderr, errMessage)

		return nil
	}

	subdirs, err := enumerateDirectories(inputRootDir)
	if err != nil {
		panic(err)
	}

	for _, subdir := range subdirs {
		err := deployAddon(filepath.Join(subdir, "addon"), outputRootDir)
		if err != nil {
			fmt.Printf("処理をスキップします %v\n", subdir)
			continue
		}
	}

	fmt.Println("アドオンの配備が完了しました")

	return nil
}

func main() {
	app := &cli.App{
		Name:    "Addon Deployer",
		Version: "v1.0.0",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "inputRootDir",
				Aliases: []string{"i"},
				Usage:   "Input root directory",
			},
			&cli.StringFlag{
				Name:    "outputRootDir",
				Aliases: []string{"o"},
				Usage:   "Output root directory",
			},
		},

		Action: appAction,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
